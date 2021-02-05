package fields

type Meta struct {
	name  string
	value interface{}
}

func NewMeta(name string, value interface{}) *Meta {
	return &Meta{name: name, value: value}
}

func (meta *Meta) Name() string {
	return meta.name
}

func (meta *Meta) Value() interface{} {
	return meta.value
}
