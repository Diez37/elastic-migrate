package fields

type Alias struct {
	name FieldName
	path string
}

func NewAlias(name string, path string) *Alias {
	return &Alias{name: FieldName(name), path: path}
}

func (field *Alias) Name() FieldName {
	return field.name
}

func (field *Alias) GetType() Type {
	return TypeAlias
}

func (field *Alias) Source() (interface{}, error) {
	// {
	//  "type": "alias",
	//  "path": "distance"
	// }

	source := map[string]interface{}{}

	source["type"] = field.GetType()
	source["path"] = field.path

	return source, nil
}
