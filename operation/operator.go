package operation

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type any interface{}

// ApplierFunc represents the function that is performed by an operator, taking a variable number of inputs
// and returning any type of result that is directly put into the JSON output of the operation
type ApplierFunc func(...gjson.Result) any

// Operator is a function with the given name *Identifier*, that is performed on input values whenever *OperatorSign* is
// between them (infix). The input arguments of that function have to be of the types specified in *ArgTypes*. The actual
// function is included in *Apply*.
//
// The types of input arguments is checked before calling the function thus it can be assumed without further type checking
// that input arguments to the ApplierFunc are of the necessary types.
type Operator struct {
	Identifier   string
	OperatorSign string
	ArgTypes     []gjson.Type
	Apply        ApplierFunc
}

// OperandTypeError is an error that indicates a wrong type of input variable given into an operation
//
// This error is caused by the user of the library and shall therefore be passed up to the CLI
type OperandTypeError struct {
	operand      gjson.Result
	expectedType gjson.Type
}

// Operators is a slice of operators registered to be used by the JSON transforming logic
//
// The operations inside the slice have to be ordered by precedence of the operator, with the most precedent
// operator (select) as the slice element at index 0
var Operators []*Operator

func (e OperandTypeError) Error() string {
	return fmt.Sprintf("argument %v is not of type %v", e.operand.Value(), e.expectedType)
}

func (o Operator) String() string {
	return o.Identifier
}

// NewOperandTypeError instantiates a new OperandTypeError with the given operator reference and expected type
// as a gjson.Type
func NewOperandTypeError(operand gjson.Result, expectedType gjson.Type) *OperandTypeError {
	return &OperandTypeError{operand: operand, expectedType: expectedType}
}

func registerOperationInOrder(op *Operator) {

	if op != nil {
		Operators = append(Operators, op)
	}

}

func init() {
	registerOperationInOrder(opSelect)
	registerOperationInOrder(opConcat)
	registerOperationInOrder(opAdd)
}
