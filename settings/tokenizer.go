package settings

const (
    TokenizerNameClassic       TokenizerName = "classic"
    TokenizerNameEdgeNgram     TokenizerName = "edge_ngram"
    TokenizerNameNgram         TokenizerName = "ngram"
    TokenizerNameKeyword       TokenizerName = "keyword"
    TokenizerNameLetter        TokenizerName = "letter"
    TokenizerNameLowercase     TokenizerName = "lowercase"
    TokenizerNamePathHierarchy TokenizerName = "path_hierarchy"
    TokenizerNamePattern       TokenizerName = "pattern"
    TokenizerNameStandard      TokenizerName = "standard"
    TokenizerNameThai          TokenizerName = "thai"
    TokenizerNameUaxUrlEmail   TokenizerName = "uax_url_email"

    TokenizerTypeCharGroup          TokenizerType = "char_group"
    TokenizerTypeClassic            TokenizerType = "classic"
    TokenizerTypeEdgeNgram          TokenizerType = "edge_ngram"
    TokenizerTypeNgram              TokenizerType = "ngram"
    TokenizerTypeKeyword            TokenizerType = "keyword"
    TokenizerTypePathHierarchy      TokenizerType = "path_hierarchy"
    TokenizerTypePattern            TokenizerType = "pattern"
    TokenizerTypeSimplePattern      TokenizerType = "simple_pattern"
    TokenizerTypeSimplePatternSplit TokenizerType = "simple_pattern_split"
    TokenizerTypeStandard           TokenizerType = "standard"
    TokenizerTypeUaxUrlEmail        TokenizerType = "uax_url_email"
    TokenizerTypeWhitespace         TokenizerType = "whitespace"

    EdgeNgramTokenCharsLetter      EdgeNgramTokenChars = "letter"
    EdgeNgramTokenCharsDigit       EdgeNgramTokenChars = "digit"
    EdgeNgramTokenCharsWhitespace  EdgeNgramTokenChars = "whitespace"
    EdgeNgramTokenCharsPunctuation EdgeNgramTokenChars = "punctuation"
    EdgeNgramTokenCharsSymbol      EdgeNgramTokenChars = "symbol"
    EdgeNgramTokenCharsCustom      EdgeNgramTokenChars = "custom"
)

type TokenizerType string
type TokenizerName string
type EdgeNgramTokenChars string

