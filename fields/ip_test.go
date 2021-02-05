package fields

import (
	"reflect"
	"testing"
)

func TestIp_GetType(t *testing.T) {
	type fields struct {
		docValues *bool
		index     *bool
		store     *bool
		boost     *float64
		nullValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "type binary",
			fields: fields{},
			want:   TypeIp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Ip{
				docValues: tt.fields.docValues,
				index:     tt.fields.index,
				store:     tt.fields.store,
				boost:     tt.fields.boost,
				nullValue: tt.fields.nullValue,
			}
			if got := field.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIp_SetBoost(t *testing.T) {
	initBoost := 9.0
	setBoost := 5.4

	type fields struct {
		docValues *bool
		index     *bool
		store     *bool
		boost     *float64
		nullValue interface{}
	}
	type args struct {
		boost float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Ip
	}{
		{
			name:   "set boost",
			fields: fields{},
			args:   args{boost: setBoost},
			want:   &Ip{boost: &setBoost},
		},
		{
			name:   "change boost",
			fields: fields{boost: &initBoost},
			args:   args{boost: setBoost},
			want:   &Ip{boost: &setBoost},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Ip{
				docValues: tt.fields.docValues,
				index:     tt.fields.index,
				store:     tt.fields.store,
				boost:     tt.fields.boost,
				nullValue: tt.fields.nullValue,
			}
			if got := field.SetBoost(tt.args.boost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBoost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIp_SetDocValues(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		docValues *bool
		index     *bool
		store     *bool
		boost     *float64
		nullValue interface{}
	}
	type args struct {
		docValues bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Ip
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{docValues: true},
			want:   &Ip{docValues: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{docValues: false},
			want:   &Ip{docValues: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{docValues: &boolTrue},
			args:   args{docValues: false},
			want:   &Ip{docValues: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Ip{
				docValues: tt.fields.docValues,
				index:     tt.fields.index,
				store:     tt.fields.store,
				boost:     tt.fields.boost,
				nullValue: tt.fields.nullValue,
			}
			if got := field.SetDocValues(tt.args.docValues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDocValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIp_SetIndex(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		docValues *bool
		index     *bool
		store     *bool
		boost     *float64
		nullValue interface{}
	}
	type args struct {
		index bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Ip
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{index: true},
			want:   &Ip{index: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{index: false},
			want:   &Ip{index: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{index: &boolTrue},
			args:   args{index: false},
			want:   &Ip{index: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Ip{
				docValues: tt.fields.docValues,
				index:     tt.fields.index,
				store:     tt.fields.store,
				boost:     tt.fields.boost,
				nullValue: tt.fields.nullValue,
			}
			if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIp_SetNullValue(t *testing.T) {
	type fields struct {
		docValues *bool
		index     *bool
		store     *bool
		boost     *float64
		nullValue interface{}
	}
	type args struct {
		nullValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Ip
	}{
		{
			name:   "empty",
			fields: fields{},
			args:   args{},
			want:   &Ip{nullValue: nil},
		},
		{
			name:   "int",
			fields: fields{},
			args:   args{nullValue: 1},
			want:   &Ip{nullValue: 1},
		},
		{
			name:   "float",
			fields: fields{},
			args:   args{nullValue: 4.5},
			want:   &Ip{nullValue: 4.5},
		},
		{
			name:   "string",
			fields: fields{},
			args:   args{nullValue: "test"},
			want:   &Ip{nullValue: "test"},
		},
		{
			name:   "bool",
			fields: fields{},
			args:   args{nullValue: true},
			want:   &Ip{nullValue: true},
		},
		{
			name:   "change",
			fields: fields{nullValue: "test"},
			args:   args{nullValue: true},
			want:   &Ip{nullValue: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Ip{
				docValues: tt.fields.docValues,
				index:     tt.fields.index,
				store:     tt.fields.store,
				boost:     tt.fields.boost,
				nullValue: tt.fields.nullValue,
			}
			if got := field.SetNullValue(tt.args.nullValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNullValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIp_SetStore(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		docValues *bool
		index     *bool
		store     *bool
		boost     *float64
		nullValue interface{}
	}
	type args struct {
		store bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Ip
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{store: true},
			want:   &Ip{store: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{store: false},
			want:   &Ip{store: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{store: &boolTrue},
			args:   args{store: false},
			want:   &Ip{store: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Ip{
				docValues: tt.fields.docValues,
				index:     tt.fields.index,
				store:     tt.fields.store,
				boost:     tt.fields.boost,
				nullValue: tt.fields.nullValue,
			}
			if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIp_Source(t *testing.T) {
	boost := 2.5
	boolTrue := true

	type fields struct {
		docValues *bool
		index     *bool
		store     *bool
		boost     *float64
		nullValue interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name:   "empty",
			fields: fields{},
			want: map[string]interface{}{
				"type": TypeIp,
			},
		},
		{
			name:   "docValues",
			fields: fields{docValues: &boolTrue},
			want: map[string]interface{}{
				"type":       TypeIp,
				"doc_values": true,
			},
		},
		{
			name:   "store",
			fields: fields{store: &boolTrue},
			want: map[string]interface{}{
				"type":  TypeIp,
				"store": true,
			},
		},
		{
			name:   "boost",
			fields: fields{boost: &boost},
			want: map[string]interface{}{
				"type":  TypeIp,
				"boost": boost,
			},
		},
		{
			name:   "index",
			fields: fields{index: &boolTrue},
			want: map[string]interface{}{
				"type":  TypeIp,
				"index": true,
			},
		},
		{
			name:   "nullValue",
			fields: fields{nullValue: 2.7},
			want: map[string]interface{}{
				"type":       TypeIp,
				"null_value": 2.7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Ip{
				docValues: tt.fields.docValues,
				index:     tt.fields.index,
				store:     tt.fields.store,
				boost:     tt.fields.boost,
				nullValue: tt.fields.nullValue,
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

func TestNewIp(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Ip
	}{
		{
			args: args{name: "test"},
			want: &Ip{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIp(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIp() = %v, want %v", got, tt.want)
			}
		})
	}
}
