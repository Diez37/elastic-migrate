package settings

const (
    TranslogDurabilityRequest TranslogDurability = "request"
    TranslogDurabilityAsync   TranslogDurability = "async"
)

type TranslogDurability string

type Translog struct {
    syncInterval       *Interval
    durability         *TranslogDurability
    flushThresholdSize *Size
}

func NewTranslog() *Translog {
    return &Translog{}
}

func (translog *Translog) SetSyncInterval(syncInterval *Interval) *Translog {
    translog.syncInterval = syncInterval

    return translog
}

func (translog *Translog) SetDurability(durability TranslogDurability) *Translog {
    translog.durability = &durability

    return translog
}

func (translog *Translog) SetFlushThresholdSize(flushThresholdSize *Size) *Translog {
    translog.flushThresholdSize = flushThresholdSize

    return translog
}

func (translog *Translog) Source() (interface{}, error) {
    source := map[string]interface{}{}

    if translog.durability != nil {
        source["durability"] = *translog.durability
    }

    if translog.syncInterval != nil {
        src, err := translog.syncInterval.Source()
        if err != nil {
            return nil, err
        }

        source["sync_interval"] = src
    }

    if translog.flushThresholdSize != nil {
        src, err := translog.flushThresholdSize.Source()
        if err != nil {
            return nil, err
        }

        source["flush_threshold_size"] = src
    }

    return source, nil
}