type Tokenizer interface {
    Type() TokenizerType
    Name() TokenizerName
    Source() (interface{}, error)
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-chargroup-tokenizer.html
type TokenizerCharGroup struct {
    name            TokenizerName
    maxTokenLength  *uint32
    tokenizeOnChars []string
}

func NewTokenizerCharGroup(name string) *TokenizerCharGroup {
    return &TokenizerCharGroup{name: TokenizerName(name)}
}

func (tokenizer *TokenizerCharGroup) SetMaxTokenLength(maxTokenLength uint32) *TokenizerCharGroup {
    tokenizer.maxTokenLength = &maxTokenLength

    return tokenizer
}

func (tokenizer *TokenizerCharGroup) AddTokenizeOnChars(tokenizeOnChars ...string) *TokenizerCharGroup {
    tokenizer.tokenizeOnChars = append(tokenizer.tokenizeOnChars, tokenizeOnChars...)

    return tokenizer
}

func (tokenizer *TokenizerCharGroup) Name() TokenizerName {
    return tokenizer.name
}

func (tokenizer *TokenizerCharGroup) Type() TokenizerType {
    return TokenizerTypeCharGroup
}

func (tokenizer *TokenizerCharGroup) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": tokenizer.Type(),
    }

    if tokenizer.maxTokenLength != nil {
        source["max_token_length"] = *tokenizer.maxTokenLength
    }

    if len(tokenizer.tokenizeOnChars) > 0 {
        source["tokenize_on_chars"] = tokenizer.tokenizeOnChars
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-classic-tokenizer.html
type TokenizerClassic struct {
    name           TokenizerName
    maxTokenLength *uint32
}

func NewTokenizerClassic(name string) *TokenizerClassic {
    return &TokenizerClassic{name: TokenizerName(name)}
}

func (tokenizer *TokenizerClassic) SetMaxTokenLength(maxTokenLength uint32) *TokenizerClassic {
    tokenizer.maxTokenLength = &maxTokenLength

    return tokenizer
}

func (tokenizer *TokenizerClassic) Name() TokenizerName {
    return tokenizer.name
}

func (tokenizer *TokenizerClassic) Type() TokenizerType {
    return TokenizerTypeClassic
}

func (tokenizer *TokenizerClassic) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": tokenizer.Type(),
    }

    if tokenizer.maxTokenLength != nil {
        source["max_token_length"] = *tokenizer.maxTokenLength
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-edgengram-tokenizer.html
type TokenizerNgram struct {
    _type            TokenizerType
    name             TokenizerName
    minGram          *uint32
    maxGram          *uint32
    tokenChars       []EdgeNgramTokenChars
    customTokenChars []string
}

func NewTokenizerNgram(name string) *TokenizerNgram {
    return &TokenizerNgram{name: TokenizerName(name), _type: TokenizerTypeNgram}
}

func NewTokenizerEdgeNgram(name string) *TokenizerNgram {
    return &TokenizerNgram{name: TokenizerName(name), _type: TokenizerTypeEdgeNgram}
}

func (tokenizer *TokenizerNgram) SetMinGram(minGram uint32) *TokenizerNgram {
    tokenizer.minGram = &minGram

    return tokenizer
}

func (tokenizer *TokenizerNgram) SetMaxGram(maxGram uint32) *TokenizerNgram {
    tokenizer.maxGram = &maxGram

    return tokenizer
}

func (tokenizer *TokenizerNgram) AddTokenChars(tokenChars ...EdgeNgramTokenChars) *TokenizerNgram {
    tokenizer.tokenChars = append(tokenizer.tokenChars, tokenChars...)

    return tokenizer
}

func (tokenizer *TokenizerNgram) AddCustomTokenChars(customTokenChars ...string) *TokenizerNgram {
    tokenizer.customTokenChars = append(tokenizer.customTokenChars, customTokenChars...)

    return tokenizer
}

func (tokenizer *TokenizerNgram) Name() TokenizerName {
    return tokenizer.name
}

func (tokenizer *TokenizerNgram) Type() TokenizerType {
    return tokenizer._type
}

func (tokenizer *TokenizerNgram) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": tokenizer.Type(),
    }

    if tokenizer.minGram != nil {
        source["min_gram"] = *tokenizer.minGram
    }

    if tokenizer.maxGram != nil {
        source["max_gram"] = *tokenizer.maxGram
    }

    if len(tokenizer.tokenChars) > 0 {
        source["token_chars"] = tokenizer.tokenChars
    }

    if len(tokenizer.customTokenChars) > 0 {
        source["custom_token_chars"] = tokenizer.customTokenChars
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-keyword-tokenizer.html
type TokenizerKeyword struct {
    name       TokenizerName
    bufferSize *uint32
}

func NewTokenizerKeyword(name string) *TokenizerKeyword {
    return &TokenizerKeyword{name: TokenizerName(name)}
}

func (tokenizer *TokenizerKeyword) SetBufferSize(bufferSize uint32) *TokenizerKeyword {
    tokenizer.bufferSize = &bufferSize

    return tokenizer
}

func (tokenizer *TokenizerKeyword) Name() TokenizerName {
    return tokenizer.name
}

func (tokenizer *TokenizerKeyword) Type() TokenizerType {
    return TokenizerTypeKeyword
}

func (tokenizer *TokenizerKeyword) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": tokenizer.Type(),
    }

    if tokenizer.bufferSize != nil {
        source["buffer_size"] = *tokenizer.bufferSize
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-pathhierarchy-tokenizer.html
type TokenizerPathHierarchy struct {
    name        TokenizerName
    reverse     *bool
    delimiter   *string
    replacement *string
    bufferSize  *uint32
    skip        *uint32
}

func NewTokenizerPathHierarchy(name string) *TokenizerPathHierarchy {
    return &TokenizerPathHierarchy{name: TokenizerName(name)}
}

func (tokenizer *TokenizerPathHierarchy) SetReverse(reverse bool) *TokenizerPathHierarchy {
    tokenizer.reverse = &reverse

    return tokenizer
}

func (tokenizer *TokenizerPathHierarchy) SetDelimiter(delimiter string) *TokenizerPathHierarchy {
    tokenizer.delimiter = &delimiter

    return tokenizer
}

func (tokenizer *TokenizerPathHierarchy) SetReplacement(replacement string) *TokenizerPathHierarchy {
    tokenizer.replacement = &replacement

    return tokenizer
}

func (tokenizer *TokenizerPathHierarchy) SetBufferSize(bufferSize uint32) *TokenizerPathHierarchy {
    tokenizer.bufferSize = &bufferSize

    return tokenizer
}

func (tokenizer *TokenizerPathHierarchy) SetSkip(skip uint32) *TokenizerPathHierarchy {
    tokenizer.skip = &skip

    return tokenizer
}

func (tokenizer *TokenizerPathHierarchy) Name() TokenizerName {
    return tokenizer.name
}

func (tokenizer *TokenizerPathHierarchy) Type() TokenizerType {
    return TokenizerTypePathHierarchy
}

func (tokenizer *TokenizerPathHierarchy) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": tokenizer.Type(),
    }

    if tokenizer.reverse != nil {
        source["reverse"] = *tokenizer.reverse
    }

    if tokenizer.delimiter != nil {
        source["delimiter"] = *tokenizer.delimiter
    }

    if tokenizer.replacement != nil {
        source["replacement"] = *tokenizer.replacement
    }

    if tokenizer.bufferSize != nil {
        source["buffer_size"] = *tokenizer.bufferSize
    }

    if tokenizer.skip != nil {
        source["skip"] = *tokenizer.skip
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-pattern-tokenizer.html
type TokenizerPattern struct {
    name    TokenizerName
    pattern *string
    group   *uint8
    flags   []JavaRegularFlag
}

func NewTokenizerPattern(name string) *TokenizerPattern {
    return &TokenizerPattern{name: TokenizerName(name)}
}

func (tokenizer *TokenizerPattern) SetPattern(pattern string) *TokenizerPattern {
    tokenizer.pattern = &pattern

    return tokenizer
}

func (tokenizer *TokenizerPattern) SetGroup(group uint8) *TokenizerPattern {
    tokenizer.group = &group

    return tokenizer
}

func (tokenizer *TokenizerPattern) AddFlags(flags ...JavaRegularFlag) *TokenizerPattern {
    tokenizer.flags = append(tokenizer.flags, flags...)

    return tokenizer
}

func (tokenizer *TokenizerPattern) Name() TokenizerName {
    return tokenizer.name
}

func (tokenizer *TokenizerPattern) Type() TokenizerType {
    return TokenizerTypePattern
}

func (tokenizer *TokenizerPattern) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": tokenizer.Type(),
    }

    if tokenizer.pattern != nil {
        source["pattern"] = *tokenizer.pattern
    }

    if tokenizer.group != nil {
        source["group"] = *tokenizer.group
    }

    if len(tokenizer.flags) > 0 {
        flags := ""
        for _, flag := range tokenizer.flags {
            flags += string(flag) + JavaRegularFlagSeparator
        }

        source["flags"] = flags[0:len(flags)-1]
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-simplepattern-tokenizer.html
type TokenizerSimplePattern struct {
    _type   TokenizerType
    name    TokenizerName
    pattern *string
}

func NewTokenizerSimplePattern(name string) *TokenizerSimplePattern {
    return &TokenizerSimplePattern{name: TokenizerName(name), _type: TokenizerTypeSimplePattern}
}

func NewTokenizerSimplePatternSplit(name string) *TokenizerSimplePattern {
    return &TokenizerSimplePattern{name: TokenizerName(name), _type: TokenizerTypeSimplePatternSplit}
}

func (tokenizer *TokenizerSimplePattern) SetPattern(pattern string) *TokenizerSimplePattern {
    tokenizer.pattern = &pattern

    return tokenizer
}

func (tokenizer *TokenizerSimplePattern) Name() TokenizerName {
    return tokenizer.name
}

func (tokenizer *TokenizerSimplePattern) Type() TokenizerType {
    return tokenizer._type
}

func (tokenizer *TokenizerSimplePattern) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": tokenizer.Type(),
    }

    if tokenizer.pattern != nil {
        source["pattern"] = *tokenizer.pattern
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-standard-tokenizer.html
type TokenizerStandard struct {
    _type          TokenizerType
    name           TokenizerName
    maxTokenLength *uint16
}

func NewTokenizerStandard(name string) *TokenizerStandard {
    return &TokenizerStandard{name: TokenizerName(name), _type: TokenizerTypeStandard}
}

func NewTokenizerUaxUrlEmail(name string) *TokenizerStandard {
    return &TokenizerStandard{name: TokenizerName(name), _type: TokenizerTypeUaxUrlEmail}
}

func NewTokenizerWhitespace(name string) *TokenizerStandard {
    return &TokenizerStandard{name: TokenizerName(name), _type: TokenizerTypeWhitespace}
}

func (tokenizer *TokenizerStandard) SetMaxTokenLength(maxTokenLength uint16) *TokenizerStandard {
    tokenizer.maxTokenLength = &maxTokenLength

    return tokenizer
}

func (tokenizer *TokenizerStandard) Name() TokenizerName {
    return tokenizer.name
}

func (tokenizer *TokenizerStandard) Type() TokenizerType {
    return tokenizer._type
}

func (tokenizer *TokenizerStandard) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": tokenizer.Type(),
    }

    if tokenizer.maxTokenLength != nil {
        source["max_token_length"] = *tokenizer.maxTokenLength
    }

    return source, nil
}
