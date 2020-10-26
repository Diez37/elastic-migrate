package fields

type Relation struct {
    question string
    answers  []string
}

func NewRelation(question string, answers ...string) *Relation {
    return &Relation{question: question, answers: answers}
}

type Join struct {
    eagerGlobalOrdinals *bool
    relations           []*Relation
}

func NewJoin(relations ...*Relation) *Join {
    return (&Join{}).Questions(relations...)
}

func (field *Join) SetEagerGlobalOrdinals(eagerGlobalOrdinals bool) *Join {
    field.eagerGlobalOrdinals = &eagerGlobalOrdinals

    return field
}

func (field *Join) Questions(relations ...*Relation) *Join {
    field.relations = append(field.relations, relations...)

    return field
}

func (field *Join) GetType() Type {
    return TypeJoin
}

func (field *Join) Source() (interface{}, error) {
    // {
    //  "type": "join",
    //  "relations": {
    //    "question": ["answers", "comment"]
    //  },
    //  "eager_global_ordinals": false
    // }

    source := map[string]interface{}{}

    source["type"] = field.GetType()

    if field.eagerGlobalOrdinals != nil {
        source["eager_global_ordinals"] = *field.eagerGlobalOrdinals
    }

    if len(field.relations) > 0 {
        relations := map[string][]string{}

        for _, relation := range field.relations {
            relations[relation.question] = relation.answers
        }

        source["relations"] = relations
    }

    return source, nil
}
