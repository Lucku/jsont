package operator

type DotOperator struct {
	BaseApplier
}

func (o *DotOperator) Identifier() string {
	return "select"
}

func (o *DotOperator) OperatorSign() string {
	return "."
}

func (o *DotOperator) ReturnTypes() []jsonType {
	return []jsonType{typeObject}
}

func Apply(operands ...interface{}) interface{} {
	return nil
}
