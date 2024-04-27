# go-self-update: Build self-updating Go programs

Package update provides functionality to implement secure, self-updating Go programs (or other single-file targets) A program can update itself by replacing its executable file with a new version.

Example of updating:

```go
import (
	"github.com/aatumaykin/go-self-update/selfupdate"
)

func doUpdate() {
	config := selfupdate.Config{
		// set repository type
		RepositoryType: selfupdate.Github,
		// set repository owner and name
		Owner: "aatumaykin",
		Repo:  "test-repository",
		// template for asset name
		Filter: &selfupdate.Filter{
			Template: "{{.Name}}-{{.OS}}-{{.Arch}}",
			// values for template
			Values: map[string]string{
				"Name": "test",
				"OS":   "", // if empty, runtime.GOOS will be used.
				"Arch": "", // if empty, runtime.GOARCH will be used.
			},
		},
	}

	sf, _ := selfupdate.New(config)
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
}
```

## Features

- features from `github.com/inconshreveable/go-update`
- work in Github and Gitea repositories