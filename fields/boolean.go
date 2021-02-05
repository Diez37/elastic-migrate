package fields

type Boolean struct {
	name      FieldName
	docValues *bool
	store     *bool
	index     *bool
	boost     *float64
	nullValue interface{}
	meta      []*Meta
}

func NewBoolean(name string) *Boolean {
	return &Boolean{name: FieldName(name)}
}

func (field *Boolean) Name() FieldName {
	return field.name
}

func (field *Boolean) SetDocValues(docValues bool) *Boolean {
	field.docValues = &docValues

	return field
}

func (field *Boolean) SetStore(store bool) *Boolean {
	field.store = &store

	return field
}

func (field *Boolean) SetIndex(index bool) *Boolean {
	field.index = &index

	return field
}

func (field *Boolean) SetBoost(boost float64) *Boolean {
	field.boost = &boost

	return field
}

func (field *Boolean) SetNullValue(nullValue interface{}) *Boolean {
	field.nullValue = nullValue

	return field
}

func (field *Boolean) Meta(metas ...*Meta) *Boolean {
	field.meta = append(field.meta, metas...)

	return field
}

func (field *Boolean) GetType() Type {
	return TypeBoolean
}

func (field *Boolean) Source() (interface{}, error) {
	// {
	//  "type": "boolean",
	//  "doc_values": false,
	//  "store": false,
	//  "boost": 2.5,
	//  "index": true,
	//  "null_value": "NULL",
	//  "meta": {
	//    "unit": "ms"
	//  }
	// }

	source := map[string]interface{}{}

	source["type"] = field.GetType()

	if field.docValues != nil {
		source["doc_values"] = *field.docValues
	}

	if field.store != nil {
		source["store"] = *field.store
	}

	if field.boost != nil {
		source["boost"] = *field.boost
	}

	if field.index != nil {
		source["index"] = *field.index
	}

	if field.nullValue != nil {
		source["null_value"] = field.nullValue
	}

	if len(field.meta) > 0 {
		body := map[string]interface{}{}

		for _, meta := range field.meta {
			body[meta.name] = meta.value
		}

		source["meta"] = body
	}

	return source, nil
}
