package index

import (
	fields2 "github.com/diez37/elastic-migrate/fields"
	"reflect"
	"testing"
)

func TestMappings_AddProperties(t *testing.T) {
	type fields struct {
		dynamic    *bool
		properties []fields2.Fielder
	}
	type args struct {
		properties []fields2.Fielder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mappings
	}{
		{
			name: "set",
			args: args{properties: []fields2.Fielder{&fields2.Keyword{}}},
			want: &Mappings{properties: []fields2.Fielder{&fields2.Keyword{}}},
		},
		{
			name:   "add",
			fields: fields{properties: []fields2.Fielder{&fields2.Keyword{}}},
			args:   args{properties: []fields2.Fielder{&fields2.Keyword{}, &fields2.Text{}}},
			want:   &Mappings{properties: []fields2.Fielder{&fields2.Keyword{}, &fields2.Keyword{}, &fields2.Text{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappings := &Mappings{
				dynamic:    tt.fields.dynamic,
				properties: tt.fields.properties,
			}
			if got := mappings.AddProperties(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMappings_SetDynamic(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		dynamic    *bool
		properties []fields2.Fielder
	}
	type args struct {
		dynamic bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mappings
	}{
		{
			name: "set",
			args: args{dynamic: true},
			want: &Mappings{dynamic: &boolTrue},
		},
		{
			name:   "set",
			fields: fields{dynamic: &boolTrue},
			args:   args{dynamic: false},
			want:   &Mappings{dynamic: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappings := &Mappings{
				dynamic:    tt.fields.dynamic,
				properties: tt.fields.properties,
			}
			if got := mappings.SetDynamic(tt.args.dynamic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDynamic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMappings_Source(t *testing.T) {
	boolTrue := true

	type fields struct {
		dynamic    *bool
		properties []fields2.Fielder
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
			name:   "dynamic",
			fields: fields{dynamic: &boolTrue},
			want: map[string]interface{}{
				"dynamic": true,
			},
		},
		{
			name:   "properties",
			fields: fields{properties: []fields2.Fielder{fields2.NewKeyword("test"), fields2.NewText("test2")}},
			want: map[string]interface{}{
				"properties": map[fields2.FieldName]interface{}{
					"test": map[string]interface{}{
						"type": fields2.Type("keyword"),
					},
					"test2": map[string]interface{}{
						"type": fields2.Type("text"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappings := &Mappings{
				dynamic:    tt.fields.dynamic,
				properties: tt.fields.properties,
			}
			got, err := mappings.Source()
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

func TestNewMappings(t *testing.T) {
	tests := []struct {
		name string
		want *Mappings
	}{
		{
			want: &Mappings{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMappings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMappings() = %v, want %v", got, tt.want)
			}
		})
	}
}
