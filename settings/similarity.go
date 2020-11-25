package settings

const (
    SimilarityNameDefault SimilarityName = "default"

    SimilarityTypeBM25            SimilarityType = "BM25"
    SimilarityTypeClassic         SimilarityType = "classic" // TF/IDF
    SimilarityTypeBoolean         SimilarityType = "boolean"
    SimilarityTypeDFR             SimilarityType = "DFR"
    SimilarityTypeDFI             SimilarityType = "DFI"
    SimilarityTypeIB              SimilarityType = "IB"
    SimilarityTypeLMDirichlet     SimilarityType = "LMDirichlet"
    SimilarityTypeLMJelinekMercer SimilarityType = "LMJelinekMercer"
    SimilarityTypeScripted        SimilarityType = "scripted"

    DFRBasicModelG   SimilarityDFRBasicModelType = "g"
    DFRBasicModelIf  SimilarityDFRBasicModelType = "if"
    DFRBasicModelIn  SimilarityDFRBasicModelType = "in"
    DFRBasicModelIne SimilarityDFRBasicModelType = "ine"

    DFRAfterEffectB SimilarityDFRAfterEffectType = "b"
    DFRAfterEffectL SimilarityDFRAfterEffectType = "l"

    DFRNormalizationNo SimilarityDFRNormalizationType = "no"
    DFRNormalizationH1 SimilarityDFRNormalizationType = "h1"
    DFRNormalizationH2 SimilarityDFRNormalizationType = "h2"
    DFRNormalizationH3 SimilarityDFRNormalizationType = "h3"
    DFRNormalizationZ  SimilarityDFRNormalizationType = "z"

    DFIIndependenceMeasureStandardized SimilarityDFIIndependenceMeasureType = "standardized"
    DFIIndependenceMeasureSaturated    SimilarityDFIIndependenceMeasureType = "saturated"
    DFIIndependenceMeasureChisquared   SimilarityDFIIndependenceMeasureType = "chisquared"

    IBDistributionLL  SimilarityIBDistributionType = "ll"
    IBDistributionSPL SimilarityIBDistributionType = "spl"

    IBLambdaDF  SimilarityIBLambdaType = "df"
    IBLambdaTTF SimilarityIBLambdaType = "ttf"
)

type SimilarityName string
type SimilarityType string
type SimilarityDFRBasicModelType string
type SimilarityDFRAfterEffectType string
type SimilarityDFRNormalizationType string
type SimilarityDFIIndependenceMeasureType string
type SimilarityIBDistributionType string
type SimilarityIBLambdaType string

type Similarity interface {
    Type() SimilarityType
    Name() SimilarityName
}

// see https://en.wikipedia.org/wiki/Okapi_BM25
type SimilarityBM25 struct {
    name             SimilarityName
    discountOverlaps *bool
    b                *float32
    k1               *float32
}

func NewSimilarityBM25(name string) *SimilarityBM25 {
    return &SimilarityBM25{name: SimilarityName(name)}
}

func (similarity *SimilarityBM25) SetB(b float32) *SimilarityBM25 {
    similarity.b = &b

    return similarity
}

func (similarity *SimilarityBM25) SetK1(k1 float32) *SimilarityBM25 {
    similarity.k1 = &k1

    return similarity
}

func (similarity *SimilarityBM25) SetDiscountOverlaps(discountOverlaps bool) *SimilarityBM25 {
    similarity.discountOverlaps = &discountOverlaps

    return similarity
}

func (similarity *SimilarityBM25) Type() SimilarityType {
    return SimilarityTypeBM25
}

func (similarity *SimilarityBM25) Name() SimilarityName {
    return similarity.name
}

func (similarity *SimilarityBM25) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": similarity.Type(),
    }

    if similarity.discountOverlaps != nil {
        source["discount_overlaps"] = *similarity.discountOverlaps
    }

    if similarity.b != nil {
        source["b"] = *similarity.b
    }

    if similarity.k1 != nil {
        source["k1"] = *similarity.k1
    }

    return source, nil
}

// see https://en.wikipedia.org/wiki/Divergence-from-randomness_model
type SimilarityDFR struct {
    name               SimilarityName
    basicModel         *SimilarityDFRBasicModelType
    afterEffect        *SimilarityDFRAfterEffectType
    normalization      *SimilarityDFRNormalizationType
    normalizationValue *float32
}

func NewSimilarityDFR(name string) *SimilarityDFR {
    return &SimilarityDFR{name: SimilarityName(name)}
}

