package fields

import "strings"

type Range struct {
    coerce *bool
    store  *bool
    index  *bool
    boost  *float64
    _type  Type
}

func NewRangeInteger() *Range {
    return &Range{_type: TypeRangeInteger}
}

func NewRangeFloat() *Range {
    return &Range{_type: TypeRangeFloat}
}

func NewRangeLong() *Range {
    return &Range{_type: TypeRangeLong}
}

func NewRangeDouble() *Range {
    return &Range{_type: TypeRangeDouble}
}

func NewRangeIp() *Range {
    return &Range{_type: TypeRangeIp}
}

func (field *Range) SetCoerce(coerce bool) *Range {
    field.coerce = &coerce

    return field
}

func (field *Range) SetStore(store bool) *Range {
    field.store = &store

    return field
}

func (field *Range) SetIndex(index bool) *Range {
    field.index = &index

    return field
}

func (field *Range) SetBoost(boost float64) *Range {
    field.boost = &boost

    return field
}

func (field *Range) GetType() Type {
    return field._type
}

func (field *Range) Source() (interface{}, error) {
    // {
    //  "type": "integer_range",
    //  "coerce": false,
    //  "boost": 2,
    //  "index": true,
    //  "store": true
    // }

    source := map[string]interface{}{}

    source["type"] = field.GetType()

    if field.store != nil {
        source["store"] = *field.store
    }

    if field.boost != nil {
        source["boost"] = *field.boost
    }

    if field.index != nil {
        source["index"] = *field.index
    }

    if field.coerce != nil {
        source["coerce"] = *field.coerce
    }

    return source, nil
}

type RangeDate struct {
    Range
    formats []string
}

func NewRangeDate() *RangeDate {
    return &RangeDate{}
}

func (field *RangeDate) GetType() Type {
    return TypeRangeDate
}

func (field *RangeDate) Formats(formats ...string) *RangeDate {
    field.formats = append(field.formats, formats...)

    return field
}

func (field *RangeDate) Source() (interface{}, error) {
    // {
    //  "type": "integer_range",
    //  "coerce": false,
    //  "boost": 2,
    //  "index": true,
    //  "store": true,
    //  "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
    // }

    sourceRange, err := field.Range.Source()
    if err != nil {
        return nil, err
    }

    source := sourceRange.(map[string]interface{})

    if len(field.formats) > 0 {
        source["format"] = strings.Join(field.formats, DateFormatSeparator)
    }

    return source, nil
}
