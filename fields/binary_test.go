package fields

import (
	"reflect"
	"testing"
)

func TestBinary_GetType(t *testing.T) {
	type fields struct {
		docValues *bool
		store     *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "type binary",
			fields: fields{},
			want:   TypeBinary,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Binary{
				docValues: tt.fields.docValues,
				store:     tt.fields.store,
			}
			if got := field.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinary_SetDocValues(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		docValues *bool
		store     *bool
	}
	type args struct {
		docValues bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Binary
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{docValues: true},
			want:   &Binary{docValues: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{docValues: false},
			want:   &Binary{docValues: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{docValues: &boolTrue},
			args:   args{docValues: false},
			want:   &Binary{docValues: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Binary{
				docValues: tt.fields.docValues,
				store:     tt.fields.store,
			}
			if got := field.SetDocValues(tt.args.docValues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDocValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinary_SetStore(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		docValues *bool
		store     *bool
	}
	type args struct {
		store bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Binary
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{store: true},
			want:   &Binary{store: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{store: false},
			want:   &Binary{store: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{store: &boolTrue},
			args:   args{store: false},
			want:   &Binary{store: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Binary{
				docValues: tt.fields.docValues,
				store:     tt.fields.store,
			}
			if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinary_Source(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		docValues *bool
		store     *bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name:   "fields empty",
			fields: fields{},
			want: map[string]interface{}{
				"type": TypeBinary,
			},
			wantErr: false,
		},
		{
			name:   "docValues true",
			fields: fields{docValues: &boolTrue},
			want: map[string]interface{}{
				"type":       TypeBinary,
				"doc_values": true,
			},
			wantErr: false,
		},
		{
			name:   "docValues false",
			fields: fields{docValues: &boolFalse},
			want: map[string]interface{}{
				"type":       TypeBinary,
				"doc_values": false,
			},
			wantErr: false,
		},
		{
			name:   "store true",
			fields: fields{store: &boolTrue},
			want: map[string]interface{}{
				"type":  TypeBinary,
				"store": true,
			},
			wantErr: false,
		},
		{
			name:   "store false",
			fields: fields{store: &boolFalse},
			want: map[string]interface{}{
				"type":  TypeBinary,
				"store": false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Binary{
				docValues: tt.fields.docValues,
				store:     tt.fields.store,
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

func TestNewBinary(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		want *Binary
		args args
	}{
		{
			name: "new",
			args: args{name: "test"},
			want: &Binary{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinary(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
