package operation

import (
	"strings"

	"github.com/tidwall/gjson"
)

var opSelect = &Operator{
	Identifier:   "select",
	ArgTypes:     []gjson.Type{gjson.JSON, gjson.String},
	OperatorSign: ".",
	Apply: func(args ...gjson.Result) any {
		return args[0].Get(args[1].String()).Value()
	},
}

var opConcat = &Operator{
	Identifier:   "concat",
	ArgTypes:     []gjson.Type{gjson.String, gjson.String},
	OperatorSign: "+",
	Apply: func(args ...gjson.Result) any {

		sb := strings.Builder{}
		sb.WriteString(args[0].String())
		sb.WriteString(args[1].String())

		return sb.String()
	},
}

var opAdd = &Operator{
	Identifier:   "add",
	ArgTypes:     []gjson.Type{gjson.Number, gjson.Number},
	OperatorSign: "+",
	Apply: func(args ...gjson.Result) any {
		return args[0].Int() + args[1].Int()
	},
}
