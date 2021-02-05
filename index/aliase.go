package index

import "github.com/olivere/elastic/v7"

type Alias struct {
	alias   string
	routing *string
	filter  elastic.Query
}

func NewAlias(alias string) *Alias {
	return &Alias{alias: alias}
}

func (alias *Alias) Alias() string {
	return alias.alias
}

func (alias *Alias) SetRouting(routing string) *Alias {
	alias.routing = &routing

	return alias
}

func (alias *Alias) SetFilter(filter elastic.Query) *Alias {
	alias.filter = filter

	return alias
}

func (alias *Alias) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if alias.routing != nil {
		source["routing"] = *alias.routing
	}

	if alias.filter != nil {
		filter, err := alias.filter.Source()

		if err != nil {
			return nil, err
		}

		source["filter"] = filter
	}

	return source, nil
}
