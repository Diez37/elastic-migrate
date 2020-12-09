package settings

import (
    "errors"
    "fmt"
    "github.com/olivere/elastic/v7"
    "strings"
)

const (
    FilterNameArabic              FilterName = "arabic_normalization"
    FilterNameAsciifolding        FilterName = "asciifolding"
    FilterNameBengali             FilterName = "bengali_normalization"
    FilterNameCjkWidth            FilterName = "cjk_width"
    FilterNameDecimalDigit        FilterName = "decimal_digit"
    FilterNameDelimitedPayload    FilterName = "delimited_payload"
    FilterNameElision             FilterName = "elision"
    FilterNameGerman              FilterName = "german_normalization"
    FilterNameHindi               FilterName = "hindi_normalization"
    FilterNameIndic               FilterName = "indic_normalization"
    FilterNameLowercase           FilterName = "lowercase"
    FilterNamePersian             FilterName = "persian_normalization"
    FilterNameScandinavianFolding FilterName = "scandinavian_folding"
    FilterNameSerbian             FilterName = "serbian_normalization"
    FilterNameSorani              FilterName = "sorani_normalization"
    FilterNameUppercase           FilterName = "uppercase"
    FilterNameApostrophe          FilterName = "apostrophe"
    FilterNameAsciiFolding        FilterName = "asciifolding"
    FilterNameCjkBigram           FilterName = "cjk_bigram"
    FilterNameClassic             FilterName = "classic"
    FilterNameEdgeNgram           FilterName = "edge_ngram"
    FilterNameFingerprint         FilterName = "fingerprint"
    FilterNameFlattenGraph        FilterName = "flatten_graph"
    FilterNameKeywordRepeat       FilterName = "keyword_repeat"
    FilterNameKstem               FilterName = "kstem"
    FilterNameLength              FilterName = "length"
    FilterNameLimit               FilterName = "limit"
    FilterNameNgram               FilterName = "ngram"
    FilterNamePorterStem          FilterName = "porter_stem"
    FilterNameRemoveDuplicates    FilterName = "remove_duplicates"
    FilterNameReverse             FilterName = "reverse"
    FilterNameShingle             FilterName = "shingle"
    FilterNameStemmer             FilterName = "stemmer"
    FilterNameStop                FilterName = "stop"
    FilterNameTrim                FilterName = "trim"
    FilterNameTruncate            FilterName = "truncate"
    FilterNameUnique              FilterName = "unique"
    FilterNameWordDelimiter       FilterName = "word_delimiter"
    FilterNameWordDelimiterGraph  FilterName = "word_delimiter_graph"

    FilterTypeAsciiFolding            FilterType = "asciifolding"
    FilterTypeCjkBigram               FilterType = "cjk_bigram"
    FilterTypeCommonGrams             FilterType = "common_grams"
    FilterTypeCondition               FilterType = "condition"
    FilterTypeDelimitedPayload        FilterType = "delimited_payload"
    FilterTypeDictionaryDecompounder  FilterType = "dictionary_decompounder"
    FilterTypeEdgeNgram               FilterType = "edge_ngram"
    FilterTypeElision                 FilterType = "elision"
    FilterTypeFingerprint             FilterType = "fingerprint"
    FilterTypeHyphenationDecompounder FilterType = "hyphenation_decompounder"
    FilterTypeKeepTypes               FilterType = "keep_types"
    FilterTypeKeep                    FilterType = "keep"
    FilterTypeKeywordMarker           FilterType = "keyword_marker"
    FilterTypeLength                  FilterType = "length"
    FilterTypeLimit                   FilterType = "limit"
    FilterTypeLowercase               FilterType = "lowercase"
    FilterTypeMinHash                 FilterType = "min_hash"
    FilterTypeMultiplexer             FilterType = "multiplexer"
    FilterTypeNgram                   FilterType = "ngram"
    FilterTypePatternCapture          FilterType = "pattern_capture"
    FilterTypePatternReplace          FilterType = "pattern_replace"
    FilterTypePhonetic                FilterType = "phonetic"
    FilterTypePredicateTokenFilter    FilterType = "predicate_token_filter"
    FilterTypeShingle                 FilterType = "shingle"
    FilterTypeSnowball                FilterType = "snowball"
    FilterTypeStemmer                 FilterType = "stemmer"
    FilterTypeStemmerOverride         FilterType = "stemmer_override"
    FilterTypeStop                    FilterType = "stop"
    FilterTypeSynonym                 FilterType = "synonyms"
    FilterTypeSynonymGraph            FilterType = "synonym_graph"
    FilterTypeTruncate                FilterType = "truncate"
    FilterTypeUnique                  FilterType = "unique"
    FilterTypeWordDelimiter           FilterType = "word_delimiter"
    FilterTypeWordDelimiterGraph      FilterType = "word_delimiter_graph"

    BigramIgnoredScriptsHan      FilterBigramIgnoredScripts = "han"
    BigramIgnoredScriptsHangul   FilterBigramIgnoredScripts = "hangul"
    BigramIgnoredScriptsHiragana FilterBigramIgnoredScripts = "hiragana"
    BigramIgnoredScriptsKatakana FilterBigramIgnoredScripts = "katakana"

    DelimitedPayloadEncodingFloat    FilterDelimitedPayloadEncoding = "float"
    DelimitedPayloadEncodingInteger  FilterDelimitedPayloadEncoding = "int"
    DelimitedPayloadEncodingIdentity FilterDelimitedPayloadEncoding = "identity"

    EdgeNgramSideFront FilterEdgeNgramSide = "front"
    EdgeNgramSideBack  FilterEdgeNgramSide = "back"

    KeepTypesModeInclude FilterKeepTypesMode = "include"
    KeepTypesModeExclude FilterKeepTypesMode = "exclude"

    LowercaseLanguageTypeGreek   FilterLowercaseLanguageType = "greek"
    LowercaseLanguageTypeIrish   FilterLowercaseLanguageType = "irish"
    LowercaseLanguageTypeTurkish FilterLowercaseLanguageType = "turkish"

    PhoneticEncoderMetaphone       FilterPhoneticEncoder = "metaphone"
    PhoneticEncoderDoubleMetaphone FilterPhoneticEncoder = "double_metaphone"
    PhoneticEncoderSoundex         FilterPhoneticEncoder = "soundex"
    PhoneticEncoderRefinedSoundex  FilterPhoneticEncoder = "refined_soundex"
    PhoneticEncoderCaverphone1     FilterPhoneticEncoder = "caverphone1"
    PhoneticEncoderCaverphone2     FilterPhoneticEncoder = "caverphone2"
    PhoneticEncoderCologne         FilterPhoneticEncoder = "cologne"
    PhoneticEncoderNysiis          FilterPhoneticEncoder = "nysiis"
    PhoneticEncoderKoelnerphonetik FilterPhoneticEncoder = "koelnerphonetik"
    PhoneticEncoderHaasephonetik   FilterPhoneticEncoder = "haasephonetik"
    PhoneticEncoderBeiderMorse     FilterPhoneticEncoder = "beider_morse"
    PhoneticEncoderDaitchMokotoff  FilterPhoneticEncoder = "daitch_mokotoff"

    PhoneticRuleTypeExact  FilterPhoneticRuleType = "exact"
    PhoneticRuleTypeApprox FilterPhoneticRuleType = "approx"

    PhoneticNameTypeAshkenazi FilterPhoneticNameType = "ashkenazi"
    PhoneticNameTypeSephardic FilterPhoneticNameType = "sephardic"
    PhoneticNameTypeGeneric   FilterPhoneticNameType = "generic"

    PhoneticLanguageSetAny       FilterPhoneticLanguageSet = "any"
    PhoneticLanguageSetCommon    FilterPhoneticLanguageSet = "common"
    PhoneticLanguageSetCyrillic  FilterPhoneticLanguageSet = "cyrillic"
    PhoneticLanguageSetEnglish   FilterPhoneticLanguageSet = "english"
    PhoneticLanguageSetFrench    FilterPhoneticLanguageSet = "french"
    PhoneticLanguageSetGerman    FilterPhoneticLanguageSet = "german"
    PhoneticLanguageSetHebrew    FilterPhoneticLanguageSet = "hebrew"
    PhoneticLanguageSetHungarian FilterPhoneticLanguageSet = "hungarian"
    PhoneticLanguageSetPolish    FilterPhoneticLanguageSet = "polish"
    PhoneticLanguageSetRomanian  FilterPhoneticLanguageSet = "romanian"
    PhoneticLanguageSetRussian   FilterPhoneticLanguageSet = "russian"
    PhoneticLanguageSetSpanish   FilterPhoneticLanguageSet = "spanish"

    SnowballLanguageArabic     FilterSnowballLanguage = "Arabic"
    SnowballLanguageArmenian   FilterSnowballLanguage = "Armenian"
    SnowballLanguageBasque     FilterSnowballLanguage = "Basque"
    SnowballLanguageCatalan    FilterSnowballLanguage = "Catalan"
    SnowballLanguageDanish     FilterSnowballLanguage = "Danish"
    SnowballLanguageDutch      FilterSnowballLanguage = "Dutch"
    SnowballLanguageEnglish    FilterSnowballLanguage = "English"
    SnowballLanguageEstonian   FilterSnowballLanguage = "Estonian"
    SnowballLanguageFinnish    FilterSnowballLanguage = "Finnish"
    SnowballLanguageFrench     FilterSnowballLanguage = "French"
    SnowballLanguageGerman     FilterSnowballLanguage = "German"
    SnowballLanguageGerman2    FilterSnowballLanguage = "German2"
    SnowballLanguageHungarian  FilterSnowballLanguage = "Hungarian"
    SnowballLanguageItalian    FilterSnowballLanguage = "Italian"
    SnowballLanguageIrish      FilterSnowballLanguage = "Irish"
    SnowballLanguageKp         FilterSnowballLanguage = "Kp"
    SnowballLanguageLithuanian FilterSnowballLanguage = "Lithuanian"
    SnowballLanguageLovins     FilterSnowballLanguage = "Lovins"
    SnowballLanguageNorwegian  FilterSnowballLanguage = "Norwegian"
    SnowballLanguagePorter     FilterSnowballLanguage = "Porter"
    SnowballLanguagePortuguese FilterSnowballLanguage = "Portuguese"
    SnowballLanguageRomanian   FilterSnowballLanguage = "Romanian"
    SnowballLanguageRussian    FilterSnowballLanguage = "Russian"
    SnowballLanguageSpanish    FilterSnowballLanguage = "Spanish"
    SnowballLanguageSwedish    FilterSnowballLanguage = "Swedish"
    SnowballLanguageTurkish    FilterSnowballLanguage = "Turkish"

    StemmerLanguageArabic            FilterStemmerLanguage = "arabic"
    StemmerLanguageArmenian          FilterStemmerLanguage = "armenian"
    StemmerLanguageBasque            FilterStemmerLanguage = "basque"
    StemmerLanguageBengali           FilterStemmerLanguage = "bengali"
    StemmerLanguageBrazilian         FilterStemmerLanguage = "brazilian"
    StemmerLanguageBulgarian         FilterStemmerLanguage = "bulgarian"
    StemmerLanguageCatalan           FilterStemmerLanguage = "catalan"
    StemmerLanguageCzech             FilterStemmerLanguage = "czech"
    StemmerLanguageDanish            FilterStemmerLanguage = "danish"
    StemmerLanguageDutch             FilterStemmerLanguage = "dutch"
    StemmerLanguageDutchKp           FilterStemmerLanguage = "dutch_kp"
    StemmerLanguageEnglish           FilterStemmerLanguage = "english"
    StemmerLanguageLightEnglish      FilterStemmerLanguage = "light_english"
    StemmerLanguageLovins            FilterStemmerLanguage = "lovins"
    StemmerLanguageMinimalEnglish    FilterStemmerLanguage = "minimal_english"
    StemmerLanguagePorter2           FilterStemmerLanguage = "porter2"
    StemmerLanguagePossessiveEnglish FilterStemmerLanguage = "possessive_english"
    StemmerLanguageEstonian          FilterStemmerLanguage = "estonian"
    StemmerLanguageFinnish           FilterStemmerLanguage = "finnish"
    StemmerLanguageLightFinnish      FilterStemmerLanguage = "light_finnish"
    StemmerLanguageLightFrench       FilterStemmerLanguage = "light_french"
    StemmerLanguageFrench            FilterStemmerLanguage = "french"
    StemmerLanguageMinimalFrench     FilterStemmerLanguage = "minimal_french"
    StemmerLanguageGalician          FilterStemmerLanguage = "galician"
    StemmerLanguageMinimalGalician   FilterStemmerLanguage = "minimal_galician"
    StemmerLanguageLightGerman       FilterStemmerLanguage = "light_german"
    StemmerLanguageGerman            FilterStemmerLanguage = "german"
    StemmerLanguageGerman2           FilterStemmerLanguage = "german2"
    StemmerLanguageMinimalGerman     FilterStemmerLanguage = "minimal_german"
    StemmerLanguageGreek             FilterStemmerLanguage = "greek"
    StemmerLanguageHindi             FilterStemmerLanguage = "hindi"
    StemmerLanguageHungarian         FilterStemmerLanguage = "hungarian"
    StemmerLanguageLightHungarian    FilterStemmerLanguage = "light_hungarian"
    StemmerLanguageIndonesian        FilterStemmerLanguage = "indonesian"
    StemmerLanguageIrish             FilterStemmerLanguage = "irish"
    StemmerLanguageLightItalian      FilterStemmerLanguage = "light_italian"
    StemmerLanguageItalian           FilterStemmerLanguage = "italian"
    StemmerLanguageSorani            FilterStemmerLanguage = "sorani"
    StemmerLanguageLatvian           FilterStemmerLanguage = "latvian"
    StemmerLanguageLithuanian        FilterStemmerLanguage = "lithuanian"
    StemmerLanguageNorwegian         FilterStemmerLanguage = "norwegian"
    StemmerLanguageLightNorwegian    FilterStemmerLanguage = "light_norwegian"
    StemmerLanguageMinimalNorwegian  FilterStemmerLanguage = "minimal_norwegian"
    StemmerLanguageLightNynorsk      FilterStemmerLanguage = "light_nynorsk"
    StemmerLanguageMinimalNynorsk    FilterStemmerLanguage = "minimal_nynorsk"
    StemmerLanguageLightPortuguese   FilterStemmerLanguage = "light_portuguese"
    StemmerLanguageMinimalPortuguese FilterStemmerLanguage = "minimal_portuguese"
    StemmerLanguagePortuguese        FilterStemmerLanguage = "portuguese"
    StemmerLanguagePortugueseRslp    FilterStemmerLanguage = "portuguese_rslp"
    StemmerLanguageRomanian          FilterStemmerLanguage = "romanian"
    StemmerLanguageRussian           FilterStemmerLanguage = "russian"
    StemmerLanguageLightRussian      FilterStemmerLanguage = "light_russian"
    StemmerLanguageLightSpanish      FilterStemmerLanguage = "light_spanish"
    StemmerLanguageSpanish           FilterStemmerLanguage = "spanish"
    StemmerLanguageSwedish           FilterStemmerLanguage = "swedish"
    StemmerLanguageLightSwedish      FilterStemmerLanguage = "light_swedish"
    StemmerLanguageTurkish           FilterStemmerLanguage = "turkish"

    StopWordsArabic     string = "_arabic_"
    StopWordsArmenian   string = "_armenian_"
    StopWordsBasque     string = "_basque_"
    StopWordsBengali    string = "_bengali_"
    StopWordsBrazilian  string = "_brazilian_"
    StopWordsBulgarian  string = "_bulgarian_"
    StopWordsCatalan    string = "_catalan_"
    StopWordsCjk        string = "_cjk_"
    StopWordsCzech      string = "_czech_"
    StopWordsDanish     string = "_danish_"
    StopWordsDutch      string = "_dutch_"
    StopWordsEnglish    string = "_english_"
    StopWordsEstonian   string = "_estonian_"
    StopWordsFinnish    string = "_finnish_"
    StopWordsFrench     string = "_french_"
    StopWordsGalician   string = "_galician_"
    StopWordsGerman     string = "_german_"
    StopWordsGreek      string = "_greek_"
    StopWordsHindi      string = "_hindi_"
    StopWordsHungarian  string = "_hungarian_"
    StopWordsIndonesian string = "_indonesian_"
    StopWordsIrish      string = "_irish_"
    StopWordsItalian    string = "_italian_"
    StopWordsLatvian    string = "_latvian_"
    StopWordsLithuanian string = "_lithuanian_"
    StopWordsNorwegian  string = "_norwegian_"
    StopWordsPersian    string = "_persian_"
    StopWordsPortuguese string = "_portuguese_"
    StopWordsRomanian   string = "_romanian_"
    StopWordsRussian    string = "_russian_"
    StopWordsSorani     string = "_sorani_"
    StopWordsSpanish    string = "_spanish_"
    StopWordsSwedish    string = "_swedish_"
    StopWordsThai       string = "_thai_"
    StopWordsTurkish    string = "_turkish_"

    SynonymFormatWordnet FilterSynonymFormat = "wordnet"
    SynonymFormatSolr    FilterSynonymFormat = "solr"

    WordDelimiterAlpha        WordDelimiter = "ALPHA"
    WordDelimiterAlphanum     WordDelimiter = "ALPHANUM"
    WordDelimiterDigit        WordDelimiter = "DIGIT"
    WordDelimiterLower        WordDelimiter = "LOWER"
    WordDelimiterSubwordDelim WordDelimiter = "SUBWORD_DELIM"
    WordDelimiterUpper        WordDelimiter = "UPPER"
)

