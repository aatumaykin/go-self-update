package selfupdate

import (
	"cmp"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/aatumaykin/go-self-update/selfupdate/gitea"
	"github.com/aatumaykin/go-self-update/selfupdate/github"
	"github.com/aatumaykin/go-self-update/selfupdate/release"
	"github.com/blang/semver"
	"github.com/inconshreveable/go-update"
)

type (
	repoClient interface {
		GetVersionUrl(version string) string
		GetRelease(ctx context.Context, version string) (release.Release, error)
	}

	RepositoryType string

	Release struct {
		Version       semver.Version
		AssetURL      string
		AssetByteSize int
		PageURL       string
		ReleaseNotes  string
		Name          string
		PublishedAt   time.Time
	}

	Filter struct {
		Template string
		Values   map[string]string
	}
	Updater struct {
		httpClient     *http.Client
		logger         *slog.Logger
		repositoryType RepositoryType
		apiBaseURL     string
		owner          string
		repo           string
		filter         *Filter
	}

	Config struct {
		HTTPClient     *http.Client
		Logger         *slog.Logger
		RepositoryType RepositoryType
		Filter         *Filter
		APIBaseURL     string
		Owner          string
		Repo           string
	}
)

const (
	Github RepositoryType = "GitHub"
	Gitea  RepositoryType = "Gitea"

	latest = "latest"
)

func New(config Config) (*Updater, error) {
	return &Updater{
		httpClient:     cmp.Or(config.HTTPClient, http.DefaultClient),
		repositoryType: cmp.Or(config.RepositoryType, Github),
		apiBaseURL:     config.APIBaseURL,
		logger:         cmp.Or(config.Logger, slog.Default()),
		owner:          config.Owner,
		repo:           config.Repo,
		filter: cmp.Or(config.Filter, &Filter{
			Template: "{{.Name}}-{{.Version}}-{{.OS}}-{{.Arch}}",
		}),
	}, nil
}

func (u *Updater) CheckVersion(ctx context.Context, version string) (*Release, error) {
	version = cmp.Or(version, latest)

	filter, err := u.getAssetNamePattern(version)
	if err != nil {
		return nil, err
	}

	var rc repoClient

	switch u.repositoryType {
	case Github:
		rc = github.New(github.Config{
			APIBaseURL: u.apiBaseURL,
			Filter:     filter,
			Owner:      u.owner,
			Repo:       u.repo,
		})
	case Gitea:
		rc = gitea.New(gitea.Config{
			APIBaseURL: u.apiBaseURL,
			Filter:     filter,
			Owner:      u.owner,
			Repo:       u.repo,
		})
	default:
		return nil, fmt.Errorf("unsupported repository type: %s", u.repositoryType)
	}

	r, err := rc.GetRelease(ctx, version)
	if err != nil {
		return nil, err
	}

	asset, found := r.FindAsset(filter)
	if !found {
		return nil, fmt.Errorf("asset not found")
	}

	result := &Release{
		Version:       r.GetVersion(),
		Name:          r.GetName(),
		PageURL:       r.GetPageURL(),
		ReleaseNotes:  r.GetReleaseNotes(),
		AssetURL:      asset.GetDownloadURL(),
		AssetByteSize: asset.GetSize(),
		PublishedAt:   r.GetPublishedAt(),
	}

	return result, nil
}

func (u *Updater) UpdateTo(ctx context.Context, rel *Release, updateOpts *update.Options) error {
	u.logger.InfoContext(ctx, "Downloading", "url", rel.AssetURL, "size", rel.AssetByteSize)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rel.AssetURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := u.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if updateOpts == nil {
		updateOpts = &update.Options{}
	}

	u.logger.InfoContext(ctx, "Applying update")
	err = update.Apply(resp.Body, *updateOpts)
	if err != nil {
		return fmt.Errorf("failed to apply update: %w", err)
	}

	u.logger.InfoContext(ctx, "Update applied")

	return nil
}

func (u *Updater) getAssetNamePattern(version string) (string, error) {
	tpl, err := template.New("name").Parse(u.filter.Template)
	if err != nil {
		return "", err
	}

	if _, ok := u.filter.Values["Name"]; !ok {
		u.filter.Values["Name"] = os.Args[0]
	}

	if _, ok := u.filter.Values["OS"]; !ok {
		u.filter.Values["OS"] = runtime.GOOS
	}

	if _, ok := u.filter.Values["Arch"]; !ok {
		u.filter.Values["Arch"] = runtime.GOARCH
	}

	if _, ok := u.filter.Values["Version"]; !ok {
		u.filter.Values["Version"] = version
	}

	var buf strings.Builder
	err = tpl.Execute(&buf, u.filter.Values)

	return buf.String(), err
}
