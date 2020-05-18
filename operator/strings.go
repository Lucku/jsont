package operator

import (
	"strings"

	"github.com/lucku/jsont/json"
	"github.com/tidwall/gjson"
)

var opConcat = &Operator{
	Identifier: "concat",
	ArgTypes:   []json.Type{json.String, json.String},
	Sign:       ":",
	Precedence: 1,
	Apply: func(args ...gjson.Result) gjson.Result {
		sb := strings.Builder{}
		sb.WriteString(args[0].String())
		sb.WriteString(args[1].String())
		return gjson.Result{Type: gjson.String, Str: sb.String()}
	},
}