func (similarity *SimilarityDFR) SetBasicModel(basicModel SimilarityDFRBasicModelType) *SimilarityDFR {
    similarity.basicModel = &basicModel

    return similarity
}

func (similarity *SimilarityDFR) SetAfterEffect(afterEffect SimilarityDFRAfterEffectType) *SimilarityDFR {
    similarity.afterEffect = &afterEffect

    return similarity
}

func (similarity *SimilarityDFR) SetNormalization(normalization SimilarityDFRNormalizationType, value float32) *SimilarityDFR {
    similarity.normalization = &normalization

    similarity.normalizationValue = &value

    return similarity
}

func (similarity *SimilarityDFR) Type() SimilarityType {
    return SimilarityTypeDFR
}

func (similarity *SimilarityDFR) Name() SimilarityName {
    return similarity.name
}

func (similarity *SimilarityDFR) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": similarity.Type(),
    }

    if similarity.basicModel != nil {
        source["basic_model"] = *similarity.basicModel
    }

    if similarity.afterEffect != nil {
        source["after_effect"] = *similarity.afterEffect
    }

    if similarity.normalization != nil {
        source["normalization"] = *similarity.normalization

        switch *similarity.normalization {
        case DFRNormalizationH1:
            source["normalization.h1.c"] = *similarity.normalizationValue
        case DFRNormalizationH2:
            source["normalization.h2.c"] = *similarity.normalizationValue
        case DFRNormalizationH3:
            source["normalization.h3.c"] = *similarity.normalizationValue
        case DFRNormalizationZ:
            source["normalization.z.c"] = *similarity.normalizationValue
        }
    }

    return source, nil
}

// see https://trec.nist.gov/pubs/trec21/papers/irra.web.nb.pdf
type SimilarityDFI struct {
    name                SimilarityName
    independenceMeasure *SimilarityDFIIndependenceMeasureType
}

func NewSimilarityDFI(name string) *SimilarityDFI {
    return &SimilarityDFI{name: SimilarityName(name)}
}

func (similarity *SimilarityDFI) SetIndependenceMeasure(independenceMeasure SimilarityDFIIndependenceMeasureType) *SimilarityDFI {
    similarity.independenceMeasure = &independenceMeasure

    return similarity
}

func (similarity *SimilarityDFI) Type() SimilarityType {
    return SimilarityTypeDFI
}

func (similarity *SimilarityDFI) Name() SimilarityName {
    return similarity.name
}

func (similarity *SimilarityDFI) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": similarity.Type(),
    }

    if similarity.independenceMeasure != nil {
        source["independence_measure"] = *similarity.independenceMeasure
    }

    return source, nil
}

// see https://lucene.apache.org/core/8_1_0/core/org/apache/lucene/search/similarities/IBSimilarity.html
type SimilarityIB struct {
    name               SimilarityName
    distribution       *SimilarityIBDistributionType
    lambda             *SimilarityIBLambdaType
    normalization      *SimilarityDFRNormalizationType
    normalizationValue *float32
}

func NewSimilarityIB(name string) *SimilarityIB {
    return &SimilarityIB{name: SimilarityName(name)}
}

func (similarity *SimilarityIB) SetDistribution(distribution SimilarityIBDistributionType) *SimilarityIB {
    similarity.distribution = &distribution

    return similarity
}

func (similarity *SimilarityIB) SetLambda(lambda SimilarityIBLambdaType) *SimilarityIB {
    similarity.lambda = &lambda

    return similarity
}

func (similarity *SimilarityIB) SetNormalization(normalization SimilarityDFRNormalizationType, value float32) *SimilarityIB {
    similarity.normalization = &normalization

    similarity.normalizationValue = &value

    return similarity
}

func (similarity *SimilarityIB) Type() SimilarityType {
    return SimilarityTypeIB
}

func (similarity *SimilarityIB) Name() SimilarityName {
    return similarity.name
}

func (similarity *SimilarityIB) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": similarity.Type(),
    }

    if similarity.distribution != nil {
        source["distribution"] = *similarity.distribution
    }

    if similarity.lambda != nil {
        source["lambda"] = *similarity.lambda
    }

    if similarity.normalization != nil {
        source["normalization"] = *similarity.normalization

        switch *similarity.normalization {
        case DFRNormalizationH1:
            source["normalization.h1.c"] = *similarity.normalizationValue
        case DFRNormalizationH2:
            source["normalization.h2.c"] = *similarity.normalizationValue
        case DFRNormalizationH3:
            source["normalization.h3.c"] = *similarity.normalizationValue
        case DFRNormalizationZ:
            source["normalization.z.c"] = *similarity.normalizationValue
        }
    }

    return source, nil
}

