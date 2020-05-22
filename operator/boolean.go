package operator

import (
	"github.com/lucku/jsont/json"
	"github.com/tidwall/gjson"
)

var opAnd = &Operator{
	Identifier: "and",
	ArgTypes:   []json.Type{json.Bool, json.Bool},
	Sign:       "&",
	Precedence: 5,
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
	ArgTypes:   []json.Type{json.Bool, json.Bool},
	Sign:       "|",
	Precedence: 4,
	Apply: func(args ...gjson.Result) gjson.Result {

		res := args[0].Bool() || args[1].Bool()

		if res {
			return gjson.Result{Type: gjson.True}
		}

		return gjson.Result{Type: gjson.False}
	},
}
