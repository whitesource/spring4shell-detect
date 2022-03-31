package gradle

import (
	"github.com/whitesource/spring4shell-detect/fs"
	"github.com/whitesource/spring4shell-detect/fs/match"
)

func Query() *fs.Query {
	return &fs.Query{
		Filename: match.NewExtensionMatcher("jar", "gem"),
	}
}
