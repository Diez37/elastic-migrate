package index

import "github.com/diez37/elastic-migrate/fields"

type Mappings struct {
	dynamic    *bool
	properties []fields.Fielder
}

func NewMappings() *Mappings {
	return &Mappings{}
}

func (mappings *Mappings) SetDynamic(dynamic bool) *Mappings {
	mappings.dynamic = &dynamic

	return mappings
}

func (mappings *Mappings) AddProperties(properties ...fields.Fielder) *Mappings {
	mappings.properties = append(mappings.properties, properties...)

	return mappings
}

func (mappings *Mappings) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if mappings.dynamic != nil {
		source["dynamic"] = *mappings.dynamic
	}

	if len(mappings.properties) > 0 {
		properties := map[fields.FieldName]interface{}{}

		for _, field := range mappings.properties {
			fieldSource, err := field.Source()

			if err != nil {
				return nil, err
			}

			properties[field.Name()] = fieldSource
		}

		source["properties"] = properties
	}

	return source, nil
}
