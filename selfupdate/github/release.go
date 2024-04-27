package github

import (
	"strings"
	"time"

	"github.com/aatumaykin/go-self-update/selfupdate/release"
	"github.com/blang/semver"
)

type Release struct {
	TagName   string    `json:"tag_name"`
	Name      string    `json:"name"`
	Body      string    `json:"body"`
	URL       string    `json:"html_url"`
	Published time.Time `json:"published_at"`
	Assets    []Asset   `json:"assets"`
}

func (r *Release) GetName() string {
	return r.Name
}

func (r *Release) GetTagName() string {
	return r.TagName
}

func (r *Release) GetVersion() semver.Version {
	version := strings.TrimPrefix(r.TagName, "v")
	semVer, _ := semver.Parse(version)

	return semVer
}

func (r *Release) GetPageURL() string {
	return r.URL
}
func (r *Release) GetReleaseNotes() string {
	return r.Body
}

func (r *Release) FindAsset(name string) (release.Asset, bool) {
	for _, asset := range r.Assets {
		if asset.Name == name {
			return &asset, true
		}
	}

	return nil, false
}

func (r *Release) GetPublishedAt() time.Time {
	return r.Published
}
