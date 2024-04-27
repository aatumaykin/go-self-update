package github

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aatumaykin/go-self-update/selfupdate/release"
)

type (
	Config struct {
		APIBaseURL string
		Filter     string
		Owner      string
		Repo       string
		HTTPClient *http.Client
	}
	Client struct {
		client     *http.Client
		apiBaseURL string
		filter     string
		owner      string
		repo       string
	}
)

const latest = "latest"

func New(config Config) *Client {
	return &Client{
		client:     http.DefaultClient,
		apiBaseURL: cmp.Or(config.APIBaseURL, "https://api.github.com"),
		filter:     config.Filter,
		owner:      config.Owner,
		repo:       config.Repo,
	}
}

func (c *Client) GetVersionUrl(version string) string {
	if version == latest {
		return fmt.Sprintf("/repos/%s/%s/releases/latest", c.owner, c.repo)
	}

	return fmt.Sprintf("/repos/%s/%s/releases/tags/%s", c.owner, c.repo, version)
}

func (c *Client) GetRelease(ctx context.Context, version string) (release.Release, error) {
	url := strings.TrimSuffix(c.apiBaseURL, "/") + c.GetVersionUrl(version)

	var r Release

	err := request(ctx, c.client, url, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func request(ctx context.Context, client *http.Client, url string, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get release: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to GET %v: %w", url, err)
	}

	if err := json.Unmarshal(body, &v); err != nil {
		return fmt.Errorf("failed to GET %v: %w", url, err)
	}

	return nil
}
