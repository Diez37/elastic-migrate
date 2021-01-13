package settings

import (
	"reflect"
	"testing"
)

func TestNewNodeLeft(t *testing.T) {
	type args struct {
		delayedTimeout *Interval
	}
	tests := []struct {
		name string
		args args
		want *NodeLeft
	}{
		{
			args: args{delayedTimeout: NewInterval("1s")},
			want: &NodeLeft{delayedTimeout: NewInterval("1s")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNodeLeft(tt.args.delayedTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNodeLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeLeft_Source(t *testing.T) {
	intervalSrc, _ := NewInterval("1s").Source()

	type fields struct {
		delayedTimeout *Interval
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{delayedTimeout: NewInterval("1s")},
			want: map[string]interface{}{
				"delayed_timeout": intervalSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodeLeft := &NodeLeft{
				delayedTimeout: tt.fields.delayedTimeout,
			}
			got, err := nodeLeft.Source()
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
