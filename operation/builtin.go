package operation

import (
	"strings"

	"github.com/tidwall/gjson"
)

var opConcat = &Operator{
	Identifier: "concat",
	ArgTypes:   []gjson.Type{gjson.String, gjson.String},
	Sign:       ":",
	Precedence: 1,
	Apply: func(args ...gjson.Result) gjson.Result {
		sb := strings.Builder{}
		sb.WriteString(args[0].String())
		sb.WriteString(args[1].String())
		return gjson.Result{Type: gjson.String, Str: sb.String()}
	},
}

var opAdd = &Operator{
	Identifier: "add",
	ArgTypes:   []gjson.Type{gjson.Number, gjson.Number},
	Sign:       "+",
	Precedence: 1,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() + args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opSub = &Operator{
	Identifier: "sub",
	ArgTypes:   []gjson.Type{gjson.Number, gjson.Number},
	Sign:       "-",
	Precedence: 1,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() - args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opMultiply = &Operator{
	Identifier: "multiply",
	ArgTypes:   []gjson.Type{gjson.Number, gjson.Number},
	Sign:       "*",
	Precedence: 2,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() * args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opDivide = &Operator{
	Identifier: "divide",
	ArgTypes:   []gjson.Type{gjson.Number, gjson.Number},
	Sign:       "/",
	Precedence: 2,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() / args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opCompare = &Operator{
	Identifier: "compare",
	ArgTypes:   []gjson.Type{Any, Any},
	Sign:       "=",
	Precedence: 1,
	Apply: func(args ...gjson.Result) gjson.Result {

		res := args[0] == args[1]

		if res {
			return gjson.Result{Type: gjson.True}
		}

		return gjson.Result{Type: gjson.False}
	},
}

var opAnd = &Operator{
	Identifier: "and",
	ArgTypes:   []gjson.Type{Bool, Bool},
	Sign:       "&",
	Precedence: 1,
	Apply: func(args ...gjson.Result) gjson.Result {

		res := args[0].Bool() && args[1].Bool()

		if res {
			return gjson.Result{Type: gjson.True}
		}

		return gjson.Result{Type: gjson.False}
	},
}

var opOr = &Operator{
	Identifier: "or",
	ArgTypes:   []gjson.Type{Bool, Bool},
	Sign:       "|",
	Precedence: 1,
	Apply: func(args ...gjson.Result) gjson.Result {

		res := args[0].Bool() || args[1].Bool()

		if res {
			return gjson.Result{Type: gjson.True}
		}

		return gjson.Result{Type: gjson.False}
	},
}
