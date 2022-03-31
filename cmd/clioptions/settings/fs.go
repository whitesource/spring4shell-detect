package settings

import (
	"github.com/go-logr/logr"
	"github.com/whitesource/spring4shell-detect/fs"
	"github.com/whitesource/spring4shell-detect/operations"
	fsop "github.com/whitesource/spring4shell-detect/operations/fs"
	rc "github.com/whitesource/spring4shell-detect/records"
	fsscreen "github.com/whitesource/spring4shell-detect/screening/fs"
	"github.com/whitesource/spring4shell-detect/utils/exec"
)

type FilesystemResolver struct {
	Disabled bool
}

func (r FilesystemResolver) Queries() map[rc.Organ]*fs.Query {
	if r.Disabled {
		return nil
	}

	return map[rc.Organ]*fs.Query{rc.OFS: fsscreen.Query()}
}

func (r FilesystemResolver) Surgeons(logger logr.Logger, commander exec.Commander) map[rc.Organ]operations.Surgeon {
	if r.Disabled {
		return nil
	}

	return map[rc.Organ]operations.Surgeon{
		rc.OFS: fsop.NewSurgeon(logger, commander),
	}
}
