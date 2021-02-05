package fields

type IndexPrefixes struct {
	name     FieldName
	minChars *int
	maxChars *int
}

func NewIndexPrefixes(name string) *IndexPrefixes {
	return &IndexPrefixes{name: FieldName(name)}
}

func (field *IndexPrefixes) Name() FieldName {
	return field.name
}

func (field *IndexPrefixes) SetMinimumChars(minimumChars int) *IndexPrefixes {
	field.minChars = &minimumChars

	return field
}

func (field *IndexPrefixes) SetMaximumChars(maximumChars int) *IndexPrefixes {
	field.maxChars = &maximumChars

	return field
}

func (field *IndexPrefixes) Source() (interface{}, error) {
	// {
	//  "min_chars" : 1,
	//  "max_chars" : 10
	// }

	source := map[string]interface{}{}

	if field.minChars != nil {
		source["min_chars"] = *field.minChars
	}

	if field.maxChars != nil {
		source["max_chars"] = *field.maxChars
	}

	return source, nil
}
