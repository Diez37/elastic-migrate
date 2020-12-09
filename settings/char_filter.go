package settings

import (
	"fmt"
	"strings"
)

const (
	CharFilterNameHtmlStrip CharFilterName = "html_strip"

	CharFilterTypeHtmlStrip      CharFilterType = "html_strip"
	CharFilterTypeMapping        CharFilterType = "mapping"
	CharFilterTypePatternReplace CharFilterType = "pattern_replace"
)

type CharFilterType string
type CharFilterName string

type CharFilter interface {
	Type() CharFilterType
	Name() CharFilterName
	Source() (interface{}, error)
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-htmlstrip-charfilter.html
type CharFilterHtmlStrip struct {
	name        CharFilterName
	escapedTags []string
}

func NewCharFilterHtmlStrip(name string) *CharFilterHtmlStrip {
	return &CharFilterHtmlStrip{name: CharFilterName(name)}
}

func (filter *CharFilterHtmlStrip) AddEscapedTags(escapedTags ...string) *CharFilterHtmlStrip {
	filter.escapedTags = append(filter.escapedTags, escapedTags...)

	return filter
}

func (filter *CharFilterHtmlStrip) Type() CharFilterType {
	return CharFilterTypeHtmlStrip
}

func (filter *CharFilterHtmlStrip) Name() CharFilterName {
	return filter.name
}

func (filter *CharFilterHtmlStrip) Source() (interface{}, error) {
	source := map[string]interface{}{
		"type": filter.Type(),
	}

	source["escaped_tags"] = filter.escapedTags

	return source, nil
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-mapping-charfilter.html
type CharFilterMapping struct {
	name         CharFilterName
	mappingsPath *string
	mappings     []*CharMapping
}

func NewCharFilterMapping(name string) *CharFilterMapping {
	return &CharFilterMapping{name: CharFilterName(name)}
}

func (filter *CharFilterMapping) SetMappingsPath(mappingsPath string) {
	filter.mappingsPath = &mappingsPath
}

func (filter *CharFilterMapping) AddMappings(mappings ...*CharMapping) *CharFilterMapping {
	filter.mappings = append(filter.mappings, mappings...)

	return filter
}

func (filter *CharFilterMapping) Type() CharFilterType {
	return CharFilterTypeMapping
}

func (filter *CharFilterMapping) Name() CharFilterName {
	return filter.name
}

func (filter *CharFilterMapping) Source() (interface{}, error) {
	source := map[string]interface{}{
		"type": filter.Type(),
	}

	if filter.mappingsPath != nil {
		source["mappings_path"] = *filter.mappingsPath
	}

	if len(filter.mappings) > 0 {
		mappings := make([]string, len(filter.mappings))

		for index, charMapping := range filter.mappings {
			mappings[index] = charMapping.String()
		}

		source["mappings"] = mappings
	}

	return source, nil
}

type CharMapping struct {
	from string
	to   string
}

func NewCharMapping(from string, to string) *CharMapping {
	return &CharMapping{from: from, to: to}
}

func (charMapping *CharMapping) String() string {
	return fmt.Sprintf("%s => %s", charMapping.from, charMapping.to)
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-pattern-replace-charfilter.html
type CharFilterPatternReplace struct {
	name        CharFilterName
	pattern     *string
	replacement *string
	flags       []JavaRegularFlag
}

func NewCharFilterPatternReplace(name string) *CharFilterPatternReplace {
	return &CharFilterPatternReplace{name: CharFilterName(name)}
}

func (filter *CharFilterPatternReplace) SetPattern(pattern string) {
	filter.pattern = &pattern
}

func (filter *CharFilterPatternReplace) SetReplacement(replacement string) {
	filter.replacement = &replacement
}

func (filter *CharFilterPatternReplace) AddFlags(flags ...JavaRegularFlag) *CharFilterPatternReplace {
	filter.flags = append(filter.flags, flags...)

	return filter
}

func (filter *CharFilterPatternReplace) Type() CharFilterType {
	return CharFilterTypePatternReplace
}

func (filter *CharFilterPatternReplace) Name() CharFilterName {
	return filter.name
}

func (filter *CharFilterPatternReplace) Source() (interface{}, error) {
	source := map[string]interface{}{
		"type": filter.Type(),
	}

	if filter.pattern != nil {
		source["pattern"] = *filter.pattern
	}

	if filter.replacement != nil {
		source["replacement"] = *filter.replacement
	}

	if len(filter.flags) > 0 {
		flags := make([]string, len(filter.flags))

		for index, flag := range filter.flags {
			flags[index] = string(flag)
		}

		source["flags"] = strings.Join(flags, JavaRegularFlagSeparator)
	}

	return source, nil
}
