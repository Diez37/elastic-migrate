package fields

import (
    "reflect"
    "testing"
)

func TestNewRange(t *testing.T) {
    tests := []struct {
        name        string
        want        *Range
        constructor func() *Range
    }{
        {
            name:        "integer",
            want:        &Range{_type: TypeRangeInteger},
            constructor: NewRangeInteger,
        },
        {
            name:        "float",
            want:        &Range{_type: TypeRangeFloat},
            constructor: NewRangeFloat,
        },
        {
            name:        "long",
            want:        &Range{_type: TypeRangeLong},
            constructor: NewRangeLong,
        },
        {
            name:        "double",
            want:        &Range{_type: TypeRangeDouble},
            constructor: NewRangeDouble,
        },
        {
            name:        "ip",
            want:        &Range{_type: TypeRangeIp},
            constructor: NewRangeIp,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.constructor(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewRangeDate() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewRangeDate(t *testing.T) {
    tests := []struct {
        name string
        want *RangeDate
    }{
        {
            name: "date",
            want: &RangeDate{Range: Range{_type: TypeRangeDate}},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewRangeDate(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewRangeDate() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRangeDate_Formats(t *testing.T) {
    type fields struct {
        Range   Range
        formats []DateFormat
    }
    type args struct {
        formats []DateFormat
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *RangeDate
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{formats: []DateFormat{DateFormatANSIC, DateFormatUnixDate}},
            want:   &RangeDate{formats: []DateFormat{DateFormatANSIC, DateFormatUnixDate}},
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
            want: &RangeDate{formats: []DateFormat{
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
            field := &RangeDate{
                Range:   tt.fields.Range,
                formats: tt.fields.formats,
            }
            if got := field.Formats(tt.args.formats...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Formats() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRangeDate_GetType(t *testing.T) {
    type fields struct {
        Range   Range
        formats []DateFormat
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "range date type",
            fields: fields{Range: Range{_type: TypeRangeDate}},
            want:   TypeRangeDate,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &RangeDate{
                Range:   tt.fields.Range,
                formats: tt.fields.formats,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRangeDate_Source(t *testing.T) {
    var initBoost float64 = 9

    type fields struct {
        Range   Range
        formats []DateFormat
    }
    tests := []struct {
        name    string
        fields  fields
        want    interface{}
        wantErr bool
    }{
        {
            name:   "empty",
            fields: fields{Range: Range{_type: TypeRangeDate}},
            want: map[string]interface{}{
                "type": TypeRangeDate,
            },
            wantErr: false,
        },
        {
            name: "coerce",
            fields: fields{Range: Range{
                coerce: &testTrue,
                _type:  TypeRangeDate,
            }},
            want: map[string]interface{}{
                "type":   TypeRangeDate,
                "coerce": true,
            },
            wantErr: false,
        },
        {
            name: "boost",
            fields: fields{Range: Range{
                boost: &initBoost,
                _type: TypeRangeDate,
            }},
            want: map[string]interface{}{
                "type":  TypeRangeDate,
                "boost": initBoost,
            },
            wantErr: false,
        },
        {
            name: "index",
            fields: fields{Range: Range{
                index: &testTrue,
                _type: TypeRangeDate,
            }},
            want: map[string]interface{}{
                "type":  TypeRangeDate,
                "index": true,
            },
            wantErr: false,
        },
        {
            name: "store",
            fields: fields{Range: Range{
                store: &testTrue,
                _type: TypeRangeDate,
            }},
            want: map[string]interface{}{
                "type":  TypeRangeDate,
                "store": true,
            },
            wantErr: false,
        },
        {
            name: "format",
            fields: fields{
                formats: []DateFormat{
                    DateFormatEpochSecond,
                    DateFormatDateOptionalTime,
                },
                Range: Range{_type: TypeRangeDate},
            },
            want: map[string]interface{}{
                "type":   TypeRangeDate,
                "format": "epoch_second||date_optional_time",
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &RangeDate{
                Range:   tt.fields.Range,
                formats: tt.fields.formats,
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

func TestRange_GetType(t *testing.T) {
    type fields struct {
        coerce *bool
        store  *bool
        index  *bool
        boost  *float64
        _type  Type
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "integer",
            fields: fields{_type: TypeRangeInteger},
            want:   TypeRangeInteger,
        },
        {
            name:   "float",
            fields: fields{_type: TypeRangeFloat},
            want:   TypeRangeFloat,
        },
        {
            name:   "long",
            fields: fields{_type: TypeRangeLong},
            want:   TypeRangeLong,
        },
        {
            name:   "double",
            fields: fields{_type: TypeRangeDouble},
            want:   TypeRangeDouble,
        },
        {
            name:   "ip",
            fields: fields{_type: TypeRangeIp},
            want:   TypeRangeIp,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Range{
                coerce: tt.fields.coerce,
                store:  tt.fields.store,
                index:  tt.fields.index,
                boost:  tt.fields.boost,
                _type:  tt.fields._type,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRange_SetBoost(t *testing.T) {
    var initBoost float64 = 9
    var setBoost = 5.4

    type fields struct {
        coerce *bool
        store  *bool
        index  *bool
        boost  *float64
        _type  Type
    }
    type args struct {
        boost float64
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Range
    }{
        {
            name:   "set boost",
            fields: fields{},
            args:   args{boost: setBoost},
            want:   &Range{boost: &setBoost},
        },
        {
            name:   "change boost",
            fields: fields{boost: &initBoost},
            args:   args{boost: setBoost},
            want:   &Range{boost: &setBoost},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Range{
                coerce: tt.fields.coerce,
                store:  tt.fields.store,
                index:  tt.fields.index,
                boost:  tt.fields.boost,
                _type:  tt.fields._type,
            }
            if got := field.SetBoost(tt.args.boost); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetBoost() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRange_SetCoerce(t *testing.T) {
    type fields struct {
        coerce *bool
        store  *bool
        index  *bool
        boost  *float64
        _type  Type
    }
    type args struct {
        coerce bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Range
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{coerce: true},
            want:   &Range{coerce: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{coerce: false},
            want:   &Range{coerce: &testFalse},
        },
        {
            name:   "change",
            fields: fields{coerce: &testTrue},
            args:   args{coerce: false},
            want:   &Range{coerce: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Range{
                coerce: tt.fields.coerce,
                store:  tt.fields.store,
                index:  tt.fields.index,
                boost:  tt.fields.boost,
                _type:  tt.fields._type,
            }
            if got := field.SetCoerce(tt.args.coerce); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetCoerce() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRange_SetIndex(t *testing.T) {
    type fields struct {
        coerce *bool
        store  *bool
        index  *bool
        boost  *float64
        _type  Type
    }
    type args struct {
        index bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Range
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{index: true},
            want:   &Range{index: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{index: false},
            want:   &Range{index: &testFalse},
        },
        {
            name:   "change",
            fields: fields{index: &testTrue},
            args:   args{index: false},
            want:   &Range{index: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Range{
                coerce: tt.fields.coerce,
                store:  tt.fields.store,
                index:  tt.fields.index,
                boost:  tt.fields.boost,
                _type:  tt.fields._type,
            }
            if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIndex() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRange_SetStore(t *testing.T) {
    type fields struct {
        coerce *bool
        store  *bool
        index  *bool
        boost  *float64
        _type  Type
    }
    type args struct {
        store bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Range
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{store: true},
            want:   &Range{store: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{store: false},
            want:   &Range{store: &testFalse},
        },
        {
            name:   "change",
            fields: fields{store: &testTrue},
            args:   args{store: false},
            want:   &Range{store: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Range{
                coerce: tt.fields.coerce,
                store:  tt.fields.store,
                index:  tt.fields.index,
                boost:  tt.fields.boost,
                _type:  tt.fields._type,
            }
            if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetStore() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRange_Source(t *testing.T) {
    var initBoost float64 = 9

    type fields struct {
        coerce *bool
        store  *bool
        index  *bool
        boost  *float64
        _type  Type
    }
    tests := []struct {
        name    string
        fields  fields
        want    interface{}
        wantErr bool
    }{
        {
            name:   "integer",
            fields: fields{_type: TypeRangeInteger},
            want: map[string]interface{}{
                "type": TypeRangeInteger,
            },
            wantErr: false,
        },
        {
            name:   "float",
            fields: fields{_type: TypeRangeFloat},
            want: map[string]interface{}{
                "type": TypeRangeFloat,
            },
            wantErr: false,
        },
        {
            name:   "long",
            fields: fields{_type: TypeRangeLong},
            want: map[string]interface{}{
                "type": TypeRangeLong,
            },
            wantErr: false,
        },
        {
            name:   "double",
            fields: fields{_type: TypeRangeDouble},
            want: map[string]interface{}{
                "type": TypeRangeDouble,
            },
            wantErr: false,
        },
        {
            name:   "ip",
            fields: fields{_type: TypeRangeIp},
            want: map[string]interface{}{
                "type": TypeRangeIp,
            },
            wantErr: false,
        },
        {
            name:   "coerce",
            fields: fields{coerce: &testTrue},
            want: map[string]interface{}{
                "type":   Type(""),
                "coerce": true,
            },
            wantErr: false,
        },
        {
            name:   "boost",
            fields: fields{boost: &initBoost},
            want: map[string]interface{}{
                "type":  Type(""),
                "boost": initBoost,
            },
            wantErr: false,
        },
        {
            name:   "index",
            fields: fields{index: &testTrue},
            want: map[string]interface{}{
                "type":  Type(""),
                "index": true,
            },
            wantErr: false,
        },
        {
            name:   "store",
            fields: fields{store: &testTrue},
            want: map[string]interface{}{
                "type":  Type(""),
                "store": true,
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Range{
                coerce: tt.fields.coerce,
                store:  tt.fields.store,
                index:  tt.fields.index,
                boost:  tt.fields.boost,
                _type:  tt.fields._type,
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
