package operator

import (
	"math"

	"github.com/lucku/jsont/json"
	"github.com/tidwall/gjson"
)

var opAdd = &Operator{
	Identifier:    "add",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          "+",
	Associativity: AssocLeft,
	Precedence:    4,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() + args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opSub = &Operator{
	Identifier:    "sub",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          "-",
	Associativity: AssocLeft,
	Precedence:    4,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() - args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opMultiply = &Operator{
	Identifier:    "multiply",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          "*",
	Associativity: AssocLeft,
	Precedence:    5,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Num * args[1].Num
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opDivide = &Operator{
	Identifier:    "divide",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          "/",
	Associativity: AssocLeft,
	Precedence:    5,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Num / args[1].Num
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opModulo = &Operator{
	Identifier:    "modulo",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          "%",
	Associativity: AssocLeft,
	Precedence:    5,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() % args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opPow = &Operator{
	Identifier:    "power",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          "^",
	Associativity: AssocRight,
	Precedence:    6,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := math.Pow(args[0].Num, args[1].Num)
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opLess = &Operator{
	Identifier:    "less",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          "<",
	Associativity: AssocLeft,
	Precedence:    3,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Num < args[1].Num

		if res {
			return gjson.Result{Type: gjson.True}
		}

		return gjson.Result{Type: gjson.False}
	},
}

var opGreater = &Operator{
	Identifier:    "greater",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          ">",
	Associativity: AssocLeft,
	Precedence:    3,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Num > args[1].Num

		if res {
			return gjson.Result{Type: gjson.True}
		}

		return gjson.Result{Type: gjson.False}
	},
}

var opLessEqual = &Operator{
	Identifier:    "lessEqual",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          "<=",
	Associativity: AssocLeft,
	Precedence:    3,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Num <= args[1].Num

		if res {
			return gjson.Result{Type: gjson.True}
		}

		return gjson.Result{Type: gjson.False}
	},
}

var opGreaterEqual = &Operator{
	Identifier:    "greaterEqual",
	ArgTypes:      []json.Type{json.Number, json.Number},
	Sign:          ">=",
	Associativity: AssocLeft,
	Precedence:    3,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Num >= args[1].Num

		if res {
			return gjson.Result{Type: gjson.True}
		}

		return gjson.Result{Type: gjson.False}
	},
}
