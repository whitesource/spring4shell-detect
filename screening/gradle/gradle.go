package gradle

import (
	"github.com/whitesource/spring4shell-detect/fs"
	"github.com/whitesource/spring4shell-detect/fs/match"
)

func Query() *fs.Query {
	return &fs.Query{
		Filename: match.NewSimpleNameMatcher(
			"build.gradle", "build.gradle.kts", "settings.gradle", "settings.gradle.kts",
		),
	}
}
