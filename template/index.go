package template

import (
	"github.com/diez37/elastic-migrate/fields"
	"github.com/diez37/elastic-migrate/index"
)

type Index struct {
	name          string
	template      *index.Index
	priority      *int
	version       *int
	indexPatterns []string
	composedOf    []string
	meta          []*fields.Meta
}

func NewIndex(name string) *Index {
	return &Index{name: name}
}

func (index *Index) SetTemplate(template *index.Index) *Index {
	index.template = template

	return index
}

func (index *Index) SetPriority(priority int) *Index {
	index.priority = &priority

	return index
}

func (index *Index) SetVersion(version int) *Index {
	index.version = &version

	return index
}

func (index *Index) AddIndexPatterns(indexPatterns []string) *Index {
	index.indexPatterns = append(index.indexPatterns, indexPatterns...)

	return index
}

func (index *Index) AddComposedOf(composedOf ...string) *Index {
	index.composedOf = append(index.composedOf, composedOf...)

	return index
}

func (index *Index) AddMeta(meta ...*fields.Meta) *Index {
	index.meta = append(index.meta, meta...)

	return index
}

func (index *Index) Name() string {
	return index.name
}

func (index *Index) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if index.template != nil {
		template, err := index.template.Source()

		if err != nil {
			return nil, err
		}

		source["template"] = template
	}

	if index.priority != nil {
		source["priority"] = *index.priority
	}

	if index.version != nil {
		source["version"] = *index.version
	}

	if len(index.indexPatterns) > 0 {
		source["index_patterns"] = index.indexPatterns
	}

	if len(index.composedOf) > 0 {
		source["composed_of"] = index.composedOf
	}

	if len(index.meta) > 0 {
		meta := map[string]interface{}{}

		for _, item := range index.meta {
			meta[item.Name()] = item.Value()
		}

		source["_meta"] = meta
	}

	return source, nil
}
