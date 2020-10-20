package fields

import (
    "reflect"
    "testing"
)

func TestNested_GetType(t *testing.T) {
    type fields struct {
        dynamic    *Dynamic
        properties []*Field
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name: "type nested",
            fields: fields{},
            want: TypeNested,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Nested{
                dynamic:    tt.fields.dynamic,
                properties: tt.fields.properties,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNested_Properties(t *testing.T) {
    type fields struct {
        dynamic    *Dynamic
        properties []*Field
    }
    type args struct {
        properties []*Field
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Nested
    }{
        {
            name:   "set",
            fields: fields{},
            args: args{properties: []*Field{
                {name: "name", field: &Keyword{}},
            }},
            want: &Nested{
                properties: []*Field{
                    {name: "name", field: &Keyword{}},
                },
            },
        },
        {
            name:   "add",
            fields: fields{properties: []*Field{
                {name: "3gram", field: &Text{}},
            }},
            args: args{properties: []*Field{
                {name: "name", field: &Keyword{}},
            }},
            want: &Nested{
                properties: []*Field{
                    {name: "3gram", field: &Text{}},
                    {name: "name", field: &Keyword{}},
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Nested{
                dynamic:    tt.fields.dynamic,
                properties: tt.fields.properties,
            }
            if got := field.Properties(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Properties() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNested_SetDynamic(t *testing.T) {
    dynamicEnabled := DynamicEnabled
    dynamicDisabled := DynamicDisabled
    dynamicStrict := DynamicStrict

    type fields struct {
        dynamic    *Dynamic
        properties []*Field
    }
    type args struct {
        dynamic Dynamic
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Nested
    }{
        {
            name:   "DynamicEnabled",
            fields: fields{},
            args:   args{dynamic: dynamicEnabled},
            want:   &Nested{dynamic: &dynamicEnabled},
        },
        {
            name:   "DynamicDisabled",
            fields: fields{},
            args:   args{dynamic: dynamicDisabled},
            want:   &Nested{dynamic: &dynamicDisabled},
        },
        {
            name:   "DynamicStrict",
            fields: fields{},
            args:   args{dynamic: dynamicStrict},
            want:   &Nested{dynamic: &dynamicStrict},
        },
        {
            name:   "change",
            fields: fields{dynamic: &dynamicDisabled},
            args:   args{dynamic: dynamicStrict},
            want:   &Nested{dynamic: &dynamicStrict},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Nested{
                dynamic:    tt.fields.dynamic,
                properties: tt.fields.properties,
            }
            if got := field.SetDynamic(tt.args.dynamic); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetDynamic() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNested_Source(t *testing.T) {
    dynamicEnabled := DynamicEnabled
    dynamicDisabled := DynamicDisabled
    dynamicStrict := DynamicStrict
    dynamicError := Dynamic("error")

    type fields struct {
        dynamic    *Dynamic
        properties []*Field
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
            want:    map[string]interface{}{
                "type": TypeNested,
            },
            wantErr: false,
        },
        {
            name:    "dynamicEnabled",
            fields:  fields{dynamic: &dynamicEnabled},
            want:    map[string]interface{}{
                "type": TypeNested,
                "dynamic": true,
            },
            wantErr: false,
        },
        {
            name:    "dynamicDisabled",
            fields:  fields{dynamic: &dynamicDisabled},
            want:    map[string]interface{}{
                "type": TypeNested,
                "dynamic": false,
            },
            wantErr: false,
        },
        {
            name:    "dynamicStrict",
            fields:  fields{dynamic: &dynamicStrict},
            want:    map[string]interface{}{
                "type": TypeNested,
                "dynamic": dynamicStrict,
            },
            wantErr: false,
        },
        {
            name:    "dynamic error",
            fields:  fields{dynamic: &dynamicError},
            want:    nil,
            wantErr: true,
        },
        {
            name:    "properties",
            fields:  fields{properties: []*Field{
                {name: "name", field: &Keyword{}},
            }},
            want:    map[string]interface{}{
                "type": TypeNested,
                "properties": map[string]interface{}{
                    "name": map[string]interface{}{
                      "type": TypeKeyword,
                    },
                },
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Nested{
                dynamic:    tt.fields.dynamic,
                properties: tt.fields.properties,
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

func TestNewNested(t *testing.T) {
    type args struct {
        properties []*Field
    }
    tests := []struct {
        name string
        args args
        want *Nested
    }{
        {
            name: "new empty",
            args: args{},
            want: &Nested{},
        },
        {
            name: "new",
            args: args{properties: []*Field{
                {name: "name", field: &Keyword{}},
            }},
            want: &Nested{properties: []*Field{
                {name: "name", field: &Keyword{}},
            }},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewNested(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewNested() = %v, want %v", got, tt.want)
            }
        })
    }
}
