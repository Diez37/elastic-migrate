package settings

const (
    AnalyzerStandard   AnalyzerName = "standard"
    AnalyzerSimple     AnalyzerName = "simple"
    AnalyzerWhitespace AnalyzerName = "whitespace"
    AnalyzerStop       AnalyzerName = "stop"

    CheckOnStartupFalse    CheckOnStartup = "false"
    CheckOnStartupTrue     CheckOnStartup = "true"
    CheckOnStartupChecksum CheckOnStartup = "checksum"

    AllocationAll          AllocationVal = "all"
    AllocationPrimaries    AllocationVal = "primaries"
    AllocationNewPrimaries AllocationVal = "new_primaries"
    AllocationNone         AllocationVal = "none"

    RebalanceAll       RebalanceVal = "all"
    RebalancePrimaries RebalanceVal = "primaries"
    RebalanceReplicas  RebalanceVal = "replicas"
    RebalanceNone      RebalanceVal = "none"

    AllocationAttributeName      AllocationAttribute = "_name"
    AllocationAttributeHostIp    AllocationAttribute = "_host_ip"
    AllocationAttributePublishIp AllocationAttribute = "_publish_ip"
    AllocationAttributeIp        AllocationAttribute = "_ip"
    AllocationAttributeHost      AllocationAttribute = "_host"
    AllocationAttributeId        AllocationAttribute = "_id"

    CodecBestCompression IndexCode = "best_compression"

    PipelineNameNone PipelineName = "_none"

    SlowWarn SlowLogLevel = "WARN"
    Trace    SlowLogLevel = "TRACE"
    Debug    SlowLogLevel = "DEBUG"
    Info     SlowLogLevel = "INFO"
)

type AnalyzerName string
type CheckOnStartup string
type IndexCode string
type AllocationVal string
type AllocationAttribute string
type RebalanceVal string
type PipelineName string
type SlowLogLevel string

type ErrorIntervalNotValid error
type ErrorSizeNotValid error
