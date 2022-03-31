package settings

import (
	"github.com/go-logr/logr"
	"github.com/whitesource/spring4shell-detect/fs"
	"github.com/whitesource/spring4shell-detect/operations"
	mavenS "github.com/whitesource/spring4shell-detect/operations/maven"
	rc "github.com/whitesource/spring4shell-detect/records"
	mavenQ "github.com/whitesource/spring4shell-detect/screening/maven"
	"github.com/whitesource/spring4shell-detect/utils/exec"
)

type MavenResolver struct {
	Disabled       bool
	AdditionalArgs []string
	Scopes         struct {
		Include []string
		Exclude []string
	}
}

func (r MavenResolver) Queries() map[rc.Organ]*fs.Query {
	if r.Disabled {
		return nil
	}

	return map[rc.Organ]*fs.Query{rc.OMaven: mavenQ.Query()}
}

func (r MavenResolver) Surgeons(logger logr.Logger, commander exec.Commander) map[rc.Organ]operations.Surgeon {
	if r.Disabled {
		return nil
	}

	return map[rc.Organ]operations.Surgeon{
		rc.OMaven: mavenS.NewSurgeon(logger, commander, r.AdditionalArgs, r.Scopes.Include, r.Scopes.Exclude),
	}
}
