package fields

type Object struct {
    enabled    *bool
    dynamic    *Dynamic
    properties []*Field
}

func NewObject(properties ...*Field) *Object {
    return (&Object{}).Properties(properties...)
}

func (field *Object) SetDynamic(dynamic Dynamic) *Object {
    field.dynamic = &dynamic

    return field
}

func (field *Object) SetEnabled(enabled bool) *Object {
    field.enabled = &enabled

    return field
}

func (field *Object) Properties(properties ...*Field) *Object {
    field.properties = append(field.properties, properties...)

    return field
}

func (field *Object) GetType() Type {
    return TypeObject
}

func (field *Object) Source() (interface{}, error) {
    // {
    //  "properties": {
    //    "first": { "type": "text" },
    //    "last":  { "type": "text" }
    //  },
    //  "dynamic": true,
    //  "enabled": false
    // }

    source := map[string]interface{}{}

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

    if field.enabled != nil {
        source["enabled"] = *field.enabled
    }

    if len(field.properties) > 0 {
        properties := map[string]interface{}{}

        for _, f := range field.properties {
            s, err := f.Source()
            if err != nil {
                return nil, err
            }

            properties[f.name] = s
        }

        source["properties"] = properties
    }

    return source, nil
}
