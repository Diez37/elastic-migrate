package settings

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

func NewPolicy() *Policy {
    return &Policy{}
}

func (policy *Policy) SetMaxMergeAtOnceExplicit(maxMergeAtOnceExplicit uint) *Policy {
    policy.maxMergeAtOnceExplicit = &maxMergeAtOnceExplicit

    return policy
}

func (policy *Policy) SetMaxMergeAtOnce(maxMergeAtOnce uint) *Policy {
    policy.maxMergeAtOnce = &maxMergeAtOnce

    return policy
}

func (policy *Policy) SetFloorSegment(floorSegment *Size) *Policy {
    policy.floorSegment = floorSegment

    return policy
}

func (policy *Policy) SetMaxMergedSegment(maxMergedSegment *Size) *Policy {
    policy.maxMergedSegment = maxMergedSegment

    return policy
}

func (policy *Policy) SetReclaimDeletesWeight(reclaimDeletesWeight float32) *Policy {
    policy.reclaimDeletesWeight = &reclaimDeletesWeight

    return policy
}

func (policy *Policy) SetExpungeDeletesAllowed(expungeDeletesAllowed float32) *Policy {
    policy.expungeDeletesAllowed = &expungeDeletesAllowed

    return policy
}

func (policy *Policy) SetSegmentsPerTier(segmentsPerTier float32) *Policy {
    policy.segmentsPerTier = &segmentsPerTier

    return policy
}

func (policy *Policy) SetDeletesPctAllowed(deletesPctAllowed float32) *Policy {
    policy.deletesPctAllowed = &deletesPctAllowed

    return policy
}

func (policy *Policy) Source() (interface{}, error) {
    source := map[string]interface{}{}

    if policy.maxMergeAtOnceExplicit != nil {
        source["max_merge_at_once_explicit"] = *policy.maxMergeAtOnceExplicit
    }

    if policy.maxMergeAtOnce != nil {
        source["max_merge_at_once"] = *policy.maxMergeAtOnce
    }

    if policy.reclaimDeletesWeight != nil {
        source["reclaim_deletes_weight"] = *policy.reclaimDeletesWeight
    }

    if policy.expungeDeletesAllowed != nil {
        source["expunge_deletes_allowed"] = *policy.expungeDeletesAllowed
    }

    if policy.segmentsPerTier != nil {
        source["segments_per_tier"] = *policy.segmentsPerTier
    }

    if policy.deletesPctAllowed != nil {
        source["deletes_pct_allowed"] = *policy.deletesPctAllowed
    }

    if policy.floorSegment != nil {
        sizeSource, err := policy.floorSegment.Source()
        if err != nil {
            return nil, err
        }

        source["floor_segment"] = sizeSource
    }

    if policy.maxMergedSegment != nil {
        sizeSource, err := policy.maxMergedSegment.Source()
        if err != nil {
            return nil, err
        }

        source["max_merged_segment"] = sizeSource
    }

    return source, nil
}
