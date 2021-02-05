package template

import (
	fields2 "github.com/diez37/elastic-migrate/fields"
	"github.com/diez37/elastic-migrate/index"
	"reflect"
	"testing"
)

func TestIndex_AddComposedOf(t *testing.T) {
	type fields struct {
		name          string
		template      *index.Index
		priority      *int
		version       *int
		indexPatterns []string
		composedOf    []string
		meta          []*fields2.Meta
	}
	type args struct {
		composedOf []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{composedOf: []string{"test"}},
			want: &Index{composedOf: []string{"test"}},
		},
		{
			name:   "add",
			fields: fields{composedOf: []string{"test"}},
			args:   args{composedOf: []string{"test2"}},
			want:   &Index{composedOf: []string{"test", "test2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				name:          tt.fields.name,
				template:      tt.fields.template,
				priority:      tt.fields.priority,
				version:       tt.fields.version,
				indexPatterns: tt.fields.indexPatterns,
				composedOf:    tt.fields.composedOf,
				meta:          tt.fields.meta,
			}
			if got := i.AddComposedOf(tt.args.composedOf...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddComposedOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_AddIndexPatterns(t *testing.T) {
	type fields struct {
		name          string
		template      *index.Index
		priority      *int
		version       *int
		indexPatterns []string
		composedOf    []string
		meta          []*fields2.Meta
	}
	type args struct {
		indexPatterns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{indexPatterns: []string{"foo*"}},
			want: &Index{indexPatterns: []string{"foo*"}},
		},
		{
			name:   "add",
			fields: fields{indexPatterns: []string{"foo*"}},
			args:   args{indexPatterns: []string{"bar"}},
			want:   &Index{indexPatterns: []string{"foo*", "bar"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				name:          tt.fields.name,
				template:      tt.fields.template,
				priority:      tt.fields.priority,
				version:       tt.fields.version,
				indexPatterns: tt.fields.indexPatterns,
				composedOf:    tt.fields.composedOf,
				meta:          tt.fields.meta,
			}
			if got := i.AddIndexPatterns(tt.args.indexPatterns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddIndexPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_AddMeta(t *testing.T) {
	type fields struct {
		name          string
		template      *index.Index
		priority      *int
		version       *int
		indexPatterns []string
		composedOf    []string
		meta          []*fields2.Meta
	}
	type args struct {
		meta []*fields2.Meta
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{meta: []*fields2.Meta{fields2.NewMeta("test", "test")}},
			want: &Index{meta: []*fields2.Meta{fields2.NewMeta("test", "test")}},
		},
		{
			name:   "change",
			fields: fields{meta: []*fields2.Meta{fields2.NewMeta("test", "test")}},
			args:   args{meta: []*fields2.Meta{fields2.NewMeta("test2", "test")}},
			want:   &Index{meta: []*fields2.Meta{fields2.NewMeta("test", "test"), fields2.NewMeta("test2", "test")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				name:          tt.fields.name,
				template:      tt.fields.template,
				priority:      tt.fields.priority,
				version:       tt.fields.version,
				indexPatterns: tt.fields.indexPatterns,
				composedOf:    tt.fields.composedOf,
				meta:          tt.fields.meta,
			}
			if got := i.AddMeta(tt.args.meta...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Name(t *testing.T) {
	type fields struct {
		name          string
		template      *index.Index
		priority      *int
		version       *int
		indexPatterns []string
		composedOf    []string
		meta          []*fields2.Meta
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				name:          tt.fields.name,
				template:      tt.fields.template,
				priority:      tt.fields.priority,
				version:       tt.fields.version,
				indexPatterns: tt.fields.indexPatterns,
				composedOf:    tt.fields.composedOf,
				meta:          tt.fields.meta,
			}
			if got := i.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetPriority(t *testing.T) {
	prioritySet := 1
	priorityChange := 2

	type fields struct {
		name          string
		template      *index.Index
		priority      *int
		version       *int
		indexPatterns []string
		composedOf    []string
		meta          []*fields2.Meta
	}
	type args struct {
		priority int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{priority: 1},
			want: &Index{priority: &prioritySet},
		},
		{
			name:   "change",
			fields: fields{priority: &prioritySet},
			args:   args{priority: 2},
			want:   &Index{priority: &priorityChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				name:          tt.fields.name,
				template:      tt.fields.template,
				priority:      tt.fields.priority,
				version:       tt.fields.version,
				indexPatterns: tt.fields.indexPatterns,
				composedOf:    tt.fields.composedOf,
				meta:          tt.fields.meta,
			}
			if got := i.SetPriority(tt.args.priority); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetTemplate(t *testing.T) {
	type fields struct {
		name          string
		template      *index.Index
		priority      *int
		version       *int
		indexPatterns []string
		composedOf    []string
		meta          []*fields2.Meta
	}
	type args struct {
		template *index.Index
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			args: args{template: &index.Index{}},
			want: &Index{template: &index.Index{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				name:          tt.fields.name,
				template:      tt.fields.template,
				priority:      tt.fields.priority,
				version:       tt.fields.version,
				indexPatterns: tt.fields.indexPatterns,
				composedOf:    tt.fields.composedOf,
				meta:          tt.fields.meta,
			}
			if got := i.SetTemplate(tt.args.template); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SetVersion(t *testing.T) {
	versionSet := 1
	versionChange := 2

	type fields struct {
		name          string
		template      *index.Index
		priority      *int
		version       *int
		indexPatterns []string
		composedOf    []string
		meta          []*fields2.Meta
	}
	type args struct {
		version int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		{
			name: "set",
			args: args{version: 1},
			want: &Index{version: &versionSet},
		},
		{
			name:   "change",
			fields: fields{version: &versionSet},
			args:   args{version: 2},
			want:   &Index{version: &versionChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				name:          tt.fields.name,
				template:      tt.fields.template,
				priority:      tt.fields.priority,
				version:       tt.fields.version,
				indexPatterns: tt.fields.indexPatterns,
				composedOf:    tt.fields.composedOf,
				meta:          tt.fields.meta,
			}
			if got := i.SetVersion(tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Source(t *testing.T) {
	template, _ := index.NewIndex().Source()
	priority := 1
	version := 1

	type fields struct {
		name          string
		template      *index.Index
		priority      *int
		version       *int
		indexPatterns []string
		composedOf    []string
		meta          []*fields2.Meta
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
			name:   "template",
			fields: fields{template: index.NewIndex()},
			want: map[string]interface{}{
				"template": template,
			},
		},
		{
			name:   "priority",
			fields: fields{priority: &priority},
			want: map[string]interface{}{
				"priority": priority,
			},
		},
		{
			name:   "version",
			fields: fields{version: &version},
			want: map[string]interface{}{
				"version": version,
			},
		},
		{
			name:   "index_patterns",
			fields: fields{indexPatterns: []string{"test"}},
			want: map[string]interface{}{
				"index_patterns": []string{"test"},
			},
		},
		{
			name:   "composed_of",
			fields: fields{composedOf: []string{"test"}},
			want: map[string]interface{}{
				"composed_of": []string{"test"},
			},
		},
		{
			name:   "meta",
			fields: fields{meta: []*fields2.Meta{fields2.NewMeta("test", "test")}},
			want: map[string]interface{}{
				"_meta": map[string]interface{}{
					"test": "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				name:          tt.fields.name,
				template:      tt.fields.template,
				priority:      tt.fields.priority,
				version:       tt.fields.version,
				indexPatterns: tt.fields.indexPatterns,
				composedOf:    tt.fields.composedOf,
				meta:          tt.fields.meta,
			}
			got, err := i.Source()
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
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Index
	}{
		{
			want: &Index{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndex(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
