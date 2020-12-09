package settings

type Scheduler struct {
    autoThrottle   *bool
    maxMergeCount  *uint32
    maxThreadCount *uint32
}

func NewScheduler() *Scheduler {
    return &Scheduler{}
}

func (scheduler *Scheduler) SetAutoThrottle(autoThrottle bool) *Scheduler {
    scheduler.autoThrottle = &autoThrottle

    return scheduler
}

func (scheduler *Scheduler) SetMaxMergeCount(maxMergeCount uint32) *Scheduler {
    scheduler.maxMergeCount = &maxMergeCount

    return scheduler
}

func (scheduler *Scheduler) SetMaxThreadCount(maxThreadCount uint32) *Scheduler {
    scheduler.maxThreadCount = &maxThreadCount

    return scheduler
}

func (scheduler *Scheduler) Source() (interface{}, error) {
    source := map[string]interface{}{}

    if scheduler.autoThrottle != nil {
        source["auto_throttle"] = *scheduler.autoThrottle
    }

    if scheduler.maxMergeCount != nil {
        source["max_merge_count"] = *scheduler.maxMergeCount
    }

    if scheduler.maxThreadCount != nil {
        source["max_thread_count"] = *scheduler.maxThreadCount
    }

    return source, nil
}
