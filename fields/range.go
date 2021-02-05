package fields

import (
	"fmt"
)

type Range struct {
	name   FieldName
	coerce *bool
	store  *bool
	index  *bool
	boost  *float64
	_type  Type
}

func NewRangeInteger(name string) *Range {
	return &Range{name: FieldName(name), _type: TypeRangeInteger}
}

func NewRangeFloat(name string) *Range {
	return &Range{name: FieldName(name), _type: TypeRangeFloat}
}

func NewRangeLong(name string) *Range {
	return &Range{name: FieldName(name), _type: TypeRangeLong}
}

func NewRangeDouble(name string) *Range {
	return &Range{name: FieldName(name), _type: TypeRangeDouble}
}

func NewRangeIp(name string) *Range {
	return &Range{name: FieldName(name), _type: TypeRangeIp}
}

func (field *Range) Name() FieldName {
	return field.name
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
	formats []DateFormat
}

func NewRangeDate(name string) *RangeDate {
	return &RangeDate{Range: Range{name: FieldName(name), _type: TypeRangeDate}}
}

func (field *RangeDate) Formats(formats ...DateFormat) *RangeDate {
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
	source["type"] = field.GetType()

	if len(field.formats) > 0 {
		source["format"] = field.formats[0]

		if len(field.formats) > 1 {
			for _, format := range field.formats[1:] {
				source["format"] = fmt.Sprintf("%s%s%s", source["format"], DateFormatSeparator, format)
			}
		}
	}

	return source, nil
}
