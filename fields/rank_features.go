package fields

type RankFeatures struct {
}

func NewRankFeatures() *RankFeatures {
    return &RankFeatures{}
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
