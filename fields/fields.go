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

    IndexOptionDocs      IndexOption = "docs"
    IndexOptionFreqs     IndexOption = "freqs"
    IndexOptionPositions IndexOption = "positions"
    IndexOptionOffsets   IndexOption = "offsets"

    SimilarityBM25    Similarity = "BM25"
    SimilarityClassic Similarity = "classic" // TF/IDF
    SimilarityBoolean Similarity = "boolean"

    NormalizerLowercase Normalizer = "lowercase"

    // @see https://www.elastic.co/guide/en/elasticsearch/reference/7.x/date.html
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

    DynamicEnabled  Dynamic = "true"
    DynamicDisabled Dynamic = "false"
    DynamicStrict   Dynamic = "strict"

    DateFormatSeparator = "||"

    // @see https://docs.oracle.com/javase/7/docs/api/java/text/SimpleDateFormat.html
    ANSIC      DateFormat = "EE MMM d HH:mm:ss yyyy"
    UnixDate   DateFormat = "EE MMM d HH:mm:ss z yyyy"
    RubyDate   DateFormat = "EE MMM dd HH:mm:ss Z yyyy"
    RFC822     DateFormat = "dd MMM yy HH:mm z"
    RFC822Z    DateFormat = "dd MMM yy HH:mm Z" // RFC822 with numeric zone
    RFC850     DateFormat = "EEEE, dd-MMM-yy HH:mm:ss z"
    RFC1123    DateFormat = "EEE, dd MMM yyyy HH:mm:ss z"
    RFC1123Z   DateFormat = "EEE, dd MMM yyyy HH:mm:ss Z" // RFC1123 with numeric zone
    RFC3339    DateFormat = "yyyy-MM-dd'T'HH:mm:ssXXX"
    Kitchen    DateFormat = "H:mma"
    Stamp      DateFormat = "MMM d HH:mm:ss"
    StampMilli DateFormat = "MMM d HH:mm:ss'.'S"
)

var ErrorScalingFactorNotScaledFloat = errors.New("scaling_factor can only be installed in field type scaled_float")
var ErrorScalingFactorNotSet = errors.New("scaling_factor required field on scaled_float type")
var ErrorDynamicUnknownType = errors.New("dynamic unknown type")

type Type string
type IndexOption string
type Similarity string
type Normalizer string
type Local string
type Dynamic string
type DateFormat string

type Fielder interface {
    Sourcer
    GetType() Type
}

type Sourcer interface {
    Source() (interface{}, error)
}
