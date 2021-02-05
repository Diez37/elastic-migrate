package index

import "github.com/diez37/elastic-migrate/settings"

type Index struct {
	mappings *Mappings
	settings *settings.Settings
	aliases  []*Alias
}

func NewIndex() *Index {
	return &Index{}
}

func (index *Index) SetMappings(mappings *Mappings) *Index {
	index.mappings = mappings

	return index
}

func (index *Index) SetSettings(settings *settings.Settings) *Index {
	index.settings = settings

	return index
}

func (index *Index) AddAliases(aliases ...*Alias) *Index {
	index.aliases = append(index.aliases, aliases...)

	return index
}

func (index *Index) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if index.mappings != nil {
		src, err := index.mappings.Source()

		if err != nil {
			return nil, err
		}

		source["mappings"] = src
	}

	if index.settings != nil {
		src, err := index.settings.Source()

		if err != nil {
			return nil, err
		}

		source["settings"] = src
	}

	if len(index.aliases) > 0 {
		aliases := map[string]interface{}{}

		for _, alias := range index.aliases {
			src, err := alias.Source()

			if err != nil {
				return nil, err
			}

			aliases[alias.Alias()] = src
		}

		source["aliases"] = aliases
	}

	return source, nil
}
