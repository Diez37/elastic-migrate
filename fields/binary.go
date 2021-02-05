package fields

type Binary struct {
	name      FieldName
	docValues *bool
	store     *bool
}

func NewBinary(name string) *Binary {
	return &Binary{name: FieldName(name)}
}

func (field *Binary) Name() FieldName {
	return field.name
}

func (field *Binary) SetDocValues(docValues bool) *Binary {
	field.docValues = &docValues

	return field
}

func (field *Binary) SetStore(store bool) *Binary {
	field.store = &store

	return field
}

func (field *Binary) GetType() Type {
	return TypeBinary
}

func (field *Binary) Source() (interface{}, error) {
	// {
	//  "type": "binary",
	//  "doc_values": false,
	//  "store": false
	// }

	source := map[string]interface{}{}

	source["type"] = field.GetType()

	if field.docValues != nil {
		source["doc_values"] = *field.docValues
	}

	if field.store != nil {
		source["store"] = *field.store
	}

	return source, nil
}
