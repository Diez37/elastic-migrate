package settings

type Limit struct {
    limit uint32
}

func NewLimit(limit uint32) *Limit {
    return &Limit{limit: limit}
}

func (limit *Limit) Source() (interface{}, error) {
    return map[string]interface{}{
        "limit": limit.limit,
    }, nil
}
