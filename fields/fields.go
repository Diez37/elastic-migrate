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
    TypeObject            Type = "object"
    TypeNested            Type = "nested"
    TypeJoin              Type = "join"
    TypeRangeInteger      Type = "integer_range"
    TypeRangeFloat        Type = "float_range"
    TypeRangeLong         Type = "long_range"
    TypeRangeDouble       Type = "double_range"
    TypeRangeIp           Type = "ip_range"
    TypeRangeDate         Type = "date_range"
    TypeIp                Type = "ip"
    TypeText              Type = "text"
    TypeSearchAsYouType   Type = "search_as_you_type"
    TypeTokenCount        Type = "token_count"
    TypeRankFeature       Type = "rank_feature"

    IndexOptionDocs      IndexOption = "docs"
    IndexOptionFreqs     IndexOption = "freqs"
    IndexOptionPositions IndexOption = "positions"
    IndexOptionOffsets   IndexOption = "offsets"

    SimilarityBM25    Similarity = "BM25"
    SimilarityClassic Similarity = "classic" // TF/IDF
    SimilarityBoolean Similarity = "boolean"

    NormalizerLowercase Normalizer = "lowercase"

    // see https://www.elastic.co/guide/en/elasticsearch/reference/7.x/date.html
    // see https://docs.oracle.com/javase/8/docs/api/java/util/Locale.html
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

    DynamicEnabled  Dynamic = "true"
    DynamicDisabled Dynamic = "false"
    DynamicStrict   Dynamic = "strict"

    DateFormatSeparator = "||"

    // see https://docs.oracle.com/javase/7/docs/api/java/text/SimpleDateFormat.html
    DateFormatANSIC      DateFormat = "EE MMM d HH:mm:ss yyyy"
    DateFormatUnixDate   DateFormat = "EE MMM d HH:mm:ss z yyyy"
    DateFormatRubyDate   DateFormat = "EE MMM dd HH:mm:ss Z yyyy"
    DateFormatRFC822     DateFormat = "dd MMM yy HH:mm z"
    DateFormatRFC822Z    DateFormat = "dd MMM yy HH:mm Z" // RFC822 with numeric zone
    DateFormatRFC850     DateFormat = "EEEE, dd-MMM-yy HH:mm:ss z"
    DateFormatRFC1123    DateFormat = "EEE, dd MMM yyyy HH:mm:ss z"
    DateFormatRFC1123Z   DateFormat = "EEE, dd MMM yyyy HH:mm:ss Z" // RFC1123 with numeric zone
    DateFormatRFC3339    DateFormat = "yyyy-MM-dd'T'HH:mm:ssXXX"
    DateFormatKitchen    DateFormat = "H:mma"
    DateFormatStamp      DateFormat = "MMM d HH:mm:ss"
    DateFormatStampMilli DateFormat = "MMM d HH:mm:ss'.'S"

    TermVectorNo                           TermVector = "no"
    TermVectorYes                          TermVector = "yes"
    TermVectorWithPositions                TermVector = "with_positions"
    TermVectorWithOffsets                  TermVector = "with_offsets"
    TermVectorWithPositionsOffsets         TermVector = "with_positions_offsets"
    TermVectorWithPositionsPayloads        TermVector = "with_positions_payloads"
    TermVectorWithPositionsOffsetsPayloads TermVector = "with_positions_offsets_payloads"

    ShingleSizeMaximum int = 4
    ShingleSizeMinimum int = 2
)

var ErrorSearchAnalyzer = errors.New("analyzer must be set when search_analyzer is set")
var ErrorShingleSizeMaximum = errors.New("max_shingle_size maximum value = 4")
var ErrorShingleSizeMinimum = errors.New("max_shingle_size minimum value = 2")
var ErrorDynamicUnknownType = errors.New("dynamic unknown type")
var ErrorSearchQuoteAnalyzer = errors.New("analyzer and search_analyzer must be set when search_quote_analyzer is set")
var ErrorScalingFactorNotSet = errors.New("scaling_factor required field on scaled_float type")
var ErrorScalingFactorNotScaledFloat = errors.New("scaling_factor can only be installed in field type scaled_float")

type Type string
type Local string
type Dynamic string
type Similarity string
type Normalizer string
type DateFormat string
type TermVector string
type IndexOption string

type Fielder interface {
    Sourcer
    GetType() Type
}

type Sourcer interface {
    Source() (interface{}, error)
}
