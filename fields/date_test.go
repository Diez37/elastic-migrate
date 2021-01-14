package fields

import (
    "reflect"
    "testing"
)

func TestDate_Formats(t *testing.T) {
    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    type args struct {
        formats []DateFormat
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Date
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{formats: []DateFormat{DateFormatANSIC, DateFormatUnixDate}},
            want:   &Date{formats: []DateFormat{DateFormatANSIC, DateFormatUnixDate}},
        },
        {
            name:   "add",
            fields: fields{formats: []DateFormat{DateFormatANSIC, DateFormatUnixDate}},
            args: args{formats: []DateFormat{
                DateFormatRubyDate,
                DateFormatRFC822,
                DateFormatRFC822Z,
                DateFormatRFC850,
                DateFormatRFC1123,
                DateFormatRFC1123Z,
                DateFormatRFC3339,
                DateFormatKitchen,
                DateFormatStamp,
                DateFormatStampMilli,
            }},
            want: &Date{formats: []DateFormat{
                DateFormatANSIC,
                DateFormatUnixDate,
                DateFormatRubyDate,
                DateFormatRFC822,
                DateFormatRFC822Z,
                DateFormatRFC850,
                DateFormatRFC1123,
                DateFormatRFC1123Z,
                DateFormatRFC3339,
                DateFormatKitchen,
                DateFormatStamp,
                DateFormatStampMilli,
            }},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.Formats(tt.args.formats...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Formats() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_GetType(t *testing.T) {
    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "type date",
            fields: fields{_type: TypeDate},
            want:   TypeDate,
        },
        {
            name:   "type date nano",
            fields: fields{_type: TypeDateNano},
            want:   TypeDateNano,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_Local(t *testing.T) {
    testLocalRoot := LocalRoot
    testLocalCanada := LocalCanada

    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    type args struct {
        locale Local
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Date
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{locale: LocalRoot},
            want:   &Date{locale: &testLocalRoot},
        },
        {
            name:   "change",
            fields: fields{locale: &testLocalRoot},
            args:   args{locale: LocalCanada},
            want:   &Date{locale: &testLocalCanada},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.Local(tt.args.locale); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Local() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_Meta(t *testing.T) {
    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    type args struct {
        metas []*Meta
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Date
    }{
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
            want: &Date{meta: []*Meta{
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
            want: &Date{meta: []*Meta{
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
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.Meta(tt.args.metas...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Meta() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_SetBoost(t *testing.T) {
    initBoost := float64(9)
    setBoost := 5.4

    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    type args struct {
        boost float64
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Date
    }{
        {
            name:   "set boost",
            fields: fields{},
            args:   args{boost: setBoost},
            want:   &Date{boost: &setBoost},
        },
        {
            name:   "change boost",
            fields: fields{boost: &initBoost},
            args:   args{boost: setBoost},
            want:   &Date{boost: &setBoost},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.SetBoost(tt.args.boost); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetBoost() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_SetDocValues(t *testing.T) {
    boolTrue := true
    boolFalse := false

    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    type args struct {
        docValues bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Date
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{docValues: true},
            want:   &Date{docValues: &boolTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{docValues: false},
            want:   &Date{docValues: &boolFalse},
        },
        {
            name:   "change",
            fields: fields{docValues: &boolTrue},
            args:   args{docValues: false},
            want:   &Date{docValues: &boolFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.SetDocValues(tt.args.docValues); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetDocValues() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_SetIgnoreMalformed(t *testing.T) {
    boolTrue := true
    boolFalse := false

    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    type args struct {
        ignoreMalformed bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Date
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{ignoreMalformed: true},
            want:   &Date{ignoreMalformed: &boolTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{ignoreMalformed: false},
            want:   &Date{ignoreMalformed: &boolFalse},
        },
        {
            name:   "change",
            fields: fields{ignoreMalformed: &boolTrue},
            args:   args{ignoreMalformed: false},
            want:   &Date{ignoreMalformed: &boolFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.SetIgnoreMalformed(tt.args.ignoreMalformed); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIgnoreMalformed() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_SetIndex(t *testing.T) {
    boolTrue := true
    boolFalse := false

    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    type args struct {
        index bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Date
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{index: true},
            want:   &Date{index: &boolTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{index: false},
            want:   &Date{index: &boolFalse},
        },
        {
            name:   "change",
            fields: fields{index: &boolTrue},
            args:   args{index: false},
            want:   &Date{index: &boolFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIndex() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_SetNullValue(t *testing.T) {
    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    type args struct {
        nullValue interface{}
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Date
    }{
        {
            name:   "empty",
            fields: fields{},
            args:   args{},
            want:   &Date{nullValue: nil},
        },
        {
            name:   "int",
            fields: fields{},
            args:   args{nullValue: 1},
            want:   &Date{nullValue: 1},
        },
        {
            name:   "float",
            fields: fields{},
            args:   args{nullValue: 4.5},
            want:   &Date{nullValue: 4.5},
        },
        {
            name:   "string",
            fields: fields{},
            args:   args{nullValue: "test"},
            want:   &Date{nullValue: "test"},
        },
        {
            name:   "bool",
            fields: fields{},
            args:   args{nullValue: true},
            want:   &Date{nullValue: true},
        },
        {
            name:   "change",
            fields: fields{nullValue: "test"},
            args:   args{nullValue: true},
            want:   &Date{nullValue: true},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.SetNullValue(tt.args.nullValue); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetNullValue() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_SetStore(t *testing.T) {
    boolTrue := true
    boolFalse := false

    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    type args struct {
        store bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Date
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{store: true},
            want:   &Date{store: &boolTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{store: false},
            want:   &Date{store: &boolFalse},
        },
        {
            name:   "change",
            fields: fields{store: &boolTrue},
            args:   args{store: false},
            want:   &Date{store: &boolFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
            }
            if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetStore() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDate_Source(t *testing.T) {
    boost := 2.5
    testLocalRoot := LocalRoot
    boolTrue := true

    type fields struct {
        docValues       *bool
        ignoreMalformed *bool
        index           *bool
        store           *bool
        boost           *float64
        locale          *Local
        formats         []DateFormat
        nullValue       interface{}
        meta            []*Meta
        _type           Type
    }
    tests := []struct {
        name    string
        fields  fields
        want    interface{}
        wantErr bool
    }{
        {
            name:   "empty date",
            fields: fields{_type: TypeDate},
            want: map[string]interface{}{
                "type": TypeDate,
            },
            wantErr: false,
        },
        {
            name:   "empty date nano",
            fields: fields{_type: TypeDateNano},
            want: map[string]interface{}{
                "type": TypeDateNano,
            },
            wantErr: false,
        },
        {
            name:   "docValues",
            fields: fields{docValues: &boolTrue},
            want: map[string]interface{}{
                "type":       Type(""),
                "doc_values": true,
            },
            wantErr: false,
        },
        {
            name:   "ignoreMalformed",
            fields: fields{ignoreMalformed: &boolTrue},
            want: map[string]interface{}{
                "type":             Type(""),
                "ignore_malformed": true,
            },
            wantErr: false,
        },
        {
            name:   "index",
            fields: fields{index: &boolTrue},
            want: map[string]interface{}{
                "type":  Type(""),
                "index": true,
            },
            wantErr: false,
        },
        {
            name:   "store",
            fields: fields{store: &boolTrue},
            want: map[string]interface{}{
                "type":  Type(""),
                "store": true,
            },
            wantErr: false,
        },
        {
            name:   "boost",
            fields: fields{boost: &boost},
            want: map[string]interface{}{
                "type":  Type(""),
                "boost": boost,
            },
            wantErr: false,
        },
        {
            name:   "locale",
            fields: fields{locale: &testLocalRoot},
            want: map[string]interface{}{
                "type":   Type(""),
                "locale": testLocalRoot,
            },
            wantErr: false,
        },
        {
            name:   "formats",
            fields: fields{formats: []DateFormat{DateFormatANSIC, DateFormatEpochMillis}},
            want: map[string]interface{}{
                "type":   Type(""),
                "format": "EE MMM d HH:mm:ss yyyy||epoch_millis",
            },
            wantErr: false,
        },
        {
            name:   "nullValue",
            fields: fields{nullValue: 2.7},
            want: map[string]interface{}{
                "type":       Type(""),
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
                "type": Type(""),
                "meta": map[string]interface{}{
                    "author":  "diez37",
                    "version": "0.0.1",
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Date{
                docValues:       tt.fields.docValues,
                ignoreMalformed: tt.fields.ignoreMalformed,
                index:           tt.fields.index,
                store:           tt.fields.store,
                boost:           tt.fields.boost,
                locale:          tt.fields.locale,
                formats:         tt.fields.formats,
                nullValue:       tt.fields.nullValue,
                meta:            tt.fields.meta,
                _type:           tt.fields._type,
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

func TestNewDate(t *testing.T) {
    tests := []struct {
        name string
        want *Date
    }{
        {
            name: "new",
            want: &Date{_type: TypeDate},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewDate(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewDate() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewDateNano(t *testing.T) {
    tests := []struct {
        name string
        want *Date
    }{
        {
            name: "new",
            want: &Date{_type: TypeDateNano},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewDateNano(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewDateNano() = %v, want %v", got, tt.want)
            }
        })
    }
}
