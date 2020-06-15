package jsont

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestEvaluateExpression(t *testing.T) {

	testFile := "testdata/input2.json"

	file, _ := ioutil.ReadFile(testFile)

	input := gjson.ParseBytes( /*json*/ file)

	e := newExpressionEvaluator( /*inputData*/ input)

	testData := []struct {
		id          string
		input       string
		expectedRes interface{}
		expectedErr error
	}{
		{
			id:          "single selector",
			input:       "$firstName",
			expectedRes: "John",
		},
		{
			id:          "single path",
			input:       "$address.streetAddress",
			expectedRes: "21 2nd Street",
		},
		{
			id:          "single operation with constant",
			input:       "$age + 5",
			expectedRes: 30.,
		},
		{
			id:          "operation with two values",
			input:       "$firstName : $lastName",
			expectedRes: "JohnSmith",
		},
		{
			id:          "malformed path",
			input:       ".data",
			expectedRes: nil,
			expectedErr: errIllegalPathSeparator,
		},
		{
			id:          "missing closing parenthesis",
			input:       "($firstName",
			expectedRes: nil,
			expectedErr: errMissingRightParenthesis,
		},
		{
			id:          "missing opening parenthesis",
			input:       "$firstName)",
			expectedRes: nil,
			expectedErr: errMissingLeftParenthesis,
		},
	}

	for _, tt := range testData {

		t.Run(tt.id, func(t *testing.T) {

			res, err := e.EvaluateExpression( /*expr*/ tt.input)

			assert.Equal(t, tt.expectedErr, err, "expected error returned by evaluation to match")

			if res == nil {
				res = &gjson.Result{}
			}

			assert.Equal(t, tt.expectedRes, res.Value(), "expected correct result of evaluation")
		})
	}
}

func jsonFromString(in string) *gjson.Result {
	return &gjson.Result{Type: gjson.String, Str: in, Index: 0, Num: 0, Raw: `"` + in + `"`}
}

func jsonFromNumber(in float64) *gjson.Result {
	return &gjson.Result{Type: gjson.Number, Num: in}
}
