package fields

type Field struct {
	name  FieldName
	field Fielder
}

func NewField(name string, field Fielder) *Field {
	return &Field{name: FieldName(name), field: field}
}

func (field *Field) Name() FieldName {
	return field.name
}

func (field *Field) Source() (interface{}, error) {
	// {
	//		...
	// }

	return field.field.Source()
}
