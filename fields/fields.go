package fields

import "errors"

const (
    TypeBinary            Type = "binary"
    TypeBoolean           Type = "boolean"
    TypeKeyword           Type = "keyword"
    TypeNumberLong        Type = "long"
    TypeNumberInteger     Type = "integer"
    TypeNumberShort       Type = "short"
    TypeNumberByte        Type = "byte"
    TypeNumberDouble      Type = "double"
    TypeNumberFloat       Type = "float"
    TypeNumberHalfFloat   Type = "half_float"
    TypeNumberScaledFloat Type = "scaled_float"
    TypeDate              Type = "date"
    TypeDateNano          Type = "date_nanos"
    TypeAlias             Type = "alias"

    DocsIndexOption      IndexOption = "docs"
    FreqsIndexOption     IndexOption = "freqs"
    PositionsIndexOption IndexOption = "positions"
    OffsetsIndexOption   IndexOption = "offsets"

    SimilarityBM25    Similarity = "BM25"
    SimilarityClassic Similarity = "classic" // TF/IDF
    SimilarityBoolean Similarity = "boolean"

    LowercaseNormalizer Normalizer = "lowercase"

    // @see https://www.elastic.co/guide/en/elasticsearch/reference/7.9/date.html
    // @see https://docs.oracle.com/javase/8/docs/api/java/util/Locale.html
    LocalRoot               Local = "ROOT" // default in elasticsearch
    LocalCanada             Local = "CANADA"
    LocalEnglish            Local = "ENGLISH"
    LocalFrench             Local = "FRENCH"
    LocalGerman             Local = "GERMAN"
    LocalItalian            Local = "ITALIAN"
    LocalJapanese           Local = "JAPANESE"
    LocalKorean             Local = "KOREAN"
    LocalChinese            Local = "CHINESE"
    LocalSimplifiedChinese  Local = "SIMPLIFIED_CHINESE"
    LocalTraditionalChinese Local = "TRADITIONAL_CHINESE"
    LocalFrance             Local = "FRANCE"
    LocalGermany            Local = "GERMANY"
    LocalItaly              Local = "ITALY"
    LocalJapan              Local = "JAPAN"
    LocalKorea              Local = "KOREA"
    LocalChina              Local = "CHINA"
    LocalPRC                Local = "PRC"
    LocalTaiwan             Local = "TAIWAN"
    LocalUK                 Local = "UK"
    LocalUS                 Local = "US"
    LocalCanadaFrench       Local = "CANADA_FRENCH"

    DateFormatSeparator = "||"
)

var ErrorScalingFactorNotScaledFloat = errors.New("scaling_factor can only be installed in field type scaled_float")
var ErrorScalingFactorNotSet = errors.New("scaling_factor required field on scaled_float type")

type Type string
type IndexOption string
type Similarity string
type Normalizer string
type Local string

type Fielder interface {
    Sourcer
    GetType() Type
}

type Sourcer interface {
    Source() (interface{}, error)
}
