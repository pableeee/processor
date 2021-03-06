package query

type Query interface {
	Build() string
}

type Operation struct {
	Type   Type
	Params interface{}
}

func AndOperation(ops []*Operation) *Operation {
	return &Operation{Type: And, Params: ops}
}

func OrOperation(ops []*Operation) *Operation {
	return &Operation{Type: Or, Params: ops}
}

func NotOperation(ops []*Operation) *Operation {
	return &Operation{Type: Not, Params: ops}
}

type FieldExpression struct {
	Field string
	Value interface{}
}

func EqualsOperation(field string, value interface{}) *Operation {
	return &Operation{Params: &FieldExpression{Field: field, Value: value}, Type: Eq}
}
