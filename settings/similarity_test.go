package settings

import (
	"github.com/olivere/elastic/v7"
	"reflect"
	"testing"
)

func TestNewSimilarityBM25(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SimilarityBM25
	}{
		{
			args: args{name: "test"},
			want: &SimilarityBM25{name: SimilarityName("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityBM25(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityBM25() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimilarityBoolean(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SimilarityDefault
	}{
		{
			args: args{name: "test"},
			want: &SimilarityDefault{name: SimilarityName("test"), _type: SimilarityTypeBoolean},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityBoolean(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityBoolean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimilarityClassic(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SimilarityDefault
	}{
		{
			args: args{name: "test"},
			want: &SimilarityDefault{name: SimilarityName("test"), _type: SimilarityTypeClassic},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityClassic(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityClassic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimilarityDFI(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SimilarityDFI
	}{
		{
			args: args{name: "test"},
			want: &SimilarityDFI{name: SimilarityName("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityDFI(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityDFI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimilarityDFR(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SimilarityDFR
	}{
		{
			args: args{name: "test"},
			want: &SimilarityDFR{name: SimilarityName("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityDFR(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityDFR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimilarityIB(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SimilarityIB
	}{
		{
			args: args{name: "test"},
			want: &SimilarityIB{name: SimilarityName("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityIB(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityIB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimilarityLMDirichlet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SimilarityLMDirichlet
	}{
		{
			args: args{name: "test"},
			want: &SimilarityLMDirichlet{name: SimilarityName("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityLMDirichlet(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityLMDirichlet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimilarityLMJelinekMercer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SimilarityLMJelinekMercer
	}{
		{
			args: args{name: "test"},
			want: &SimilarityLMJelinekMercer{name: SimilarityName("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityLMJelinekMercer(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityLMJelinekMercer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimilarityScript(t *testing.T) {
	type args struct {
		script *elastic.Script
	}
	tests := []struct {
		name string
		args args
		want *SimilarityScript
	}{
		{
			args: args{script: elastic.NewScript("test")},
			want: &SimilarityScript{script: elastic.NewScript("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityScript(tt.args.script); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimilarityScripted(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SimilarityScripted
	}{
		{
			args: args{name: "test"},
			want: &SimilarityScripted{name: SimilarityName("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimilarityScripted(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimilarityScripted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityBM25_Name(t *testing.T) {
	type fields struct {
		name             SimilarityName
		discountOverlaps *bool
		b                *float32
		k1               *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityName
	}{
		{
			fields: fields{name: SimilarityName("test")},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityBM25{
				name:             tt.fields.name,
				discountOverlaps: tt.fields.discountOverlaps,
				b:                tt.fields.b,
				k1:               tt.fields.k1,
			}
			if got := similarity.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityBM25_SetB(t *testing.T) {
	numberSet := float32(1)
	numberChange := float32(2)

	type fields struct {
		name             SimilarityName
		discountOverlaps *bool
		b                *float32
		k1               *float32
	}
	type args struct {
		b float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityBM25
	}{
		{
			name: "set",
			args: args{b: 1},
			want: &SimilarityBM25{b: &numberSet},
		},
		{
			name: "change",
			fields: fields{b: &numberSet},
			args: args{b: 2},
			want: &SimilarityBM25{b: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityBM25{
				name:             tt.fields.name,
				discountOverlaps: tt.fields.discountOverlaps,
				b:                tt.fields.b,
				k1:               tt.fields.k1,
			}
			if got := similarity.SetB(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityBM25_SetDiscountOverlaps(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name             SimilarityName
		discountOverlaps *bool
		b                *float32
		k1               *float32
	}
	type args struct {
		discountOverlaps bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityBM25
	}{
		{
			name: "set",
			args: args{discountOverlaps: true},
			want: &SimilarityBM25{discountOverlaps: &boolTrue},
		},
		{
			name: "change",
			fields: fields{discountOverlaps: &boolTrue},
			args: args{discountOverlaps: false},
			want: &SimilarityBM25{discountOverlaps: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityBM25{
				name:             tt.fields.name,
				discountOverlaps: tt.fields.discountOverlaps,
				b:                tt.fields.b,
				k1:               tt.fields.k1,
			}
			if got := similarity.SetDiscountOverlaps(tt.args.discountOverlaps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDiscountOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityBM25_SetK1(t *testing.T) {
	numberSet := float32(1)
	numberChange := float32(2)

	type fields struct {
		name             SimilarityName
		discountOverlaps *bool
		b                *float32
		k1               *float32
	}
	type args struct {
		k1 float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityBM25
	}{
		{
			name: "set",
			args: args{k1: 1},
			want: &SimilarityBM25{k1: &numberSet},
		},
		{
			name: "change",
			fields: fields{k1: &numberSet},
			args: args{k1: 2},
			want: &SimilarityBM25{k1: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityBM25{
				name:             tt.fields.name,
				discountOverlaps: tt.fields.discountOverlaps,
				b:                tt.fields.b,
				k1:               tt.fields.k1,
			}
			if got := similarity.SetK1(tt.args.k1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetK1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityBM25_Source(t *testing.T) {
	boolTrue := true
	number := float32(2.5)

	type fields struct {
		name             SimilarityName
		discountOverlaps *bool
		b                *float32
		k1               *float32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": SimilarityTypeBM25,
			},
		},
		{
			name: "discount_overlaps",
			fields: fields{discountOverlaps: &boolTrue},
			want: map[string]interface{}{
				"type": SimilarityTypeBM25,
				"discount_overlaps": true,
			},
		},
		{
			name: "b",
			fields: fields{b: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeBM25,
				"b": float32(2.5),
			},
		},
		{
			name: "k1",
			fields: fields{k1: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeBM25,
				"k1": float32(2.5),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityBM25{
				name:             tt.fields.name,
				discountOverlaps: tt.fields.discountOverlaps,
				b:                tt.fields.b,
				k1:               tt.fields.k1,
			}
			got, err := similarity.Source()
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

func TestSimilarityBM25_Type(t *testing.T) {
	type fields struct {
		name             SimilarityName
		discountOverlaps *bool
		b                *float32
		k1               *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityType
	}{
		{
			want: SimilarityTypeBM25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityBM25{
				name:             tt.fields.name,
				discountOverlaps: tt.fields.discountOverlaps,
				b:                tt.fields.b,
				k1:               tt.fields.k1,
			}
			if got := similarity.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDFI_Name(t *testing.T) {
	type fields struct {
		name                SimilarityName
		independenceMeasure *SimilarityDFIIndependenceMeasureType
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFI{
				name:                tt.fields.name,
				independenceMeasure: tt.fields.independenceMeasure,
			}
			if got := similarity.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDFI_SetIndependenceMeasure(t *testing.T) {
	dFIIndependenceMeasureStandardized := DFIIndependenceMeasureStandardized
	dFIIndependenceMeasureSaturated := DFIIndependenceMeasureSaturated
	dFIIndependenceMeasureChisquared := DFIIndependenceMeasureChisquared

	type fields struct {
		name                SimilarityName
		independenceMeasure *SimilarityDFIIndependenceMeasureType
	}
	type args struct {
		independenceMeasure SimilarityDFIIndependenceMeasureType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityDFI
	}{
		{
			name: "dfi_independence_measure_standardized",
			args: args{independenceMeasure: DFIIndependenceMeasureStandardized},
			want: &SimilarityDFI{independenceMeasure: &dFIIndependenceMeasureStandardized},
		},
		{
			name: "dfi_independence_measure_saturated",
			args: args{independenceMeasure: DFIIndependenceMeasureSaturated},
			want: &SimilarityDFI{independenceMeasure: &dFIIndependenceMeasureSaturated},
		},
		{
			name: "dfi_independence_measure_chisquared",
			args: args{independenceMeasure: DFIIndependenceMeasureChisquared},
			want: &SimilarityDFI{independenceMeasure: &dFIIndependenceMeasureChisquared},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFI{
				name:                tt.fields.name,
				independenceMeasure: tt.fields.independenceMeasure,
			}
			if got := similarity.SetIndependenceMeasure(tt.args.independenceMeasure); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndependenceMeasure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDFI_Source(t *testing.T) {
	dFIIndependenceMeasureStandardized := DFIIndependenceMeasureStandardized
	dFIIndependenceMeasureSaturated := DFIIndependenceMeasureSaturated
	dFIIndependenceMeasureChisquared := DFIIndependenceMeasureChisquared

	type fields struct {
		name                SimilarityName
		independenceMeasure *SimilarityDFIIndependenceMeasureType
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": SimilarityTypeDFI,
			},
		},
		{
			name: "dfi_independence_measure_standardized",
			fields: fields{independenceMeasure: &dFIIndependenceMeasureStandardized},
			want: map[string]interface{}{
				"type": SimilarityTypeDFI,
				"independence_measure": DFIIndependenceMeasureStandardized,
			},
		},
		{
			name: "dfi_independence_measure_saturated",
			fields: fields{independenceMeasure: &dFIIndependenceMeasureSaturated},
			want: map[string]interface{}{
				"type": SimilarityTypeDFI,
				"independence_measure": DFIIndependenceMeasureSaturated,
			},
		},
		{
			name: "dfi_independence_measure_chisquared",
			fields: fields{independenceMeasure: &dFIIndependenceMeasureChisquared},
			want: map[string]interface{}{
				"type": SimilarityTypeDFI,
				"independence_measure": DFIIndependenceMeasureChisquared,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFI{
				name:                tt.fields.name,
				independenceMeasure: tt.fields.independenceMeasure,
			}
			got, err := similarity.Source()
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

func TestSimilarityDFI_Type(t *testing.T) {
	type fields struct {
		name                SimilarityName
		independenceMeasure *SimilarityDFIIndependenceMeasureType
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityType
	}{
		{
			want: SimilarityTypeDFI,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFI{
				name:                tt.fields.name,
				independenceMeasure: tt.fields.independenceMeasure,
			}
			if got := similarity.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDFR_Name(t *testing.T) {
	type fields struct {
		name               SimilarityName
		basicModel         *SimilarityDFRBasicModelType
		afterEffect        *SimilarityDFRAfterEffectType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFR{
				name:               tt.fields.name,
				basicModel:         tt.fields.basicModel,
				afterEffect:        tt.fields.afterEffect,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDFR_SetAfterEffect(t *testing.T) {
	dFRAfterEffectB := DFRAfterEffectB
	dFRAfterEffectL := DFRAfterEffectL

	type fields struct {
		name               SimilarityName
		basicModel         *SimilarityDFRBasicModelType
		afterEffect        *SimilarityDFRAfterEffectType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	type args struct {
		afterEffect SimilarityDFRAfterEffectType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityDFR
	}{
		{
			name: "dfr_after_effect_b",
			args: args{afterEffect: DFRAfterEffectB},
			want: &SimilarityDFR{afterEffect: &dFRAfterEffectB},
		},
		{
			name: "dfr_after_effect_l",
			args: args{afterEffect: DFRAfterEffectL},
			want: &SimilarityDFR{afterEffect: &dFRAfterEffectL},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFR{
				name:               tt.fields.name,
				basicModel:         tt.fields.basicModel,
				afterEffect:        tt.fields.afterEffect,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.SetAfterEffect(tt.args.afterEffect); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAfterEffect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDFR_SetBasicModel(t *testing.T) {
	dFRBasicModelG := DFRBasicModelG
	dFRBasicModelIf := DFRBasicModelIf
	dFRBasicModelIn := DFRBasicModelIn
	dFRBasicModelIne := DFRBasicModelIne

	type fields struct {
		name               SimilarityName
		basicModel         *SimilarityDFRBasicModelType
		afterEffect        *SimilarityDFRAfterEffectType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	type args struct {
		basicModel SimilarityDFRBasicModelType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityDFR
	}{
		{
			name: "dfr_basic_model_g",
			args: args{basicModel: DFRBasicModelG},
			want: &SimilarityDFR{basicModel: &dFRBasicModelG},
		},
		{
			name: "dfr_basic_model_if",
			args: args{basicModel: DFRBasicModelIf},
			want: &SimilarityDFR{basicModel: &dFRBasicModelIf},
		},
		{
			name: "dfr_basic_model_in",
			args: args{basicModel: DFRBasicModelIn},
			want: &SimilarityDFR{basicModel: &dFRBasicModelIn},
		},
		{
			name: "dfr_basic_model_ine",
			args: args{basicModel: DFRBasicModelIne},
			want: &SimilarityDFR{basicModel: &dFRBasicModelIne},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFR{
				name:               tt.fields.name,
				basicModel:         tt.fields.basicModel,
				afterEffect:        tt.fields.afterEffect,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.SetBasicModel(tt.args.basicModel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBasicModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDFR_SetNormalization(t *testing.T) {
	dFRNormalizationNo := DFRNormalizationNo
	dFRNormalizationH1 := DFRNormalizationH1
	dFRNormalizationH2 := DFRNormalizationH2
	dFRNormalizationH3 := DFRNormalizationH3
	dFRNormalizationZ := DFRNormalizationZ
	number := float32(2.5)

	type fields struct {
		name               SimilarityName
		basicModel         *SimilarityDFRBasicModelType
		afterEffect        *SimilarityDFRAfterEffectType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	type args struct {
		normalization SimilarityDFRNormalizationType
		value         float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityDFR
	}{
		{
			name: "dfr_normalization_no",
			args: args{normalization: DFRNormalizationNo, value: 2.5},
			want: &SimilarityDFR{normalization: &dFRNormalizationNo, normalizationValue: &number},
		},
		{
			name: "dfr_normalization_h1",
			args: args{normalization: DFRNormalizationH1, value: 2.5},
			want: &SimilarityDFR{normalization: &dFRNormalizationH1, normalizationValue: &number},
		},
		{
			name: "dfr_normalization_h2",
			args: args{normalization: DFRNormalizationH2, value: 2.5},
			want: &SimilarityDFR{normalization: &dFRNormalizationH2, normalizationValue: &number},
		},
		{
			name: "dfr_normalization_h3",
			args: args{normalization: DFRNormalizationH3, value: 2.5},
			want: &SimilarityDFR{normalization: &dFRNormalizationH3, normalizationValue: &number},
		},
		{
			name: "dfr_normalization_z",
			args: args{normalization: DFRNormalizationZ, value: 2.5},
			want: &SimilarityDFR{normalization: &dFRNormalizationZ, normalizationValue: &number},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFR{
				name:               tt.fields.name,
				basicModel:         tt.fields.basicModel,
				afterEffect:        tt.fields.afterEffect,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.SetNormalization(tt.args.normalization, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNormalization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDFR_Source(t *testing.T) {
	dFRNormalizationNo := DFRNormalizationNo
	dFRNormalizationH1 := DFRNormalizationH1
	dFRNormalizationH2 := DFRNormalizationH2
	dFRNormalizationH3 := DFRNormalizationH3
	dFRNormalizationZ := DFRNormalizationZ
	dFRBasicModelG := DFRBasicModelG
	dFRAfterEffectB := DFRAfterEffectB
	number := float32(2.5)

	type fields struct {
		name               SimilarityName
		basicModel         *SimilarityDFRBasicModelType
		afterEffect        *SimilarityDFRAfterEffectType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": SimilarityTypeDFR,
			},
		},
		{
			name: "basic_model",
			fields: fields{basicModel: &dFRBasicModelG},
			want: map[string]interface{}{
				"type": SimilarityTypeDFR,
				"basic_model": DFRBasicModelG,
			},
		},
		{
			name: "after_effect",
			fields: fields{afterEffect: &dFRAfterEffectB},
			want: map[string]interface{}{
				"type": SimilarityTypeDFR,
				"after_effect": DFRAfterEffectB,
			},
		},
		{
			name: "dfr_normalization_no",
			fields: fields{normalization: &dFRNormalizationNo},
			want: map[string]interface{}{
				"type": SimilarityTypeDFR,
				"normalization": DFRNormalizationNo,
			},
		},
		{
			name: "dfr_normalization_h1",
			fields: fields{normalization: &dFRNormalizationH1, normalizationValue: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeDFR,
				"normalization": DFRNormalizationH1,
				"normalization.h1.c": float32(2.5),
			},
		},
		{
			name: "dfr_normalization_h2",
			fields: fields{normalization: &dFRNormalizationH2, normalizationValue: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeDFR,
				"normalization": DFRNormalizationH2,
				"normalization.h2.c": float32(2.5),
			},
		},
		{
			name: "dfr_normalization_h3",
			fields: fields{normalization: &dFRNormalizationH3, normalizationValue: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeDFR,
				"normalization": DFRNormalizationH3,
				"normalization.h3.c": float32(2.5),
			},
		},
		{
			name: "dfr_normalization_z",
			fields: fields{normalization: &dFRNormalizationZ, normalizationValue: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeDFR,
				"normalization": DFRNormalizationZ,
				"normalization.z.c": float32(2.5),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFR{
				name:               tt.fields.name,
				basicModel:         tt.fields.basicModel,
				afterEffect:        tt.fields.afterEffect,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			got, err := similarity.Source()
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

func TestSimilarityDFR_Type(t *testing.T) {
	type fields struct {
		name               SimilarityName
		basicModel         *SimilarityDFRBasicModelType
		afterEffect        *SimilarityDFRAfterEffectType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityType
	}{
		{
			want: SimilarityTypeDFR,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDFR{
				name:               tt.fields.name,
				basicModel:         tt.fields.basicModel,
				afterEffect:        tt.fields.afterEffect,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDefault_Name(t *testing.T) {
	type fields struct {
		name  SimilarityName
		_type SimilarityType
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDefault{
				name:  tt.fields.name,
				_type: tt.fields._type,
			}
			if got := similarity.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityDefault_Source(t *testing.T) {
	type fields struct {
		name  SimilarityName
		_type SimilarityType
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "similarity_type_boolean",
			fields: fields{_type: SimilarityTypeBoolean},
			want: map[string]interface{}{
				"type": SimilarityTypeBoolean,
			},
		},
		{
			name: "similarity_type_classic",
			fields: fields{_type: SimilarityTypeClassic},
			want: map[string]interface{}{
				"type": SimilarityTypeClassic,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDefault{
				name:  tt.fields.name,
				_type: tt.fields._type,
			}
			got, err := similarity.Source()
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

func TestSimilarityDefault_Type(t *testing.T) {
	type fields struct {
		name  SimilarityName
		_type SimilarityType
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityType
	}{
		{
			name: "similarity_type_boolean",
			fields: fields{_type: SimilarityTypeBoolean},
			want: SimilarityTypeBoolean,
		},
		{
			name: "similarity_type_classic",
			fields: fields{_type: SimilarityTypeClassic},
			want: SimilarityTypeClassic,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityDefault{
				name:  tt.fields.name,
				_type: tt.fields._type,
			}
			if got := similarity.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityIB_Name(t *testing.T) {
	type fields struct {
		name               SimilarityName
		distribution       *SimilarityIBDistributionType
		lambda             *SimilarityIBLambdaType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityIB{
				name:               tt.fields.name,
				distribution:       tt.fields.distribution,
				lambda:             tt.fields.lambda,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityIB_SetDistribution(t *testing.T) {
	iBDistributionLL := IBDistributionLL
	iBDistributionSPL := IBDistributionSPL

	type fields struct {
		name               SimilarityName
		distribution       *SimilarityIBDistributionType
		lambda             *SimilarityIBLambdaType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	type args struct {
		distribution SimilarityIBDistributionType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityIB
	}{
		{
			name: "i_b_distribution_l_l",
			args: args{distribution: IBDistributionLL},
			want: &SimilarityIB{distribution: &iBDistributionLL},
		},
		{
			name: "i_b_distribution_s_p_l",
			args: args{distribution: IBDistributionSPL},
			want: &SimilarityIB{distribution: &iBDistributionSPL},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityIB{
				name:               tt.fields.name,
				distribution:       tt.fields.distribution,
				lambda:             tt.fields.lambda,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.SetDistribution(tt.args.distribution); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDistribution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityIB_SetLambda(t *testing.T) {
	iBLambdaDF := IBLambdaDF
	iBLambdaTTF := IBLambdaTTF

	type fields struct {
		name               SimilarityName
		distribution       *SimilarityIBDistributionType
		lambda             *SimilarityIBLambdaType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	type args struct {
		lambda SimilarityIBLambdaType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityIB
	}{
		{
			name: "i_b_lambda_d_f",
			args: args{lambda: IBLambdaDF},
			want: &SimilarityIB{lambda: &iBLambdaDF},
		},
		{
			name: "i_b_lambda_t_t_f",
			args: args{lambda: IBLambdaTTF},
			want: &SimilarityIB{lambda: &iBLambdaTTF},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityIB{
				name:               tt.fields.name,
				distribution:       tt.fields.distribution,
				lambda:             tt.fields.lambda,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.SetLambda(tt.args.lambda); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLambda() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityIB_SetNormalization(t *testing.T) {
	dFRNormalizationNo := DFRNormalizationNo
	dFRNormalizationH1 := DFRNormalizationH1
	dFRNormalizationH2 := DFRNormalizationH2
	dFRNormalizationH3 := DFRNormalizationH3
	dFRNormalizationZ := DFRNormalizationZ
	number := float32(2.5)

	type fields struct {
		name               SimilarityName
		distribution       *SimilarityIBDistributionType
		lambda             *SimilarityIBLambdaType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	type args struct {
		normalization SimilarityDFRNormalizationType
		value         float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityIB
	}{
		{
			name: "dfr_normalization_no",
			args: args{normalization: DFRNormalizationNo, value: 2.5},
			want: &SimilarityIB{normalization: &dFRNormalizationNo, normalizationValue: &number},
		},
		{
			name: "dfr_normalization_h1",
			args: args{normalization: DFRNormalizationH1, value: 2.5},
			want: &SimilarityIB{normalization: &dFRNormalizationH1, normalizationValue: &number},
		},
		{
			name: "dfr_normalization_h2",
			args: args{normalization: DFRNormalizationH2, value: 2.5},
			want: &SimilarityIB{normalization: &dFRNormalizationH2, normalizationValue: &number},
		},
		{
			name: "dfr_normalization_h3",
			args: args{normalization: DFRNormalizationH3, value: 2.5},
			want: &SimilarityIB{normalization: &dFRNormalizationH3, normalizationValue: &number},
		},
		{
			name: "dfr_normalization_z",
			args: args{normalization: DFRNormalizationZ, value: 2.5},
			want: &SimilarityIB{normalization: &dFRNormalizationZ, normalizationValue: &number},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityIB{
				name:               tt.fields.name,
				distribution:       tt.fields.distribution,
				lambda:             tt.fields.lambda,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.SetNormalization(tt.args.normalization, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNormalization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityIB_Source(t *testing.T) {
	iBDistributionLL := IBDistributionLL
	iBLambdaDF := IBLambdaDF
	dFRNormalizationNo := DFRNormalizationNo
	dFRNormalizationH1 := DFRNormalizationH1
	dFRNormalizationH2 := DFRNormalizationH2
	dFRNormalizationH3 := DFRNormalizationH3
	dFRNormalizationZ := DFRNormalizationZ
	number := float32(2.5)

	type fields struct {
		name               SimilarityName
		distribution       *SimilarityIBDistributionType
		lambda             *SimilarityIBLambdaType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": SimilarityTypeIB,
			},
		},
		{
			name: "distribution",
			fields: fields{distribution: &iBDistributionLL},
			want: map[string]interface{}{
				"type": SimilarityTypeIB,
				"distribution": IBDistributionLL,
			},
		},
		{
			name: "lambda",
			fields: fields{lambda: &iBLambdaDF},
			want: map[string]interface{}{
				"type": SimilarityTypeIB,
				"lambda": IBLambdaDF,
			},
		},
		{
			name: "dfr_normalization_no",
			fields: fields{normalization: &dFRNormalizationNo},
			want: map[string]interface{}{
				"type": SimilarityTypeIB,
				"normalization": DFRNormalizationNo,
			},
		},
		{
			name: "dfr_normalization_h1",
			fields: fields{normalization: &dFRNormalizationH1, normalizationValue: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeIB,
				"normalization": DFRNormalizationH1,
				"normalization.h1.c": float32(2.5),
			},
		},
		{
			name: "dfr_normalization_h2",
			fields: fields{normalization: &dFRNormalizationH2, normalizationValue: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeIB,
				"normalization": DFRNormalizationH2,
				"normalization.h2.c": float32(2.5),
			},
		},
		{
			name: "dfr_normalization_h3",
			fields: fields{normalization: &dFRNormalizationH3, normalizationValue: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeIB,
				"normalization": DFRNormalizationH3,
				"normalization.h3.c": float32(2.5),
			},
		},
		{
			name: "dfr_normalization_z",
			fields: fields{normalization: &dFRNormalizationZ, normalizationValue: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeIB,
				"normalization": DFRNormalizationZ,
				"normalization.z.c": float32(2.5),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityIB{
				name:               tt.fields.name,
				distribution:       tt.fields.distribution,
				lambda:             tt.fields.lambda,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			got, err := similarity.Source()
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

func TestSimilarityIB_Type(t *testing.T) {
	type fields struct {
		name               SimilarityName
		distribution       *SimilarityIBDistributionType
		lambda             *SimilarityIBLambdaType
		normalization      *SimilarityDFRNormalizationType
		normalizationValue *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityType
	}{
		{
			want: SimilarityTypeIB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityIB{
				name:               tt.fields.name,
				distribution:       tt.fields.distribution,
				lambda:             tt.fields.lambda,
				normalization:      tt.fields.normalization,
				normalizationValue: tt.fields.normalizationValue,
			}
			if got := similarity.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityLMDirichlet_Name(t *testing.T) {
	type fields struct {
		name SimilarityName
		mu   *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityLMDirichlet{
				name: tt.fields.name,
				mu:   tt.fields.mu,
			}
			if got := similarity.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityLMDirichlet_SetMU(t *testing.T) {
	numberSet := float32(1)
	numberChange := float32(2)

	type fields struct {
		name SimilarityName
		mu   *float32
	}
	type args struct {
		mu float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityLMDirichlet
	}{
		{
			name: "set",
			args: args{mu: 1},
			want: &SimilarityLMDirichlet{mu: &numberSet},
		},
		{
			name: "change",
			fields: fields{mu: &numberSet},
			args: args{mu: 2},
			want: &SimilarityLMDirichlet{mu: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityLMDirichlet{
				name: tt.fields.name,
				mu:   tt.fields.mu,
			}
			if got := similarity.SetMU(tt.args.mu); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityLMDirichlet_Source(t *testing.T) {
	number := float32(2.5)

	type fields struct {
		name SimilarityName
		mu   *float32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": SimilarityTypeLMDirichlet,
			},
		},
		{
			name: "mu",
			fields: fields{mu: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeLMDirichlet,
				"mu": float32(2.5),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityLMDirichlet{
				name: tt.fields.name,
				mu:   tt.fields.mu,
			}
			got, err := similarity.Source()
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

func TestSimilarityLMDirichlet_Type(t *testing.T) {
	type fields struct {
		name SimilarityName
		mu   *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityType
	}{
		{
			want: SimilarityTypeLMDirichlet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityLMDirichlet{
				name: tt.fields.name,
				mu:   tt.fields.mu,
			}
			if got := similarity.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityLMJelinekMercer_Name(t *testing.T) {
	type fields struct {
		name   SimilarityName
		lambda *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityLMJelinekMercer{
				name:   tt.fields.name,
				lambda: tt.fields.lambda,
			}
			if got := similarity.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityLMJelinekMercer_SetLambda(t *testing.T) {
	numberSet := float32(1)
	numberChange := float32(2)

	type fields struct {
		name   SimilarityName
		lambda *float32
	}
	type args struct {
		lambda float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityLMJelinekMercer
	}{
		{
			name: "set",
			args: args{lambda: 1},
			want: &SimilarityLMJelinekMercer{lambda: &numberSet},
		},
		{
			name: "change",
			fields: fields{lambda: &numberSet},
			args: args{lambda: 2},
			want: &SimilarityLMJelinekMercer{lambda: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityLMJelinekMercer{
				name:   tt.fields.name,
				lambda: tt.fields.lambda,
			}
			if got := similarity.SetLambda(tt.args.lambda); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLambda() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityLMJelinekMercer_Source(t *testing.T) {
	number := float32(2.5)

	type fields struct {
		name   SimilarityName
		lambda *float32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": SimilarityTypeLMJelinekMercer,
			},
		},
		{
			name: "lambda",
			fields: fields{lambda: &number},
			want: map[string]interface{}{
				"type": SimilarityTypeLMJelinekMercer,
				"lambda": float32(2.5),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityLMJelinekMercer{
				name:   tt.fields.name,
				lambda: tt.fields.lambda,
			}
			got, err := similarity.Source()
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

func TestSimilarityLMJelinekMercer_Type(t *testing.T) {
	type fields struct {
		name   SimilarityName
		lambda *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityType
	}{
		{
			want: SimilarityTypeLMJelinekMercer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityLMJelinekMercer{
				name:   tt.fields.name,
				lambda: tt.fields.lambda,
			}
			if got := similarity.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityScript_Source(t *testing.T) {
	source, _ := elastic.NewScript("test").Source()

	type fields struct {
		script *elastic.Script
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{script: elastic.NewScript("test")},
			want: source,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			script := &SimilarityScript{
				script: tt.fields.script,
			}
			got, err := script.Source()
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

func TestSimilarityScripted_Name(t *testing.T) {
	type fields struct {
		name         SimilarityName
		weightScript *SimilarityScript
		script       *SimilarityScript
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityScripted{
				name:         tt.fields.name,
				weightScript: tt.fields.weightScript,
				script:       tt.fields.script,
			}
			if got := similarity.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityScripted_SetScript(t *testing.T) {
	type fields struct {
		name         SimilarityName
		weightScript *SimilarityScript
		script       *SimilarityScript
	}
	type args struct {
		script *SimilarityScript
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityScripted
	}{
		{
			args: args{script: &SimilarityScript{}},
			want: &SimilarityScripted{script: &SimilarityScript{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityScripted{
				name:         tt.fields.name,
				weightScript: tt.fields.weightScript,
				script:       tt.fields.script,
			}
			if got := similarity.SetScript(tt.args.script); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityScripted_SetWeightScript(t *testing.T) {
	type fields struct {
		name         SimilarityName
		weightScript *SimilarityScript
		script       *SimilarityScript
	}
	type args struct {
		script *SimilarityScript
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SimilarityScripted
	}{
		{
			args: args{script: &SimilarityScript{}},
			want: &SimilarityScripted{weightScript: &SimilarityScript{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityScripted{
				name:         tt.fields.name,
				weightScript: tt.fields.weightScript,
				script:       tt.fields.script,
			}
			if got := similarity.SetWeightScript(tt.args.script); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWeightScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityScripted_Source(t *testing.T) {
	scriptSrc, _ := (&SimilarityScript{script: elastic.NewScript("test")}).Source()

	type fields struct {
		name         SimilarityName
		weightScript *SimilarityScript
		script       *SimilarityScript
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": SimilarityTypeScripted,
			},
		},
		{
			name: "weight_script",
			fields: fields{weightScript: &SimilarityScript{script: elastic.NewScript("test")}},
			want: map[string]interface{}{
				"type":          SimilarityTypeScripted,
				"weight_script": scriptSrc,
			},
		},
		{
			name: "script",
			fields: fields{script: &SimilarityScript{script: elastic.NewScript("test")}},
			want: map[string]interface{}{
				"type":   SimilarityTypeScripted,
				"script": scriptSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityScripted{
				name:         tt.fields.name,
				weightScript: tt.fields.weightScript,
				script:       tt.fields.script,
			}
			got, err := similarity.Source()
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

func TestSimilarityScripted_Type(t *testing.T) {
	type fields struct {
		name         SimilarityName
		weightScript *SimilarityScript
		script       *SimilarityScript
	}
	tests := []struct {
		name   string
		fields fields
		want   SimilarityType
	}{
		{
			want: SimilarityTypeScripted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := &SimilarityScripted{
				name:         tt.fields.name,
				weightScript: tt.fields.weightScript,
				script:       tt.fields.script,
			}
			if got := similarity.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
