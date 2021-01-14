package fields

import (
    "reflect"
    "testing"
)

func TestNewObject(t *testing.T) {
    type args struct {
        properties []*Field
    }
    tests := []struct {
        name string
        args args
        want *Object
    }{
        {
            name: "object",
            args: args{},
            want: &Object{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewObject(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewObject() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestObject_GetType(t *testing.T) {
    type fields struct {
        enabled    *bool
        dynamic    *Dynamic
        properties []*Field
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "object type",
            fields: fields{},
            want:   TypeObject,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Object{
                enabled:    tt.fields.enabled,
                dynamic:    tt.fields.dynamic,
                properties: tt.fields.properties,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestObject_Properties(t *testing.T) {
    type fields struct {
        enabled    *bool
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
        want   *Object
    }{
        {
            name:   "set",
            fields: fields{},
            args: args{properties: []*Field{{
                name:  "test",
                field: &Keyword{},
            }}},
            want: &Object{
                properties: []*Field{{
                    name:  "test",
                    field: &Keyword{},
                }},
            },
        },
        {
            name: "add",
            fields: fields{properties: []*Field{{
                name:  "test1",
                field: &Keyword{},
            }}},
            args: args{properties: []*Field{{
                name:  "test",
                field: &Keyword{},
            }}},
            want: &Object{
                properties: []*Field{
                    {
                        name:  "test1",
                        field: &Keyword{},
                    },
                    {
                        name:  "test",
                        field: &Keyword{},
                    },
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Object{
                enabled:    tt.fields.enabled,
                dynamic:    tt.fields.dynamic,
                properties: tt.fields.properties,
            }
            if got := field.Properties(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Properties() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestObject_SetDynamic(t *testing.T) {
    initDynamic := DynamicEnabled
    changeDynamic := DynamicDisabled

    type fields struct {
        enabled    *bool
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
        want   *Object
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{dynamic: initDynamic},
            want:   &Object{dynamic: &initDynamic},
        },
        {
            name:   "change",
            fields: fields{dynamic: &initDynamic},
            args:   args{dynamic: changeDynamic},
            want:   &Object{dynamic: &changeDynamic},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Object{
                enabled:    tt.fields.enabled,
                dynamic:    tt.fields.dynamic,
                properties: tt.fields.properties,
            }
            if got := field.SetDynamic(tt.args.dynamic); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetDynamic() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestObject_SetEnabled(t *testing.T) {
    boolTrue := true
    boolFalse := false

    type fields struct {
        enabled    *bool
        dynamic    *Dynamic
        properties []*Field
    }
    type args struct {
        enabled bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Object
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{enabled: boolTrue},
            want:   &Object{enabled: &boolTrue},
        },
        {
            name:   "change",
            fields: fields{enabled: &boolFalse},
            args:   args{enabled: boolTrue},
            want:   &Object{enabled: &boolTrue},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Object{
                enabled:    tt.fields.enabled,
                dynamic:    tt.fields.dynamic,
                properties: tt.fields.properties,
            }
            if got := field.SetEnabled(tt.args.enabled); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetEnabled() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestObject_Source(t *testing.T) {
    dynamic := DynamicEnabled
    dynamicTest := Dynamic("test")
    boolTrue := true

    type fields struct {
        enabled    *bool
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
            want:    map[string]interface{}{},
            wantErr: false,
        },
        {
            name:    "dynamic",
            fields:  fields{dynamic: &dynamic},
            want:    map[string]interface{}{
                "dynamic": true,
            },
            wantErr: false,
        },
        {
            name:    "dynamicError",
            fields:  fields{dynamic: &dynamicTest},
            want:    nil,
            wantErr: true,
        },
        {
            name:    "enabled",
            fields:  fields{enabled: &boolTrue},
            want:    map[string]interface{}{
                "enabled": true,
            },
            wantErr: false,
        },
        {
            name:    "properties",
            fields:  fields{properties: []*Field{{
                name: "test",
                field: &Keyword{},
            }}},
            want:    map[string]interface{}{
                "properties": map[string]interface{}{
                    "test": map[string]interface{}{
                        "type": TypeKeyword,
                    },
                },
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Object{
                enabled:    tt.fields.enabled,
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
