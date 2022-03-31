package ruby

import (
	"github.com/whitesource/spring4shell-detect/fs"
	"github.com/whitesource/spring4shell-detect/fs/match"
)

func Query() *fs.Query {
	return &fs.Query{
		Filename: match.NewSimpleNameMatcher("Gemfile.lock", "gems.locked"),
	}
}
