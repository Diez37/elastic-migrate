package fields

import (
	"reflect"
	"testing"
)

func TestNewRange(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		want        *Range
		args        args
		constructor func(name string) *Range
	}{
		{
			name:        "integer",
			want:        &Range{_type: TypeRangeInteger, name: "test"},
			args:        args{name: "test"},
			constructor: NewRangeInteger,
		},
		{
			name:        "float",
			args:        args{name: "test"},
			want:        &Range{_type: TypeRangeFloat, name: "test"},
			constructor: NewRangeFloat,
		},
		{
			name:        "long",
			args:        args{name: "test"},
			want:        &Range{_type: TypeRangeLong, name: "test"},
			constructor: NewRangeLong,
		},
		{
			name:        "double",
			args:        args{name: "test"},
			want:        &Range{_type: TypeRangeDouble, name: "test"},
			constructor: NewRangeDouble,
		},
		{
			name:        "ip",
			args:        args{name: "test"},
			want:        &Range{_type: TypeRangeIp, name: "test"},
			constructor: NewRangeIp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.constructor(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRangeDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRangeDate(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *RangeDate
	}{
		{
			name: "date",
			args: args{name: "test"},
			want: &RangeDate{Range: Range{_type: TypeRangeDate, name: "test"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRangeDate(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRangeDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeDate_Formats(t *testing.T) {
	type fields struct {
		Range   Range
		formats []DateFormat
	}
	type args struct {
		formats []DateFormat
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RangeDate
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{formats: []DateFormat{DateFormatANSIC, DateFormatUnixDate}},
			want:   &RangeDate{formats: []DateFormat{DateFormatANSIC, DateFormatUnixDate}},
		},
		{
			name:   "add",
			fields: fields{formats: []DateFormat{DateFormatANSIC, DateFormatUnixDate}},
			args: args{formats: []DateFormat{
				DateFormatRubyDate,
				DateFormatRFC822,
				DateFormatRFC822Z,
				DateFormatRFC850,
				DateFormatRFC1123,
				DateFormatRFC1123Z,
				DateFormatRFC3339,
				DateFormatKitchen,
				DateFormatStamp,
				DateFormatStampMilli,
			}},
			want: &RangeDate{formats: []DateFormat{
				DateFormatANSIC,
				DateFormatUnixDate,
				DateFormatRubyDate,
				DateFormatRFC822,
				DateFormatRFC822Z,
				DateFormatRFC850,
				DateFormatRFC1123,
				DateFormatRFC1123Z,
				DateFormatRFC3339,
				DateFormatKitchen,
				DateFormatStamp,
				DateFormatStampMilli,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &RangeDate{
				Range:   tt.fields.Range,
				formats: tt.fields.formats,
			}
			if got := field.Formats(tt.args.formats...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeDate_GetType(t *testing.T) {
	type fields struct {
		Range   Range
		formats []DateFormat
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "range date type",
			fields: fields{Range: Range{_type: TypeRangeDate}},
			want:   TypeRangeDate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &RangeDate{
				Range:   tt.fields.Range,
				formats: tt.fields.formats,
			}
			if got := field.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeDate_Source(t *testing.T) {
	initBoost := float64(9)
	boolTrue := true

	type fields struct {
		Range   Range
		formats []DateFormat
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name:   "empty",
			fields: fields{Range: Range{_type: TypeRangeDate}},
			want: map[string]interface{}{
				"type": TypeRangeDate,
			},
			wantErr: false,
		},
		{
			name: "coerce",
			fields: fields{Range: Range{
				coerce: &boolTrue,
				_type:  TypeRangeDate,
			}},
			want: map[string]interface{}{
				"type":   TypeRangeDate,
				"coerce": true,
			},
			wantErr: false,
		},
		{
			name: "boost",
			fields: fields{Range: Range{
				boost: &initBoost,
				_type: TypeRangeDate,
			}},
			want: map[string]interface{}{
				"type":  TypeRangeDate,
				"boost": initBoost,
			},
			wantErr: false,
		},
		{
			name: "index",
			fields: fields{Range: Range{
				index: &boolTrue,
				_type: TypeRangeDate,
			}},
			want: map[string]interface{}{
				"type":  TypeRangeDate,
				"index": true,
			},
			wantErr: false,
		},
		{
			name: "store",
			fields: fields{Range: Range{
				store: &boolTrue,
				_type: TypeRangeDate,
			}},
			want: map[string]interface{}{
				"type":  TypeRangeDate,
				"store": true,
			},
			wantErr: false,
		},
		{
			name: "format",
			fields: fields{
				formats: []DateFormat{
					DateFormatEpochSecond,
					DateFormatDateOptionalTime,
				},
				Range: Range{_type: TypeRangeDate},
			},
			want: map[string]interface{}{
				"type":   TypeRangeDate,
				"format": "epoch_second||date_optional_time",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &RangeDate{
				Range:   tt.fields.Range,
				formats: tt.fields.formats,
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

func TestRange_GetType(t *testing.T) {
	type fields struct {
		coerce *bool
		store  *bool
		index  *bool
		boost  *float64
		_type  Type
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "integer",
			fields: fields{_type: TypeRangeInteger},
			want:   TypeRangeInteger,
		},
		{
			name:   "float",
			fields: fields{_type: TypeRangeFloat},
			want:   TypeRangeFloat,
		},
		{
			name:   "long",
			fields: fields{_type: TypeRangeLong},
			want:   TypeRangeLong,
		},
		{
			name:   "double",
			fields: fields{_type: TypeRangeDouble},
			want:   TypeRangeDouble,
		},
		{
			name:   "ip",
			fields: fields{_type: TypeRangeIp},
			want:   TypeRangeIp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Range{
				coerce: tt.fields.coerce,
				store:  tt.fields.store,
				index:  tt.fields.index,
				boost:  tt.fields.boost,
				_type:  tt.fields._type,
			}
			if got := field.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_SetBoost(t *testing.T) {
	initBoost := float64(9)
	setBoost := 5.4

	type fields struct {
		coerce *bool
		store  *bool
		index  *bool
		boost  *float64
		_type  Type
	}
	type args struct {
		boost float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Range
	}{
		{
			name:   "set boost",
			fields: fields{},
			args:   args{boost: setBoost},
			want:   &Range{boost: &setBoost},
		},
		{
			name:   "change boost",
			fields: fields{boost: &initBoost},
			args:   args{boost: setBoost},
			want:   &Range{boost: &setBoost},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Range{
				coerce: tt.fields.coerce,
				store:  tt.fields.store,
				index:  tt.fields.index,
				boost:  tt.fields.boost,
				_type:  tt.fields._type,
			}
			if got := field.SetBoost(tt.args.boost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBoost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_SetCoerce(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce *bool
		store  *bool
		index  *bool
		boost  *float64
		_type  Type
	}
	type args struct {
		coerce bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Range
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{coerce: true},
			want:   &Range{coerce: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{coerce: false},
			want:   &Range{coerce: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{coerce: &boolTrue},
			args:   args{coerce: false},
			want:   &Range{coerce: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Range{
				coerce: tt.fields.coerce,
				store:  tt.fields.store,
				index:  tt.fields.index,
				boost:  tt.fields.boost,
				_type:  tt.fields._type,
			}
			if got := field.SetCoerce(tt.args.coerce); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCoerce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_SetIndex(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce *bool
		store  *bool
		index  *bool
		boost  *float64
		_type  Type
	}
	type args struct {
		index bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Range
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{index: true},
			want:   &Range{index: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{index: false},
			want:   &Range{index: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{index: &boolTrue},
			args:   args{index: false},
			want:   &Range{index: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Range{
				coerce: tt.fields.coerce,
				store:  tt.fields.store,
				index:  tt.fields.index,
				boost:  tt.fields.boost,
				_type:  tt.fields._type,
			}
			if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_SetStore(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce *bool
		store  *bool
		index  *bool
		boost  *float64
		_type  Type
	}
	type args struct {
		store bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Range
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{store: true},
			want:   &Range{store: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{store: false},
			want:   &Range{store: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{store: &boolTrue},
			args:   args{store: false},
			want:   &Range{store: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Range{
				coerce: tt.fields.coerce,
				store:  tt.fields.store,
				index:  tt.fields.index,
				boost:  tt.fields.boost,
				_type:  tt.fields._type,
			}
			if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_Source(t *testing.T) {
	initBoost := float64(9)
	boolTrue := true

	type fields struct {
		coerce *bool
		store  *bool
		index  *bool
		boost  *float64
		_type  Type
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name:   "integer",
			fields: fields{_type: TypeRangeInteger},
			want: map[string]interface{}{
				"type": TypeRangeInteger,
			},
			wantErr: false,
		},
		{
			name:   "float",
			fields: fields{_type: TypeRangeFloat},
			want: map[string]interface{}{
				"type": TypeRangeFloat,
			},
			wantErr: false,
		},
		{
			name:   "long",
			fields: fields{_type: TypeRangeLong},
			want: map[string]interface{}{
				"type": TypeRangeLong,
			},
			wantErr: false,
		},
		{
			name:   "double",
			fields: fields{_type: TypeRangeDouble},
			want: map[string]interface{}{
				"type": TypeRangeDouble,
			},
			wantErr: false,
		},
		{
			name:   "ip",
			fields: fields{_type: TypeRangeIp},
			want: map[string]interface{}{
				"type": TypeRangeIp,
			},
			wantErr: false,
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
			name:   "boost",
			fields: fields{boost: &initBoost},
			want: map[string]interface{}{
				"type":  Type(""),
				"boost": initBoost,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Range{
				coerce: tt.fields.coerce,
				store:  tt.fields.store,
				index:  tt.fields.index,
				boost:  tt.fields.boost,
				_type:  tt.fields._type,
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
