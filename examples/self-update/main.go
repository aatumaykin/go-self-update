package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/aatumaykin/go-self-update/selfupdate"
)

type Flags struct {
	Update      bool
	CheckUpdate bool
}

func main() {
	var f Flags

	flag.BoolVar(&f.Update, "update", false, "Update bot and exit")
	flag.BoolVar(&f.CheckUpdate, "check-update", false, "Check for updates and exit")

	flag.Parse()

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
	if f.CheckUpdate {
		release, err := sf.CheckVersion(ctx, "")
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		slog.Info("Available release: ", "version", release.Version.String())
		os.Exit(0)
	}

	if f.Update {
		release, err := sf.CheckVersion(ctx, "")
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		slog.Info("Available release: ", "version", release.Version.String())

		err = sf.UpdateTo(ctx, release, nil)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		os.Exit(0)
	}

	fmt.Println("Nothing to do")
}
