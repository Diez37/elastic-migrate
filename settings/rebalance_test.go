package settings

import (
	"reflect"
	"testing"
)

func TestNewRebalance(t *testing.T) {
	type args struct {
		enable RebalanceVal
	}
	tests := []struct {
		name string
		args args
		want *Rebalance
	}{
		{
			name: "rebalance_all",
			args: args{enable: RebalanceAll},
			want: &Rebalance{enable: RebalanceAll},
		},
		{
			name: "rebalance_primaries",
			args: args{enable: RebalancePrimaries},
			want: &Rebalance{enable: RebalancePrimaries},
		},
		{
			name: "rebalance_replicas",
			args: args{enable: RebalanceReplicas},
			want: &Rebalance{enable: RebalanceReplicas},
		},
		{
			name: "rebalance_none",
			args: args{enable: RebalanceNone},
			want: &Rebalance{enable: RebalanceNone},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRebalance(tt.args.enable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRebalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRebalance_Source(t *testing.T) {
	type fields struct {
		enable RebalanceVal
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "rebalance_all",
			fields: fields{enable: RebalanceAll},
			want: map[string]interface{}{
				"enable": RebalanceAll,
			},
		},
		{
			name: "rebalance_primaries",
			fields: fields{enable: RebalancePrimaries},
			want: map[string]interface{}{
				"enable": RebalancePrimaries,
			},
		},
		{
			name: "rebalance_replicas",
			fields: fields{enable: RebalanceReplicas},
			want: map[string]interface{}{
				"enable": RebalanceReplicas,
			},
		},
		{
			name: "rebalance_none",
			fields: fields{enable: RebalanceNone},
			want: map[string]interface{}{
				"enable": RebalanceNone,
			},
		},
		{
			name: "empty",
			fields: fields{},
			want: map[string]interface{}{
				"enable": RebalanceVal(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rebalance := &Rebalance{
				enable: tt.fields.enable,
			}
			got, err := rebalance.Source()
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
