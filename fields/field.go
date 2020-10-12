package fields

type Field struct {
    properties bool
    name       string
    field      Fielder
}

func NewField(name string, field Fielder) *Field {
    return &Field{name: name, field: field}
}

func (field *Field) SetProperties(properties bool) *Field {
    field.properties = properties

    return field
}

func (field *Field) Source() (interface{}, error) {
    // {
    //		"properties": {
    //			...
    //		}
    // }
    // or
    // {
    //		...
    // }

    body, err := field.field.Source()
    if err != nil {
        return nil, err
    }

    if !field.properties {
        return body, nil
    }

    return map[string]interface{}{
        "properties": body,
    }, nil
}
