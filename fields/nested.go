package fields

type Nested struct {
	name       FieldName
	dynamic    *Dynamic
	properties []*Field
}

func NewNested(name string, properties ...*Field) *Nested {
	return (&Nested{name: FieldName(name)}).Properties(properties...)
}

func (field *Nested) Name() FieldName {
	return field.name
}

func (field *Nested) SetDynamic(dynamic Dynamic) *Nested {
	field.dynamic = &dynamic

	return field
}

func (field *Nested) Properties(properties ...*Field) *Nested {
	field.properties = append(field.properties, properties...)

	return field
}

func (field *Nested) GetType() Type {
	return TypeNested
}

func (field *Nested) Source() (interface{}, error) {
	// {
	//  "type": "nested",
	//  "dynamic": true,
	//  "properties": {
	//    "age":  { "type": "integer" },
	//    "name": { "type": "text"  }
	//  }
	// }

	source := map[string]interface{}{}

	source["type"] = field.GetType()

	if field.dynamic != nil {
		switch *field.dynamic {
		case DynamicEnabled:
			source["dynamic"] = true
		case DynamicDisabled:
			source["dynamic"] = false
		case DynamicStrict:
			source["dynamic"] = DynamicStrict
		default:
			return nil, ErrorDynamicUnknownType
		}
	}

	if len(field.properties) > 0 {
		properties := map[FieldName]interface{}{}

		for _, f := range field.properties {
			s, err := f.Source()
			if err != nil {
				return nil, err
			}

			properties[f.Name()] = s
		}

		source["properties"] = properties
	}

	return source, nil
}
