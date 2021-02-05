package template

import (
	fields2 "github.com/diez37/elastic-migrate/fields"
	"github.com/diez37/elastic-migrate/index"
	"reflect"
	"testing"
)

func TestComponent_Name(t *testing.T) {
	type fields struct {
		name            string
		template        *index.Index
		version         *int
		allowAutoCreate *bool
		meta            []*fields2.Meta
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
			component := &Component{
				name:            tt.fields.name,
				template:        tt.fields.template,
				version:         tt.fields.version,
				allowAutoCreate: tt.fields.allowAutoCreate,
				meta:            tt.fields.meta,
			}
			if got := component.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetAllowAutoCreate(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name            string
		template        *index.Index
		version         *int
		allowAutoCreate *bool
		meta            []*fields2.Meta
	}
	type args struct {
		allowAutoCreate bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name: "set",
			args: args{allowAutoCreate: true},
			want: &Component{allowAutoCreate: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{allowAutoCreate: &boolTrue},
			args:   args{allowAutoCreate: false},
			want:   &Component{allowAutoCreate: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := &Component{
				name:            tt.fields.name,
				template:        tt.fields.template,
				version:         tt.fields.version,
				allowAutoCreate: tt.fields.allowAutoCreate,
				meta:            tt.fields.meta,
			}
			if got := component.SetAllowAutoCreate(tt.args.allowAutoCreate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAllowAutoCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetMeta(t *testing.T) {
	type fields struct {
		name            string
		template        *index.Index
		version         *int
		allowAutoCreate *bool
		meta            []*fields2.Meta
	}
	type args struct {
		meta []*fields2.Meta
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name: "set",
			args: args{meta: []*fields2.Meta{fields2.NewMeta("test", "test")}},
			want: &Component{meta: []*fields2.Meta{fields2.NewMeta("test", "test")}},
		},
		{
			name:   "add",
			fields: fields{meta: []*fields2.Meta{fields2.NewMeta("test", "test")}},
			args:   args{meta: []*fields2.Meta{fields2.NewMeta("test2", "test")}},
			want:   &Component{meta: []*fields2.Meta{fields2.NewMeta("test", "test"), fields2.NewMeta("test2", "test")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := &Component{
				name:            tt.fields.name,
				template:        tt.fields.template,
				version:         tt.fields.version,
				allowAutoCreate: tt.fields.allowAutoCreate,
				meta:            tt.fields.meta,
			}
			if got := component.SetMeta(tt.args.meta...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetTemplate(t *testing.T) {
	type fields struct {
		name            string
		template        *index.Index
		version         *int
		allowAutoCreate *bool
		meta            []*fields2.Meta
	}
	type args struct {
		template *index.Index
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			args: args{template: &index.Index{}},
			want: &Component{template: &index.Index{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := &Component{
				name:            tt.fields.name,
				template:        tt.fields.template,
				version:         tt.fields.version,
				allowAutoCreate: tt.fields.allowAutoCreate,
				meta:            tt.fields.meta,
			}
			if got := component.SetTemplate(tt.args.template); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetVersion(t *testing.T) {
	versionSet := 1
	versionChange := 2

	type fields struct {
		name            string
		template        *index.Index
		version         *int
		allowAutoCreate *bool
		meta            []*fields2.Meta
	}
	type args struct {
		version int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name: "set",
			args: args{version: 1},
			want: &Component{version: &versionSet},
		},
		{
			name:   "change",
			fields: fields{version: &versionSet},
			args:   args{version: 2},
			want:   &Component{version: &versionChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := &Component{
				name:            tt.fields.name,
				template:        tt.fields.template,
				version:         tt.fields.version,
				allowAutoCreate: tt.fields.allowAutoCreate,
				meta:            tt.fields.meta,
			}
			if got := component.SetVersion(tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_Source(t *testing.T) {
	version := 1
	boolTrue := true
	indexSource, _ := index.NewIndex().Source()

	type fields struct {
		name            string
		template        *index.Index
		version         *int
		allowAutoCreate *bool
		meta            []*fields2.Meta
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
				"template": indexSource,
			},
		},
		{
			name:   "version",
			fields: fields{version: &version},
			want: map[string]interface{}{
				"version": 1,
			},
		},
		{
			name:   "allowAutoCreate",
			fields: fields{allowAutoCreate: &boolTrue},
			want: map[string]interface{}{
				"allow_auto_create": true,
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
			component := &Component{
				name:            tt.fields.name,
				template:        tt.fields.template,
				version:         tt.fields.version,
				allowAutoCreate: tt.fields.allowAutoCreate,
				meta:            tt.fields.meta,
			}
			got, err := component.Source()
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

func TestNewComponent(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Component
	}{
		{
			want: &Component{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComponent(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComponent() = %v, want %v", got, tt.want)
			}
		})
	}
}
