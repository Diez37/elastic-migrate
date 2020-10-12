package fields

type Keyword struct {
    docValues                *bool
    store                    *bool
    index                    *bool
    eagerGlobalOrdinals      *bool
    norms                    *bool
    splitQueriesOnWhitespace *bool
    boost                    *float64
    ignoreAbove              *int
    indexOptions             *IndexOption
    similarity               *Similarity
    normalizer               *Normalizer
    nullValue                interface{}
    fields                   []*Field
    meta                     []*Meta
}

func NewKeyword() *Keyword {
    return &Keyword{}
}

func (field *Keyword) SetDocValues(docValues bool) *Keyword {
    field.docValues = &docValues

    return field
}

func (field *Keyword) SetStore(store bool) *Keyword {
    field.store = &store

    return field
}

func (field *Keyword) SetIndex(index bool) *Keyword {
    field.index = &index

    return field
}

func (field *Keyword) SetEagerGlobalOrdinals(eagerGlobalOrdinals bool) *Keyword {
    field.eagerGlobalOrdinals = &eagerGlobalOrdinals

    return field
}

func (field *Keyword) SetNorms(norms bool) *Keyword {
    field.norms = &norms

    return field
}

func (field *Keyword) SetSplitQueriesOnWhitespace(splitQueriesOnWhitespace bool) *Keyword {
    field.splitQueriesOnWhitespace = &splitQueriesOnWhitespace

    return field
}

func (field *Keyword) SetBoost(boost float64) *Keyword {
    field.boost = &boost

    return field
}

func (field *Keyword) SetIgnoreAbove(ignoreAbove int) *Keyword {
    field.ignoreAbove = &ignoreAbove

    return field
}

func (field *Keyword) SetIndexOptions(indexOptions IndexOption) *Keyword {
    field.indexOptions = &indexOptions

    return field
}

func (field *Keyword) SetSimilarity(similarity Similarity) *Keyword {
    field.similarity = &similarity

    return field
}

func (field *Keyword) SetNormalizer(normalizer Normalizer) *Keyword {
    field.normalizer = &normalizer

    return field
}

func (field *Keyword) SetNullValue(nullValue interface{}) *Keyword {
    field.nullValue = nullValue

    return field
}

func (field *Keyword) Meta(metas ...*Meta) *Keyword {
    field.meta = append(field.meta, metas...)

    return field
}

func (field *Keyword) Fields(fields ...*Field) *Keyword {
    field.fields = append(field.fields, fields...)

    return field
}

func (field *Keyword) GetType() Type {
    return TypeKeyword
}

func (field *Keyword) Source() (interface{}, error) {
    //	{
    //		"type": "keyword",
    //		"doc_values": false,
    //		"store": false,
    //		"boost": 2.5,
    //		"index": true,
    //		"null_value": "NULL",
    //		"meta": {
    //          "unit": "ms"
    //        },
    //		"eager_global_ordinals": true,
    //		"fields": {
    //          "raw": {
    //            "type":  "keyword"
    //          }
    //        },
    //		"ignore_above": 256
    //		"index_options": "docs",
    //		"norms": false,
    //		"similarity": "BM25",
    //		"normalizer": "lowercase",
    //		"split_queries_on_whitespace": false
    //	}

    source := map[string]interface{}{}

    source["type"] = field.GetType()

    if field.docValues != nil {
        source["doc_values"] = *field.docValues
    }

    if field.store != nil {
        source["store"] = *field.store
    }

    if field.boost != nil {
        source["boost"] = *field.boost
    }

    if field.index != nil {
        source["index"] = *field.index
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

    if field.eagerGlobalOrdinals != nil {
        source["eager_global_ordinals"] = *field.eagerGlobalOrdinals
    }

    if len(field.fields) > 0 {
        body := map[string]interface{}{}

        for _, subField := range field.fields {
            s, err := subField.Source()
            if err != nil {
                return nil, err
            }

            body[subField.name] = s
        }

        source["fields"] = body
    }

    if field.ignoreAbove != nil {
        source["ignore_above"] = *field.ignoreAbove
    }

    if field.indexOptions != nil {
        source["index_options"] = field.indexOptions
    }

    if field.norms != nil {
        source["norms"] = *field.norms
    }

    if field.similarity != nil {
        source["similarity"] = *field.similarity
    }

    if field.normalizer != nil {
        source["normalizer"] = *field.normalizer
    }

    if field.splitQueriesOnWhitespace != nil {
        source["split_queries_on_whitespace"] = *field.splitQueriesOnWhitespace
    }

    return source, nil
}
