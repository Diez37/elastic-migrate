package fields

type Ip struct {
    docValues *bool
    index     *bool
    store     *bool
    boost     *float64
    nullValue interface{}
}

func NewIp() *Ip {
    return &Ip{}
}

func (field *Ip) SetDocValues(docValues bool) *Ip {
    field.docValues = &docValues

    return field
}

func (field *Ip) SetStore(store bool) *Ip {
    field.store = &store

    return field
}

func (field *Ip) SetIndex(index bool) *Ip {
    field.index = &index

    return field
}

func (field *Ip) SetBoost(boost float64) *Ip {
    field.boost = &boost

    return field
}

func (field *Ip) SetNullValue(nullValue interface{}) *Ip {
    field.nullValue = nullValue

    return field
}

func (field *Ip) GetType() Type {
    return TypeIp
}

func (field *Ip) Source() (interface{}, error) {
    // {
    //  "type": "ip",
    //  "boost": 2,
    //  "doc_values": false,
    //  "index": true,
    //  "null_value": "NULL",
    //  "store": true
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

    return source, nil
}
