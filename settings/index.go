package settings

const (
	CodecBestCompression IndexCode    = "best_compression"
	PipelineNameNone     PipelineName = "_none"
)

type IndexCode string
type PipelineName string

type Index struct {
	loadFixedBitsetFiltersEagerly *bool
	hidden                        *bool
	verifiedBeforeClose           *bool
	sourceOnly                    *bool
	numberOfShards                *uint32
	numberOfRoutingShards         *uint32
	routingPartitionSize          *uint32
	numberOfReplicas              *uint32
	maxResultWindow               *uint32
	maxInnerResultWindow          *uint32
	maxRescoreWindow              *uint32
	maxDocvalueFieldsSearch       *uint32
	maxScriptFields               *uint32
	maxNgramDiff                  *uint32
	maxShingleDiff                *uint32
	maxRefreshListeners           *uint32
	maxTermsCount                 *uint32
	maxRegexLength                *uint32
	maxAdjacencyMatrixFilters     *uint32
	maxSlicesPerScroll            *uint32
	autoExpandReplicas            *string
	defaultPipeline               *PipelineName
	finalPipeline                 *PipelineName
	codec                         *IndexCode
	refreshInterval               *Interval
	gcDeletes                     *Interval
	flushAfterMerge               *Size
	search                        *Search
	analyze                       *Analyze
	highlight                     *Highlight
	routing                       *Routing
	shard                         *Shard
	softDeletes                   *SoftDeletes
	unassigned                    *Unassigned
	mapping                       *Mapping
	merge                         *Merge
	write                         *Write
	indexing                      *Slowlog
	blocks                        *Blocks
	analysis                      *Analysis
	similarity                    []Similarity
}

func NewIndex() *Index {
	return &Index{}
}

func (index *Index) SetLoadFixedBitsetFiltersEagerly(loadFixedBitsetFiltersEagerly bool) *Index {
	index.loadFixedBitsetFiltersEagerly = &loadFixedBitsetFiltersEagerly

	return index
}

func (index *Index) SetHidden(hidden bool) *Index {
	index.hidden = &hidden

	return index
}

func (index *Index) SetVerifiedBeforeClose(verifiedBeforeClose bool) *Index {
	index.verifiedBeforeClose = &verifiedBeforeClose

	return index
}

func (index *Index) SetSourceOnly(sourceOnly bool) *Index {
	index.sourceOnly = &sourceOnly

	return index
}

func (index *Index) SetNumberOfShards(numberOfShards uint32) *Index {
	index.numberOfShards = &numberOfShards

	return index
}

func (index *Index) SetNumberOfRoutingShards(numberOfRoutingShards uint32) *Index {
	index.numberOfRoutingShards = &numberOfRoutingShards

	return index
}

func (index *Index) SetRoutingPartitionSize(routingPartitionSize uint32) *Index {
	index.routingPartitionSize = &routingPartitionSize

	return index
}

func (index *Index) SetNumberOfReplicas(numberOfReplicas uint32) *Index {
	index.numberOfReplicas = &numberOfReplicas

	return index
}

func (index *Index) SetMaxResultWindow(maxResultWindow uint32) *Index {
	index.maxResultWindow = &maxResultWindow

	return index
}

func (index *Index) SetMaxInnerResultWindow(maxInnerResultWindow uint32) *Index {
	index.maxInnerResultWindow = &maxInnerResultWindow

	return index
}

func (index *Index) SetMaxRescoreWindow(maxRescoreWindow uint32) *Index {
	index.maxRescoreWindow = &maxRescoreWindow

	return index
}

func (index *Index) SetMaxDocvalueFieldsSearch(maxDocvalueFieldsSearch uint32) *Index {
	index.maxDocvalueFieldsSearch = &maxDocvalueFieldsSearch

	return index
}

func (index *Index) SetMaxScriptFields(maxScriptFields uint32) *Index {
	index.maxScriptFields = &maxScriptFields

	return index
}

func (index *Index) SetMaxNgramDiff(maxNgramDiff uint32) *Index {
	index.maxNgramDiff = &maxNgramDiff

	return index
}

func (index *Index) SetMaxShingleDiff(maxShingleDiff uint32) *Index {
	index.maxShingleDiff = &maxShingleDiff

	return index
}

func (index *Index) SetMaxRefreshListeners(maxRefreshListeners uint32) *Index {
	index.maxRefreshListeners = &maxRefreshListeners

	return index
}

