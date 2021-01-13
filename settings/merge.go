package settings

type Merge struct {
    scheduler *Scheduler
    policy    *Policy
}

func NewMerge() *Merge {
    return &Merge{}
}

func (merge *Merge) SetScheduler(scheduler *Scheduler) *Merge {
    merge.scheduler = scheduler
    
    return merge
}

func (merge *Merge) SetPolicy(policy *Policy) *Merge {
    merge.policy = policy
    
    return merge
}

func (merge *Merge) Source() (interface{}, error) {
    source := map[string]interface{}{}

    if merge.scheduler != nil {
        schedulerSource, err := merge.scheduler.Source()
        if err != nil {
            return nil, err
        }

        source["scheduler"] = schedulerSource
    }

    if merge.policy != nil {
        policySource, err := merge.policy.Source()
        if err != nil {
            return nil, err
        }

        source["policy"] = policySource
    }

    return source, nil
}
