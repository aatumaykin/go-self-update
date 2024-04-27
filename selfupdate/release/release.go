package release

import (
	"time"

	"github.com/blang/semver"
)

type (
	Release interface {
		GetName() string
		GetTagName() string
		GetVersion() semver.Version
		GetPageURL() string
		GetReleaseNotes() string
		GetPublishedAt() time.Time
		FindAsset(name string) (Asset, bool)
	}

	Asset interface {
		GetName() string
		GetSize() int
		GetDownloadURL() string
	}
)
