package fields

import (
	"reflect"
	"testing"
)

func TestNewNumber(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		args        args
		constructor func(name string) *Number
		want        *Number
	}{
		{
			name:        "long",
			constructor: NewLong,
			args:        args{name: "test"},
			want:        &Number{name: "test", _type: TypeNumberLong},
		},
		{
			name:        "integer",
			constructor: NewInteger,
			args:        args{name: "test"},
			want:        &Number{name: "test", _type: TypeNumberInteger},
		},
		{
			name:        "short",
			constructor: NewShort,
			args:        args{name: "test"},
			want:        &Number{name: "test", _type: TypeNumberShort},
		},
		{
			name:        "byte",
			constructor: NewByte,
			args:        args{name: "test"},
			want:        &Number{name: "test", _type: TypeNumberByte},
		},
		{
			name:        "double",
			constructor: NewDouble,
			args:        args{name: "test"},
			want:        &Number{name: "test", _type: TypeNumberDouble},
		},
		{
			name:        "float",
			constructor: NewFloat,
			args:        args{name: "test"},
			want:        &Number{name: "test", _type: TypeNumberFloat},
		},
		{
			name:        "half_float",
			constructor: NewHalfFloat,
			args:        args{name: "test"},
			want:        &Number{name: "test", _type: TypeNumberHalfFloat},
		},
		{
			name:        "scaled_float",
			constructor: NewScaledFloat,
			args:        args{name: "test"},
			want:        &Number{name: "test", _type: TypeNumberScaledFloat},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.constructor(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_GetType(t *testing.T) {
	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "long",
			fields: fields{_type: TypeNumberLong},
			want:   TypeNumberLong,
		},
		{
			name:   "integer",
			fields: fields{_type: TypeNumberInteger},
			want:   TypeNumberInteger,
		},
		{
			name:   "short",
			fields: fields{_type: TypeNumberShort},
			want:   TypeNumberShort,
		},
		{
			name:   "byte",
			fields: fields{_type: TypeNumberByte},
			want:   TypeNumberByte,
		},
		{
			name:   "double",
			fields: fields{_type: TypeNumberDouble},
			want:   TypeNumberDouble,
		},
		{
			name:   "float",
			fields: fields{_type: TypeNumberFloat},
			want:   TypeNumberFloat,
		},
		{
			name:   "half_float",
			fields: fields{_type: TypeNumberHalfFloat},
			want:   TypeNumberHalfFloat,
		},
		{
			name:   "scaled_float",
			fields: fields{_type: TypeNumberScaledFloat},
			want:   TypeNumberScaledFloat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_Meta(t *testing.T) {
	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	type args struct {
		metas []*Meta
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Number
	}{
		{
			name:   "empty",
			fields: fields{},
			args:   args{},
			want:   &Number{},
		},
		{
			name:   "set meta",
			fields: fields{},
			args: args{metas: []*Meta{
				{
					name:  "version",
					value: "0.0.1",
				},
				{
					name:  "boost",
					value: 3.5,
				},
			}},
			want: &Number{meta: []*Meta{
				{
					name:  "version",
					value: "0.0.1",
				},
				{
					name:  "boost",
					value: 3.5,
				},
			}},
		},
		{
			name: "add meta",
			fields: fields{meta: []*Meta{
				{
					name:  "author",
					value: "diez37",
				},
				{
					name:  "year",
					value: 2020,
				},
			}},
			args: args{metas: []*Meta{
				{
					name:  "version",
					value: "0.0.1",
				},
				{
					name:  "boost",
					value: 3.5,
				},
			}},
			want: &Number{meta: []*Meta{
				{
					name:  "author",
					value: "diez37",
				},
				{
					name:  "year",
					value: 2020,
				},
				{
					name:  "version",
					value: "0.0.1",
				},
				{
					name:  "boost",
					value: 3.5,
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.Meta(tt.args.metas...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Meta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_SetBoost(t *testing.T) {
	initBoost := 9.0
	setBoost := 5.4

	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	type args struct {
		boost float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Number
	}{
		{
			name:   "set boost",
			fields: fields{},
			args:   args{boost: setBoost},
			want:   &Number{boost: &setBoost},
		},
		{
			name:   "change boost",
			fields: fields{boost: &initBoost},
			args:   args{boost: setBoost},
			want:   &Number{boost: &setBoost},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.SetBoost(tt.args.boost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBoost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_SetCoerce(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	type args struct {
		coerce bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Number
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{coerce: true},
			want:   &Number{coerce: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{coerce: false},
			want:   &Number{coerce: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{coerce: &boolTrue},
			args:   args{coerce: false},
			want:   &Number{coerce: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.SetCoerce(tt.args.coerce); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCoerce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_SetDocValues(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	type args struct {
		docValues bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Number
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{docValues: true},
			want:   &Number{docValues: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{docValues: false},
			want:   &Number{docValues: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{docValues: &boolTrue},
			args:   args{docValues: false},
			want:   &Number{docValues: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.SetDocValues(tt.args.docValues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDocValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_SetIgnoreMalformed(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	type args struct {
		ignoreMalformed bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Number
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{ignoreMalformed: true},
			want:   &Number{ignoreMalformed: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{ignoreMalformed: false},
			want:   &Number{ignoreMalformed: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{ignoreMalformed: &boolTrue},
			args:   args{ignoreMalformed: false},
			want:   &Number{ignoreMalformed: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.SetIgnoreMalformed(tt.args.ignoreMalformed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIgnoreMalformed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_SetIndex(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	type args struct {
		index bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Number
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{index: true},
			want:   &Number{index: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{index: false},
			want:   &Number{index: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{index: &boolTrue},
			args:   args{index: false},
			want:   &Number{index: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_SetNullValue(t *testing.T) {
	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	type args struct {
		nullValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Number
	}{
		{
			name:   "empty",
			fields: fields{},
			args:   args{},
			want:   &Number{nullValue: nil},
		},
		{
			name:   "int",
			fields: fields{},
			args:   args{nullValue: 1},
			want:   &Number{nullValue: 1},
		},
		{
			name:   "float",
			fields: fields{},
			args:   args{nullValue: 4.5},
			want:   &Number{nullValue: 4.5},
		},
		{
			name:   "string",
			fields: fields{},
			args:   args{nullValue: "test"},
			want:   &Number{nullValue: "test"},
		},
		{
			name:   "bool",
			fields: fields{},
			args:   args{nullValue: true},
			want:   &Number{nullValue: true},
		},
		{
			name:   "change",
			fields: fields{nullValue: "test"},
			args:   args{nullValue: true},
			want:   &Number{nullValue: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.SetNullValue(tt.args.nullValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNullValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_SetScalingFactor(t *testing.T) {
	initValue := 9
	setValue := 5

	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	type args struct {
		scalingFactor int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Number
	}{
		{
			name:   "set boost",
			fields: fields{},
			args:   args{scalingFactor: setValue},
			want:   &Number{scalingFactor: &setValue},
		},
		{
			name:   "change boost",
			fields: fields{scalingFactor: &initValue},
			args:   args{scalingFactor: setValue},
			want:   &Number{scalingFactor: &setValue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.SetScalingFactor(tt.args.scalingFactor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScalingFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_SetStore(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	type args struct {
		store bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Number
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{store: true},
			want:   &Number{store: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{store: false},
			want:   &Number{store: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{store: &boolTrue},
			args:   args{store: false},
			want:   &Number{store: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_Source(t *testing.T) {
	scalingFactor := 9
	boost := 9.0
	boolTrue := true

	type fields struct {
		coerce          *bool
		docValues       *bool
		ignoreMalformed *bool
		index           *bool
		store           *bool
		boost           *float64
		scalingFactor   *int
		nullValue       interface{}
		meta            []*Meta
		_type           Type
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name:   "long",
			fields: fields{_type: TypeNumberLong},
			want: map[string]interface{}{
				"type": TypeNumberLong,
			},
			wantErr: false,
		},
		{
			name:   "integer",
			fields: fields{_type: TypeNumberInteger},
			want: map[string]interface{}{
				"type": TypeNumberInteger,
			},
			wantErr: false,
		},
		{
			name:   "short",
			fields: fields{_type: TypeNumberShort},
			want: map[string]interface{}{
				"type": TypeNumberShort,
			},
			wantErr: false,
		},
		{
			name:   "byte",
			fields: fields{_type: TypeNumberByte},
			want: map[string]interface{}{
				"type": TypeNumberByte,
			},
			wantErr: false,
		},
		{
			name:   "double",
			fields: fields{_type: TypeNumberDouble},
			want: map[string]interface{}{
				"type": TypeNumberDouble,
			},
			wantErr: false,
		},
		{
			name:   "float",
			fields: fields{_type: TypeNumberFloat},
			want: map[string]interface{}{
				"type": TypeNumberFloat,
			},
			wantErr: false,
		},
		{
			name:   "halfFloat",
			fields: fields{_type: TypeNumberHalfFloat},
			want: map[string]interface{}{
				"type": TypeNumberHalfFloat,
			},
			wantErr: false,
		},
		{
			name:   "scaledFloat",
			fields: fields{_type: TypeNumberScaledFloat, scalingFactor: &scalingFactor},
			want: map[string]interface{}{
				"type":           TypeNumberScaledFloat,
				"scaling_factor": scalingFactor,
			},
			wantErr: false,
		},
		{
			name:    "scaledFloatError",
			fields:  fields{_type: TypeNumberScaledFloat},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "coerce",
			fields: fields{coerce: &boolTrue},
			want: map[string]interface{}{
				"type":   Type(""),
				"coerce": true,
			},
			wantErr: false,
		},
		{
			name:   "docValues",
			fields: fields{docValues: &boolTrue},
			want: map[string]interface{}{
				"type":       Type(""),
				"doc_values": true,
			},
			wantErr: false,
		},
		{
			name:   "ignoreMalformed",
			fields: fields{ignoreMalformed: &boolTrue},
			want: map[string]interface{}{
				"type":             Type(""),
				"ignore_malformed": true,
			},
			wantErr: false,
		},
		{
			name:   "index",
			fields: fields{index: &boolTrue},
			want: map[string]interface{}{
				"type":  Type(""),
				"index": true,
			},
			wantErr: false,
		},
		{
			name:   "store",
			fields: fields{store: &boolTrue},
			want: map[string]interface{}{
				"type":  Type(""),
				"store": true,
			},
			wantErr: false,
		},
		{
			name:   "boost",
			fields: fields{boost: &boost},
			want: map[string]interface{}{
				"type":  Type(""),
				"boost": boost,
			},
			wantErr: false,
		},
		{
			name:    "scalingFactorError",
			fields:  fields{scalingFactor: &scalingFactor},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "scalingFactor",
			fields: fields{scalingFactor: &scalingFactor, _type: TypeNumberScaledFloat},
			want: map[string]interface{}{
				"type":           TypeNumberScaledFloat,
				"scaling_factor": scalingFactor,
			},
			wantErr: false,
		},
		{
			name:   "nullValue",
			fields: fields{nullValue: scalingFactor},
			want: map[string]interface{}{
				"type":       Type(""),
				"null_value": scalingFactor,
			},
			wantErr: false,
		},
		{
			name: "meta",
			fields: fields{meta: []*Meta{
				{
					name:  "author",
					value: "diez37",
				},
				{
					name:  "version",
					value: "0.0.1",
				},
			}},
			want: map[string]interface{}{
				"type": Type(""),
				"meta": map[string]interface{}{
					"author":  "diez37",
					"version": "0.0.1",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Number{
				coerce:          tt.fields.coerce,
				docValues:       tt.fields.docValues,
				ignoreMalformed: tt.fields.ignoreMalformed,
				index:           tt.fields.index,
				store:           tt.fields.store,
				boost:           tt.fields.boost,
				scalingFactor:   tt.fields.scalingFactor,
				nullValue:       tt.fields.nullValue,
				meta:            tt.fields.meta,
				_type:           tt.fields._type,
			}
			got, err := field.Source()
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
