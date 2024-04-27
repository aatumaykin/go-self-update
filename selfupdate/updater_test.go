package selfupdate

import (
	"context"
	"testing"
	"time"

	"github.com/blang/semver"
)

func TestUpdater_CheckVersion(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		slug    string
		version string
		want    *Release
		wantErr bool
	}{
		{
			name: "gitea: latest should return test-darwin-arm64",
			config: Config{
				RepositoryType: Gitea,
				Owner:          "tumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "darwin",
						"Arch": "arm64",
					},
				},
			},
			version: "",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://gitea.com/tumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://gitea.com/tumaykin/test-repository/releases/download/1.0.0/test-darwin-arm64",
				AssetByteSize: 2029586,
				PublishedAt:   time.Date(2024, 4, 26, 22, 45, 00, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "gitea: latest should return test-darwin-amd64",
			config: Config{
				RepositoryType: Gitea,
				Owner:          "tumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "darwin",
						"Arch": "amd64",
					},
				},
			},
			version: "",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://gitea.com/tumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://gitea.com/tumaykin/test-repository/releases/download/1.0.0/test-darwin-amd64",
				AssetByteSize: 2026240,
				PublishedAt:   time.Date(2024, 4, 26, 22, 45, 00, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "gitea: latest should return test-linux-amd64",
			config: Config{
				RepositoryType: Gitea,
				Owner:          "tumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "linux",
						"Arch": "amd64",
					},
				},
			},
			version: "",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://gitea.com/tumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://gitea.com/tumaykin/test-repository/releases/download/1.0.0/test-linux-amd64",
				AssetByteSize: 1897953,
				PublishedAt:   time.Date(2024, 4, 26, 22, 45, 00, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "gitea: latest should return test-linux-arm64",
			config: Config{
				RepositoryType: Gitea,
				Owner:          "tumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "linux",
						"Arch": "arm64",
					},
				},
			},
			version: "",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://gitea.com/tumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://gitea.com/tumaykin/test-repository/releases/download/1.0.0/test-linux-arm64",
				AssetByteSize: 1946270,
				PublishedAt:   time.Date(2024, 4, 26, 22, 45, 00, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "gitea: 1.0.0 should return test-darwin-arm64",
			config: Config{
				RepositoryType: Gitea,
				Owner:          "tumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "darwin",
						"Arch": "arm64",
					},
				},
			},
			version: "1.0.0",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://gitea.com/tumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://gitea.com/tumaykin/test-repository/releases/download/1.0.0/test-darwin-arm64",
				AssetByteSize: 2029586,
				PublishedAt:   time.Date(2024, 4, 26, 22, 45, 00, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "github: latest should return test-darwin-arm64",
			config: Config{
				RepositoryType: Github,
				Owner:          "aatumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "darwin",
						"Arch": "arm64",
					},
				},
			},
			version: "",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://github.com/aatumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://github.com/aatumaykin/test-repository/releases/download/1.0.0/test-darwin-arm64",
				AssetByteSize: 2029586,
				PublishedAt:   time.Date(2024, 4, 27, 9, 54, 23, 0, time.UTC),
				ReleaseNotes:  "test release",
			},
			wantErr: false,
		},
		{
			name: "github: latest should return test-darwin-amd64",
			config: Config{
				RepositoryType: Github,
				Owner:          "aatumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "darwin",
						"Arch": "amd64",
					},
				},
			},
			version: "",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://github.com/aatumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://github.com/aatumaykin/test-repository/releases/download/1.0.0/test-darwin-amd64",
				AssetByteSize: 2026240,
				PublishedAt:   time.Date(2024, 4, 27, 9, 54, 23, 0, time.UTC),
				ReleaseNotes:  "test release",
			},
			wantErr: false,
		},
		{
			name: "github: latest should return test-linux-amd64",
			config: Config{
				RepositoryType: Github,
				Owner:          "aatumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "linux",
						"Arch": "amd64",
					},
				},
			},
			version: "",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://github.com/aatumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://github.com/aatumaykin/test-repository/releases/download/1.0.0/test-linux-amd64",
				AssetByteSize: 1897953,
				PublishedAt:   time.Date(2024, 4, 27, 9, 54, 23, 0, time.UTC),
				ReleaseNotes:  "test release",
			},
			wantErr: false,
		},
		{
			name: "github: latest should return test-linux-arm64",
			config: Config{
				RepositoryType: Github,
				Owner:          "aatumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "linux",
						"Arch": "arm64",
					},
				},
			},
			version: "",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://github.com/aatumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://github.com/aatumaykin/test-repository/releases/download/1.0.0/test-linux-arm64",
				AssetByteSize: 1946270,
				PublishedAt:   time.Date(2024, 4, 27, 9, 54, 23, 0, time.UTC),
				ReleaseNotes:  "test release",
			},
			wantErr: false,
		},
		{
			name: "github: 1.0.0 should return test-linux-arm64",
			config: Config{
				RepositoryType: Github,
				Owner:          "aatumaykin",
				Repo:           "test-repository",
				Filter: &Filter{
					Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
					Values: map[string]string{
						"Name": "test",
						"OS":   "linux",
						"Arch": "arm64",
					},
				},
			},
			version: "1.0.0",
			want: &Release{
				Version: func() semver.Version {
					v, _ := semver.New("1.0.0")
					return *v
				}(),
				Name:          "1.0.0",
				PageURL:       "https://github.com/aatumaykin/test-repository/releases/tag/1.0.0",
				AssetURL:      "https://github.com/aatumaykin/test-repository/releases/download/1.0.0/test-linux-arm64",
				AssetByteSize: 1946270,
				PublishedAt:   time.Date(2024, 4, 27, 9, 54, 23, 0, time.UTC),
				ReleaseNotes:  "test release",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := New(tt.config)
			ctx := context.Background()
			r, err := u.CheckVersion(ctx, tt.version)

			if (err != nil) != tt.wantErr {
				t.Errorf("CheckVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				if r == nil {
					t.Errorf("CheckVersion() got = %v, want %v", r, tt.want)
					return
				}

				if r.Name != tt.want.Name {
					t.Errorf("CheckVersion() got = %v, want %v", r, tt.want)
				}

				if r.ReleaseNotes != tt.want.ReleaseNotes {
					t.Errorf("CheckVersion() got = %v, want %v", r, tt.want)
				}

				if r.PageURL != tt.want.PageURL {
					t.Errorf("CheckVersion() got = %v, want %v", r, tt.want)
				}

				if r.Version.String() != tt.want.Version.String() {
					t.Errorf("CheckVersion() got = %v, want %v", r, tt.want)
				}

				if r.AssetURL != tt.want.AssetURL {
					t.Errorf("CheckVersion() got = %v, want %v", r, tt.want)
				}

				if r.AssetByteSize != tt.want.AssetByteSize {
					t.Errorf("CheckVersion() got = %v, want %v", r, tt.want)
				}

				if r.PublishedAt != tt.want.PublishedAt {
					t.Errorf("CheckVersion() got = %v, want %v", r, tt.want)
				}
			}
		})
	}
}
