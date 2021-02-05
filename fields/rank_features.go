package fields

type RankFeatures struct {
	name FieldName
}

func NewRankFeatures(name string) *RankFeatures {
	return &RankFeatures{name: FieldName(name)}
}

func (field *RankFeatures) Name() FieldName {
	return field.name
}

func (field *RankFeatures) GetType() Type {
	return TypeRankFeatures
}

func (field *RankFeatures) Source() (interface{}, error) {
	// {
	//  "type": "rank_features"
	// }

	source := map[string]interface{}{}

	source["type"] = field.GetType()

	return source, nil
}
