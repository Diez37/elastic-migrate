package fields

type Field struct {
    name  string
    field Fielder
}

func NewField(name string, field Fielder) *Field {
    return &Field{name: name, field: field}
}

func (field *Field) Source() (interface{}, error) {
    // {
    //		...
    // }

    return field.field.Source()
}
