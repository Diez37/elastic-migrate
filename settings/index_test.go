package settings

import (
	"reflect"
	"testing"
)

func TestIndex_AddSimilarity(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		similarity []Similarity
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{similarity: []Similarity{&SimilarityBM25{}}},
			want: &Index{similarity: []Similarity{&SimilarityBM25{}}},
		},
		{
			name: "change",
			fields: fields{similarity: []Similarity{&SimilarityBM25{}}},
			args: args{similarity: []Similarity{&SimilarityDFI{}}},
			want: &Index{similarity: []Similarity{&SimilarityBM25{}, &SimilarityDFI{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.AddSimilarity(tt.args.similarity...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetAnalysis(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		analysis *Analysis
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{analysis: &Analysis{}},
			want: &Index{analysis: &Analysis{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetAnalysis(tt.args.analysis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAnalysis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetAnalyze(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		analyze *Analyze
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{analyze: &Analyze{}},
			want: &Index{analyze: &Analyze{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetAnalyze(tt.args.analyze); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAnalyze() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetAutoExpandReplicas(t *testing.T) {
	stringSet := "set"
	stringChange := "change"

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		autoExpandReplicas string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{autoExpandReplicas: "set"},
			want: &Index{autoExpandReplicas: &stringSet},
		},
		{
			name: "change",
			fields: fields{autoExpandReplicas: &stringSet},
			args: args{autoExpandReplicas: "change"},
			want: &Index{autoExpandReplicas: &stringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetAutoExpandReplicas(tt.args.autoExpandReplicas); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAutoExpandReplicas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetBlocks(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		blocks *Blocks
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{blocks: &Blocks{}},
			want: &Index{blocks: &Blocks{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetBlocks(tt.args.blocks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBlocks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetCodec(t *testing.T) {
	codecBestCompression := CodecBestCompression

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		codec IndexCode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{codec: CodecBestCompression},
			want: &Index{codec: &codecBestCompression},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetCodec(tt.args.codec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCodec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetDefaultPipeline(t *testing.T) {
	pipelineNameNone := PipelineNameNone

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		defaultPipeline PipelineName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{defaultPipeline: PipelineNameNone},
			want: &Index{defaultPipeline: &pipelineNameNone},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetDefaultPipeline(tt.args.defaultPipeline); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDefaultPipeline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetFinalPipeline(t *testing.T) {
	pipelineNameNone := PipelineNameNone

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		finalPipeline PipelineName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{finalPipeline: PipelineNameNone},
			want: &Index{finalPipeline: &pipelineNameNone},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetFinalPipeline(tt.args.finalPipeline); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFinalPipeline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetFlushAfterMerge(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		flushAfterMerge *Size
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{flushAfterMerge: &Size{value: "1mb"}},
			want: &Index{flushAfterMerge: &Size{value: "1mb"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetFlushAfterMerge(tt.args.flushAfterMerge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFlushAfterMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetGcDeletes(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		gcDeletes *Interval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{gcDeletes: &Interval{value: "1s"}},
			want: &Index{gcDeletes: &Interval{value: "1s"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetGcDeletes(tt.args.gcDeletes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetGcDeletes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetHidden(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		hidden bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{hidden: true},
			want: &Index{hidden: &boolTrue},
		},
		{
			name: "change",
			fields: fields{hidden: &boolTrue},
			args: args{hidden: false},
			want: &Index{hidden: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetHidden(tt.args.hidden); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHidden() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetHighlight(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		highlight *Highlight
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{highlight: &Highlight{}},
			want: &Index{highlight: &Highlight{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetHighlight(tt.args.highlight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHighlight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetIndexing(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		indexing *Slowlog
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{indexing: &Slowlog{}},
			want: &Index{indexing: &Slowlog{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetIndexing(tt.args.indexing); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndexing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetLoadFixedBitsetFiltersEagerly(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		loadFixedBitsetFiltersEagerly bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{loadFixedBitsetFiltersEagerly: true},
			want: &Index{loadFixedBitsetFiltersEagerly: &boolTrue},
		},
		{
			name: "change",
			fields: fields{loadFixedBitsetFiltersEagerly: &boolTrue},
			args: args{loadFixedBitsetFiltersEagerly: false},
			want: &Index{loadFixedBitsetFiltersEagerly: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetLoadFixedBitsetFiltersEagerly(tt.args.loadFixedBitsetFiltersEagerly); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLoadFixedBitsetFiltersEagerly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMapping(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		mapping *Mapping
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{mapping: &Mapping{}},
			want: &Index{mapping: &Mapping{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMapping(tt.args.mapping); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxAdjacencyMatrixFilters(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxAdjacencyMatrixFilters uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxAdjacencyMatrixFilters: 1},
			want: &Index{maxAdjacencyMatrixFilters: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxAdjacencyMatrixFilters: &numberSet},
			args: args{maxAdjacencyMatrixFilters: 2},
			want: &Index{maxAdjacencyMatrixFilters: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxAdjacencyMatrixFilters(tt.args.maxAdjacencyMatrixFilters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxAdjacencyMatrixFilters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxDocvalueFieldsSearch(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxDocvalueFieldsSearch uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxDocvalueFieldsSearch: 1},
			want: &Index{maxDocvalueFieldsSearch: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxDocvalueFieldsSearch: &numberSet},
			args: args{maxDocvalueFieldsSearch: 2},
			want: &Index{maxDocvalueFieldsSearch: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxDocvalueFieldsSearch(tt.args.maxDocvalueFieldsSearch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxDocvalueFieldsSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxInnerResultWindow(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxInnerResultWindow uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxInnerResultWindow: 1},
			want: &Index{maxInnerResultWindow: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxInnerResultWindow: &numberSet},
			args: args{maxInnerResultWindow: 2},
			want: &Index{maxInnerResultWindow: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxInnerResultWindow(tt.args.maxInnerResultWindow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxInnerResultWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxNgramDiff(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxNgramDiff uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxNgramDiff: 1},
			want: &Index{maxNgramDiff: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxNgramDiff: &numberSet},
			args: args{maxNgramDiff: 2},
			want: &Index{maxNgramDiff: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxNgramDiff(tt.args.maxNgramDiff); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxNgramDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxRefreshListeners(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxRefreshListeners uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxRefreshListeners: 1},
			want: &Index{maxRefreshListeners: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxRefreshListeners: &numberSet},
			args: args{maxRefreshListeners: 2},
			want: &Index{maxRefreshListeners: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxRefreshListeners(tt.args.maxRefreshListeners); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxRefreshListeners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxRegexLength(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxRegexLength uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxRegexLength: 1},
			want: &Index{maxRegexLength: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxRegexLength: &numberSet},
			args: args{maxRegexLength: 2},
			want: &Index{maxRegexLength: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxRegexLength(tt.args.maxRegexLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxRegexLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxRescoreWindow(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxRescoreWindow uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxRescoreWindow: 1},
			want: &Index{maxRescoreWindow: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxRescoreWindow: &numberSet},
			args: args{maxRescoreWindow: 2},
			want: &Index{maxRescoreWindow: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxRescoreWindow(tt.args.maxRescoreWindow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxRescoreWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxResultWindow(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxResultWindow uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxResultWindow: 1},
			want: &Index{maxResultWindow: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxResultWindow: &numberSet},
			args: args{maxResultWindow: 2},
			want: &Index{maxResultWindow: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxResultWindow(tt.args.maxResultWindow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxResultWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxScriptFields(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxScriptFields uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxScriptFields: 1},
			want: &Index{maxScriptFields: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxScriptFields: &numberSet},
			args: args{maxScriptFields: 2},
			want: &Index{maxScriptFields: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxScriptFields(tt.args.maxScriptFields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxScriptFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxShingleDiff(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxShingleDiff uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxShingleDiff: 1},
			want: &Index{maxShingleDiff: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxShingleDiff: &numberSet},
			args: args{maxShingleDiff: 2},
			want: &Index{maxShingleDiff: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxShingleDiff(tt.args.maxShingleDiff); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxShingleDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxSlicesPerScroll(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxSlicesPerScroll uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxSlicesPerScroll: 1},
			want: &Index{maxSlicesPerScroll: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxSlicesPerScroll: &numberSet},
			args: args{maxSlicesPerScroll: 2},
			want: &Index{maxSlicesPerScroll: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxSlicesPerScroll(tt.args.maxSlicesPerScroll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxSlicesPerScroll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMaxTermsCount(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		maxTermsCount uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{maxTermsCount: 1},
			want: &Index{maxTermsCount: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxTermsCount: &numberSet},
			args: args{maxTermsCount: 2},
			want: &Index{maxTermsCount: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMaxTermsCount(tt.args.maxTermsCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxTermsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMerge(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		merge *Merge
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{merge: &Merge{}},
			want: &Index{merge: &Merge{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetMerge(tt.args.merge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetNumberOfReplicas(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		numberOfReplicas uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{numberOfReplicas: 1},
			want: &Index{numberOfReplicas: &numberSet},
		},
		{
			name: "change",
			fields: fields{numberOfReplicas: &numberSet},
			args: args{numberOfReplicas: 2},
			want: &Index{numberOfReplicas: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetNumberOfReplicas(tt.args.numberOfReplicas); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNumberOfReplicas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetNumberOfRoutingShards(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		numberOfRoutingShards uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{numberOfRoutingShards: 1},
			want: &Index{numberOfRoutingShards: &numberSet},
		},
		{
			name: "change",
			fields: fields{numberOfRoutingShards: &numberSet},
			args: args{numberOfRoutingShards: 2},
			want: &Index{numberOfRoutingShards: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetNumberOfRoutingShards(tt.args.numberOfRoutingShards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNumberOfRoutingShards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetNumberOfShards(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		numberOfShards uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{numberOfShards: 1},
			want: &Index{numberOfShards: &numberSet},
		},
		{
			name: "change",
			fields: fields{numberOfShards: &numberSet},
			args: args{numberOfShards: 2},
			want: &Index{numberOfShards: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetNumberOfShards(tt.args.numberOfShards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNumberOfShards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetRefreshInterval(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		refreshInterval *Interval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{refreshInterval: &Interval{value: "1s"}},
			want: &Index{refreshInterval: &Interval{value: "1s"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetRefreshInterval(tt.args.refreshInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRefreshInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetRouting(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		routing *Routing
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{routing: &Routing{}},
			want: &Index{routing: &Routing{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetRouting(tt.args.routing); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRouting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetRoutingPartitionSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		routingPartitionSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{routingPartitionSize: 1},
			want: &Index{routingPartitionSize: &numberSet},
		},
		{
			name: "change",
			fields: fields{routingPartitionSize: &numberSet},
			args: args{routingPartitionSize: 2},
			want: &Index{routingPartitionSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetRoutingPartitionSize(tt.args.routingPartitionSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRoutingPartitionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetSearch(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		search *Search
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{search: &Search{}},
			want: &Index{search: &Search{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetSearch(tt.args.search); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetShard(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		shard *Shard
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{shard: &Shard{}},
			want: &Index{shard: &Shard{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetShard(tt.args.shard); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetShard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetSoftDeletes(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		softDeletes *SoftDeletes
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{softDeletes: &SoftDeletes{}},
			want: &Index{softDeletes: &SoftDeletes{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetSoftDeletes(tt.args.softDeletes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSoftDeletes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetSourceOnly(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		sourceOnly bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{sourceOnly: true},
			want: &Index{sourceOnly: &boolTrue},
		},
		{
			name: "change",
			fields: fields{sourceOnly: &boolTrue},
			args: args{sourceOnly: false},
			want: &Index{sourceOnly: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetSourceOnly(tt.args.sourceOnly); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSourceOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetTranslog(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		translog *Translog
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{translog: &Translog{}},
			want: &Index{translog: &Translog{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetTranslog(tt.args.translog); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTranslog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetUnassigned(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		unassigned *Unassigned
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{unassigned: &Unassigned{}},
			want: &Index{unassigned: &Unassigned{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetUnassigned(tt.args.unassigned); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUnassigned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetVerifiedBeforeClose(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		verifiedBeforeClose bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{verifiedBeforeClose: true},
			want: &Index{verifiedBeforeClose: &boolTrue},
		},
		{
			name: "change",
			fields: fields{verifiedBeforeClose: &boolTrue},
			args: args{verifiedBeforeClose: false},
			want: &Index{verifiedBeforeClose: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetVerifiedBeforeClose(tt.args.verifiedBeforeClose); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetVerifiedBeforeClose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetWrite(t *testing.T) {
	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	type args struct {
		write *Write
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{write: &Write{}},
			want: &Index{write: &Write{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			if got := index.SetWrite(tt.args.write); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWrite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Source(t *testing.T) {
	boolTrue := true
	number := uint32(1)
	stringTest := "test"
	pipelineNameNone := PipelineNameNone
	codecBestCompression := CodecBestCompression
	intervalSrc, _ := (&Interval{value: "1s"}).Source()
	sizeSrc, _ := (&Size{value: "1mb"}).Source()
	searchSrc, _ := (&Search{}).Source()
	analyzeSrc, _ := (&Analyze{maxTokenCount: 1}).Source()
	highlightSrc, _ := (&Highlight{maxAnalyzedOffset: 1}).Source()
	routingSrc, _ := (&Routing{}).Source()
	shardSrc, _ := (&Shard{checkOnStartup: CheckOnStartupFalse}).Source()
	softDeletesSrc, _ := (&SoftDeletes{}).Source()
	unassignedSrc, _ := (&Unassigned{nodeLeft: &NodeLeft{delayedTimeout: &Interval{value: "1s"}}}).Source()
	mappingSrc, _ := (&Mapping{}).Source()
	mergeSrc, _ := (&Merge{}).Source()
	writeSrc, _ := (&Write{}).Source()
	slowLogSrc, _ := (&Slowlog{}).Source()
	blocksSrc, _ := (&Blocks{}).Source()
	analysisSrc, _ := (&Analysis{}).Source()
	translogSrc, _ := (&Translog{}).Source()
	similarityBM25Src, _ := (&SimilarityBM25{}).Source()

	type fields struct {
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
		translog                      *Translog
		similarity                    []Similarity
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{},
		},
		{
			name: "load_fixed_bitset_filters_eagerly",
			fields: fields{loadFixedBitsetFiltersEagerly: &boolTrue},
			want: map[string]interface{}{
				"load_fixed_bitset_filters_eagerly": true,
			},
		},
		{
			name: "hidden",
			fields: fields{hidden: &boolTrue},
			want: map[string]interface{}{
				"hidden": true,
			},
		},
		{
			name: "verified_before_close",
			fields: fields{verifiedBeforeClose: &boolTrue},
			want: map[string]interface{}{
				"verified_before_close": true,
			},
		},
		{
			name: "source_only",
			fields: fields{sourceOnly: &boolTrue},
			want: map[string]interface{}{
				"source_only": true,
			},
		},
		{
			name: "number_of_shards",
			fields: fields{numberOfShards: &number},
			want: map[string]interface{}{
				"number_of_shards": uint32(1),
			},
		},
		{
			name: "number_of_routing_shards",
			fields: fields{numberOfRoutingShards: &number},
			want: map[string]interface{}{
				"number_of_routing_shards": uint32(1),
			},
		},
		{
			name: "routing_partition_size",
			fields: fields{routingPartitionSize: &number},
			want: map[string]interface{}{
				"routing_partition_size": uint32(1),
			},
		},
		{
			name: "number_of_replicas",
			fields: fields{numberOfReplicas: &number},
			want: map[string]interface{}{
				"number_of_replicas": uint32(1),
			},
		},
		{
			name: "max_result_window",
			fields: fields{maxResultWindow: &number},
			want: map[string]interface{}{
				"max_result_window": uint32(1),
			},
		},
		{
			name: "max_inner_result_window",
			fields: fields{maxInnerResultWindow: &number},
			want: map[string]interface{}{
				"max_inner_result_window": uint32(1),
			},
		},
		{
			name: "max_rescore_window",
			fields: fields{maxRescoreWindow: &number},
			want: map[string]interface{}{
				"max_rescore_window": uint32(1),
			},
		},
		{
			name: "max_docvalue_fields_search",
			fields: fields{maxDocvalueFieldsSearch: &number},
			want: map[string]interface{}{
				"max_docvalue_fields_search": uint32(1),
			},
		},
		{
			name: "max_script_fields",
			fields: fields{maxScriptFields: &number},
			want: map[string]interface{}{
				"max_script_fields": uint32(1),
			},
		},
		{
			name: "max_ngram_diff",
			fields: fields{maxNgramDiff: &number},
			want: map[string]interface{}{
				"max_ngram_diff": uint32(1),
			},
		},
		{
			name: "max_shingle_diff",
			fields: fields{maxShingleDiff: &number},
			want: map[string]interface{}{
				"max_shingle_diff": uint32(1),
			},
		},
		{
			name: "max_refresh_listeners",
			fields: fields{maxRefreshListeners: &number},
			want: map[string]interface{}{
				"max_refresh_listeners": uint32(1),
			},
		},
		{
			name: "max_terms_count",
			fields: fields{maxTermsCount: &number},
			want: map[string]interface{}{
				"max_terms_count": uint32(1),
			},
		},
		{
			name: "max_regex_length",
			fields: fields{maxRegexLength: &number},
			want: map[string]interface{}{
				"max_regex_length": uint32(1),
			},
		},
		{
			name: "max_adjacency_matrix_filters",
			fields: fields{maxAdjacencyMatrixFilters: &number},
			want: map[string]interface{}{
				"max_adjacency_matrix_filters": uint32(1),
			},
		},
		{
			name: "max_slices_per_scroll",
			fields: fields{maxSlicesPerScroll: &number},
			want: map[string]interface{}{
				"max_slices_per_scroll": uint32(1),
			},
		},
		{
			name: "auto_expand_replicas",
			fields: fields{autoExpandReplicas: &stringTest},
			want: map[string]interface{}{
				"auto_expand_replicas": "test",
			},
		},
		{
			name: "default_pipeline",
			fields: fields{defaultPipeline: &pipelineNameNone},
			want: map[string]interface{}{
				"default_pipeline": PipelineNameNone,
			},
		},
		{
			name: "final_pipeline",
			fields: fields{finalPipeline: &pipelineNameNone},
			want: map[string]interface{}{
				"final_pipeline": PipelineNameNone,
			},
		},
		{
			name: "codec",
			fields: fields{codec: &codecBestCompression},
			want: map[string]interface{}{
				"codec": CodecBestCompression,
			},
		},
		{
			name: "refresh_interval",
			fields: fields{refreshInterval: &Interval{value: "1s"}},
			want: map[string]interface{}{
				"refresh_interval": intervalSrc,
			},
		},
		{
			name: "gc_deletes",
			fields: fields{gcDeletes: &Interval{value: "1s"}},
			want: map[string]interface{}{
				"gc_deletes": intervalSrc,
			},
		},
		{
			name: "flush_after_merge",
			fields: fields{flushAfterMerge: &Size{value: "1mb"}},
			want: map[string]interface{}{
				"flush_after_merge": sizeSrc,
			},
		},
		{
			name: "search",
			fields: fields{search: &Search{}},
			want: map[string]interface{}{
				"search": searchSrc,
			},
		},
		{
			name: "analyze",
			fields: fields{analyze: &Analyze{maxTokenCount: 1}},
			want: map[string]interface{}{
				"analyze": analyzeSrc,
			},
		},
		{
			name: "highlight",
			fields: fields{highlight: &Highlight{maxAnalyzedOffset: 1}},
			want: map[string]interface{}{
				"highlight": highlightSrc,
			},
		},
		{
			name: "routing",
			fields: fields{routing: &Routing{}},
			want: map[string]interface{}{
				"routing": routingSrc,
			},
		},
		{
			name: "shard",
			fields: fields{shard: &Shard{checkOnStartup: CheckOnStartupFalse}},
			want: map[string]interface{}{
				"shard": shardSrc,
			},
		},
		{
			name: "soft_deletes",
			fields: fields{softDeletes: &SoftDeletes{}},
			want: map[string]interface{}{
				"soft_deletes": softDeletesSrc,
			},
		},
		{
			name: "unassigned",
			fields: fields{unassigned: &Unassigned{nodeLeft: &NodeLeft{delayedTimeout: &Interval{value: "1s"}}}},
			want: map[string]interface{}{
				"unassigned": unassignedSrc,
			},
		},
		{
			name: "mapping",
			fields: fields{mapping: &Mapping{}},
			want: map[string]interface{}{
				"mapping": mappingSrc,
			},
		},
		{
			name: "merge",
			fields: fields{merge: &Merge{}},
			want: map[string]interface{}{
				"merge": mergeSrc,
			},
		},
		{
			name: "write",
			fields: fields{write: &Write{}},
			want: map[string]interface{}{
				"write": writeSrc,
			},
		},
		{
			name: "indexing",
			fields: fields{indexing: &Slowlog{}},
			want: map[string]interface{}{
				"indexing": slowLogSrc,
			},
		},
		{
			name: "blocks",
			fields: fields{blocks: &Blocks{}},
			want: map[string]interface{}{
				"blocks": blocksSrc,
			},
		},
		{
			name: "analysis",
			fields: fields{analysis: &Analysis{}},
			want: map[string]interface{}{
				"analysis": analysisSrc,
			},
		},
		{
			name: "translog",
			fields: fields{translog: &Translog{}},
			want: map[string]interface{}{
				"translog": translogSrc,
			},
		},
		{
			name: "similarity",
			fields: fields{similarity: []Similarity{&SimilarityBM25{name: "test"}}},
			want: map[string]interface{}{
				"similarity": map[SimilarityName]interface{}{
					"test": similarityBM25Src,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				loadFixedBitsetFiltersEagerly: tt.fields.loadFixedBitsetFiltersEagerly,
				hidden:                        tt.fields.hidden,
				verifiedBeforeClose:           tt.fields.verifiedBeforeClose,
				sourceOnly:                    tt.fields.sourceOnly,
				numberOfShards:                tt.fields.numberOfShards,
				numberOfRoutingShards:         tt.fields.numberOfRoutingShards,
				routingPartitionSize:          tt.fields.routingPartitionSize,
				numberOfReplicas:              tt.fields.numberOfReplicas,
				maxResultWindow:               tt.fields.maxResultWindow,
				maxInnerResultWindow:          tt.fields.maxInnerResultWindow,
				maxRescoreWindow:              tt.fields.maxRescoreWindow,
				maxDocvalueFieldsSearch:       tt.fields.maxDocvalueFieldsSearch,
				maxScriptFields:               tt.fields.maxScriptFields,
				maxNgramDiff:                  tt.fields.maxNgramDiff,
				maxShingleDiff:                tt.fields.maxShingleDiff,
				maxRefreshListeners:           tt.fields.maxRefreshListeners,
				maxTermsCount:                 tt.fields.maxTermsCount,
				maxRegexLength:                tt.fields.maxRegexLength,
				maxAdjacencyMatrixFilters:     tt.fields.maxAdjacencyMatrixFilters,
				maxSlicesPerScroll:            tt.fields.maxSlicesPerScroll,
				autoExpandReplicas:            tt.fields.autoExpandReplicas,
				defaultPipeline:               tt.fields.defaultPipeline,
				finalPipeline:                 tt.fields.finalPipeline,
				codec:                         tt.fields.codec,
				refreshInterval:               tt.fields.refreshInterval,
				gcDeletes:                     tt.fields.gcDeletes,
				flushAfterMerge:               tt.fields.flushAfterMerge,
				search:                        tt.fields.search,
				analyze:                       tt.fields.analyze,
				highlight:                     tt.fields.highlight,
				routing:                       tt.fields.routing,
				shard:                         tt.fields.shard,
				softDeletes:                   tt.fields.softDeletes,
				unassigned:                    tt.fields.unassigned,
				mapping:                       tt.fields.mapping,
				merge:                         tt.fields.merge,
				write:                         tt.fields.write,
				indexing:                      tt.fields.indexing,
				blocks:                        tt.fields.blocks,
				analysis:                      tt.fields.analysis,
				translog:                      tt.fields.translog,
				similarity:                    tt.fields.similarity,
			}
			got, err := index.Source()
			if (err != nil) != tt.wantErr {
				t.Errorf("Source() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Source() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIndex(t *testing.T) {
	tests := []struct {
		name string
		want *Index
	}{
		{
			want: &Index{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
