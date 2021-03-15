package version

import "context"

type Manager interface {
	Exists(ctx context.Context, number int32) (bool, error)
	Save(ctx context.Context, version *Version) error
}
