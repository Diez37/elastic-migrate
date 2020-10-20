package fields

import (
    "reflect"
    "testing"
)

func TestGeoPoint_GetType(t *testing.T) {
    type fields struct {
        ignoreMalformed *bool
        ignoreZValue    *bool
        nullValue       interface{}
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "type geo point",
            fields: fields{},
            want:   TypeGeoPoint,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &GeoPoint{
                ignoreMalformed: tt.fields.ignoreMalformed,
                ignoreZValue:    tt.fields.ignoreZValue,
                nullValue:       tt.fields.nullValue,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestGeoPoint_SetIgnoreMalformed(t *testing.T) {
    type fields struct {
        ignoreMalformed *bool
        ignoreZValue    *bool
        nullValue       interface{}
    }
    type args struct {
        ignoreMalformed bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *GeoPoint
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{ignoreMalformed: true},
            want:   &GeoPoint{ignoreMalformed: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{ignoreMalformed: false},
            want:   &GeoPoint{ignoreMalformed: &testFalse},
        },
        {
            name:   "change",
            fields: fields{ignoreMalformed: &testTrue},
            args:   args{ignoreMalformed: false},
            want:   &GeoPoint{ignoreMalformed: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &GeoPoint{
                ignoreMalformed: tt.fields.ignoreMalformed,
                ignoreZValue:    tt.fields.ignoreZValue,
                nullValue:       tt.fields.nullValue,
            }
            if got := field.SetIgnoreMalformed(tt.args.ignoreMalformed); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIgnoreMalformed() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestGeoPoint_SetIgnoreZValue(t *testing.T) {
    type fields struct {
        ignoreMalformed *bool
        ignoreZValue    *bool
        nullValue       interface{}
    }
    type args struct {
        ignoreZValue bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *GeoPoint
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{ignoreZValue: true},
            want:   &GeoPoint{ignoreZValue: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{ignoreZValue: false},
            want:   &GeoPoint{ignoreZValue: &testFalse},
        },
        {
            name:   "change",
            fields: fields{ignoreZValue: &testTrue},
            args:   args{ignoreZValue: false},
            want:   &GeoPoint{ignoreZValue: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &GeoPoint{
                ignoreMalformed: tt.fields.ignoreMalformed,
                ignoreZValue:    tt.fields.ignoreZValue,
                nullValue:       tt.fields.nullValue,
            }
            if got := field.SetIgnoreZValue(tt.args.ignoreZValue); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIgnoreZValue() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestGeoPoint_SetNullValue(t *testing.T) {
    type fields struct {
        ignoreMalformed *bool
        ignoreZValue    *bool
        nullValue       interface{}
    }
    type args struct {
        nullValue interface{}
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *GeoPoint
    }{
        {
            name:   "empty",
            fields: fields{},
            args:   args{},
            want:   &GeoPoint{nullValue: nil},
        },
        {
            name:   "int",
            fields: fields{},
            args:   args{nullValue: 1},
            want:   &GeoPoint{nullValue: 1},
        },
        {
            name:   "float",
            fields: fields{},
            args:   args{nullValue: 4.5},
            want:   &GeoPoint{nullValue: 4.5},
        },
        {
            name:   "string",
            fields: fields{},
            args:   args{nullValue: "test"},
            want:   &GeoPoint{nullValue: "test"},
        },
        {
            name:   "bool",
            fields: fields{},
            args:   args{nullValue: true},
            want:   &GeoPoint{nullValue: true},
        },
        {
            name:   "change",
            fields: fields{nullValue: "test"},
            args:   args{nullValue: true},
            want:   &GeoPoint{nullValue: true},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &GeoPoint{
                ignoreMalformed: tt.fields.ignoreMalformed,
                ignoreZValue:    tt.fields.ignoreZValue,
                nullValue:       tt.fields.nullValue,
            }
            if got := field.SetNullValue(tt.args.nullValue); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetNullValue() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestGeoPoint_Source(t *testing.T) {
    type fields struct {
        ignoreMalformed *bool
        ignoreZValue    *bool
        nullValue       interface{}
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
                "type": TypeGeoPoint,
            },
            wantErr: false,
        },
        {
            name:   "ignoreMalformed",
            fields: fields{ignoreMalformed: &testTrue},
            want: map[string]interface{}{
                "type":             TypeGeoPoint,
                "ignore_malformed": true,
            },
            wantErr: false,
        },
        {
            name:   "ignoreZValue",
            fields: fields{ignoreZValue: &testTrue},
            want: map[string]interface{}{
                "type":           TypeGeoPoint,
                "ignore_z_value": true,
            },
            wantErr: false,
        },
        {
            name:   "nullValue",
            fields: fields{nullValue: 2.7},
            want: map[string]interface{}{
                "type":       TypeGeoPoint,
                "null_value": 2.7,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &GeoPoint{
                ignoreMalformed: tt.fields.ignoreMalformed,
                ignoreZValue:    tt.fields.ignoreZValue,
                nullValue:       tt.fields.nullValue,
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

func TestNewGeoPoint(t *testing.T) {
    tests := []struct {
        name string
        want *GeoPoint
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewGeoPoint(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewGeoPoint() = %v, want %v", got, tt.want)
            }
        })
    }
}
