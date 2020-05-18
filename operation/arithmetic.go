package operation

import (
	"github.com/lucku/jsont/json"
	"github.com/tidwall/gjson"
)

var opAdd = &Operator{
	Identifier: "add",
	ArgTypes:   []json.Type{json.Number, json.Number},
	Sign:       "+",
	Precedence: 1,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() + args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opSub = &Operator{
	Identifier: "sub",
	ArgTypes:   []json.Type{json.Number, json.Number},
	Sign:       "-",
	Precedence: 1,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() - args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opMultiply = &Operator{
	Identifier: "multiply",
	ArgTypes:   []json.Type{json.Number, json.Number},
	Sign:       "*",
	Precedence: 2,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() * args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}

var opDivide = &Operator{
	Identifier: "divide",
	ArgTypes:   []json.Type{json.Number, json.Number},
	Sign:       "/",
	Precedence: 2,
	Apply: func(args ...gjson.Result) gjson.Result {
		res := args[0].Int() / args[1].Int()
		return gjson.Result{Type: gjson.Number, Num: float64(res)}
	},
}
