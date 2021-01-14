package fields

import (
    "reflect"
    "testing"
)

func TestJoin_GetType(t *testing.T) {
    type fields struct {
        eagerGlobalOrdinals *bool
        relations           []*Relation
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "type join",
            fields: fields{},
            want:   TypeJoin,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Join{
                eagerGlobalOrdinals: tt.fields.eagerGlobalOrdinals,
                relations:           tt.fields.relations,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestJoin_Questions(t *testing.T) {
    type fields struct {
        eagerGlobalOrdinals *bool
        relations           []*Relation
    }
    type args struct {
        relations []*Relation
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Join
    }{
        {
            name:   "set",
            fields: fields{},
            args: args{relations: []*Relation{
                {
                    question: "test",
                    answers:  []string{"test1", "test2"},
                },
            }},
            want: &Join{relations: []*Relation{
                {
                    question: "test",
                    answers:  []string{"test1", "test2"},
                },
            }},
        },
        {
            name: "add",
            fields: fields{relations: []*Relation{
                {
                    question: "test2",
                    answers:  []string{"test3", "test4"},
                },
            }},
            args: args{relations: []*Relation{
                {
                    question: "test",
                    answers:  []string{"test1", "test2"},
                },
            }},
            want: &Join{relations: []*Relation{
                {
                    question: "test2",
                    answers:  []string{"test3", "test4"},
                },
                {
                    question: "test",
                    answers:  []string{"test1", "test2"},
                },
            }},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Join{
                eagerGlobalOrdinals: tt.fields.eagerGlobalOrdinals,
                relations:           tt.fields.relations,
            }
            if got := field.Questions(tt.args.relations...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Questions() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestJoin_SetEagerGlobalOrdinals(t *testing.T) {
    boolTrue := true
    boolFalse := false

    type fields struct {
        eagerGlobalOrdinals *bool
        relations           []*Relation
    }
    type args struct {
        eagerGlobalOrdinals bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Join
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{eagerGlobalOrdinals: true},
            want:   &Join{eagerGlobalOrdinals: &boolTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{eagerGlobalOrdinals: false},
            want:   &Join{eagerGlobalOrdinals: &boolFalse},
        },
        {
            name:   "change",
            fields: fields{eagerGlobalOrdinals: &boolTrue},
            args:   args{eagerGlobalOrdinals: false},
            want:   &Join{eagerGlobalOrdinals: &boolFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Join{
                eagerGlobalOrdinals: tt.fields.eagerGlobalOrdinals,
                relations:           tt.fields.relations,
            }
            if got := field.SetEagerGlobalOrdinals(tt.args.eagerGlobalOrdinals); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetEagerGlobalOrdinals() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestJoin_Source(t *testing.T) {
    boolTrue := true

    type fields struct {
        eagerGlobalOrdinals *bool
        relations           []*Relation
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
                "type": TypeJoin,
            },
            wantErr: false,
        },
        {
            name:   "eagerGlobalOrdinals",
            fields: fields{eagerGlobalOrdinals: &boolTrue},
            want: map[string]interface{}{
                "type":                  TypeJoin,
                "eager_global_ordinals": boolTrue,
            },
            wantErr: false,
        },
        {
            name: "relations",
            fields: fields{relations: []*Relation{
                {
                    question: "test",
                    answers:  []string{"test1", "test2"},
                },
            }},
            want: map[string]interface{}{
                "type": TypeJoin,
                "relations": map[string][]string{
                    "test": {"test1", "test2"},
                },
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Join{
                eagerGlobalOrdinals: tt.fields.eagerGlobalOrdinals,
                relations:           tt.fields.relations,
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

func TestNewJoin(t *testing.T) {
    type args struct {
        relations []*Relation
    }
    tests := []struct {
        name string
        args args
        want *Join
    }{
        {
            name: "new",
            args: args{},
            want: &Join{},
        },
        {
            name: "new relations",
            args: args{relations: []*Relation{
                {
                    question: "test",
                    answers:  []string{"test1", "test2"},
                },
            }},
            want: &Join{relations: []*Relation{
                {
                    question: "test",
                    answers:  []string{"test1", "test2"},
                },
            }},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewJoin(tt.args.relations...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewJoin() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewRelation(t *testing.T) {
    type args struct {
        question string
        answers  []string
    }
    tests := []struct {
        name string
        args args
        want *Relation
    }{
        {
            name: "new",
            args: args{
                question: "test",
                answers:  []string{"test1", "test2"},
            },
            want: &Relation{
                question: "test",
                answers:  []string{"test1", "test2"},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewRelation(tt.args.question, tt.args.answers...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewRelation() = %v, want %v", got, tt.want)
            }
        })
    }
}
