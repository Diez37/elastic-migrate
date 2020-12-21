package settings

type Settings struct {
	index *Index
}

func NewSettings(index *Index) *Settings {
	return &Settings{index: index}
}

func (settings *Settings) Source() (interface{}, error) {
	indexSource, err := settings.index.Source()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"index": indexSource,
	}, nil
}
