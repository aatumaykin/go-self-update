package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aatumaykin/go-self-update/selfupdate"
)

func main() {
	sf, _ := selfupdate.New(selfupdate.Config{
		RepositoryType: selfupdate.Github,
		Owner:          "aatumaykin",
		Repo:           "test-repository",
		Filter: &selfupdate.Filter{
			Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
			Values: map[string]string{
				"Name": "test",
			},
		},
	})

	ctx := context.Background()
	release, err := sf.CheckVersion(ctx, "")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Available release: ", "version", release.Version.String(), "url", release.AssetURL)
}