func (index *Index) SetMaxTermsCount(maxTermsCount uint32) *Index {
	index.maxTermsCount = &maxTermsCount

	return index
}

func (index *Index) SetMaxRegexLength(maxRegexLength uint32) *Index {
	index.maxRegexLength = &maxRegexLength

	return index
}

func (index *Index) SetMaxAdjacencyMatrixFilters(maxAdjacencyMatrixFilters uint32) *Index {
	index.maxAdjacencyMatrixFilters = &maxAdjacencyMatrixFilters

	return index
}

func (index *Index) SetMaxSlicesPerScroll(maxSlicesPerScroll uint32) *Index {
	index.maxSlicesPerScroll = &maxSlicesPerScroll

	return index
}

func (index *Index) SetAutoExpandReplicas(autoExpandReplicas string) *Index {
	index.autoExpandReplicas = &autoExpandReplicas

	return index
}

func (index *Index) SetRefreshInterval(refreshInterval *Interval) *Index {
	index.refreshInterval = refreshInterval

	return index
}

func (index *Index) SetGcDeletes(gcDeletes *Interval) *Index {
	index.gcDeletes = gcDeletes

	return index
}

func (index *Index) SetFlushAfterMerge(flushAfterMerge *Size) *Index {
	index.flushAfterMerge = flushAfterMerge

	return index
}

func (index *Index) SetDefaultPipeline(defaultPipeline *PipelineName) *Index {
	index.defaultPipeline = defaultPipeline

	return index
}

func (index *Index) SetFinalPipeline(finalPipeline *PipelineName) *Index {
	index.finalPipeline = finalPipeline

	return index
}

func (index *Index) SetSearch(search *Search) *Index {
	index.search = search

	return index
}

func (index *Index) SetAnalyze(analyze *Analyze) *Index {
	index.analyze = analyze

	return index
}

func (index *Index) SetHighlight(highlight *Highlight) *Index {
	index.highlight = highlight

	return index
}

func (index *Index) SetRouting(routing *Routing) *Index {
	index.routing = routing

	return index
}

func (index *Index) SetShard(shard *Shard) *Index {
	index.shard = shard

	return index
}

func (index *Index) SetSoftDeletes(softDeletes *SoftDeletes) *Index {
	index.softDeletes = softDeletes

	return index
}

func (index *Index) SetCodec(codec *IndexCode) *Index {
	index.codec = codec

	return index
}

func (index *Index) SetUnassigned(unassigned *Unassigned) *Index {
	index.unassigned = unassigned

	return index
}

func (index *Index) SetMapping(mapping *Mapping) *Index {
	index.mapping = mapping

	return index
}

func (index *Index) SetMerge(merge *Merge) *Index {
	index.merge = merge

	return index
}

func (index *Index) SetWrite(write *Write) *Index {
	index.write = write

	return index
}

func (index *Index) SetIndexing(indexing *Slowlog) *Index {
	index.indexing = indexing

	return index
}

func (index *Index) SetBlocks(blocks *Blocks) *Index {
	index.blocks = blocks

	return index
}

func (index *Index) SetAnalysis(analysis *Analysis) *Index {
	index.analysis = analysis

	return index
}

func (index *Index) AddSimilarity(similarity ...Similarity) *Index {
	index.similarity = append(index.similarity, similarity...)

	return index
}

