package fields

import "strings"

type Date struct {
    docValues       *bool
    ignoreMalformed *bool
    index           *bool
    store           *bool
    boost           *float64
    locale          *Local
    formats         []string
    nullValue       interface{}
    meta            []*Meta
    _type           Type
}

func NewDate() *Date {
    return &Date{_type: TypeDate}
}

func NewDateNano() *Date {
    return &Date{_type: TypeDateNano}
}

func (field *Date) SetDocValues(docValues bool) *Date {
    field.docValues = &docValues

    return field
}

func (field *Date) SetIgnoreMalformed(ignoreMalformed bool) *Date {
    field.ignoreMalformed = &ignoreMalformed

    return field
}

func (field *Date) SetIndex(index bool) *Date {
    field.index = &index

    return field
}

func (field *Date) SetStore(store bool) *Date {
    field.store = &store

    return field
}

func (field *Date) SetBoost(boost float64) *Date {
    field.boost = &boost

    return field
}

func (field *Date) SetNullValue(nullValue interface{}) *Date {
    field.nullValue = nullValue

    return field
}

func (field *Date) Meta(metas ...*Meta) *Date {
    field.meta = append(field.meta, metas...)

    return field
}

func (field *Date) Formats(formats ...string) *Date {
    field.formats = append(field.formats, formats...)

    return field
}

func (field *Date) Local(locale Local) *Date {
    field.locale = &locale

    return field
}

func (field *Date) GetType() Type {
    return field._type
}

func (field *Date) Source() (interface{}, error) {
    // {
    //		"type":   "date",
    //		"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
    // }

    source := map[string]interface{}{}

    source["type"] = field.GetType()

    if field.docValues != nil {
        source["doc_values"] = *field.docValues
    }

    if field.index != nil {
        source["index"] = *field.index
    }

    if field.store != nil {
        source["store"] = *field.store
    }

    if field.boost != nil {
        source["boost"] = *field.boost
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

    if field.ignoreMalformed != nil {
        source["ignore_malformed"] = *field.ignoreMalformed
    }

    if field.locale != nil {
        source["locale"] = *field.locale
    }

    if len(field.formats) > 0 {
        source["format"] = strings.Join(field.formats, DateFormatSeparator)
    }

    return source, nil
}
