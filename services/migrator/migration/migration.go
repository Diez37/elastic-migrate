package migration

import (
	"github.com/diez37/elastic-migrate/services/migrator/actions"
	"github.com/diez37/elastic-migrate/services/migrator/version"
)

type Migration struct {
	version *version.Version
	up      []actions.Action
	down    []actions.Action
}

func (migration *Migration) Version() *version.Version {
	return migration.version
}

func (migration *Migration) SetVersion(version *version.Version) *Migration {
	migration.version = version

	return migration
}

func (migration *Migration) Up() []actions.Action {
	return migration.up
}

func (migration *Migration) AddUp(up ...actions.Action) *Migration {
	migration.up = append(migration.up, up...)

	return migration
}

func (migration *Migration) Down() []actions.Action {
	return migration.down
}

func (migration *Migration) AddDown(down ...actions.Action) *Migration {
	migration.down = append(migration.down, down...)

	return migration
}
