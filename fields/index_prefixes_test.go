package fields

import (
    "reflect"
    "testing"
)

func TestIndexPrefixes_SetMaximumChars(t *testing.T) {
    initValue := 1
    value := 2

    type fields struct {
        minChars *int
        maxChars *int
    }
    type args struct {
        maximumChars int
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *IndexPrefixes
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{maximumChars: value},
            want:   &IndexPrefixes{maxChars: &value},
        },
        {
            name:   "change",
            fields: fields{maxChars: &initValue},
            args:   args{maximumChars: value},
            want:   &IndexPrefixes{maxChars: &value},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            prefix := &IndexPrefixes{
                minChars: tt.fields.minChars,
                maxChars: tt.fields.maxChars,
            }
            if got := prefix.SetMaximumChars(tt.args.maximumChars); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetMaximumChars() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestIndexPrefixes_SetMinimumChars(t *testing.T) {
    initValue := 1
    value := 2

    type fields struct {
        minChars *int
        maxChars *int
    }
    type args struct {
        minimumChars int
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *IndexPrefixes
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{minimumChars: value},
            want:   &IndexPrefixes{minChars: &value},
        },
        {
            name:   "change",
            fields: fields{minChars: &initValue},
            args:   args{minimumChars: value},
            want:   &IndexPrefixes{minChars: &value},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            prefix := &IndexPrefixes{
                minChars: tt.fields.minChars,
                maxChars: tt.fields.maxChars,
            }
            if got := prefix.SetMinimumChars(tt.args.minimumChars); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetMinimumChars() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestIndexPrefixes_Source(t *testing.T) {
    value := 2

    type fields struct {
        minChars *int
        maxChars *int
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
            name:   "minChars",
            fields: fields{minChars: &value},
            want: map[string]interface{}{
                "min_chars": value,
            },
            wantErr: false,
        },
        {
            name:   "maxChars",
            fields: fields{maxChars: &value},
            want: map[string]interface{}{
                "max_chars": value,
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            prefix := &IndexPrefixes{
                minChars: tt.fields.minChars,
                maxChars: tt.fields.maxChars,
            }
            got, err := prefix.Source()
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

func TestNewIndexPrefixes(t *testing.T) {
    tests := []struct {
        name string
        want *IndexPrefixes
    }{
        {
            name: "new",
            want: &IndexPrefixes{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewIndexPrefixes(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewIndexPrefixes() = %v, want %v", got, tt.want)
            }
        })
    }
}
