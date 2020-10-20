package fields

import (
    "reflect"
    "testing"
)

func TestNewMeta(t *testing.T) {
    type args struct {
        name  string
        value interface{}
    }
    tests := []struct {
        name string
        args args
        want *Meta
    }{
        {
            name: "new",
            args: args{
                name:  "author",
                value: "diez37",
            },
            want: &Meta{
                name:  "author",
                value: "diez37",
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewMeta(tt.args.name, tt.args.value); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewMeta() = %v, want %v", got, tt.want)
            }
        })
    }
}
