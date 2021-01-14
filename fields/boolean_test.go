package fields

import (
    "reflect"
    "testing"
)

func TestBoolean_GetType(t *testing.T) {
    type fields struct {
        docValues *bool
        store     *bool
        index     *bool
        boost     *float64
        nullValue interface{}
        meta      []*Meta
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "type boolean",
            fields: fields{},
            want:   TypeBoolean,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Boolean{
                docValues: tt.fields.docValues,
                store:     tt.fields.store,
                index:     tt.fields.index,
                boost:     tt.fields.boost,
                nullValue: tt.fields.nullValue,
                meta:      tt.fields.meta,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBoolean_Meta(t *testing.T) {
    type fields struct {
        docValues *bool
        store     *bool
        index     *bool
        boost     *float64
        nullValue interface{}
        meta      []*Meta
    }
    type args struct {
        metas []*Meta
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Boolean
    }{
        {
            name:   "empty",
            fields: fields{},
            args:   args{},
            want:   &Boolean{},
        },
        {
            name:   "set meta",
            fields: fields{},
            args: args{metas: []*Meta{
                {
                    name:  "version",
                    value: "0.0.1",
                },
                {
                    name:  "boost",
                    value: 3.5,
                },
            }},
            want: &Boolean{meta: []*Meta{
                {
                    name:  "version",
                    value: "0.0.1",
                },
                {
                    name:  "boost",
                    value: 3.5,
                },
            }},
        },
        {
            name: "add meta",
            fields: fields{meta: []*Meta{
                {
                    name:  "author",
                    value: "diez37",
                },
                {
                    name:  "year",
                    value: 2020,
                },
            }},
            args: args{metas: []*Meta{
                {
                    name:  "version",
                    value: "0.0.1",
                },
                {
                    name:  "boost",
                    value: 3.5,
                },
            }},
            want: &Boolean{meta: []*Meta{
                {
                    name:  "author",
                    value: "diez37",
                },
                {
                    name:  "year",
                    value: 2020,
                },
                {
                    name:  "version",
                    value: "0.0.1",
                },
                {
                    name:  "boost",
                    value: 3.5,
                },
            }},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Boolean{
                docValues: tt.fields.docValues,
                store:     tt.fields.store,
                index:     tt.fields.index,
                boost:     tt.fields.boost,
                nullValue: tt.fields.nullValue,
                meta:      tt.fields.meta,
            }
            if got := field.Meta(tt.args.metas...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Meta() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBoolean_SetBoost(t *testing.T) {
    initBoost := 9.0
    setBoost := 5.4

    type fields struct {
        docValues *bool
        store     *bool
        index     *bool
        boost     *float64
        nullValue interface{}
        meta      []*Meta
    }
    type args struct {
        boost float64
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Boolean
    }{
        {
            name:   "set boost",
            fields: fields{},
            args:   args{boost: setBoost},
            want:   &Boolean{boost: &setBoost},
        },
        {
            name:   "change boost",
            fields: fields{boost: &initBoost},
            args:   args{boost: setBoost},
            want:   &Boolean{boost: &setBoost},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Boolean{
                docValues: tt.fields.docValues,
                store:     tt.fields.store,
                index:     tt.fields.index,
                boost:     tt.fields.boost,
                nullValue: tt.fields.nullValue,
                meta:      tt.fields.meta,
            }
            if got := field.SetBoost(tt.args.boost); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetBoost() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBoolean_SetDocValues(t *testing.T) {
    boolTrue := true
    boolFalse := false

    type fields struct {
        docValues *bool
        store     *bool
        index     *bool
        boost     *float64
        nullValue interface{}
        meta      []*Meta
    }
    type args struct {
        docValues bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Boolean
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{docValues: true},
            want:   &Boolean{docValues: &boolTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{docValues: false},
            want:   &Boolean{docValues: &boolFalse},
        },
        {
            name:   "change",
            fields: fields{docValues: &boolTrue},
            args:   args{docValues: false},
            want:   &Boolean{docValues: &boolFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Boolean{
                docValues: tt.fields.docValues,
                store:     tt.fields.store,
                index:     tt.fields.index,
                boost:     tt.fields.boost,
                nullValue: tt.fields.nullValue,
                meta:      tt.fields.meta,
            }
            if got := field.SetDocValues(tt.args.docValues); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetDocValues() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBoolean_SetIndex(t *testing.T) {
    boolTrue := true
    boolFalse := false

    type fields struct {
        docValues *bool
        store     *bool
        index     *bool
        boost     *float64
        nullValue interface{}
        meta      []*Meta
    }
    type args struct {
        index bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Boolean
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{index: true},
            want:   &Boolean{index: &boolTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{index: false},
            want:   &Boolean{index: &boolFalse},
        },
        {
            name:   "change",
            fields: fields{index: &boolTrue},
            args:   args{index: false},
            want:   &Boolean{index: &boolFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Boolean{
                docValues: tt.fields.docValues,
                store:     tt.fields.store,
                index:     tt.fields.index,
                boost:     tt.fields.boost,
                nullValue: tt.fields.nullValue,
                meta:      tt.fields.meta,
            }
            if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIndex() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBoolean_SetNullValue(t *testing.T) {
    type fields struct {
        docValues *bool
        store     *bool
        index     *bool
        boost     *float64
        nullValue interface{}
        meta      []*Meta
    }
    type args struct {
        nullValue interface{}
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Boolean
    }{
        {
            name:   "empty",
            fields: fields{},
            args:   args{},
            want:   &Boolean{nullValue: nil},
        },
        {
            name:   "int",
            fields: fields{},
            args:   args{nullValue: 1},
            want:   &Boolean{nullValue: 1},
        },
        {
            name:   "float",
            fields: fields{},
            args:   args{nullValue: 4.5},
            want:   &Boolean{nullValue: 4.5},
        },
        {
            name:   "string",
            fields: fields{},
            args:   args{nullValue: "test"},
            want:   &Boolean{nullValue: "test"},
        },
        {
            name:   "bool",
            fields: fields{},
            args:   args{nullValue: true},
            want:   &Boolean{nullValue: true},
        },
        {
            name:   "change",
            fields: fields{nullValue: "test"},
            args:   args{nullValue: true},
            want:   &Boolean{nullValue: true},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Boolean{
                docValues: tt.fields.docValues,
                store:     tt.fields.store,
                index:     tt.fields.index,
                boost:     tt.fields.boost,
                nullValue: tt.fields.nullValue,
                meta:      tt.fields.meta,
            }
            if got := field.SetNullValue(tt.args.nullValue); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetNullValue() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBoolean_SetStore(t *testing.T) {
    boolTrue := true
    boolFalse := false

    type fields struct {
        docValues *bool
        store     *bool
        index     *bool
        boost     *float64
        nullValue interface{}
        meta      []*Meta
    }
    type args struct {
        store bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Boolean
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{store: true},
            want:   &Boolean{store: &boolTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{store: false},
            want:   &Boolean{store: &boolFalse},
        },
        {
            name:   "change",
            fields: fields{store: &boolTrue},
            args:   args{store: false},
            want:   &Boolean{store: &boolFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Boolean{
                docValues: tt.fields.docValues,
                store:     tt.fields.store,
                index:     tt.fields.index,
                boost:     tt.fields.boost,
                nullValue: tt.fields.nullValue,
                meta:      tt.fields.meta,
            }
            if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetStore() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBoolean_Source(t *testing.T) {
    boost := 2.5
    boolTrue := true

    type fields struct {
        docValues *bool
        store     *bool
        index     *bool
        boost     *float64
        nullValue interface{}
        meta      []*Meta
    }
    tests := []struct {
        name    string
        fields  fields
        want    interface{}
        wantErr bool
    }{
        {
            name:   "empty",
            fields: fields{},
            want: map[string]interface{}{
                "type": TypeBoolean,
            },
        },
        {
            name:   "docValues",
            fields: fields{docValues: &boolTrue},
            want: map[string]interface{}{
                "type":       TypeBoolean,
                "doc_values": true,
            },
        },
        {
            name:   "store",
            fields: fields{store: &boolTrue},
            want: map[string]interface{}{
                "type":  TypeBoolean,
                "store": true,
            },
        },
        {
            name:   "boost",
            fields: fields{boost: &boost},
            want: map[string]interface{}{
                "type":  TypeBoolean,
                "boost": boost,
            },
        },
        {
            name:   "index",
            fields: fields{index: &boolTrue},
            want: map[string]interface{}{
                "type":  TypeBoolean,
                "index": true,
            },
        },
        {
            name:   "nullValue",
            fields: fields{nullValue: 2.7},
            want: map[string]interface{}{
                "type":       TypeBoolean,
                "null_value": 2.7,
            },
        },
        {
            name: "meta",
            fields: fields{meta: []*Meta{
                {
                    name:  "author",
                    value: "diez37",
                },
                {
                    name:  "version",
                    value: "0.0.1",
                },
            }},
            want: map[string]interface{}{
                "type": TypeBoolean,
                "meta": map[string]interface{}{
                    "author":  "diez37",
                    "version": "0.0.1",
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Boolean{
                docValues: tt.fields.docValues,
                store:     tt.fields.store,
                index:     tt.fields.index,
                boost:     tt.fields.boost,
                nullValue: tt.fields.nullValue,
                meta:      tt.fields.meta,
            }
            got, err := field.Source()
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

func TestNewBoolean(t *testing.T) {
    tests := []struct {
        name string
        want *Boolean
    }{
        {
            name: "new",
            want: &Boolean{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewBoolean(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewBoolean() = %v, want %v", got, tt.want)
            }
        })
    }
}
