package settings

const (
    AnalyzerStandard AnalyzerName = "standard"
)

type AnalyzerName string

func NewAnalyzerName(name string) AnalyzerName {
    return AnalyzerName(name)
}
