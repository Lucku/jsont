package operator

type jsonType uint8

const (
	typeString jsonType = iota
	typeInt
	typeFloat
	typeBool
	typeObject
	typeArray
)

type Applier interface {
	ApplyUnary(operand interface{}) (interface{}, error)
	ApplyBinary(firstOperand, secondOperand interface{}) (interface{}, error)
}

type Operator interface {
	Identifier() string
	OperatorSign() string
	OperandTypes() []jsonType
	ReturnTypes() []jsonType
	Applier
}

type BaseApplier struct {
}

func (o *BaseApplier) ApplyUnary(operand interface{}) (interface{}, error) {
	return nil, nil
}

func (o *BaseApplier) ApplyBinary(firstOperand, secondOperand interface{}) (interface{}, error) {
	return nil, nil
}
