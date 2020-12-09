package settings

import "strings"

const (
    AnalyzerNameStandard   AnalyzerName = "standard"
    AnalyzerNameSimple     AnalyzerName = "simple"
    AnalyzerNameWhitespace AnalyzerName = "whitespace"
    AnalyzerNameStop       AnalyzerName = "stop"
    AnalyzerNameKeyword    AnalyzerName = "keyword"

    // see https://www.elastic.co/guide/en/elasticsearch/reference/7.x/analysis-lang-analyzer.html
    AnalyzerNameArabic     AnalyzerName = "arabic"
    AnalyzerNameArmenian   AnalyzerName = "armenian"
    AnalyzerNameBasque     AnalyzerName = "basque"
    AnalyzerNameBengali    AnalyzerName = "bengali"
    AnalyzerNameBulgarian  AnalyzerName = "bulgarian"
    AnalyzerNameCatalan    AnalyzerName = "catalan"
    AnalyzerNameCzech      AnalyzerName = "czech"
    AnalyzerNameDutch      AnalyzerName = "dutch"
    AnalyzerNameEnglish    AnalyzerName = "english"
    AnalyzerNameFinnish    AnalyzerName = "finnish"
    AnalyzerNameFrench     AnalyzerName = "french"
    AnalyzerNameGalician   AnalyzerName = "galician"
    AnalyzerNameGerman     AnalyzerName = "german"
    AnalyzerNameHindi      AnalyzerName = "hindi"
    AnalyzerNameHungarian  AnalyzerName = "hungarian"
    AnalyzerNameIndonesian AnalyzerName = "indonesian"
    AnalyzerNameIrish      AnalyzerName = "irish"
    AnalyzerNameItalian    AnalyzerName = "italian"
    AnalyzerNameLatvian    AnalyzerName = "latvian"
    AnalyzerNameLithuanian AnalyzerName = "lithuanian"
    AnalyzerNameNorwegian  AnalyzerName = "norwegian"
    AnalyzerNamePortuguese AnalyzerName = "portuguese"
    AnalyzerNameRomanian   AnalyzerName = "romanian"
    AnalyzerNameRussian    AnalyzerName = "russian"
    AnalyzerNameSorani     AnalyzerName = "sorani"
    AnalyzerNameSpanish    AnalyzerName = "spanish"
    AnalyzerNameSwedish    AnalyzerName = "swedish"
    AnalyzerNameTurkish    AnalyzerName = "turkish"

    AnalyzerTypeStandard    AnalyzerType = "standard"
    AnalyzerTypeStop        AnalyzerType = "stop"
    AnalyzerTypePattern     AnalyzerType = "pattern"
    AnalyzerTypeFingerprint AnalyzerType = "fingerprint"
)

type AnalyzerType string
type AnalyzerName string

