package fields

type Alias struct {
    path string
}

func NewAlias(path string) *Alias {
    return &Alias{path: path}
}

func (field *Alias) GetType() Type {
    return TypeAlias
}

func (field *Alias) Source() (interface{}, error) {
    // {
    //      "type": "alias",
    //		"path": "distance"
    // }

    source := map[string]interface{}{}

    source["type"] = field.GetType()
    source["path"] = field.path

    return source, nil
}
