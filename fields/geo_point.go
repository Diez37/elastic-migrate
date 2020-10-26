package fields

type GeoPoint struct {
    ignoreMalformed *bool
    ignoreZValue    *bool
    nullValue       interface{}
}

func NewGeoPoint() *GeoPoint {
    return &GeoPoint{}
}

func (field *GeoPoint) SetIgnoreMalformed(ignoreMalformed bool) *GeoPoint {
    field.ignoreMalformed = &ignoreMalformed

    return field
}

func (field *GeoPoint) SetIgnoreZValue(ignoreZValue bool) *GeoPoint {
    field.ignoreZValue = &ignoreZValue

    return field
}

func (field *GeoPoint) SetNullValue(nullValue interface{}) *GeoPoint {
    field.nullValue = nullValue

    return field
}

func (field *GeoPoint) GetType() Type {
    return TypeGeoPoint
}

func (field *GeoPoint) Source() (interface{}, error) {
    // {
    //  "type": "geo_point",
    //  "ignore_malformed": true,
    //  "ignore_z_value": true,
    //  "null_value": "NULL"
    // }

    source := map[string]interface{}{}

    source["type"] = field.GetType()

    if field.ignoreZValue != nil {
        source["ignore_z_value"] = *field.ignoreZValue
    }

    if field.ignoreMalformed != nil {
        source["ignore_malformed"] = *field.ignoreMalformed
    }

    if field.nullValue != nil {
        source["null_value"] = field.nullValue
    }

    return source, nil
}
