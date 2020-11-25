package settings

type Merge struct {
    scheduler *Scheduler
    policy    *Policy
}

type Scheduler struct {
    autoThrottle   *bool
    maxMergeCount  *uint
    maxThreadCount *uint
}

type Policy struct {
    maxMergeAtOnceExplicit *uint
    maxMergeAtOnce         *uint
    floorSegment           *Size
    maxMergedSegment       *Size
    reclaimDeletesWeight   *float32
    expungeDeletesAllowed  *float32
    segmentsPerTier        *float32
    deletesPctAllowed      *float32
}
