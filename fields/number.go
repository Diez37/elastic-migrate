package fields

import "errors"

var ErrorScalingFactorNotSet = errors.New("scaling_factor required field on scaled_float type")
var ErrorScalingFactorNotScaledFloat = errors.New("scaling_factor can only be installed in field type scaled_float")

type Number struct {
    coerce          *bool
    docValues       *bool
    ignoreMalformed *bool
    index           *bool
    store           *bool
    boost           *float64
    scalingFactor   *int
    nullValue       interface{}
    meta            []*Meta
    _type           Type
}

func NewLong() *Number {
    return &Number{_type: TypeNumberLong}
}

func NewInteger() *Number {
    return &Number{_type: TypeNumberInteger}
}

func NewShort() *Number {
    return &Number{_type: TypeNumberShort}
}

func NewByte() *Number {
    return &Number{_type: TypeNumberByte}
}

func NewDouble() *Number {
    return &Number{_type: TypeNumberDouble}
}

func NewFloat() *Number {
    return &Number{_type: TypeNumberFloat}
}

func NewHalfFloat() *Number {
    return &Number{_type: TypeNumberHalfFloat}
}

func NewScaledFloat() *Number {
    return &Number{_type: TypeNumberScaledFloat}
}

func (field *Number) SetDocValues(docValues bool) *Number {
    field.docValues = &docValues

    return field
}

func (field *Number) SetIndex(index bool) *Number {
    field.index = &index

    return field
}

func (field *Number) SetStore(store bool) *Number {
    field.store = &store

    return field
}

func (field *Number) SetBoost(boost float64) *Number {
    field.boost = &boost

    return field
}

func (field *Number) SetNullValue(nullValue interface{}) *Number {
    field.nullValue = nullValue

    return field
}

func (field *Number) Meta(metas ...*Meta) *Number {
    field.meta = append(field.meta, metas...)

    return field
}

func (field *Number) SetCoerce(coerce bool) *Number {
    field.coerce = &coerce

    return field
}

func (field *Number) SetIgnoreMalformed(ignoreMalformed bool) *Number {
    field.ignoreMalformed = &ignoreMalformed

    return field
}

func (field *Number) SetScalingFactor(scalingFactor int) *Number {
    field.scalingFactor = &scalingFactor

    return field
}

func (field *Number) GetType() Type {
    return field._type
}

func (field *Number) Source() (interface{}, error) {
    // {
    //  "type": "long",
    //  "coerce": false,
    //  "boost": 2,
    //  "doc_values": false,
    //  "ignore_malformed": true,
    //  "index": true,
    //  "null_value": "NULL",
    //  "store": true,
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

    if field.coerce != nil {
        source["coerce"] = *field.coerce
    }

    if field.ignoreMalformed != nil {
        source["ignore_malformed"] = *field.ignoreMalformed
    }

    if field.scalingFactor == nil && field.GetType() == TypeNumberScaledFloat {
        return nil, ErrorScalingFactorNotSet
    }

    if field.scalingFactor != nil {
        if field.GetType() != TypeNumberScaledFloat {
            return nil, ErrorScalingFactorNotScaledFloat
        }

        source["scaling_factor"] = *field.scalingFactor
    }

    return source, nil
}
