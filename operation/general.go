package operation

import (
	"github.com/lucku/jsont/json"
	"github.com/tidwall/gjson"
)

var opCompare = &Operator{
	Identifier: "compare",
	ArgTypes:   []json.Type{json.Any, json.Any},
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

var opIfNull = &Operator{
	Identifier: "ifnull",
	ArgTypes:   []json.Type{json.Any, json.Any},
	Sign:       "?",
	Precedence: 3,
	Apply: func(args ...gjson.Result) gjson.Result {

		if json.CheckType(args[0], json.Null) {
			return args[1]
		}

		return args[0]
	},
}
