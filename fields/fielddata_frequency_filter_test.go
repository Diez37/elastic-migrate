package fields

import (
	"reflect"
	"testing"
)

func TestFielddataFrequencyFilter_SetMaximum(t *testing.T) {
	value := 10.0
	initValue := -1.0

	type fields struct {
		min            *float64
		max            *float64
		minSegmentSize *int
	}
	type args struct {
		maximum float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FieldDataFrequencyFilter
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{maximum: value},
			want: &FieldDataFrequencyFilter{
				max: &value,
			},
		},
		{
			name:   "change",
			fields: fields{max: &initValue},
			args:   args{maximum: value},
			want: &FieldDataFrequencyFilter{
				max: &value,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fielddata := &FieldDataFrequencyFilter{
				min:            tt.fields.min,
				max:            tt.fields.max,
				minSegmentSize: tt.fields.minSegmentSize,
			}
			if got := fielddata.SetMaximum(tt.args.maximum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFielddataFrequencyFilter_SetMinSegmentSize(t *testing.T) {
	value := 10
	initValue := -1

	type fields struct {
		min            *float64
		max            *float64
		minSegmentSize *int
	}
	type args struct {
		minSegmentSize int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FieldDataFrequencyFilter
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{minSegmentSize: value},
			want: &FieldDataFrequencyFilter{
				minSegmentSize: &value,
			},
		},
		{
			name:   "change",
			fields: fields{minSegmentSize: &initValue},
			args:   args{minSegmentSize: value},
			want: &FieldDataFrequencyFilter{
				minSegmentSize: &value,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fielddata := &FieldDataFrequencyFilter{
				min:            tt.fields.min,
				max:            tt.fields.max,
				minSegmentSize: tt.fields.minSegmentSize,
			}
			if got := fielddata.SetMinSegmentSize(tt.args.minSegmentSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinSegmentSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFielddataFrequencyFilter_SetMinimum(t *testing.T) {
	value := 10.0
	initValue := -1.0

	type fields struct {
		min            *float64
		max            *float64
		minSegmentSize *int
	}
	type args struct {
		minimum float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FieldDataFrequencyFilter
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{minimum: value},
			want: &FieldDataFrequencyFilter{
				min: &value,
			},
		},
		{
			name:   "change",
			fields: fields{min: &initValue},
			args:   args{minimum: value},
			want: &FieldDataFrequencyFilter{
				min: &value,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fielddata := &FieldDataFrequencyFilter{
				min:            tt.fields.min,
				max:            tt.fields.max,
				minSegmentSize: tt.fields.minSegmentSize,
			}
			if got := fielddata.SetMinimum(tt.args.minimum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFielddataFrequencyFilter_Source(t *testing.T) {
	value := 10.0
	valueInt := 10

	type fields struct {
		min            *float64
		max            *float64
		minSegmentSize *int
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name:    "empty",
			fields:  fields{},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name:   "max",
			fields: fields{max: &value},
			want: map[string]interface{}{
				"max": value,
			},
			wantErr: false,
		},
		{
			name:   "min",
			fields: fields{min: &value},
			want: map[string]interface{}{
				"min": value,
			},
			wantErr: false,
		},
		{
			name:   "minSegmentSize",
			fields: fields{minSegmentSize: &valueInt},
			want: map[string]interface{}{
				"min_segment_size": valueInt,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fielddata := &FieldDataFrequencyFilter{
				min:            tt.fields.min,
				max:            tt.fields.max,
				minSegmentSize: tt.fields.minSegmentSize,
			}
			got, err := fielddata.Source()
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

func TestNewFielddataFrequencyFilter(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FieldDataFrequencyFilter
	}{
		{
			args: args{name: "test"},
			want: &FieldDataFrequencyFilter{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFieldDataFrequencyFilter(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFieldDataFrequencyFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
