package fields

type FieldDataFrequencyFilter struct {
	name           FieldName
	min            *float64
	max            *float64
	minSegmentSize *int
}

func NewFieldDataFrequencyFilter(name string) *FieldDataFrequencyFilter {
	return &FieldDataFrequencyFilter{name: FieldName(name)}
}

func (field *FieldDataFrequencyFilter) Name() FieldName {
	return field.name
}

func (field *FieldDataFrequencyFilter) SetMinimum(minimum float64) *FieldDataFrequencyFilter {
	field.min = &minimum

	return field
}

func (field *FieldDataFrequencyFilter) SetMaximum(maximum float64) *FieldDataFrequencyFilter {
	field.max = &maximum

	return field
}

func (field *FieldDataFrequencyFilter) SetMinSegmentSize(minSegmentSize int) *FieldDataFrequencyFilter {
	field.minSegmentSize = &minSegmentSize

	return field
}

func (field *FieldDataFrequencyFilter) Source() (interface{}, error) {
	// {
	//  "min": 0.001,
	//  "max": 0.1,
	//  "min_segment_size": 500
	// }

	source := map[string]interface{}{}

	if field.min != nil {
		source["min"] = *field.min
	}

	if field.max != nil {
		source["max"] = *field.max
	}

	if field.minSegmentSize != nil {
		source["min_segment_size"] = *field.minSegmentSize
	}

	return source, nil
}
