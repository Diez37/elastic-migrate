package settings

type Retention struct {
    operations uint32
}

func NewRetention(operations uint32) *Retention {
    return &Retention{operations: operations}
}

func (retention *Retention) Source() (interface{}, error) {
    return map[string]interface{}{
        "operations": retention.operations,
    }, nil
}
