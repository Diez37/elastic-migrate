package settings

import (
	"reflect"
	"testing"
)

func TestNewSoftDeletes(t *testing.T) {
	tests := []struct {
		name string
		want *SoftDeletes
	}{
		{
			want: &SoftDeletes{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSoftDeletes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSoftDeletes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSoftDeletes_SetEnabled(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		enabled        *bool
		retention      *Retention
		retentionLease *RetentionLease
	}
	type args struct {
		enabled bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SoftDeletes
	}{
		{
			name: "set",
			args: args{enabled: true},
			want: &SoftDeletes{enabled: &boolTrue},
		},
		{
			name: "set",
			fields: fields{enabled: &boolTrue},
			args: args{enabled: false},
			want: &SoftDeletes{enabled: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			softDeletes := &SoftDeletes{
				enabled:        tt.fields.enabled,
				retention:      tt.fields.retention,
				retentionLease: tt.fields.retentionLease,
			}
			if got := softDeletes.SetEnabled(tt.args.enabled); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSoftDeletes_SetRetention(t *testing.T) {
	type fields struct {
		enabled        *bool
		retention      *Retention
		retentionLease *RetentionLease
	}
	type args struct {
		retention *Retention
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SoftDeletes
	}{
		{
			args: args{retention: &Retention{}},
			want: &SoftDeletes{retention: &Retention{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			softDeletes := &SoftDeletes{
				enabled:        tt.fields.enabled,
				retention:      tt.fields.retention,
				retentionLease: tt.fields.retentionLease,
			}
			if got := softDeletes.SetRetention(tt.args.retention); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRetention() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSoftDeletes_SetRetentionLease(t *testing.T) {
	type fields struct {
		enabled        *bool
		retention      *Retention
		retentionLease *RetentionLease
	}
	type args struct {
		retentionLease *RetentionLease
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SoftDeletes
	}{
		{
			args: args{retentionLease: &RetentionLease{}},
			want: &SoftDeletes{retentionLease: &RetentionLease{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			softDeletes := &SoftDeletes{
				enabled:        tt.fields.enabled,
				retention:      tt.fields.retention,
				retentionLease: tt.fields.retentionLease,
			}
			if got := softDeletes.SetRetentionLease(tt.args.retentionLease); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRetentionLease() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSoftDeletes_Source(t *testing.T) {
	boolTrue := true
	retentionSrc, _ := (&Retention{}).Source()
	retentionLeaseSrc, _ := (&RetentionLease{period: &Interval{value: "1s"}}).Source()

	type fields struct {
		enabled        *bool
		retention      *Retention
		retentionLease *RetentionLease
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
			name: "enabled",
			fields: fields{enabled: &boolTrue},
			want: map[string]interface{}{
				"enabled": true,
			},
		},
		{
			name: "retention",
			fields: fields{retention: &Retention{}},
			want: map[string]interface{}{
				"retention": retentionSrc,
			},
		},
		{
			name: "retention_lease",
			fields: fields{retentionLease: &RetentionLease{period: &Interval{value: "1s"}}},
			want: map[string]interface{}{
				"retention_lease": retentionLeaseSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			softDeletes := &SoftDeletes{
				enabled:        tt.fields.enabled,
				retention:      tt.fields.retention,
				retentionLease: tt.fields.retentionLease,
			}
			got, err := softDeletes.Source()
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
