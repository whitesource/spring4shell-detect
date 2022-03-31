package settings

import (
	"github.com/go-logr/logr"
	"github.com/whitesource/spring4shell-detect/fs"
	"github.com/whitesource/spring4shell-detect/operations"
	rubyS "github.com/whitesource/spring4shell-detect/operations/ruby"
	rc "github.com/whitesource/spring4shell-detect/records"
	rubyQ "github.com/whitesource/spring4shell-detect/screening/ruby"
	"github.com/whitesource/spring4shell-detect/utils/exec"
)

type RubyResolver struct {
	Disabled bool
}

func (r RubyResolver) Queries() map[rc.Organ]*fs.Query {
	if r.Disabled {
		return nil
	}

	return map[rc.Organ]*fs.Query{rc.ORuby: rubyQ.Query()}
}

func (r RubyResolver) Surgeons(logger logr.Logger, commander exec.Commander) map[rc.Organ]operations.Surgeon {
	if r.Disabled {
		return nil
	}

	return map[rc.Organ]operations.Surgeon{
		rc.ORuby: rubyS.NewSurgeon(logger, commander),
	}
}
