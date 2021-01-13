package settings

import (
	"reflect"
	"testing"
)

func TestNewPolicy(t *testing.T) {
	tests := []struct {
		name string
		want *Policy
	}{
		{
			want: &Policy{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPolicy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_SetDeletesPctAllowed(t *testing.T) {
	numberSet := float32(1)
	numberChange := float32(2)

	type fields struct {
		maxMergeAtOnceExplicit *uint
		maxMergeAtOnce         *uint
		floorSegment           *Size
		maxMergedSegment       *Size
		reclaimDeletesWeight   *float32
		expungeDeletesAllowed  *float32
		segmentsPerTier        *float32
		deletesPctAllowed      *float32
	}
	type args struct {
		deletesPctAllowed float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Policy
	}{
		{
			name: "set",
			args: args{deletesPctAllowed: 1},
			want: &Policy{deletesPctAllowed: &numberSet},
		},
		{
			name: "change",
			fields: fields{deletesPctAllowed: &numberSet},
			args: args{deletesPctAllowed: 2},
			want: &Policy{deletesPctAllowed: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := &Policy{
				maxMergeAtOnceExplicit: tt.fields.maxMergeAtOnceExplicit,
				maxMergeAtOnce:         tt.fields.maxMergeAtOnce,
				floorSegment:           tt.fields.floorSegment,
				maxMergedSegment:       tt.fields.maxMergedSegment,
				reclaimDeletesWeight:   tt.fields.reclaimDeletesWeight,
				expungeDeletesAllowed:  tt.fields.expungeDeletesAllowed,
				segmentsPerTier:        tt.fields.segmentsPerTier,
				deletesPctAllowed:      tt.fields.deletesPctAllowed,
			}
			if got := policy.SetDeletesPctAllowed(tt.args.deletesPctAllowed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDeletesPctAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_SetExpungeDeletesAllowed(t *testing.T) {
	numberSet := float32(1)
	numberChange := float32(2)

	type fields struct {
		maxMergeAtOnceExplicit *uint
		maxMergeAtOnce         *uint
		floorSegment           *Size
		maxMergedSegment       *Size
		reclaimDeletesWeight   *float32
		expungeDeletesAllowed  *float32
		segmentsPerTier        *float32
		deletesPctAllowed      *float32
	}
	type args struct {
		expungeDeletesAllowed float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Policy
	}{
		{
			name: "set",
			args: args{expungeDeletesAllowed: 1},
			want: &Policy{expungeDeletesAllowed: &numberSet},
		},
		{
			name: "change",
			fields: fields{expungeDeletesAllowed: &numberSet},
			args: args{expungeDeletesAllowed: 2},
			want: &Policy{expungeDeletesAllowed: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := &Policy{
				maxMergeAtOnceExplicit: tt.fields.maxMergeAtOnceExplicit,
				maxMergeAtOnce:         tt.fields.maxMergeAtOnce,
				floorSegment:           tt.fields.floorSegment,
				maxMergedSegment:       tt.fields.maxMergedSegment,
				reclaimDeletesWeight:   tt.fields.reclaimDeletesWeight,
				expungeDeletesAllowed:  tt.fields.expungeDeletesAllowed,
				segmentsPerTier:        tt.fields.segmentsPerTier,
				deletesPctAllowed:      tt.fields.deletesPctAllowed,
			}
			if got := policy.SetExpungeDeletesAllowed(tt.args.expungeDeletesAllowed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExpungeDeletesAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_SetFloorSegment(t *testing.T) {
	type fields struct {
		maxMergeAtOnceExplicit *uint
		maxMergeAtOnce         *uint
		floorSegment           *Size
		maxMergedSegment       *Size
		reclaimDeletesWeight   *float32
		expungeDeletesAllowed  *float32
		segmentsPerTier        *float32
		deletesPctAllowed      *float32
	}
	type args struct {
		floorSegment *Size
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Policy
	}{
		{
			name: "set",
			args: args{floorSegment: NewSize("1mb")},
			want: &Policy{floorSegment: NewSize("1mb")},
		},
		{
			name: "change",
			fields: fields{floorSegment: NewSize("1mb")},
			args: args{floorSegment: NewSize("2mb")},
			want: &Policy{floorSegment: NewSize("2mb")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := &Policy{
				maxMergeAtOnceExplicit: tt.fields.maxMergeAtOnceExplicit,
				maxMergeAtOnce:         tt.fields.maxMergeAtOnce,
				floorSegment:           tt.fields.floorSegment,
				maxMergedSegment:       tt.fields.maxMergedSegment,
				reclaimDeletesWeight:   tt.fields.reclaimDeletesWeight,
				expungeDeletesAllowed:  tt.fields.expungeDeletesAllowed,
				segmentsPerTier:        tt.fields.segmentsPerTier,
				deletesPctAllowed:      tt.fields.deletesPctAllowed,
			}
			if got := policy.SetFloorSegment(tt.args.floorSegment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFloorSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_SetMaxMergeAtOnce(t *testing.T) {
	numberSet := uint(1)
	numberChange := uint(2)

	type fields struct {
		maxMergeAtOnceExplicit *uint
		maxMergeAtOnce         *uint
		floorSegment           *Size
		maxMergedSegment       *Size
		reclaimDeletesWeight   *float32
		expungeDeletesAllowed  *float32
		segmentsPerTier        *float32
		deletesPctAllowed      *float32
	}
	type args struct {
		maxMergeAtOnce uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Policy
	}{
		{
			name: "set",
			args: args{maxMergeAtOnce: 1},
			want: &Policy{maxMergeAtOnce: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxMergeAtOnce: &numberSet},
			args: args{maxMergeAtOnce: 2},
			want: &Policy{maxMergeAtOnce: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := &Policy{
				maxMergeAtOnceExplicit: tt.fields.maxMergeAtOnceExplicit,
				maxMergeAtOnce:         tt.fields.maxMergeAtOnce,
				floorSegment:           tt.fields.floorSegment,
				maxMergedSegment:       tt.fields.maxMergedSegment,
				reclaimDeletesWeight:   tt.fields.reclaimDeletesWeight,
				expungeDeletesAllowed:  tt.fields.expungeDeletesAllowed,
				segmentsPerTier:        tt.fields.segmentsPerTier,
				deletesPctAllowed:      tt.fields.deletesPctAllowed,
			}
			if got := policy.SetMaxMergeAtOnce(tt.args.maxMergeAtOnce); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxMergeAtOnce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_SetMaxMergeAtOnceExplicit(t *testing.T) {
	numberSet := uint(1)
	numberChange := uint(2)

	type fields struct {
		maxMergeAtOnceExplicit *uint
		maxMergeAtOnce         *uint
		floorSegment           *Size
		maxMergedSegment       *Size
		reclaimDeletesWeight   *float32
		expungeDeletesAllowed  *float32
		segmentsPerTier        *float32
		deletesPctAllowed      *float32
	}
	type args struct {
		maxMergeAtOnceExplicit uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Policy
	}{
		{
			name: "set",
			args: args{maxMergeAtOnceExplicit: 1},
			want: &Policy{maxMergeAtOnceExplicit: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxMergeAtOnceExplicit: &numberSet},
			args: args{maxMergeAtOnceExplicit: 2},
			want: &Policy{maxMergeAtOnceExplicit: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := &Policy{
				maxMergeAtOnceExplicit: tt.fields.maxMergeAtOnceExplicit,
				maxMergeAtOnce:         tt.fields.maxMergeAtOnce,
				floorSegment:           tt.fields.floorSegment,
				maxMergedSegment:       tt.fields.maxMergedSegment,
				reclaimDeletesWeight:   tt.fields.reclaimDeletesWeight,
				expungeDeletesAllowed:  tt.fields.expungeDeletesAllowed,
				segmentsPerTier:        tt.fields.segmentsPerTier,
				deletesPctAllowed:      tt.fields.deletesPctAllowed,
			}
			if got := policy.SetMaxMergeAtOnceExplicit(tt.args.maxMergeAtOnceExplicit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxMergeAtOnceExplicit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_SetMaxMergedSegment(t *testing.T) {
	type fields struct {
		maxMergeAtOnceExplicit *uint
		maxMergeAtOnce         *uint
		floorSegment           *Size
		maxMergedSegment       *Size
		reclaimDeletesWeight   *float32
		expungeDeletesAllowed  *float32
		segmentsPerTier        *float32
		deletesPctAllowed      *float32
	}
	type args struct {
		maxMergedSegment *Size
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Policy
	}{
		{
			name: "set",
			args: args{maxMergedSegment: NewSize("1mb")},
			want: &Policy{maxMergedSegment: NewSize("1mb")},
		},
		{
			name: "change",
			fields: fields{maxMergedSegment: NewSize("1mb")},
			args: args{maxMergedSegment: NewSize("2mb")},
			want: &Policy{maxMergedSegment: NewSize("2mb")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := &Policy{
				maxMergeAtOnceExplicit: tt.fields.maxMergeAtOnceExplicit,
				maxMergeAtOnce:         tt.fields.maxMergeAtOnce,
				floorSegment:           tt.fields.floorSegment,
				maxMergedSegment:       tt.fields.maxMergedSegment,
				reclaimDeletesWeight:   tt.fields.reclaimDeletesWeight,
				expungeDeletesAllowed:  tt.fields.expungeDeletesAllowed,
				segmentsPerTier:        tt.fields.segmentsPerTier,
				deletesPctAllowed:      tt.fields.deletesPctAllowed,
			}
			if got := policy.SetMaxMergedSegment(tt.args.maxMergedSegment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxMergedSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_SetReclaimDeletesWeight(t *testing.T) {
	numberSet := float32(1)
	numberChange := float32(2)

	type fields struct {
		maxMergeAtOnceExplicit *uint
		maxMergeAtOnce         *uint
		floorSegment           *Size
		maxMergedSegment       *Size
		reclaimDeletesWeight   *float32
		expungeDeletesAllowed  *float32
		segmentsPerTier        *float32
		deletesPctAllowed      *float32
	}
	type args struct {
		reclaimDeletesWeight float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Policy
	}{
		{
			name: "set",
			args: args{reclaimDeletesWeight: 1},
			want: &Policy{reclaimDeletesWeight: &numberSet},
		},
		{
			name: "change",
			fields: fields{reclaimDeletesWeight: &numberSet},
			args: args{reclaimDeletesWeight: 2},
			want: &Policy{reclaimDeletesWeight: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := &Policy{
				maxMergeAtOnceExplicit: tt.fields.maxMergeAtOnceExplicit,
				maxMergeAtOnce:         tt.fields.maxMergeAtOnce,
				floorSegment:           tt.fields.floorSegment,
				maxMergedSegment:       tt.fields.maxMergedSegment,
				reclaimDeletesWeight:   tt.fields.reclaimDeletesWeight,
				expungeDeletesAllowed:  tt.fields.expungeDeletesAllowed,
				segmentsPerTier:        tt.fields.segmentsPerTier,
				deletesPctAllowed:      tt.fields.deletesPctAllowed,
			}
			if got := policy.SetReclaimDeletesWeight(tt.args.reclaimDeletesWeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReclaimDeletesWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_SetSegmentsPerTier(t *testing.T) {
	numberSet := float32(1)
	numberChange := float32(2)

	type fields struct {
		maxMergeAtOnceExplicit *uint
		maxMergeAtOnce         *uint
		floorSegment           *Size
		maxMergedSegment       *Size
		reclaimDeletesWeight   *float32
		expungeDeletesAllowed  *float32
		segmentsPerTier        *float32
		deletesPctAllowed      *float32
	}
	type args struct {
		segmentsPerTier float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Policy
	}{
		{
			name: "set",
			args: args{segmentsPerTier: 1},
			want: &Policy{segmentsPerTier: &numberSet},
		},
		{
			name: "change",
			fields: fields{segmentsPerTier: &numberSet},
			args: args{segmentsPerTier: 2},
			want: &Policy{segmentsPerTier: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := &Policy{
				maxMergeAtOnceExplicit: tt.fields.maxMergeAtOnceExplicit,
				maxMergeAtOnce:         tt.fields.maxMergeAtOnce,
				floorSegment:           tt.fields.floorSegment,
				maxMergedSegment:       tt.fields.maxMergedSegment,
				reclaimDeletesWeight:   tt.fields.reclaimDeletesWeight,
				expungeDeletesAllowed:  tt.fields.expungeDeletesAllowed,
				segmentsPerTier:        tt.fields.segmentsPerTier,
				deletesPctAllowed:      tt.fields.deletesPctAllowed,
			}
			if got := policy.SetSegmentsPerTier(tt.args.segmentsPerTier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSegmentsPerTier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_Source(t *testing.T) {
	numberFloat := float32(1)
	numberUint := uint(1)
	sizeSrc, _ := (&Size{value: "1mb"}).Source()

	type fields struct {
		maxMergeAtOnceExplicit *uint
		maxMergeAtOnce         *uint
		floorSegment           *Size
		maxMergedSegment       *Size
		reclaimDeletesWeight   *float32
		expungeDeletesAllowed  *float32
		segmentsPerTier        *float32
		deletesPctAllowed      *float32
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
			name: "max_merge_at_once_explicit",
			fields: fields{maxMergeAtOnceExplicit: &numberUint},
			want: map[string]interface{}{
				"max_merge_at_once_explicit": uint(1),
			},
		},
		{
			name: "max_merge_at_once",
			fields: fields{maxMergeAtOnce: &numberUint},
			want: map[string]interface{}{
				"max_merge_at_once": uint(1),
			},
		},
		{
			name: "reclaim_deletes_weight",
			fields: fields{reclaimDeletesWeight: &numberFloat},
			want: map[string]interface{}{
				"reclaim_deletes_weight": float32(1),
			},
		},
		{
			name: "expunge_deletes_allowed",
			fields: fields{expungeDeletesAllowed: &numberFloat},
			want: map[string]interface{}{
				"expunge_deletes_allowed": float32(1),
			},
		},
		{
			name: "segments_per_tier",
			fields: fields{segmentsPerTier: &numberFloat},
			want: map[string]interface{}{
				"segments_per_tier": float32(1),
			},
		},
		{
			name: "deletes_pct_allowed",
			fields: fields{deletesPctAllowed: &numberFloat},
			want: map[string]interface{}{
				"deletes_pct_allowed": float32(1),
			},
		},
		{
			name: "floor_segment",
			fields: fields{floorSegment: &Size{value: "1mb"}},
			want: map[string]interface{}{
				"floor_segment": sizeSrc,
			},
		},
		{
			name: "max_merged_segment",
			fields: fields{maxMergedSegment: &Size{value: "1mb"}},
			want: map[string]interface{}{
				"max_merged_segment": sizeSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policy := &Policy{
				maxMergeAtOnceExplicit: tt.fields.maxMergeAtOnceExplicit,
				maxMergeAtOnce:         tt.fields.maxMergeAtOnce,
				floorSegment:           tt.fields.floorSegment,
				maxMergedSegment:       tt.fields.maxMergedSegment,
				reclaimDeletesWeight:   tt.fields.reclaimDeletesWeight,
				expungeDeletesAllowed:  tt.fields.expungeDeletesAllowed,
				segmentsPerTier:        tt.fields.segmentsPerTier,
				deletesPctAllowed:      tt.fields.deletesPctAllowed,
			}
			got, err := policy.Source()
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
