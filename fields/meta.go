package fields

type Meta struct {
    name  string
    value interface{}
}

func NewMeta(name string, value interface{}) *Meta {
    return &Meta{name: name, value: value}
}
