package jsont

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestEvalExpression(t *testing.T) {

	expr := "$select.example:($test.something:literal)"

	input := gjson.Parse(`{"select":{"example":"test"},"test":{"something":"value"}}`)

	e := newExpressionEvaluator(input)

	res, err := e.EvaluateExpression(expr)

	t.Log(err)
	t.Log(res)
}
