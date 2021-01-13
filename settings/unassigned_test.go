package settings

import (
	"reflect"
	"testing"
)

func TestNewUnassigned(t *testing.T) {
	type args struct {
		nodeLeft *NodeLeft
	}
	tests := []struct {
		name string
		args args
		want *Unassigned
	}{
		{
			args: args{nodeLeft: &NodeLeft{}},
			want: &Unassigned{nodeLeft: &NodeLeft{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUnassigned(tt.args.nodeLeft); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnassigned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnassigned_Source(t *testing.T) {
	nodeLeftSrc, _ := (&NodeLeft{delayedTimeout: &Interval{value: "1s"}}).Source()

	type fields struct {
		nodeLeft *NodeLeft
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{nodeLeft: &NodeLeft{delayedTimeout: &Interval{value: "1s"}}},
			want: map[string]interface{}{
				"node_left": nodeLeftSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unassigned := &Unassigned{
				nodeLeft: tt.fields.nodeLeft,
			}
			got, err := unassigned.Source()
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
