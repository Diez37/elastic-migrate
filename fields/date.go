package fields

import "fmt"

const (
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
)

type Local string

type Date struct {
	name            FieldName
	docValues       *bool
	ignoreMalformed *bool
	index           *bool
	store           *bool
	boost           *float64
	locale          *Local
	formats         []DateFormat
	nullValue       interface{}
	meta            []*Meta
	_type           Type
}

func NewDate(name string) *Date {
	return &Date{name: FieldName(name), _type: TypeDate}
}

func NewDateNano(name string) *Date {
	return &Date{name: FieldName(name), _type: TypeDateNano}
}

func (field *Date) Name() FieldName {
	return field.name
}

func (field *Date) SetDocValues(docValues bool) *Date {
	field.docValues = &docValues

	return field
}

func (field *Date) SetIgnoreMalformed(ignoreMalformed bool) *Date {
	field.ignoreMalformed = &ignoreMalformed

	return field
}

func (field *Date) SetIndex(index bool) *Date {
	field.index = &index

	return field
}

func (field *Date) SetStore(store bool) *Date {
	field.store = &store

	return field
}

func (field *Date) SetBoost(boost float64) *Date {
	field.boost = &boost

	return field
}

func (field *Date) SetNullValue(nullValue interface{}) *Date {
	field.nullValue = nullValue

	return field
}

func (field *Date) Meta(metas ...*Meta) *Date {
	field.meta = append(field.meta, metas...)

	return field
}

func (field *Date) Formats(formats ...DateFormat) *Date {
	field.formats = append(field.formats, formats...)

	return field
}

func (field *Date) Local(locale Local) *Date {
	field.locale = &locale

	return field
}

func (field *Date) GetType() Type {
	return field._type
}

func (field *Date) Source() (interface{}, error) {
	// {
	//  "type":   "date",
	//  "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
	// }

	source := map[string]interface{}{}

	source["type"] = field.GetType()

	if field.docValues != nil {
		source["doc_values"] = *field.docValues
	}

	if field.index != nil {
		source["index"] = *field.index
	}

	if field.store != nil {
		source["store"] = *field.store
	}

	if field.boost != nil {
		source["boost"] = *field.boost
	}

	if field.nullValue != nil {
		source["null_value"] = field.nullValue
	}

	if len(field.meta) > 0 {
		body := map[string]interface{}{}

		for _, meta := range field.meta {
			body[meta.name] = meta.value
		}

		source["meta"] = body
	}

	if field.ignoreMalformed != nil {
		source["ignore_malformed"] = *field.ignoreMalformed
	}

	if field.locale != nil {
		source["locale"] = *field.locale
	}

	if len(field.formats) > 0 {
		source["format"] = field.formats[0]

		if len(field.formats) > 1 {
			for _, format := range field.formats[1:] {
				source["format"] = fmt.Sprintf("%s%s%s", source["format"], DateFormatSeparator, format)
			}
		}
	}

	return source, nil
}