func (index *Index) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if index.loadFixedBitsetFiltersEagerly != nil {
		source["load_fixed_bitset_filters_eagerly"] = *index.loadFixedBitsetFiltersEagerly
	}

	if index.hidden != nil {
		source["hidden"] = *index.hidden
	}

	if index.verifiedBeforeClose != nil {
		source["verified_before_close"] = *index.verifiedBeforeClose
	}

	if index.sourceOnly != nil {
		source["source_only"] = *index.sourceOnly
	}

	if index.numberOfShards != nil {
		source["number_of_shards"] = *index.numberOfShards
	}

	if index.numberOfRoutingShards != nil {
		source["number_of_routing_shards"] = *index.numberOfRoutingShards
	}

	if index.routingPartitionSize != nil {
		source["routing_partition_size"] = *index.routingPartitionSize
	}

	if index.numberOfReplicas != nil {
		source["number_of_replicas"] = *index.numberOfReplicas
	}

	if index.maxResultWindow != nil {
		source["max_result_window"] = *index.maxResultWindow
	}

	if index.maxInnerResultWindow != nil {
		source["max_inner_result_window"] = *index.maxInnerResultWindow
	}

	if index.maxRescoreWindow != nil {
		source["max_rescore_window"] = *index.maxRescoreWindow
	}

	if index.maxDocvalueFieldsSearch != nil {
		source["max_docvalue_fields_search"] = *index.maxDocvalueFieldsSearch
	}

	if index.maxScriptFields != nil {
		source["max_script_fields"] = *index.maxScriptFields
	}

	if index.maxNgramDiff != nil {
		source["max_ngram_diff"] = *index.maxNgramDiff
	}

	if index.maxShingleDiff != nil {
		source["max_shingle_diff"] = *index.maxShingleDiff
	}

	if index.maxRefreshListeners != nil {
		source["max_refresh_listeners"] = *index.maxRefreshListeners
	}

	if index.maxTermsCount != nil {
		source["max_terms_count"] = *index.maxTermsCount
	}

	if index.maxRegexLength != nil {
		source["max_regex_length"] = *index.maxRegexLength
	}

	if index.maxAdjacencyMatrixFilters != nil {
		source["max_adjacency_matrix_filters"] = *index.maxAdjacencyMatrixFilters
	}

	if index.maxSlicesPerScroll != nil {
		source["max_slices_per_scroll"] = *index.maxSlicesPerScroll
	}

	if index.autoExpandReplicas != nil {
		source["auto_expand_replicas"] = *index.autoExpandReplicas
	}

	if index.defaultPipeline != nil {
		source["default_pipeline"] = *index.defaultPipeline
	}

	if index.finalPipeline != nil {
		source["final_pipeline"] = *index.finalPipeline
	}

	if index.codec != nil {
		source["codec"] = *index.codec
	}

	if index.refreshInterval != nil {
		src, err := index.refreshInterval.Source()
		if err != nil {
			return nil, err
		}

		source["refresh_interval"] = src
	}

	if index.gcDeletes != nil {
		src, err := index.gcDeletes.Source()
		if err != nil {
			return nil, err
		}

		source["gc_deletes"] = src
	}

	if index.flushAfterMerge != nil {
		src, err := index.flushAfterMerge.Source()
		if err != nil {
			return nil, err
		}

		source["flush_after_merge"] = src
	}

	if index.search != nil {
		src, err := index.search.Source()
		if err != nil {
			return nil, err
		}

		source["search"] = src
	}

	if index.analyze != nil {
		src, err := index.analyze.Source()
		if err != nil {
			return nil, err
		}

		source["analyze"] = src
	}

	if index.highlight != nil {
		src, err := index.highlight.Source()
		if err != nil {
			return nil, err
		}

		source["highlight"] = src
	}

	if index.routing != nil {
		src, err := index.routing.Source()
		if err != nil {
			return nil, err
		}

		source["routing"] = src
	}

	if index.shard != nil {
		src, err := index.shard.Source()
		if err != nil {
			return nil, err
		}

		source["shard"] = src
	}

	if index.softDeletes != nil {
		src, err := index.softDeletes.Source()
		if err != nil {
			return nil, err
		}

		source["soft_deletes"] = src
	}

	if index.unassigned != nil {
		src, err := index.unassigned.Source()
		if err != nil {
			return nil, err
		}

		source["unassigned"] = src
	}

	if index.mapping != nil {
		src, err := index.mapping.Source()
		if err != nil {
			return nil, err
		}

		source["mapping"] = src
	}

	if index.merge != nil {
		src, err := index.merge.Source()
		if err != nil {
			return nil, err
		}

		source["merge"] = src
	}

	if index.write != nil {
		src, err := index.write.Source()
		if err != nil {
			return nil, err
		}

		source["write"] = src
	}

	if index.indexing != nil {
		src, err := index.indexing.Source()
		if err != nil {
			return nil, err
		}

		source["indexing"] = src
	}

	if index.blocks != nil {
		src, err := index.blocks.Source()
		if err != nil {
			return nil, err
		}

		source["blocks"] = src
	}

	if index.analysis != nil {
		src, err := index.analysis.Source()
		if err != nil {
			return nil, err
		}

		source["analysis"] = src
	}

	if len(index.similarity) > 0 {
		similarities := map[SimilarityName]interface{}{}

		for _, similarity := range index.similarity {
			src, err := similarity.Source()
			if err != nil {
				return nil, err
			}

			similarities[similarity.Name()] = src
		}

		source["similarity"] = similarities
	}

	return source, nil
}
