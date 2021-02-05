package index

import (
	"github.com/olivere/elastic/v7"
	"reflect"
	"testing"
)

func TestAlias_Alias(t *testing.T) {
	type fields struct {
		alias   string
		routing *string
		filter  elastic.Query
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{alias: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alias := &Alias{
				alias:   tt.fields.alias,
				routing: tt.fields.routing,
				filter:  tt.fields.filter,
			}
			if got := alias.Alias(); got != tt.want {
				t.Errorf("Alias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlias_SetFilter(t *testing.T) {
	type fields struct {
		alias   string
		routing *string
		filter  elastic.Query
	}
	type args struct {
		filter elastic.Query
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Alias
	}{
		{
			name: "set",
			args: args{filter: elastic.NewTermQuery("test", "test")},
			want: &Alias{filter: elastic.NewTermQuery("test", "test")},
		},
		{
			name:   "change",
			fields: fields{filter: elastic.NewTermQuery("test", "test")},
			args:   args{filter: elastic.NewTermQuery("test", "test2")},
			want:   &Alias{filter: elastic.NewTermQuery("test", "test2")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alias := &Alias{
				alias:   tt.fields.alias,
				routing: tt.fields.routing,
				filter:  tt.fields.filter,
			}
			if got := alias.SetFilter(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlias_SetRouting(t *testing.T) {
	setString := "set"
	changeString := "change"

	type fields struct {
		alias   string
		routing *string
		filter  elastic.Query
	}
	type args struct {
		routing string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Alias
	}{
		{
			name: "set",
			args: args{routing: "set"},
			want: &Alias{routing: &setString},
		},
		{
			name:   "change",
			fields: fields{routing: &setString},
			args:   args{routing: "change"},
			want:   &Alias{routing: &changeString},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alias := &Alias{
				alias:   tt.fields.alias,
				routing: tt.fields.routing,
				filter:  tt.fields.filter,
			}
			if got := alias.SetRouting(tt.args.routing); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRouting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlias_Source(t *testing.T) {
	setString := "set"

	type fields struct {
		alias   string
		routing *string
		filter  elastic.Query
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
			name:   "routing",
			fields: fields{routing: &setString},
			want: map[string]interface{}{
				"routing": "set",
			},
		},
		{
			name:   "routing",
			fields: fields{filter: elastic.NewTermQuery("test", "test")},
			want: map[string]interface{}{
				"filter": map[string]interface{}{
					"term": map[string]interface{}{
						"test": "test",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alias := &Alias{
				alias:   tt.fields.alias,
				routing: tt.fields.routing,
				filter:  tt.fields.filter,
			}
			got, err := alias.Source()
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

func TestNewAlias(t *testing.T) {
	type args struct {
		alias string
	}
	tests := []struct {
		name string
		args args
		want *Alias
	}{
		{
			args: args{alias: "test"},
			want: &Alias{alias: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlias(tt.args.alias); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}
