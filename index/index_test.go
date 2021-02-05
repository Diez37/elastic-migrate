package index

import (
	fields2 "github.com/diez37/elastic-migrate/fields"
	"github.com/diez37/elastic-migrate/settings"
	"reflect"
	"testing"
)

func TestIndex_AddAliases(t *testing.T) {
	type fields struct {
		mappings *Mappings
		settings *settings.Settings
		aliases  []*Alias
	}
	type args struct {
		aliases []*Alias
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{aliases: []*Alias{NewAlias("test")}},
			want: &Index{aliases: []*Alias{NewAlias("test")}},
		},
		{
			name:   "change",
			fields: fields{aliases: []*Alias{NewAlias("test")}},
			args:   args{aliases: []*Alias{NewAlias("test2")}},
			want:   &Index{aliases: []*Alias{NewAlias("test"), NewAlias("test2")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				mappings: tt.fields.mappings,
				settings: tt.fields.settings,
				aliases:  tt.fields.aliases,
			}
			if got := index.AddAliases(tt.args.aliases...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddAliases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetMappings(t *testing.T) {
	type fields struct {
		mappings *Mappings
		settings *settings.Settings
		aliases  []*Alias
	}
	type args struct {
		mappings *Mappings
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{mappings: NewMappings()},
			want: &Index{mappings: NewMappings()},
		},
		{
			name:   "change",
			fields: fields{mappings: NewMappings()},
			args:   args{mappings: NewMappings().AddProperties(fields2.NewKeyword("test"))},
			want:   &Index{mappings: NewMappings().AddProperties(fields2.NewKeyword("test"))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				mappings: tt.fields.mappings,
				settings: tt.fields.settings,
				aliases:  tt.fields.aliases,
			}
			if got := index.SetMappings(tt.args.mappings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMappings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetSettings(t *testing.T) {
	type fields struct {
		mappings *Mappings
		settings *settings.Settings
		aliases  []*Alias
	}
	type args struct {
		settings *settings.Settings
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{settings: settings.NewSettings(settings.NewIndex())},
			want: &Index{settings: settings.NewSettings(settings.NewIndex())},
		},
		{
			name:   "change",
			fields: fields{settings: settings.NewSettings(settings.NewIndex())},
			args:   args{settings: settings.NewSettings(settings.NewIndex().SetNumberOfReplicas(1))},
			want:   &Index{settings: settings.NewSettings(settings.NewIndex().SetNumberOfReplicas(1))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				mappings: tt.fields.mappings,
				settings: tt.fields.settings,
				aliases:  tt.fields.aliases,
			}
			if got := index.SetSettings(tt.args.settings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Source(t *testing.T) {
	type fields struct {
		mappings *Mappings
		settings *settings.Settings
		aliases  []*Alias
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
			name:   "mappings",
			fields: fields{mappings: NewMappings().AddProperties(fields2.NewKeyword("test"))},
			want: map[string]interface{}{
				"mappings": map[string]interface{}{
					"properties": map[fields2.FieldName]interface{}{
						"test": map[string]interface{}{
							"type": fields2.Type("keyword"),
						},
					},
				},
			},
		},
		{
			name:   "settings",
			fields: fields{settings: settings.NewSettings(settings.NewIndex().SetNumberOfReplicas(1))},
			want: map[string]interface{}{
				"settings": map[string]interface{}{
					"index": map[string]interface{}{
						"number_of_replicas": uint32(1),
					},
				},
			},
		},
		{
			name:   "aliases",
			fields: fields{aliases: []*Alias{NewAlias("test")}},
			want: map[string]interface{}{
				"aliases": map[string]interface{}{
					"test": map[string]interface{}{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				mappings: tt.fields.mappings,
				settings: tt.fields.settings,
				aliases:  tt.fields.aliases,
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
