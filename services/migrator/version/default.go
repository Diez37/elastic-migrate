package version

import "context"

type DefaultVersionManager struct {
}

func NewDefaultVersionManager() *DefaultVersionManager {
	return &DefaultVersionManager{}
}

func (v *DefaultVersionManager) Exists(_ context.Context, _ int32) (bool, error) {
	return false, nil
}

func (v *DefaultVersionManager) Save(_ context.Context, _ *Version) error {
	return nil
}
