package migrator

import (
	"context"
	"errors"
	loaders2 "github.com/diez37/elastic-migrate/services/migrator/loaders"
	"github.com/diez37/elastic-migrate/services/migrator/migration"
	"github.com/diez37/elastic-migrate/services/migrator/version"
	"github.com/hashicorp/go-multierror"
	"github.com/olivere/elastic/v7"
)

type Migrator interface {
	Migrate() error
}

type Option func(migrator *DefaultMigrator) *DefaultMigrator

type DefaultMigrator struct {
	client         *elastic.Client
	ctx            context.Context
	versionManager version.Manager
	loaders        []loaders2.Loader
}

func NewDefaultMigrator(options ...Option) *DefaultMigrator {
	migrator := &DefaultMigrator{}

	for _, option := range options {
		migrator = option(migrator)
	}

	if migrator.ctx == nil {
		migrator.ctx = context.Background()
	}

	if migrator.versionManager == nil {
		migrator.versionManager = version.NewDefaultVersionManager()
	}

	return migrator
}

func (migrator *DefaultMigrator) Migrate() error {
	if migrator.client == nil {
		return errors.New("elastic client cannot be null")
	}

	errs := &multierror.Error{}

	var migrations []*migration.Migration

	for _, loader := range migrator.loaders {
		m, err := loader.Load()
		if err != nil {
			errs = multierror.Append(errs, err)

			continue
		}

		migrations = append(migrations, m...)
	}

	if len(errs.WrappedErrors()) > 0 {
		return errs
	}

	for _, m := range migrations {
		exists, err := migrator.versionManager.Exists(migrator.ctx, m.Version().Number())

		if err != nil {
			errs = multierror.Append(errs, err)
		}

		if exists {
			for _, action := range m.Up() {
				_, err := migrator.client.PerformRequest(migrator.ctx, action.PerformRequestOptions())
				if err != nil {
					errs = multierror.Append(errs, err)

					break
				}
			}

			if len(errs.WrappedErrors()) > 0 {
				for _, action := range m.Down() {
					_, err := migrator.client.PerformRequest(migrator.ctx, action.PerformRequestOptions())
					if err != nil {
						errs = multierror.Append(errs, err)
					}
				}
			} else {
				err := migrator.versionManager.Save(migrator.ctx, m.Version())
				if err != nil {
					errs = multierror.Append(errs, err)
				}
			}
		}
	}

	if len(errs.WrappedErrors()) > 0 {
		return errs
	}

	return nil
}

func OptionSetClient(client *elastic.Client) Option {
	return func(migrator *DefaultMigrator) *DefaultMigrator {
		migrator.client = client

		return migrator
	}
}

func OptionSetCtx(ctx context.Context) Option {
	return func(migrator *DefaultMigrator) *DefaultMigrator {
		migrator.ctx = ctx

		return migrator
	}
}

func OptionSetVersionManager(versionManager version.Manager) Option {
	return func(migrator *DefaultMigrator) *DefaultMigrator {
		migrator.versionManager = versionManager

		return migrator
	}
}

func OptionAddMigrationsLoader(loader loaders2.Loader) Option {
	return func(migrator *DefaultMigrator) *DefaultMigrator {
		migrator.loaders = append(migrator.loaders, loader)

		return migrator
	}
}
