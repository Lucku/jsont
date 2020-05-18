package operation

import (
	"fmt"

	"github.com/tidwall/gjson"
)

// Any represents any JSON type and can be used by operator definitions to declare that any kind of JSON
// type is accepted as argument
var Any gjson.Type = 6

// Bool is the group of JSON types True and False
var Bool gjson.Type = 7

// ApplierFunc represents the function that is performed by an operator, taking a variable number of inputs
// as gjson.Result and returning again a gjson.Result
type ApplierFunc func(...gjson.Result) gjson.Result

// Operator is a function with the given name *Identifier*, that is performed on input values whenever *OperatorSign* is
// between them (infix). The input arguments of that function have to be of the types specified in *ArgTypes*. The precedence
// of the operator is defined in the corresponding attribute, with 1 being the lowest precedence (usually the select operator).
// The actual function is included in *Apply*.
//
// The types of input arguments is checked before calling the function thus it can be assumed without further type checking
// that input arguments to the ApplierFunc are of the necessary types.
type Operator struct {
	Identifier string
	Sign       string
	Precedence int
	ArgTypes   []gjson.Type
	Apply      ApplierFunc
}

// OperandTypeError is an error that indicates a wrong type of input variable given into an operation
//
// This error is caused by the user of the library and shall therefore be passed up to the CLI
type OperandTypeError struct {
	operand      gjson.Result
	expectedType gjson.Type
}

// Operators is a map of operators registered to be used by the JSON transforming logic
//
// The operator sign acts as the map key for an operator, making it easy to look up an operator by
// its sign
var Operators = make(map[string]*Operator)

func registerOperator(op *Operator) {

	if op != nil {
		Operators[op.Sign] = op
	}
}

// IsOperator checks if a given literal is associated with a registered operator and returns true is that is
// the case
func IsOperator(literal string) bool {

	if _, ok := Operators[literal]; ok {
		return true
	}

	return false
}

// GetOperator returns the operator for a given literal or nil if there is no registered operator with the given sign
func GetOperator(literal string) *Operator {

	if o, ok := Operators[literal]; ok {
		return o
	}

	return nil
}

func init() {
	registerOperator(opConcat)
	registerOperator(opAdd)

	fmt.Println(Operators)
}
