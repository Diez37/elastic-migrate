package settings

type Mapping struct {
    coerce          *bool
    ignoreMalformed *bool
    totalFields     *Limit
    depth           *Limit
    nestedFields    *Limit
    nestedObjects   *Limit
    fieldNameLength *Limit
}

func NewMapping() *Mapping {
    return &Mapping{}
}

func (mapping *Mapping) SetCoerce(coerce bool) *Mapping {
    mapping.coerce = &coerce

    return mapping
}

func (mapping *Mapping) SetIgnoreMalformed(ignoreMalformed bool) *Mapping {
    mapping.ignoreMalformed = &ignoreMalformed

    return mapping
}

func (mapping *Mapping) SetTotalFields(totalFields *Limit) *Mapping {
    mapping.totalFields = totalFields

    return mapping
}

func (mapping *Mapping) SetDepth(depth *Limit) *Mapping {
    mapping.depth = depth

    return mapping
}

func (mapping *Mapping) SetNestedFields(nestedFields *Limit) *Mapping {
    mapping.nestedFields = nestedFields

    return mapping
}

func (mapping *Mapping) SetNestedObjects(nestedObjects *Limit) *Mapping {
    mapping.nestedObjects = nestedObjects

    return mapping
}

func (mapping *Mapping) SetFieldNameLength(fieldNameLength *Limit) *Mapping {
    mapping.fieldNameLength = fieldNameLength

    return mapping
}

func (mapping *Mapping) Source() (interface{}, error) {
    source := map[string]interface{}{}

    if mapping.coerce != nil {
        source["coerce"] = *mapping.coerce
    }

    if mapping.ignoreMalformed != nil {
        source["ignore_malformed"] = *mapping.ignoreMalformed
    }

    if mapping.totalFields != nil {
        limitSource, err := mapping.totalFields.Source()
        if err != nil {
            return nil, err
        }

        source["total_fields"] = limitSource
    }

    if mapping.depth != nil {
        limitSource, err := mapping.depth.Source()
        if err != nil {
            return nil, err
        }

        source["depth"] = limitSource
    }

    if mapping.nestedFields != nil {
        limitSource, err := mapping.nestedFields.Source()
        if err != nil {
            return nil, err
        }

        source["nested_fields"] = limitSource
    }

    if mapping.nestedObjects != nil {
        limitSource, err := mapping.nestedObjects.Source()
        if err != nil {
            return nil, err
        }

        source["nested_objects"] = limitSource
    }

    if mapping.fieldNameLength != nil {
        limitSource, err := mapping.fieldNameLength.Source()
        if err != nil {
            return nil, err
        }

        source["field_name_length"] = limitSource
    }

    return source, nil
}