type Analyzer interface {
    Type() AnalyzerType
    Name() AnalyzerName
    Source() (interface{}, error)
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/7.x/analysis-standard-analyzer.html
type AnalyzerStandard struct {
    name           AnalyzerName
    maxTokenLength *uint8
    stopWordsPath  *string
    stopWords      []string
}

func NewAnalyzerStandard(name string) *AnalyzerStandard {
    return &AnalyzerStandard{name: AnalyzerName(name)}
}

func (analyzer *AnalyzerStandard) SetMaxTokenLength(maxTokenLength uint8) *AnalyzerStandard {
    analyzer.maxTokenLength = &maxTokenLength

    return analyzer
}

func (analyzer *AnalyzerStandard) SetStopWordsPath(stopWordsPath string) *AnalyzerStandard {
    analyzer.stopWordsPath = &stopWordsPath

    return analyzer
}

func (analyzer *AnalyzerStandard) AddStopWords(stopWords ...string) *AnalyzerStandard {
    analyzer.stopWords = append(analyzer.stopWords, stopWords...)

    return analyzer
}

func (analyzer *AnalyzerStandard) Name() AnalyzerName {
    return analyzer.name
}

func (analyzer *AnalyzerStandard) Type() AnalyzerType {
    return AnalyzerTypeStandard
}

func (analyzer *AnalyzerStandard) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": analyzer.Type(),
    }

    if analyzer.maxTokenLength != nil {
        source["max_token_length"] = *analyzer.maxTokenLength
    }

    if analyzer.stopWordsPath != nil {
        source["stopwords_path"] = *analyzer.stopWordsPath
    }

    if len(analyzer.stopWords) > 0 {
        source["stopwords"] = analyzer.stopWords
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/7.x/analysis-stop-analyzer.html
type AnalyzerStop struct {
    name          AnalyzerName
    stopWordsPath *string
    stopWords     []string
}

func NewAnalyzerStop(name string) *AnalyzerStop {
    return &AnalyzerStop{name: AnalyzerName(name)}
}

func (analyzer *AnalyzerStop) SetStopWordsPath(stopWordsPath string) *AnalyzerStop {
    analyzer.stopWordsPath = &stopWordsPath

    return analyzer
}

func (analyzer *AnalyzerStop) AddStopWords(stopWords ...string) *AnalyzerStop {
    analyzer.stopWords = append(analyzer.stopWords, stopWords...)

    return analyzer
}

func (analyzer *AnalyzerStop) Name() AnalyzerName {
    return analyzer.name
}

func (analyzer *AnalyzerStop) Type() AnalyzerType {
    return AnalyzerTypeStop
}

func (analyzer *AnalyzerStop) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": analyzer.Type(),
    }

    if analyzer.stopWordsPath != nil {
        source["stopwords_path"] = *analyzer.stopWordsPath
    }

    if len(analyzer.stopWords) > 0 {
        source["stopwords"] = analyzer.stopWords
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/7.x/analysis-pattern-analyzer.html
type AnalyzerPattern struct {
    name          AnalyzerName
    lowercase     *bool
    pattern       *string
    stopWordsPath *string
    stopWords     []string
    flags         []JavaRegularFlag
}

func NewAnalyzerPattern(name string) *AnalyzerPattern {
    return &AnalyzerPattern{name: AnalyzerName(name)}
}

func (analyzer *AnalyzerPattern) SetLowercase(lowercase bool) *AnalyzerPattern {
    analyzer.lowercase = &lowercase

    return analyzer
}

func (analyzer *AnalyzerPattern) SetPattern(pattern string) *AnalyzerPattern {
    analyzer.pattern = &pattern

    return analyzer
}

func (analyzer *AnalyzerPattern) SetStopWordsPath(stopWordsPath string) *AnalyzerPattern {
    analyzer.stopWordsPath = &stopWordsPath

    return analyzer
}

func (analyzer *AnalyzerPattern) AddStopWords(stopWords ...string) *AnalyzerPattern {
    analyzer.stopWords = append(analyzer.stopWords, stopWords...)

    return analyzer
}

func (analyzer *AnalyzerPattern) AddFlags(flags ...JavaRegularFlag) *AnalyzerPattern {
    analyzer.flags = append(analyzer.flags, flags...)

    return analyzer
}

func (analyzer *AnalyzerPattern) Name() AnalyzerName {
    return analyzer.name
}

func (analyzer *AnalyzerPattern) Type() AnalyzerType {
    return AnalyzerTypePattern
}

func (analyzer *AnalyzerPattern) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": analyzer.Type(),
    }

    if analyzer.pattern != nil {
        source["pattern"] = *analyzer.pattern
    }

    if len(analyzer.flags) > 0 {
        var flags []string
        for _, flag := range analyzer.flags {
            flags = append(flags, string(flag))
        }

        source["flags"] = strings.Join(flags, JavaRegularFlagSeparator)
    }

    if analyzer.lowercase != nil {
        source["lowercase"] = *analyzer.lowercase
    }

    if analyzer.stopWordsPath != nil {
        source["stopwords_path"] = *analyzer.stopWordsPath
    }

    if len(analyzer.stopWords) > 0 {
        source["stopwords"] = analyzer.stopWords
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/7.x/analysis-fingerprint-analyzer.html
type AnalyzerFingerprint struct {
    name          AnalyzerName
    maxOutputSize *uint8
    separator     *string
    stopWordsPath *string
    stopWords     []string
}

func NewAnalyzerFingerprint(name string) *AnalyzerFingerprint {
    return &AnalyzerFingerprint{name: AnalyzerName(name)}
}

func (analyzer *AnalyzerFingerprint) SetMaxOutputSize(maxOutputSize uint8) *AnalyzerFingerprint {
    analyzer.maxOutputSize = &maxOutputSize

    return analyzer
}

func (analyzer *AnalyzerFingerprint) SetSeparator(separator string) *AnalyzerFingerprint {
    analyzer.separator = &separator

    return analyzer
}

func (analyzer *AnalyzerFingerprint) SetStopWordsPath(stopWordsPath string) *AnalyzerFingerprint {
    analyzer.stopWordsPath = &stopWordsPath

    return analyzer
}

func (analyzer *AnalyzerFingerprint) AddStopWords(stopWords ...string) *AnalyzerFingerprint {
    analyzer.stopWords = append(analyzer.stopWords, stopWords...)

    return analyzer
}

func (analyzer *AnalyzerFingerprint) Name() AnalyzerName {
    return analyzer.name
}

func (analyzer *AnalyzerFingerprint) Type() AnalyzerType {
    return AnalyzerTypeFingerprint
}

func (analyzer *AnalyzerFingerprint) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": analyzer.Type(),
    }

    if analyzer.separator != nil {
        source["separator"] = *analyzer.separator
    }

    if analyzer.maxOutputSize != nil {
        source["max_output_size"] = *analyzer.maxOutputSize
    }

    if analyzer.stopWordsPath != nil {
        source["stopwords_path"] = *analyzer.stopWordsPath
    }

    if len(analyzer.stopWords) > 0 {
        source["stopwords"] = analyzer.stopWords
    }

    return source, nil
}
