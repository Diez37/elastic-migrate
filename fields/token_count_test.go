package fields

import (
    "github.com/diez37/elastic-migrate/settings"
    "reflect"
    "testing"
)

func TestNewTokenCount(t *testing.T) {
    tests := []struct {
        name string
        want *TokenCount
    }{
        {
            name: "new",
            want: &TokenCount{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewTokenCount(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewTokenCount() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTokenCount_GetType(t *testing.T) {
    type fields struct {
        enablePositionIncrements *bool
        docValues                *bool
        index                    *bool
        store                    *bool
        boost                    *float64
        analyzer                 *settings.AnalyzerName
        nullValue                interface{}
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "type token count",
            fields: fields{},
            want:   TypeTokenCount,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &TokenCount{
                enablePositionIncrements: tt.fields.enablePositionIncrements,
                docValues:                tt.fields.docValues,
                index:                    tt.fields.index,
                store:                    tt.fields.store,
                boost:                    tt.fields.boost,
                analyzer:                 tt.fields.analyzer,
                nullValue:                tt.fields.nullValue,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTokenCount_SetAnalyzer(t *testing.T) {
    initAnalyzerName := settings.AnalyzerStandard
    changeAnalyzerName := settings.AnalyzerSimple

    type fields struct {
        enablePositionIncrements *bool
        docValues                *bool
        index                    *bool
        store                    *bool
        boost                    *float64
        analyzer                 *settings.AnalyzerName
        nullValue                interface{}
    }
    type args struct {
        analyzer settings.AnalyzerName
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *TokenCount
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{analyzer: initAnalyzerName},
            want:   &TokenCount{analyzer: &initAnalyzerName},
        },
        {
            name:   "change",
            fields: fields{analyzer: &changeAnalyzerName},
            args:   args{analyzer: initAnalyzerName},
            want:   &TokenCount{analyzer: &initAnalyzerName},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &TokenCount{
                enablePositionIncrements: tt.fields.enablePositionIncrements,
                docValues:                tt.fields.docValues,
                index:                    tt.fields.index,
                store:                    tt.fields.store,
                boost:                    tt.fields.boost,
                analyzer:                 tt.fields.analyzer,
                nullValue:                tt.fields.nullValue,
            }
            if got := field.SetAnalyzer(tt.args.analyzer); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetAnalyzer() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTokenCount_SetBoost(t *testing.T) {
    initBoost := 9.0
    setBoost := 5.4

    type fields struct {
        enablePositionIncrements *bool
        docValues                *bool
        index                    *bool
        store                    *bool
        boost                    *float64
        analyzer                 *settings.AnalyzerName
        nullValue                interface{}
    }
    type args struct {
        boost float64
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *TokenCount
    }{
        {
            name:   "set boost",
            fields: fields{},
            args:   args{boost: setBoost},
            want:   &TokenCount{boost: &setBoost},
        },
        {
            name:   "change boost",
            fields: fields{boost: &initBoost},
            args:   args{boost: setBoost},
            want:   &TokenCount{boost: &setBoost},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &TokenCount{
                enablePositionIncrements: tt.fields.enablePositionIncrements,
                docValues:                tt.fields.docValues,
                index:                    tt.fields.index,
                store:                    tt.fields.store,
                boost:                    tt.fields.boost,
                analyzer:                 tt.fields.analyzer,
                nullValue:                tt.fields.nullValue,
            }
            if got := field.SetBoost(tt.args.boost); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetBoost() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTokenCount_SetDocValues(t *testing.T) {
    type fields struct {
        enablePositionIncrements *bool
        docValues                *bool
        index                    *bool
        store                    *bool
        boost                    *float64
        analyzer                 *settings.AnalyzerName
        nullValue                interface{}
    }
    type args struct {
        docValues bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *TokenCount
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{docValues: true},
            want:   &TokenCount{docValues: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{docValues: false},
            want:   &TokenCount{docValues: &testFalse},
        },
        {
            name:   "change",
            fields: fields{docValues: &testTrue},
            args:   args{docValues: false},
            want:   &TokenCount{docValues: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &TokenCount{
                enablePositionIncrements: tt.fields.enablePositionIncrements,
                docValues:                tt.fields.docValues,
                index:                    tt.fields.index,
                store:                    tt.fields.store,
                boost:                    tt.fields.boost,
                analyzer:                 tt.fields.analyzer,
                nullValue:                tt.fields.nullValue,
            }
            if got := field.SetDocValues(tt.args.docValues); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetDocValues() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTokenCount_SetEnablePositionIncrements(t *testing.T) {
    type fields struct {
        enablePositionIncrements *bool
        docValues                *bool
        index                    *bool
        store                    *bool
        boost                    *float64
        analyzer                 *settings.AnalyzerName
        nullValue                interface{}
    }
    type args struct {
        enablePositionIncrements bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *TokenCount
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{enablePositionIncrements: true},
            want:   &TokenCount{enablePositionIncrements: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{enablePositionIncrements: false},
            want:   &TokenCount{enablePositionIncrements: &testFalse},
        },
        {
            name:   "change",
            fields: fields{enablePositionIncrements: &testTrue},
            args:   args{enablePositionIncrements: false},
            want:   &TokenCount{enablePositionIncrements: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &TokenCount{
                enablePositionIncrements: tt.fields.enablePositionIncrements,
                docValues:                tt.fields.docValues,
                index:                    tt.fields.index,
                store:                    tt.fields.store,
                boost:                    tt.fields.boost,
                analyzer:                 tt.fields.analyzer,
                nullValue:                tt.fields.nullValue,
            }
            if got := field.SetEnablePositionIncrements(tt.args.enablePositionIncrements); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetEnablePositionIncrements() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTokenCount_SetIndex(t *testing.T) {
    type fields struct {
        enablePositionIncrements *bool
        docValues                *bool
        index                    *bool
        store                    *bool
        boost                    *float64
        analyzer                 *settings.AnalyzerName
        nullValue                interface{}
    }
    type args struct {
        index bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *TokenCount
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{index: true},
            want:   &TokenCount{index: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{index: false},
            want:   &TokenCount{index: &testFalse},
        },
        {
            name:   "change",
            fields: fields{index: &testTrue},
            args:   args{index: false},
            want:   &TokenCount{index: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &TokenCount{
                enablePositionIncrements: tt.fields.enablePositionIncrements,
                docValues:                tt.fields.docValues,
                index:                    tt.fields.index,
                store:                    tt.fields.store,
                boost:                    tt.fields.boost,
                analyzer:                 tt.fields.analyzer,
                nullValue:                tt.fields.nullValue,
            }
            if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIndex() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTokenCount_SetNullValue(t *testing.T) {
    type fields struct {
        enablePositionIncrements *bool
        docValues                *bool
        index                    *bool
        store                    *bool
        boost                    *float64
        analyzer                 *settings.AnalyzerName
        nullValue                interface{}
    }
    type args struct {
        nullValue interface{}
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *TokenCount
    }{
        {
            name:   "empty",
            fields: fields{},
            args:   args{},
            want:   &TokenCount{nullValue: nil},
        },
        {
            name:   "int",
            fields: fields{},
            args:   args{nullValue: 1},
            want:   &TokenCount{nullValue: 1},
        },
        {
            name:   "float",
            fields: fields{},
            args:   args{nullValue: 4.5},
            want:   &TokenCount{nullValue: 4.5},
        },
        {
            name:   "string",
            fields: fields{},
            args:   args{nullValue: "test"},
            want:   &TokenCount{nullValue: "test"},
        },
        {
            name:   "bool",
            fields: fields{},
            args:   args{nullValue: true},
            want:   &TokenCount{nullValue: true},
        },
        {
            name:   "change",
            fields: fields{nullValue: "test"},
            args:   args{nullValue: true},
            want:   &TokenCount{nullValue: true},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &TokenCount{
                enablePositionIncrements: tt.fields.enablePositionIncrements,
                docValues:                tt.fields.docValues,
                index:                    tt.fields.index,
                store:                    tt.fields.store,
                boost:                    tt.fields.boost,
                analyzer:                 tt.fields.analyzer,
                nullValue:                tt.fields.nullValue,
            }
            if got := field.SetNullValue(tt.args.nullValue); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetNullValue() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTokenCount_SetStore(t *testing.T) {
    type fields struct {
        enablePositionIncrements *bool
        docValues                *bool
        index                    *bool
        store                    *bool
        boost                    *float64
        analyzer                 *settings.AnalyzerName
        nullValue                interface{}
    }
    type args struct {
        store bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *TokenCount
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{store: true},
            want:   &TokenCount{store: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{store: false},
            want:   &TokenCount{store: &testFalse},
        },
        {
            name:   "change",
            fields: fields{store: &testTrue},
            args:   args{store: false},
            want:   &TokenCount{store: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &TokenCount{
                enablePositionIncrements: tt.fields.enablePositionIncrements,
                docValues:                tt.fields.docValues,
                index:                    tt.fields.index,
                store:                    tt.fields.store,
                boost:                    tt.fields.boost,
                analyzer:                 tt.fields.analyzer,
                nullValue:                tt.fields.nullValue,
            }
            if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetStore() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTokenCount_Source(t *testing.T) {
    initAnalyzerName := settings.AnalyzerStandard
    initBoost := 9.0

    type fields struct {
        enablePositionIncrements *bool
        docValues                *bool
        index                    *bool
        store                    *bool
        boost                    *float64
        analyzer                 *settings.AnalyzerName
        nullValue                interface{}
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
                "type": TypeTokenCount,
            },
            wantErr: false,
        },
        {
            name:    "analyzer",
            fields:  fields{analyzer: &initAnalyzerName},
            want:    map[string]interface{}{
                "type": TypeTokenCount,
                "analyzer": initAnalyzerName,
            },
            wantErr: false,
        },
        {
            name:    "enable position increments",
            fields:  fields{enablePositionIncrements: &testTrue},
            want:    map[string]interface{}{
                "type": TypeTokenCount,
                "enable_position_increments": true,
            },
            wantErr: false,
        },
        {
            name:    "boost",
            fields:  fields{boost: &initBoost},
            want:    map[string]interface{}{
                "type": TypeTokenCount,
                "boost": initBoost,
            },
            wantErr: false,
        },
        {
            name:    "doc_values",
            fields:  fields{docValues: &testTrue},
            want:    map[string]interface{}{
                "type": TypeTokenCount,
                "doc_values": true,
            },
            wantErr: false,
        },
        {
            name:    "index",
            fields:  fields{index: &testTrue},
            want:    map[string]interface{}{
                "type": TypeTokenCount,
                "index": true,
            },
            wantErr: false,
        },
        {
            name:    "nullValue",
            fields:  fields{nullValue: testTrue},
            want:    map[string]interface{}{
                "type": TypeTokenCount,
                "null_value": true,
            },
            wantErr: false,
        },
        {
            name:    "store",
            fields:  fields{store: &testTrue},
            want:    map[string]interface{}{
                "type": TypeTokenCount,
                "store": true,
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &TokenCount{
                enablePositionIncrements: tt.fields.enablePositionIncrements,
                docValues:                tt.fields.docValues,
                index:                    tt.fields.index,
                store:                    tt.fields.store,
                boost:                    tt.fields.boost,
                analyzer:                 tt.fields.analyzer,
                nullValue:                tt.fields.nullValue,
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
