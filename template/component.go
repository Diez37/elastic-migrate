package template

import (
	"github.com/diez37/elastic-migrate/fields"
	"github.com/diez37/elastic-migrate/index"
)

type Component struct {
	name            string
	template        *index.Index
	version         *int
	allowAutoCreate *bool
	meta            []*fields.Meta
}

func NewComponent(name string) *Component {
	return &Component{name: name}
}

func (component *Component) Name() string {
	return component.name
}

func (component *Component) SetTemplate(template *index.Index) *Component {
	component.template = template

	return component
}

func (component *Component) SetVersion(version int) *Component {
	component.version = &version

	return component
}

func (component *Component) SetAllowAutoCreate(allowAutoCreate bool) *Component {
	component.allowAutoCreate = &allowAutoCreate

	return component
}

func (component *Component) SetMeta(meta ...*fields.Meta) *Component {
	component.meta = append(component.meta, meta...)

	return component
}

func (component *Component) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if component.template != nil {
		template, err := component.template.Source()

		if err != nil {
			return nil, err
		}

		source["template"] = template
	}

	if component.version != nil {
		source["version"] = *component.version
	}

	if component.allowAutoCreate != nil {
		source["allow_auto_create"] = *component.allowAutoCreate
	}

	if len(component.meta) > 0 {
		meta := map[string]interface{}{}

		for _, item := range component.meta {
			meta[item.Name()] = item.Value()
		}

		source["_meta"] = meta
	}

	return source, nil
}
