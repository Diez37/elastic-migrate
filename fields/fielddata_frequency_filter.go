package fields

type FielddataFrequencyFilter struct {
    min            *float64
    max            *float64
    minSegmentSize *int
}

func NewFielddataFrequencyFilter() *FielddataFrequencyFilter {
    return &FielddataFrequencyFilter{}
}

func (fielddata *FielddataFrequencyFilter) SetMinimum(minimum float64) *FielddataFrequencyFilter {
    fielddata.min = &minimum

    return fielddata
}

func (fielddata *FielddataFrequencyFilter) SetMaximum(maximum float64) *FielddataFrequencyFilter {
    fielddata.max = &maximum

    return fielddata
}

func (fielddata *FielddataFrequencyFilter) SetMinSegmentSize(minSegmentSize int) *FielddataFrequencyFilter {
    fielddata.minSegmentSize = &minSegmentSize

    return fielddata
}

func (fielddata *FielddataFrequencyFilter) Source() (interface{}, error) {
    // {
    //  "min": 0.001,
    //  "max": 0.1,
    //  "min_segment_size": 500
    // }

    source := map[string]interface{}{}

    if fielddata.min != nil {
        source["min"] = *fielddata.min
    }

    if fielddata.max != nil {
        source["max"] = *fielddata.max
    }

    if fielddata.minSegmentSize != nil {
        source["min_segment_size"] = *fielddata.minSegmentSize
    }

    return source, nil
}

