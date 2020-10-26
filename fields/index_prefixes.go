package fields

type IndexPrefixes struct {
    minChars *int
    maxChars *int
}

func NewIndexPrefixes() *IndexPrefixes {
    return &IndexPrefixes{}
}

func (prefix *IndexPrefixes) SetMinimumChars(minimumChars int) *IndexPrefixes {
    prefix.minChars = &minimumChars

    return prefix
}

func (prefix *IndexPrefixes) SetMaximumChars(maximumChars int) *IndexPrefixes {
    prefix.maxChars = &maximumChars

    return prefix
}

func (prefix *IndexPrefixes) Source() (interface{}, error) {
    // {
    //  "min_chars" : 1,
    //  "max_chars" : 10
    // }

    source := map[string]interface{}{}

    if prefix.minChars != nil {
        source["min_chars"] = *prefix.minChars
    }

    if prefix.maxChars != nil {
        source["max_chars"] = *prefix.maxChars
    }

    return source, nil
}