// see https://lucene.apache.org/core/8_0_0/core/org/apache/lucene/search/similarities/LMDirichletSimilarity.html
type SimilarityLMDirichlet struct {
    name SimilarityName
    mu   *float32
}

func NewSimilarityLMDirichlet(name string) *SimilarityLMDirichlet {
    return &SimilarityLMDirichlet{name: SimilarityName(name)}
}

func (similarity *SimilarityLMDirichlet) SetMU(mu float32) *SimilarityLMDirichlet {
    similarity.mu = &mu

    return similarity
}

func (similarity *SimilarityLMDirichlet) Type() SimilarityType {
    return SimilarityTypeLMDirichlet
}

func (similarity *SimilarityLMDirichlet) Name() SimilarityName {
    return similarity.name
}

func (similarity *SimilarityLMDirichlet) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": similarity.Type(),
    }

    if similarity.mu != nil {
        source["mu"] = *similarity.mu
    }

    return source, nil
}

// see https://lucene.apache.org/core/8_6_2/core/org/apache/lucene/search/similarities/LMJelinekMercerSimilarity.html
// or see https://www.researchgate.net/publication/221152474_A_Comparative_Study_of_Probabalistic_and_Language_Models_for_Information_Retrieval
type SimilarityLMJelinekMercer struct {
    name   SimilarityName
    lambda *float32
}

func NewSimilarityLMJelinekMercer(name string) *SimilarityLMJelinekMercer {
    return &SimilarityLMJelinekMercer{name: SimilarityName(name)}
}

func (similarity *SimilarityLMJelinekMercer) SetLambda(lambda float32) *SimilarityLMJelinekMercer {
    similarity.lambda = &lambda

    return similarity
}

func (similarity *SimilarityLMJelinekMercer) Type() SimilarityType {
    return SimilarityTypeLMJelinekMercer
}

func (similarity *SimilarityLMJelinekMercer) Name() SimilarityName {
    return similarity.name
}

func (similarity *SimilarityLMJelinekMercer) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": similarity.Type(),
    }

    if similarity.lambda != nil {
        source["lambda"] = *similarity.lambda
    }

    return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules-similarity.html#scripted_similarity
type SimilarityScripted struct {
    name         SimilarityName
    weightScript *SimilarityScript
    script       *SimilarityScript
}

func NewSimilarityScripted(name string) *SimilarityScripted {
    return &SimilarityScripted{name: SimilarityName(name)}
}

func (similarity *SimilarityScripted) SetScript(script *SimilarityScript) *SimilarityScripted {
    similarity.script = script

    return similarity
}

func (similarity *SimilarityScripted) SetWeightScript(script *SimilarityScript) *SimilarityScripted {
    similarity.weightScript = script

    return similarity
}

func (similarity *SimilarityScripted) Type() SimilarityType {
    return SimilarityTypeScripted
}

func (similarity *SimilarityScripted) Name() SimilarityName {
    return similarity.name
}

func (similarity *SimilarityScripted) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": similarity.Type(),
    }

    if similarity.weightScript != nil {
        script, _ := similarity.weightScript.Source()

        source["weight_script"] = script
    }

    if similarity.script != nil {
        script, _ := similarity.script.Source()

        source["script"] = script
    }

    return source, nil
}

type SimilarityScript struct {
    source string
}

func NewSimilarityScript(source string) *SimilarityScript {
    return &SimilarityScript{source: source}
}

func (script *SimilarityScript) Source() (interface{}, error) {
    return map[string]interface{}{
        "source": script.source,
    }, nil
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules-similarity.html#default-base
type SimilarityDefault struct {
    name  SimilarityName
    _type SimilarityType
}

func NewSimilarityBoolean(name string) *SimilarityDefault {
    return &SimilarityDefault{name: SimilarityName(name), _type: SimilarityTypeBoolean}
}

func NewSimilarityClassic(name string) *SimilarityDefault {
    return &SimilarityDefault{name: SimilarityName(name), _type: SimilarityTypeClassic}
}

func (similarity *SimilarityDefault) Type() SimilarityType {
    return similarity._type
}

func (similarity *SimilarityDefault) Name() SimilarityName {
    return similarity.name
}

func (similarity *SimilarityDefault) Source() (interface{}, error) {
    return map[string]interface{}{
        "type": similarity.Type(),
    }, nil
}
