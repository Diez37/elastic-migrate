package loaders

import "github.com/diez37/elastic-migrate/services/migrator/migration"

type Loader interface {
	Load() ([]*migration.Migration, error)
}
