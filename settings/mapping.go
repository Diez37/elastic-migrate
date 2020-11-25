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
