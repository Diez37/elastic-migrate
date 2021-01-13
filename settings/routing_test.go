package settings

import (
	"reflect"
	"testing"
)

func TestNewRouting(t *testing.T) {
	tests := []struct {
		name string
		want *Routing
	}{
		{
			want: &Routing{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouting(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouting_SetAllocation(t *testing.T) {
	type fields struct {
		allocation *Allocation
		rebalance  *Rebalance
	}
	type args struct {
		allocation *Allocation
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Routing
	}{
		{
			name: "set",
			args: args{allocation: &Allocation{}},
			want: &Routing{allocation: &Allocation{}},
		},
		{
			name: "change",
			fields: fields{allocation: &Allocation{}},
			args: args{allocation: &Allocation{include: AllocationAttributes{AllocationAttributeName: []interface{}{"test"}}}},
			want: &Routing{allocation: &Allocation{include: AllocationAttributes{AllocationAttributeName: []interface{}{"test"}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routing := &Routing{
				allocation: tt.fields.allocation,
				rebalance:  tt.fields.rebalance,
			}
			if got := routing.SetAllocation(tt.args.allocation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAllocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouting_SetRebalance(t *testing.T) {
	type fields struct {
		allocation *Allocation
		rebalance  *Rebalance
	}
	type args struct {
		rebalance *Rebalance
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Routing
	}{
		{
			name: "set",
			args: args{rebalance: &Rebalance{enable: RebalanceAll}},
			want: &Routing{rebalance: &Rebalance{enable: RebalanceAll}},
		},
		{
			name: "change",
			fields: fields{rebalance: &Rebalance{enable: RebalancePrimaries}},
			args: args{rebalance: &Rebalance{enable: RebalanceAll}},
			want: &Routing{rebalance: &Rebalance{enable: RebalanceAll}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routing := &Routing{
				allocation: tt.fields.allocation,
				rebalance:  tt.fields.rebalance,
			}
			if got := routing.SetRebalance(tt.args.rebalance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRebalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouting_Source(t *testing.T) {
	allocationSrc, _ := (&Allocation{include: AllocationAttributes{AllocationAttributeName: []interface{}{"test"}}}).Source()
	rebalanceSrc, _ := (&Rebalance{enable: RebalanceAll}).Source()

	type fields struct {
		allocation *Allocation
		rebalance  *Rebalance
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "allocation",
			fields: fields{allocation: &Allocation{include: AllocationAttributes{AllocationAttributeName: []interface{}{"test"}}}},
			want: map[string]interface{}{
				"allocation": allocationSrc,
			},
		},
		{
			name: "rebalance",
			fields: fields{rebalance: &Rebalance{enable: RebalanceAll}},
			want: map[string]interface{}{
				"rebalance": rebalanceSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routing := &Routing{
				allocation: tt.fields.allocation,
				rebalance:  tt.fields.rebalance,
			}
			got, err := routing.Source()
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
