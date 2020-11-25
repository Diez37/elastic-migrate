package settings

type Index struct {
    loadFixedBitsetFiltersEagerly *bool
    hidden                        *bool
    numberOfShards                *uint
    numberOfRoutingShards         *uint
    routingPartitionSize          *uint
    numberOfReplicas              *uint
    maxResultWindow               *uint
    maxInnerResultWindow          *uint
    maxRescoreWindow              *uint
    maxDocvalueFieldsSearch       *uint
    maxScriptFields               *uint
    maxNgramDiff                  *uint
    maxShingleDiff                *uint
    maxRefreshListeners           *uint
    maxTermsCount                 *uint
    maxRegexLength                *uint
    maxAdjacencyMatrixFilters     *uint
    maxSlicesPerScroll            *uint
    autoExpandReplicas            *string
    verifiedBeforeClose           *bool
    sourceOnly                    *bool
    refreshInterval               *Interval
    gcDeletes                     *Interval
    flushAfterMerge               *Size
    defaultPipeline               *PipelineName
    finalPipeline                 *PipelineName
    search                        *Search
    analyze                       *Analyze
    highlight                     *Highlight
    routing                       *Routing
    shard                         *Shard
    softDeletes                   *SoftDeletes
    codec                         *IndexCode
    unassigned                    *Unassigned
    mapping                       *Mapping
    merge                         *Merge
    write                         *Write
    indexing                      *Slowlog
    blocks                        *Blocks
    similarity                    []*Similarity
}

type Blocks struct {
    metadata            *bool
    read                *bool
    readOnlyAllowDelete *bool
    readOnly            *bool
    write               *bool
}

type Shard struct {
    checkOnStartup *CheckOnStartup
}

type SoftDeletes struct {
    enabled        *bool
    retention      *Retention
    retentionLease *RetentionLease
}

type Retention struct {
    operations *uint
}

type RetentionLease struct {
    period *Interval
}

type Search struct {
    throttled *bool
    slowlog   *Slowlog
    idle      *Idle
}

type Slowlog struct {
    reformat  *bool
    source    *uint
    level     *SlowLogLevel
    threshold *Threshold
}

type Threshold struct {
    fetch *LogLevel
    query *LogLevel
}

type LogLevel struct {
    warn  *int8
    trace *int8
    debug *int8
    info  *int8
}

type Idle struct {
    after *Interval
}

type Analyze struct {
    maxTokenCount *uint
}

type Highlight struct {
    maxAnalyzedOffset *uint
}

type Routing struct {
    allocation *Allocation
}

type Unassigned struct {
    nodeLeft *NodeLeft
}

type NodeLeft struct {
    delayedTimeout *Interval
}

type Write struct {
    waitForActiveShards *uint
}