type FilterType string
type FilterName string
type FilterBigramIgnoredScripts string
type FilterDelimitedPayloadEncoding string
type FilterEdgeNgramSide string
type FilterKeepTypesMode string
type FilterLowercaseLanguageType string
type FilterPhoneticEncoder string
type FilterPhoneticRuleType string
type FilterPhoneticNameType string
type FilterPhoneticLanguageSet string
type FilterSnowballLanguage string
type FilterStemmerLanguage string
type FilterSynonymFormat string
type WordDelimiter string

var FilterScriptConditionError = errors.New("script is empty")
var FilterFiltersConditionError = errors.New("filters is empty")

type Filter interface {
    Type() FilterType
    Name() FilterName
    Source() (interface{}, error)
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-asciifolding-tokenfilter.html
type FilterAsciiFolding struct {
    name             FilterName
    preserveOriginal *bool
}

func NewFilterAsciiFolding(name string) *FilterAsciiFolding {
    return &FilterAsciiFolding{name: FilterName(name)}
}

func (filter *FilterAsciiFolding) SetPreserveOriginal(preserveOriginal bool) *FilterAsciiFolding {
    filter.preserveOriginal = &preserveOriginal

    return filter
}

func (filter *FilterAsciiFolding) Type() FilterType {
    return FilterTypeAsciiFolding
}

func (filter *FilterAsciiFolding) Name() FilterName {
    return filter.name
}

func (filter *FilterAsciiFolding) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.preserveOriginal != nil {
        source["preserve_original"] = *filter.preserveOriginal
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-cjk-bigram-tokenfilter.html
type FilterCjkBigRam struct {
    name           FilterName
    outputUnigrams *bool
    ignoredScripts []FilterBigramIgnoredScripts
}

func NewFilterCjkBigRam(name string) *FilterCjkBigRam {
    return &FilterCjkBigRam{name: FilterName(name)}
}

func (filter *FilterCjkBigRam) SetOutputUnigrams(outputUnigrams bool) *FilterCjkBigRam {
    filter.outputUnigrams = &outputUnigrams

    return filter
}

func (filter *FilterCjkBigRam) AddIgnoredScripts(ignoredScripts ...FilterBigramIgnoredScripts) *FilterCjkBigRam {
    filter.ignoredScripts = append(filter.ignoredScripts, ignoredScripts...)

    return filter
}

func (filter *FilterCjkBigRam) Type() FilterType {
    return FilterTypeCjkBigram
}

func (filter *FilterCjkBigRam) Name() FilterName {
    return filter.name
}

func (filter *FilterCjkBigRam) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.outputUnigrams != nil {
        source["output_unigrams"] = *filter.outputUnigrams
    }

    if len(filter.ignoredScripts) > 0 {
        source["ignored_scripts"] = filter.ignoredScripts
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-common-grams-tokenfilter.html
type FilterCommonGrams struct {
    name            FilterName
    ignoreCase      *bool
    queryMode       *bool
    commonWordsPath *string
    commonWords     []string
}

func NewFilterCommonGrams(name string) *FilterCommonGrams {
    return &FilterCommonGrams{name: FilterName(name)}
}

func (filter *FilterCommonGrams) SetIgnoreCase(ignoreCase bool) *FilterCommonGrams {
    filter.ignoreCase = &ignoreCase

    return filter
}

func (filter *FilterCommonGrams) SetQueryMode(queryMode bool) *FilterCommonGrams {
    filter.queryMode = &queryMode

    return filter
}

func (filter *FilterCommonGrams) SetCommonWordsPath(commonWordsPath string) *FilterCommonGrams {
    filter.commonWordsPath = &commonWordsPath

    return filter
}

func (filter *FilterCommonGrams) AddCommonWords(commonWords ...string) *FilterCommonGrams {
    filter.commonWords = append(filter.commonWords, commonWords...)

    return filter
}

func (filter *FilterCommonGrams) Type() FilterType {
    return FilterTypeCommonGrams
}

func (filter *FilterCommonGrams) Name() FilterName {
    return filter.name
}

func (filter *FilterCommonGrams) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.ignoreCase != nil {
        source["ignore_case"] = *filter.ignoreCase
    }

    if filter.queryMode != nil {
        source["query_mode"] = *filter.queryMode
    }

    if filter.commonWordsPath != nil {
        source["common_words_path"] = *filter.commonWordsPath
    }

    if len(filter.commonWords) > 0 {
        source["common_words"] = filter.commonWords
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-condition-tokenfilter.html
type FilterCondition struct {
    name   FilterName
    script *elastic.Script
    filter []FilterName
}

func NewFilterCondition(name string) *FilterCondition {
    return &FilterCondition{name: FilterName(name)}
}

func (filter *FilterCondition) SetScript(script *elastic.Script) *FilterCondition {
    filter.script = script

    return filter
}

func (filter *FilterCondition) AddFilter(filters ...FilterName) *FilterCondition {
    filter.filter = append(filter.filter, filters...)

    return filter
}

func (filter *FilterCondition) Type() FilterType {
    return FilterTypeCondition
}

func (filter *FilterCondition) Name() FilterName {
    return filter.name
}

func (filter *FilterCondition) Source() (interface{}, error) {
    if filter.script == nil {
        return nil, FilterScriptConditionError
    }

    if len(filter.filter) == 0 {
        return nil, FilterFiltersConditionError
    }

    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.script != nil {
        script, err := filter.script.Source()
        if err != nil {
            return nil, err
        }

        source["script"] = script
    }

    if len(filter.filter) > 0 {
        source["filter"] = filter.filter
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-delimited-payload-tokenfilter.html
type FilterDelimitedPayload struct {
    name      FilterName
    delimiter *string
    encoding  *FilterDelimitedPayloadEncoding
}

func NewFilterDelimitedPayload(name string) *FilterDelimitedPayload {
	return &FilterDelimitedPayload{name: FilterName(name)}
}

func (filter *FilterDelimitedPayload) SetDelimiter(delimiter string) *FilterDelimitedPayload {
	filter.delimiter = &delimiter
	
	return filter
}

func (filter *FilterDelimitedPayload) SetEncoding(encoding FilterDelimitedPayloadEncoding) *FilterDelimitedPayload {
	filter.encoding = &encoding
	
	return filter
}

func (filter *FilterDelimitedPayload) Type() FilterType {
    return FilterTypeDelimitedPayload
}

func (filter *FilterDelimitedPayload) Name() FilterName {
    return filter.name
}

func (filter *FilterDelimitedPayload) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.delimiter != nil {
        source["delimiter"] = *filter.delimiter
    }

    if filter.encoding != nil {
        source["encoding"] = *filter.encoding
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-dict-decomp-tokenfilter.html
type FilterDictionaryDecompounder struct {
    name             FilterName
    wordListPath     *string
    maxSubwordSize   *uint32
    minSubwordSize   *uint32
    minWordSize      *uint32
    onlyLongestMatch *bool
    wordList         []string
}

func NewFilterDictionaryDecompounder(name string) *FilterDictionaryDecompounder {
	return &FilterDictionaryDecompounder{name: FilterName(name)}
}

func (filter *FilterDictionaryDecompounder) SetWordListPath(wordListPath string) *FilterDictionaryDecompounder {
	filter.wordListPath = &wordListPath

	return filter
}

func (filter *FilterDictionaryDecompounder) SetMaxSubwordSize(maxSubwordSize uint32) *FilterDictionaryDecompounder {
	filter.maxSubwordSize = &maxSubwordSize

	return filter
}

func (filter *FilterDictionaryDecompounder) SetMinSubwordSize(minSubwordSize uint32) *FilterDictionaryDecompounder {
	filter.minSubwordSize = &minSubwordSize

	return filter
}

func (filter *FilterDictionaryDecompounder) SetMinWordSize(minWordSize uint32) *FilterDictionaryDecompounder {
	filter.minWordSize = &minWordSize

	return filter
}

func (filter *FilterDictionaryDecompounder) SetOnlyLongestMatch(onlyLongestMatch bool) *FilterDictionaryDecompounder {
	filter.onlyLongestMatch = &onlyLongestMatch

	return filter
}

func (filter *FilterDictionaryDecompounder) AddWordList(wordList ...string) *FilterDictionaryDecompounder {
	filter.wordList = append(filter.wordList, wordList...)

	return filter
}

func (filter *FilterDictionaryDecompounder) Type() FilterType {
    return FilterTypeDictionaryDecompounder
}

func (filter *FilterDictionaryDecompounder) Name() FilterName {
    return filter.name
}

func (filter *FilterDictionaryDecompounder) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.wordListPath != nil {
        source["word_list_path"] = *filter.wordListPath
    }

    if filter.maxSubwordSize != nil {
        source["max_subword_size"] = *filter.maxSubwordSize
    }

    if filter.minSubwordSize != nil {
        source["min_subword_size"] = *filter.minSubwordSize
    }

    if filter.minWordSize != nil {
        source["min_word_size"] = *filter.minWordSize
    }

    if filter.onlyLongestMatch != nil {
        source["only_longest_match"] = *filter.onlyLongestMatch
    }

    if len(filter.wordList) > 0 {
        source["word_list"] = filter.wordList
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-edgengram-tokenfilter.html
type FilterEdgeNgram struct {
    name             FilterName
    maxGram          *uint32
    minGram          *uint32
    preserveOriginal *bool
    side             *FilterEdgeNgramSide
}

func NewFilterEdgeNgram(name string) *FilterEdgeNgram {
	return &FilterEdgeNgram{name: FilterName(name)}
}

func (filter *FilterEdgeNgram) SetMaxGram(maxGram uint32) *FilterEdgeNgram {
	filter.maxGram = &maxGram

	return filter
}

func (filter *FilterEdgeNgram) SetMinGram(minGram uint32) *FilterEdgeNgram {
	filter.minGram = &minGram

	return filter
}

func (filter *FilterEdgeNgram) SetPreserveOriginal(preserveOriginal bool) *FilterEdgeNgram {
	filter.preserveOriginal = &preserveOriginal

	return filter
}

func (filter *FilterEdgeNgram) SetSide(side FilterEdgeNgramSide) *FilterEdgeNgram {
	filter.side = &side

	return filter
}

func (filter *FilterEdgeNgram) Type() FilterType {
    return FilterTypeEdgeNgram
}

func (filter *FilterEdgeNgram) Name() FilterName {
    return filter.name
}

func (filter *FilterEdgeNgram) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.maxGram != nil {
        source["max_gram"] = *filter.maxGram
    }

    if filter.minGram != nil {
        source["min_gram"] = *filter.minGram
    }

    if filter.preserveOriginal != nil {
        source["preserve_original"] = *filter.preserveOriginal
    }

    if filter.side != nil {
        source["side"] = *filter.side
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-elision-tokenfilter.html
type FilterElision struct {
    name         FilterName
    articlesPath *string
    articlesCase *bool
    articles     []string
}

func NewFilterElision(name string) *FilterElision {
	return &FilterElision{name: FilterName(name)}
}

func (filter *FilterElision) SetArticlesPath(articlesPath string) *FilterElision {
	filter.articlesPath = &articlesPath

	return filter
}

func (filter *FilterElision) SetArticlesCase(articlesCase bool) *FilterElision {
	filter.articlesCase = &articlesCase

	return filter
}

func (filter *FilterElision) AddArticles(articles ...string) *FilterElision {
	filter.articles = append(filter.articles, articles...)

	return filter
}

func (filter *FilterElision) Type() FilterType {
    return FilterTypeElision
}

func (filter *FilterElision) Name() FilterName {
    return filter.name
}

func (filter *FilterElision) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.articlesPath != nil {
        source["articles_path"] = *filter.articlesPath
    }

    if filter.articlesCase != nil {
        source["articles_case"] = *filter.articlesCase
    }

    if len(filter.articles) > 0 {
        source["articles"] = filter.articles
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-fingerprint-tokenfilter.html
type FilterFingerprint struct {
    name          FilterName
    maxOutputSize *uint32
    separator     *string
}

func NewFilterFingerprint(name string) *FilterFingerprint {
	return &FilterFingerprint{name: FilterName(name)}
}

func (filter *FilterFingerprint) SetMaxOutputSize(maxOutputSize uint32) *FilterFingerprint {
	filter.maxOutputSize = &maxOutputSize

	return filter
}

func (filter *FilterFingerprint) SetSeparator(separator string) *FilterFingerprint {
	filter.separator = &separator

	return filter
}

func (filter *FilterFingerprint) Type() FilterType {
    return FilterTypeFingerprint
}

func (filter *FilterFingerprint) Name() FilterName {
    return filter.name
}

func (filter *FilterFingerprint) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.maxOutputSize != nil {
        source["max_output_size"] = *filter.maxOutputSize
    }

    if filter.separator != nil {
        source["separator"] = *filter.separator
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-hyp-decomp-tokenfilter.html
type FilterHyphenationDecompounder struct {
    name                    FilterName
    hyphenationPatternsPath *string
    wordListPath            *string
    maxSubwordSize          *uint32
    minSubwordSize          *uint32
    minWordSize             *uint32
    onlyLongestMatch        *bool
    wordList                []string
}

func NewFilterHyphenationDecompounder(name string) *FilterHyphenationDecompounder {
	return &FilterHyphenationDecompounder{name: FilterName(name)}
}

func (filter *FilterHyphenationDecompounder) SetHyphenationPatternsPath(hyphenationPatternsPath string) *FilterHyphenationDecompounder {
	filter.hyphenationPatternsPath = &hyphenationPatternsPath

	return filter
}

func (filter *FilterHyphenationDecompounder) SetWordListPath(wordListPath string) *FilterHyphenationDecompounder {
	filter.wordListPath = &wordListPath

	return filter
}

func (filter *FilterHyphenationDecompounder) SetMaxSubwordSize(maxSubwordSize uint32) *FilterHyphenationDecompounder {
	filter.maxSubwordSize = &maxSubwordSize

	return filter
}

func (filter *FilterHyphenationDecompounder) SetMinSubwordSize(minSubwordSize uint32) *FilterHyphenationDecompounder {
	filter.minSubwordSize = &minSubwordSize

	return filter
}

func (filter *FilterHyphenationDecompounder) SetMinWordSize(minWordSize uint32) *FilterHyphenationDecompounder {
	filter.minWordSize = &minWordSize

	return filter
}

func (filter *FilterHyphenationDecompounder) SetOnlyLongestMatch(onlyLongestMatch bool) *FilterHyphenationDecompounder {
	filter.onlyLongestMatch = &onlyLongestMatch

	return filter
}

func (filter *FilterHyphenationDecompounder) AddWordList(wordList string) *FilterHyphenationDecompounder {
	filter.wordList = append(filter.wordList, wordList)

	return filter
}

func (filter *FilterHyphenationDecompounder) Type() FilterType {
    return FilterTypeHyphenationDecompounder
}

func (filter *FilterHyphenationDecompounder) Name() FilterName {
    return filter.name
}

func (filter *FilterHyphenationDecompounder) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.hyphenationPatternsPath != nil {
        source["hyphenation_patterns_path"] = *filter.hyphenationPatternsPath
    }

    if filter.wordListPath != nil {
        source["word_list_path"] = *filter.wordListPath
    }

    if filter.maxSubwordSize != nil {
        source["max_subword_size"] = *filter.maxSubwordSize
    }

    if filter.minSubwordSize != nil {
        source["min_subword_size"] = *filter.minSubwordSize
    }

    if filter.minWordSize != nil {
        source["min_word_size"] = *filter.minWordSize
    }

    if filter.onlyLongestMatch != nil {
        source["only_longest_match"] = *filter.onlyLongestMatch
    }

    if len(filter.wordList) > 0 {
        source["word_list"] = filter.wordList
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-keep-types-tokenfilter.html
type FilterKeepTypes struct {
    name  FilterName
    mode  *FilterKeepTypesMode
    types []string
}

func NewFilterKeepTypes(name string) *FilterKeepTypes {
	return &FilterKeepTypes{name: FilterName(name)}
}

func (filter *FilterKeepTypes) SetMode(mode FilterKeepTypesMode) *FilterKeepTypes {
	filter.mode = &mode

	return filter
}

func (filter *FilterKeepTypes) AddTypes(types ...string) *FilterKeepTypes {
	filter.types = append(filter.types, types...)

	return filter
}

func (filter *FilterKeepTypes) Type() FilterType {
    return FilterTypeKeepTypes
}

func (filter *FilterKeepTypes) Name() FilterName {
    return filter.name
}

func (filter *FilterKeepTypes) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.mode != nil {
        source["mode"] = *filter.mode
    }

    if len(filter.types) > 0 {
        source["types"] = filter.types
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-keep-words-tokenfilter.html
type FilterKeep struct {
    name          FilterName
    keepWordsCase *bool
    keepWordsPath *string
    keepWords     []string
}

func NewFilterKeep(name string) *FilterKeep {
	return &FilterKeep{name: FilterName(name)}
}

func (filter *FilterKeep) SetKeepWordsCase(keepWordsCase bool) *FilterKeep {
	filter.keepWordsCase = &keepWordsCase

	return filter
}

func (filter *FilterKeep) SetKeepWordsPath(keepWordsPath string) *FilterKeep {
	filter.keepWordsPath = &keepWordsPath

	return filter
}

func (filter *FilterKeep) AddKeepWords(keepWords ...string) *FilterKeep {
	filter.keepWords = append(filter.keepWords, keepWords...)

	return filter
}

func (filter *FilterKeep) Type() FilterType {
    return FilterTypeKeep
}

func (filter *FilterKeep) Name() FilterName {
    return filter.name
}

func (filter *FilterKeep) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.keepWordsCase != nil {
        source["keep_words_case"] = *filter.keepWordsCase
    }

    if filter.keepWordsPath != nil {
        source["keep_words_path"] = *filter.keepWordsPath
    }

    if len(filter.keepWords) > 0 {
        source["keep_words"] = filter.keepWords
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-keyword-marker-tokenfilter.html
type FilterKeywordMarker struct {
    name            FilterName
    ignoreCase      *bool
    keywordsPath    *string
    keywordsPattern *string
    keywords        []string
}

func NewFilterKeywordMarker(name string) *FilterKeywordMarker {
	return &FilterKeywordMarker{name: FilterName(name)}
}

func (filter *FilterKeywordMarker) SetIgnoreCase(ignoreCase bool) *FilterKeywordMarker {
	filter.ignoreCase = &ignoreCase

	return filter
}

func (filter *FilterKeywordMarker) SetKeywordsPath(keywordsPath string) *FilterKeywordMarker {
	filter.keywordsPath = &keywordsPath

	return filter
}

func (filter *FilterKeywordMarker) SetKeywordsPattern(keywordsPattern string) *FilterKeywordMarker {
	filter.keywordsPattern = &keywordsPattern

	return filter
}

func (filter *FilterKeywordMarker) AddKeywords(keywords ...string) *FilterKeywordMarker {
	filter.keywords = append(filter.keywords, keywords...)

	return filter
}

func (filter *FilterKeywordMarker) Type() FilterType {
    return FilterTypeKeywordMarker
}

func (filter *FilterKeywordMarker) Name() FilterName {
    return filter.name
}

func (filter *FilterKeywordMarker) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.ignoreCase != nil {
        source["ignore_case"] = *filter.ignoreCase
    }

    if filter.keywordsPath != nil {
        source["keywords_path"] = *filter.keywordsPath
    }

    if filter.keywordsPattern != nil {
        source["keywords_pattern"] = *filter.keywordsPattern
    }

    if len(filter.keywords) > 0 {
        source["keywords"] = filter.keywords
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-length-tokenfilter.html
type FilterLength struct {
    name FilterName
    min  *uint32
    max  *uint32
}

func NewFilterLength(name string) *FilterLength {
	return &FilterLength{name: FilterName(name)}
}

func (filter *FilterLength) SetMin(min uint32) *FilterLength {
	filter.min = &min

	return filter
}

func (filter *FilterLength) SetMax(max uint32) *FilterLength {
	filter.max = &max

	return filter
}

func (filter *FilterLength) Type() FilterType {
    return FilterTypeLength
}

func (filter *FilterLength) Name() FilterName {
    return filter.name
}

func (filter *FilterLength) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.min != nil {
        source["min"] = *filter.min
    }

    if filter.max != nil {
        source["max"] = *filter.max
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-limit-token-count-tokenfilter.html
type FilterLimit struct {
    name             FilterName
    maxTokenCount    *uint32
    consumeAllTokens *bool
}

func NewFilterLimit(name string) *FilterLimit {
	return &FilterLimit{name: FilterName(name)}
}

func (filter *FilterLimit) SetMaxTokenCount(maxTokenCount uint32) *FilterLimit {
	filter.maxTokenCount = &maxTokenCount

	return filter
}

func (filter *FilterLimit) SetConsumeAllTokens(consumeAllTokens bool) *FilterLimit {
	filter.consumeAllTokens = &consumeAllTokens

	return filter
}

func (filter *FilterLimit) Type() FilterType {
    return FilterTypeLimit
}

func (filter *FilterLimit) Name() FilterName {
    return filter.name
}

func (filter *FilterLimit) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.maxTokenCount != nil {
        source["max_token_count"] = *filter.maxTokenCount
    }

    if filter.consumeAllTokens != nil {
        source["consume_all_tokens"] = *filter.consumeAllTokens
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-lowercase-tokenfilter.html
type FilterLowercase struct {
    name     FilterName
    language *FilterLowercaseLanguageType
}

func NewFilterLowercase(name string) *FilterLowercase {
	return &FilterLowercase{name: FilterName(name)}
}

func (filter *FilterLowercase) SetLanguage(language FilterLowercaseLanguageType) *FilterLowercase {
	filter.language = &language

	return filter
}

func (filter *FilterLowercase) Type() FilterType {
    return FilterTypeLowercase
}

func (filter *FilterLowercase) Name() FilterName {
    return filter.name
}

func (filter *FilterLowercase) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.language != nil {
        source["language"] = *filter.language
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-minhash-tokenfilter.html
type FilterMinHash struct {
    name         FilterName
    bucketCount  *uint32
    hashCount    *uint32
    hashSetSize  *uint32
    withRotation *bool
}

func NewFilterMinHash(name string) *FilterMinHash {
	return &FilterMinHash{name: FilterName(name)}
}

func (filter *FilterMinHash) SetBucketCount(bucketCount uint32) *FilterMinHash {
	filter.bucketCount = &bucketCount

	return filter
}

func (filter *FilterMinHash) SetHashCount(hashCount uint32) *FilterMinHash {
	filter.hashCount = &hashCount

	return filter
}

func (filter *FilterMinHash) SetHashSetSize(hashSetSize uint32) *FilterMinHash {
	filter.hashSetSize = &hashSetSize

	return filter
}

func (filter *FilterMinHash) SetWithRotation(withRotation bool) *FilterMinHash {
	filter.withRotation = &withRotation

	return filter
}

func (filter *FilterMinHash) Type() FilterType {
    return FilterTypeMinHash
}

func (filter *FilterMinHash) Name() FilterName {
    return filter.name
}

func (filter *FilterMinHash) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.bucketCount != nil {
        source["bucket_count"] = *filter.bucketCount
    }

    if filter.hashCount != nil {
        source["hash_count"] = *filter.hashCount
    }

    if filter.hashSetSize != nil {
        source["hash_set_size"] = *filter.hashSetSize
    }

    if filter.withRotation != nil {
        source["with_rotation"] = *filter.withRotation
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-multiplexer-tokenfilter.html
type FilterMultiplexer struct {
    name    FilterName
    filters []FilterName
}

func NewFilterMultiplexer(name string) *FilterMultiplexer {
	return &FilterMultiplexer{name: FilterName(name)}
}

func (filter *FilterMultiplexer) AddFilters(filters ...FilterName) *FilterMultiplexer {
	filter.filters = append(filter.filters, filters...)

	return filter
}

func (filter *FilterMultiplexer) Type() FilterType {
    return FilterTypeMultiplexer
}

func (filter *FilterMultiplexer) Name() FilterName {
    return filter.name
}

func (filter *FilterMultiplexer) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if len(filter.filters) > 0 {
        source["filters"] = filter.filters
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-ngram-tokenfilter.html
type FilterNgram struct {
    name             FilterName
    maxGram          *uint8
    minGram          *uint8
    preserveOriginal *bool
}

func NewFilterNgram(name string) *FilterNgram {
	return &FilterNgram{name: FilterName(name)}
}

func (filter *FilterNgram) SetMaxGram(maxGram uint8) *FilterNgram {
	filter.maxGram = &maxGram

	return filter
}

func (filter *FilterNgram) SetMinGram(minGram uint8) *FilterNgram {
	filter.minGram = &minGram

	return filter
}

func (filter *FilterNgram) SetPreserveOriginal(preserveOriginal bool) *FilterNgram {
	filter.preserveOriginal = &preserveOriginal

	return filter
}

func (filter *FilterNgram) Type() FilterType {
    return FilterTypeNgram
}

func (filter *FilterNgram) Name() FilterName {
    return filter.name
}

func (filter *FilterNgram) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.maxGram != nil {
        source["max_gram"] = *filter.maxGram
    }

    if filter.minGram != nil {
        source["min_gram"] = *filter.minGram
    }

    if filter.preserveOriginal != nil {
        source["preserve_original"] = *filter.preserveOriginal
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-pattern-capture-tokenfilter.html
type FilterPatternCapture struct {
    name             FilterName
    preserveOriginal *bool
    patterns         []string
}

func NewFilterPatternCapture(name string) *FilterPatternCapture {
	return &FilterPatternCapture{name: FilterName(name)}
}

func (filter *FilterPatternCapture) SetPreserveOriginal(preserveOriginal bool) *FilterPatternCapture {
	filter.preserveOriginal = &preserveOriginal

	return filter
}

func (filter *FilterPatternCapture) AddPatterns(patterns ...string) *FilterPatternCapture {
	filter.patterns = append(filter.patterns, patterns...)

	return filter
}

func (filter *FilterPatternCapture) Type() FilterType {
    return FilterTypePatternCapture
}

func (filter *FilterPatternCapture) Name() FilterName {
    return filter.name
}

func (filter *FilterPatternCapture) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.preserveOriginal != nil {
        source["preserve_original"] = *filter.preserveOriginal
    }

    if len(filter.patterns) > 0 {
        source["patterns"] = filter.patterns
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-pattern_replace-tokenfilter.html
type FilterPatternReplace struct {
    name        FilterName
    all         *bool
    pattern     *string
    replacement *string
}

func NewFilterPatternReplace(name string) *FilterPatternReplace {
	return &FilterPatternReplace{name: FilterName(name)}
}

func (filter *FilterPatternReplace) SetAll(all bool) *FilterPatternReplace {
	filter.all = &all

	return filter
}

func (filter *FilterPatternReplace) SetPattern(pattern string) *FilterPatternReplace {
	filter.pattern = &pattern

	return filter
}

func (filter *FilterPatternReplace) SetReplacement(replacement string) *FilterPatternReplace {
	filter.replacement = &replacement

	return filter
}

func (filter *FilterPatternReplace) Type() FilterType {
    return FilterTypePatternReplace
}

func (filter *FilterPatternReplace) Name() FilterName {
    return filter.name
}

func (filter *FilterPatternReplace) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.all != nil {
        source["all"] = *filter.all
    }

    if filter.pattern != nil {
        source["pattern"] = *filter.pattern
    }

    if filter.replacement != nil {
        source["replacement"] = *filter.replacement
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/plugins/7.10/analysis-phonetic-token-filter.html
type FilterPhonetic struct {
    name        FilterName
    encoder     *FilterPhoneticEncoder
    ruleType    *FilterPhoneticRuleType
    nameType    *FilterPhoneticNameType
    languageSet *FilterPhoneticLanguageSet
    replace     *bool
    maxCodeLen  *uint32
}

func NewFilterPhonetic(name string) *FilterPhonetic {
	return &FilterPhonetic{name: FilterName(name)}
}

func (filter *FilterPhonetic) SetEncoder(encoder FilterPhoneticEncoder) *FilterPhonetic {
	filter.encoder = &encoder

	return filter
}

func (filter *FilterPhonetic) SetRuleType(ruleType FilterPhoneticRuleType) *FilterPhonetic {
	filter.ruleType = &ruleType

	return filter
}

func (filter *FilterPhonetic) SetNameType(nameType FilterPhoneticNameType) *FilterPhonetic {
	filter.nameType = &nameType

	return filter
}

func (filter *FilterPhonetic) SetLanguageSet(languageSet FilterPhoneticLanguageSet) *FilterPhonetic {
	filter.languageSet = &languageSet

	return filter
}

func (filter *FilterPhonetic) SetReplace(replace bool) *FilterPhonetic {
	filter.replace = &replace

	return filter
}

func (filter *FilterPhonetic) SetMaxCodeLen(maxCodeLen uint32) *FilterPhonetic {
	filter.maxCodeLen = &maxCodeLen

	return filter
}

func (filter *FilterPhonetic) Type() FilterType {
    return FilterTypePhonetic
}

func (filter *FilterPhonetic) Name() FilterName {
    return filter.name
}

func (filter *FilterPhonetic) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.encoder != nil {
        source["encoder"] = *filter.encoder
    }

    if filter.ruleType != nil {
        source["rule_type"] = *filter.ruleType
    }

    if filter.nameType != nil {
        source["name_type"] = *filter.nameType
    }

    if filter.languageSet != nil {
        source["language_set"] = *filter.languageSet
    }

    if filter.replace != nil {
        source["replace"] = *filter.replace
    }

    if filter.maxCodeLen != nil {
        source["max_code_len"] = *filter.maxCodeLen
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-predicatefilter-tokenfilter.html
type FilterPredicateTokenFilter struct {
    name   FilterName
    script *elastic.Script
}

func NewFilterPredicateTokenFilter(name string) *FilterPredicateTokenFilter {
	return &FilterPredicateTokenFilter{name: FilterName(name)}
}

func (filter *FilterPredicateTokenFilter) SetScript(script *elastic.Script) *FilterPredicateTokenFilter {
	filter.script = script

	return filter
}

func (filter *FilterPredicateTokenFilter) Type() FilterType {
    return FilterTypePredicateTokenFilter
}

func (filter *FilterPredicateTokenFilter) Name() FilterName {
    return filter.name
}

func (filter *FilterPredicateTokenFilter) Source() (interface{}, error) {
    if filter.script == nil {
        return nil, FilterScriptConditionError
    }

    source := map[string]interface{}{
        "type": filter.Type(),
    }

    script, err := filter.script.Source()
    if err != nil {
        return nil, err
    }

    source["script"] = script

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-shingle-tokenfilter.html
type FilterShingle struct {
    name                       FilterName
    maxShingleSize             *uint32
    minShingleSize             *uint32
    outputUnigrams             *bool
    outputUnigramsIfNoShingles *bool
    tokenSeparator             *string
    fillerToken                *string
}

func NewFilterShingle(name string) *FilterShingle {
	return &FilterShingle{name: FilterName(name)}
}

func (filter *FilterShingle) SetMaxShingleSize(maxShingleSize uint32) *FilterShingle {
	filter.maxShingleSize = &maxShingleSize

	return filter
}

func (filter *FilterShingle) SetMinShingleSize(minShingleSize uint32) *FilterShingle {
	filter.minShingleSize = &minShingleSize

	return filter
}

func (filter *FilterShingle) SetOutputUnigrams(outputUnigrams bool) *FilterShingle {
	filter.outputUnigrams = &outputUnigrams

	return filter
}

func (filter *FilterShingle) SetOutputUnigramsIfNoShingles(outputUnigramsIfNoShingles bool) *FilterShingle {
	filter.outputUnigramsIfNoShingles = &outputUnigramsIfNoShingles

	return filter
}

func (filter *FilterShingle) SetTokenSeparator(tokenSeparator string) *FilterShingle {
	filter.tokenSeparator = &tokenSeparator

	return filter
}

func (filter *FilterShingle) SetFillerToken(fillerToken string) *FilterShingle {
	filter.fillerToken = &fillerToken

	return filter
}

func (filter *FilterShingle) Type() FilterType {
    return FilterTypeShingle
}

func (filter *FilterShingle) Name() FilterName {
    return filter.name
}

func (filter *FilterShingle) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.maxShingleSize != nil {
        source["max_shingle_size"] = *filter.maxShingleSize
    }

    if filter.minShingleSize != nil {
        source["min_shingle_size"] = *filter.minShingleSize
    }

    if filter.outputUnigrams != nil {
        source["output_unigrams"] = *filter.outputUnigrams
    }

    if filter.outputUnigramsIfNoShingles != nil {
        source["output_unigrams_if_no_shingles"] = *filter.outputUnigramsIfNoShingles
    }

    if filter.tokenSeparator != nil {
        source["token_separator"] = *filter.tokenSeparator
    }

    if filter.fillerToken != nil {
        source["filler_token"] = *filter.fillerToken
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-snowball-tokenfilter.html
type FilterSnowball struct {
    name     FilterName
    language *FilterSnowballLanguage
}

func NewFilterSnowball(name string) *FilterSnowball {
	return &FilterSnowball{name: FilterName(name)}
}

func (filter *FilterSnowball) SetLanguage(language FilterSnowballLanguage) *FilterSnowball {
	filter.language = &language

	return filter
}

func (filter *FilterSnowball) Type() FilterType {
    return FilterTypeSnowball
}

func (filter *FilterSnowball) Name() FilterName {
    return filter.name
}

func (filter *FilterSnowball) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.language != nil {
        source["language"] = *filter.language
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-stemmer-tokenfilter.html
type FilterStemmer struct {
    name     FilterName
    language *FilterStemmerLanguage
}

func NewFilterStemmer(name string) *FilterStemmer {
    return &FilterStemmer{name: FilterName(name)}
}

func (filter *FilterStemmer) SetLanguage(language FilterStemmerLanguage) *FilterStemmer {
	filter.language = &language

	return filter
}

func (filter *FilterStemmer) Type() FilterType {
    return FilterTypeStemmer
}

func (filter *FilterStemmer) Name() FilterName {
    return filter.name
}

func (filter *FilterStemmer) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.language != nil {
        source["language"] = *filter.language
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-stemmer-override-tokenfilter.html
type FilterStemmerOverride struct {
    name      FilterName
    rulesPath *string
    rules     []*StemmerOverrideRule
}

func NewFilterStemmerOverride(name string) *FilterStemmerOverride {
	return &FilterStemmerOverride{name: FilterName(name)}
}

func (filter *FilterStemmerOverride) SetRulesPath(rulesPath string) *FilterStemmerOverride {
	filter.rulesPath = &rulesPath

	return filter
}

func (filter *FilterStemmerOverride) AddRules(rules ...*StemmerOverrideRule) *FilterStemmerOverride {
	filter.rules = append(filter.rules, rules...)

	return filter
}

func (filter *FilterStemmerOverride) Type() FilterType {
    return FilterTypeStemmerOverride
}

func (filter *FilterStemmerOverride) Name() FilterName {
    return filter.name
}

func (filter *FilterStemmerOverride) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.rulesPath != nil {
        source["rules_path"] = *filter.rulesPath
    }

    if len(filter.rules) > 0 {
        var rulesSource []string

        for _, rule := range filter.rules {
            rulesSource = append(rulesSource, rule.String())
        }

        source["rules"] = rulesSource
    }

    return source, nil
}

type StemmerOverrideRule struct {
    stemmer string
    worlds  []string
}

func NewStemmerOverrideRule(stemmer string, worlds ...string) *StemmerOverrideRule {
    return &StemmerOverrideRule{stemmer: stemmer, worlds: worlds}
}

func (rule *StemmerOverrideRule) String() string {
    return fmt.Sprintf("%s => %s", strings.Join(rule.worlds, ","), rule.stemmer)
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-stop-tokenfilter.html
type FilterStop struct {
    name           FilterName
    stopWordsPath  *string
    ignoreCase     *bool
    removeTrailing *bool
    stopWords      []string
}

func NewFilterStop(name string) *FilterStop {
	return &FilterStop{name: FilterName(name)}
}

func (filter *FilterStop) SetStopWordsPath(stopWordsPath string) *FilterStop {
	filter.stopWordsPath = &stopWordsPath

	return filter
}

func (filter *FilterStop) SetIgnoreCase(ignoreCase bool) *FilterStop {
	filter.ignoreCase = &ignoreCase

	return filter
}

func (filter *FilterStop) SetRemoveTrailing(removeTrailing bool) *FilterStop {
	filter.removeTrailing = &removeTrailing

	return filter
}

func (filter *FilterStop) AddStopWords(stopWords ...string) *FilterStop {
	filter.stopWords = append(filter.stopWords, stopWords...)

	return filter
}

func (filter *FilterStop) Type() FilterType {
    return FilterTypeStop
}

func (filter *FilterStop) Name() FilterName {
    return filter.name
}

func (filter *FilterStop) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.stopWordsPath != nil {
        source["stop_words_path"] = *filter.stopWordsPath
    }

    if filter.ignoreCase != nil {
        source["ignore_case"] = *filter.ignoreCase
    }

    if filter.removeTrailing != nil {
        source["remove_trailing"] = *filter.removeTrailing
    }

    if len(filter.stopWords) > 0 {
        source["stop_words"] = filter.stopWords
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-synonym-tokenfilter.html
// or https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-synonym-graph-tokenfilter.html
type FilterSynonym struct {
    _type        FilterType
    name         FilterName
    format       FilterSynonymFormat
    synonymsPath *string
    expand       *bool
    lenient      *bool
    synonyms     []Synonym
}

func NewFilterSynonym(name string) *FilterSynonym {
	return &FilterSynonym{name: FilterName(name), _type: FilterTypeSynonym}
}

func NewFilterSynonymGraph(name string) *FilterSynonym {
	return &FilterSynonym{name: FilterName(name), _type: FilterTypeSynonymGraph}
}

func (filter *FilterSynonym) SetFormat(format FilterSynonymFormat) *FilterSynonym {
	filter.format = format

	return filter
}

func (filter *FilterSynonym) SetSynonymsPath(synonymsPath string) *FilterSynonym {
	filter.synonymsPath = &synonymsPath

	return filter
}

func (filter *FilterSynonym) SetExpand(expand bool) *FilterSynonym {
	filter.expand = &expand

	return filter
}

func (filter *FilterSynonym) SetLenient(lenient bool) *FilterSynonym {
	filter.lenient = &lenient

	return filter
}

func (filter *FilterSynonym) AddSynonyms(synonyms ...Synonym) *FilterSynonym {
	filter.synonyms = append(filter.synonyms, synonyms...)

	return filter
}

func (filter *FilterSynonym) Type() FilterType {
    return filter._type
}

func (filter *FilterSynonym) Name() FilterName {
    return filter.name
}

func (filter *FilterSynonym) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.format == SynonymFormatWordnet {
        source["format"] = SynonymFormatWordnet
    }

    if filter.synonymsPath != nil {
        source["synonyms_path"] = *filter.synonymsPath
    }

    if filter.expand != nil {
        source["expand"] = *filter.expand
    }

    if filter.lenient != nil {
        source["lenient"] = *filter.lenient
    }

    if len(filter.synonyms) > 0 {
        var synonyms []string

        for _, synonym := range filter.synonyms {
            synonyms = append(synonyms, synonym.Synonym())
        }

        source["synonyms"] = synonyms
    }

    return source, nil
}

type Synonym interface {
    Synonym() string
}

type SynonymSolr struct {
    synonyms []string
    words    []string
}

func NewSynonymSolr() *SynonymSolr {
	return &SynonymSolr{}
}

func (synonym *SynonymSolr) AddSynonyms(synonyms ...string) *SynonymSolr {
	synonym.synonyms = append(synonym.synonyms, synonyms...)
	
	return synonym
}

func (synonym *SynonymSolr) AddWords(words ...string) *SynonymSolr {
	synonym.words = append(synonym.words, words...)
	
	return synonym
}

func (synonym *SynonymSolr) Synonym() string {
    return fmt.Sprintf("%s => %s", strings.Join(synonym.words, ","), strings.Join(synonym.synonyms, ","))
}

type SynonymWordnet struct {
    synsetId    uint32
    wordNumber  uint32
    word        string
    synsetType  string
    senseNumber uint32
    tagCount    uint32
}

func NewSynonymWordnet() *SynonymWordnet {
	return &SynonymWordnet{}
}

func (synonym *SynonymWordnet) SetSynsetId(synsetId uint32) *SynonymWordnet {
	synonym.synsetId = synsetId
	
	return synonym
}

func (synonym *SynonymWordnet) SetWordNumber(wordNumber uint32) *SynonymWordnet {
	synonym.wordNumber = wordNumber
	
	return synonym
}

func (synonym *SynonymWordnet) SetWord(word string) *SynonymWordnet {
	synonym.word = word

	return synonym
}

func (synonym *SynonymWordnet) SetSynsetType(synsetType string) *SynonymWordnet {
	synonym.synsetType = synsetType

	return synonym
}

func (synonym *SynonymWordnet) SetSenseNumber(senseNumber uint32) *SynonymWordnet {
	synonym.senseNumber = senseNumber

	return synonym
}

func (synonym *SynonymWordnet) SetTagCount(tagCount uint32) *SynonymWordnet {
	synonym.tagCount = tagCount

	return synonym
}

func (synonym *SynonymWordnet) Synonym() string {
    return fmt.Sprintf(
        "s(%d,%d,'%s',%s,%d,%d).",
        synonym.synsetId,
        synonym.wordNumber,
        synonym.word,
        synonym.synsetType,
        synonym.senseNumber,
        synonym.tagCount,
    )
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-truncate-tokenfilter.html
type FilterTruncate struct {
    name   FilterName
    length *uint32
}

func NewFilterTruncate(name string) *FilterTruncate {
    return &FilterTruncate{name: FilterName(name)}
}

func (filter *FilterTruncate) SetLength(length uint32) *FilterTruncate {
    filter.length = &length

    return filter
}

func (filter *FilterTruncate) Type() FilterType {
    return FilterTypeTruncate
}

func (filter *FilterTruncate) Name() FilterName {
    return filter.name
}

func (filter *FilterTruncate) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.length != nil {
        source["length"] = *filter.length
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-unique-tokenfilter.html
type FilterUnique struct {
    name               FilterName
    onlyOnSamePosition *bool
}

func NewFilterUnique(name string) *FilterUnique {
	return &FilterUnique{name: FilterName(name)}
}

func (filter *FilterUnique) SetOnlyOnSamePosition(onlyOnSamePosition bool) *FilterUnique {
	filter.onlyOnSamePosition = &onlyOnSamePosition

	return filter
}

func (filter *FilterUnique) Type() FilterType {
    return FilterTypeUnique
}

func (filter *FilterUnique) Name() FilterName {
    return filter.name
}

func (filter *FilterUnique) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.onlyOnSamePosition != nil {
        source["only_on_same_position"] = *filter.onlyOnSamePosition
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-word-delimiter-tokenfilter.html
type FilterWordDelimiter struct {
    name                  FilterName
    catenateAll           *bool
    catenateNumbers       *bool
    catenateWords         *bool
    generateNumberParts   *bool
    generateWordParts     *bool
    preserveOriginal      *bool
    splitOnCaseChange     *bool
    splitOnNumerics       *bool
    stemEnglishPossessive *bool
    protectedWordsPath    *string
    typeTablePath         *string
    protectedWords        []string
    typeTable             []*WordDelimiterType
}

func NewFilterWordDelimiter(name string) *FilterWordDelimiter {
	return &FilterWordDelimiter{name: FilterName(name)}
}

func (filter *FilterWordDelimiter) SetCatenateAll(catenateAll bool) *FilterWordDelimiter {
	filter.catenateAll = &catenateAll

	return filter
}

func (filter *FilterWordDelimiter) SetCatenateNumbers(catenateNumbers bool) *FilterWordDelimiter {
	filter.catenateNumbers = &catenateNumbers

	return filter
}

func (filter *FilterWordDelimiter) SetCatenateWords(catenateWords bool) *FilterWordDelimiter {
	filter.catenateWords = &catenateWords

	return filter
}

func (filter *FilterWordDelimiter) SetGenerateNumberParts(generateNumberParts bool) *FilterWordDelimiter {
	filter.generateNumberParts = &generateNumberParts

	return filter
}

func (filter *FilterWordDelimiter) SetGenerateWordParts(generateWordParts bool) *FilterWordDelimiter {
	filter.generateWordParts = &generateWordParts

	return filter
}

func (filter *FilterWordDelimiter) SetPreserveOriginal(preserveOriginal bool) *FilterWordDelimiter {
	filter.preserveOriginal = &preserveOriginal

	return filter
}

func (filter *FilterWordDelimiter) SetSplitOnCaseChange(splitOnCaseChange bool) *FilterWordDelimiter {
	filter.splitOnCaseChange = &splitOnCaseChange

	return filter
}

func (filter *FilterWordDelimiter) SetSplitOnNumerics(splitOnNumerics bool) *FilterWordDelimiter {
	filter.splitOnNumerics = &splitOnNumerics

	return filter
}

func (filter *FilterWordDelimiter) SetStemEnglishPossessive(stemEnglishPossessive bool) *FilterWordDelimiter {
	filter.stemEnglishPossessive = &stemEnglishPossessive

	return filter
}

func (filter *FilterWordDelimiter) SetProtectedWordsPath(protectedWordsPath string) *FilterWordDelimiter {
	filter.protectedWordsPath = &protectedWordsPath

	return filter
}

func (filter *FilterWordDelimiter) SetTypeTablePath(typeTablePath string) *FilterWordDelimiter {
	filter.typeTablePath = &typeTablePath

	return filter
}

func (filter *FilterWordDelimiter) AddProtectedWords(protectedWords ...string) *FilterWordDelimiter {
	filter.protectedWords = append(filter.protectedWords, protectedWords...)

	return filter
}

func (filter *FilterWordDelimiter) AddTypeTable(typeTable ...*WordDelimiterType) *FilterWordDelimiter {
	filter.typeTable = append(filter.typeTable, typeTable...)

	return filter
}

func (filter *FilterWordDelimiter) Type() FilterType {
    return FilterTypeWordDelimiter
}

func (filter *FilterWordDelimiter) Name() FilterName {
    return filter.name
}

func (filter *FilterWordDelimiter) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": filter.Type(),
    }

    if filter.catenateAll != nil {
        source["catenate_all"] = *filter.catenateAll
    }

    if filter.catenateNumbers != nil {
        source["catenate_numbers"] = *filter.catenateNumbers
    }

    if filter.catenateWords != nil {
        source["catenate_words"] = *filter.catenateWords
    }

    if filter.generateNumberParts != nil {
        source["generate_number_parts"] = *filter.generateNumberParts
    }

    if filter.generateWordParts != nil {
        source["generate_word_parts"] = *filter.generateWordParts
    }

    if filter.preserveOriginal != nil {
        source["preserve_original"] = *filter.preserveOriginal
    }

    if filter.splitOnCaseChange != nil {
        source["split_on_case_change"] = *filter.splitOnCaseChange
    }

    if filter.splitOnNumerics != nil {
        source["split_on_numerics"] = *filter.splitOnNumerics
    }

    if filter.stemEnglishPossessive != nil {
        source["stem_english_possessive"] = *filter.stemEnglishPossessive
    }

    if filter.protectedWordsPath != nil {
        source["protected_words_path"] = *filter.protectedWordsPath
    }

    if filter.typeTablePath != nil {
        source["type_table_path"] = *filter.typeTablePath
    }

    if len(filter.protectedWords) > 0 {
        source["protected_words"] = filter.protectedWords
    }

    if len(filter.typeTable) > 0 {
        typeTable := make([]string, len(filter.typeTable))

        for index, item := range filter.typeTable {
            typeTable[index] = item.String()
        }

        source["type_table"] = typeTable
    }

    return source, nil
}

type WordDelimiterType struct {
    word  string
    _type WordDelimiter
}

func NewWordDelimiterType(word string, _type WordDelimiter) *WordDelimiterType {
	return &WordDelimiterType{word: word, _type: _type}
}

func (wordDelimiter *WordDelimiterType) String() string {
    return fmt.Sprintf("%s => %s", wordDelimiter.word, wordDelimiter._type)
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-word-delimiter-graph-tokenfilter.html
type FilterWordDelimiterGraph struct {
    FilterWordDelimiter

    adjustOffsets  *bool
    ignoreKeywords *bool
}

func NewFilterWordDelimiterGraph(name string) *FilterWordDelimiterGraph {
	return &FilterWordDelimiterGraph{FilterWordDelimiter: FilterWordDelimiter{name: FilterName(name)}}
}

func (filter *FilterWordDelimiterGraph) SetAdjustOffsets(adjustOffsets bool) *FilterWordDelimiterGraph {
	filter.adjustOffsets = &adjustOffsets

	return filter
}

func (filter *FilterWordDelimiterGraph) SetIgnoreKeywords(ignoreKeywords bool) *FilterWordDelimiterGraph {
	filter.ignoreKeywords = &ignoreKeywords

	return filter
}

func (filter *FilterWordDelimiterGraph) Type() FilterType {
    return FilterTypeWordDelimiterGraph
}

func (filter *FilterWordDelimiterGraph) Source() (interface{}, error) {
    sourceTemp, err := filter.FilterWordDelimiter.Source()
    if err != nil {
        return nil, err
    }

    source := sourceTemp.(map[string]interface{})

    source["type"] = filter.Type()

    if filter.adjustOffsets != nil {
        source["adjust_offsets"] = *filter.adjustOffsets
    }

    if filter.ignoreKeywords != nil {
        source["ignore_keywords"] = *filter.ignoreKeywords
    }

    return source, nil
}
