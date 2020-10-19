package fields

type RankFeature struct {
    positiveScoreImpact *bool
}

func NewRankFeature() *RankFeature {
    return &RankFeature{}
}

func (field *RankFeature) SetPositiveScoreImpact(positiveScoreImpact bool) *RankFeature {
    field.positiveScoreImpact = &positiveScoreImpact

    return field
}

func (field *RankFeature) GetType() Type {
    return TypeRankFeature
}

func (field *RankFeature) Source() (interface{}, error) {
    // {
    //  "type": "rank_feature",
    //  "positive_score_impact": false
    // }

    source := map[string]interface{}{}

    source["type"] = field.GetType()

    if field.positiveScoreImpact != nil {
        source["positive_score_impact"] = *field.positiveScoreImpact
    }

    return source, nil
}
