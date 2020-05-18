package transform

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/lucku/jsont/json"
	"github.com/lucku/jsont/transform"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// JSONTransformer transforms JSON input data of a file into output data using a transformation policy
type JSONTransformer struct {
	options       options
	inputFile     string
	transformData gjson.Result
}

func NewJSONTransformer(transformBytes []byte, opts ...Option) (*JSONTransformer, error) {

	options := newOptions(opts...)

	if options.validate {
		if ok := validateJSON(transformBytes); !ok {
			return nil, errors.New("Transformation file has invalid JSON syntax")
		}
	}

	return &JSONTransformer{options: options, transformData: gjson.ParseBytes(transformBytes)}, nil
}

func NewJSONTransformerWithFile(transformFile string, opts ...Option) (*JSONTransformer, error) {

	fileData, err := ioutil.ReadFile(transformFile)

	if err != nil {
		return nil, err
	}

	return NewJSONTransformer(fileData, opts...)
}

func (j *JSONTransformer) Transform(inData []byte) ([]byte, error) {

	if j.options.validate {
		if ok := validateJSON(inData); !ok {
			return nil, errors.New("Input file has invalid JSON syntax")
		}
	}

	input := gjson.ParseBytes(inData)

	// var leafIndices map[string]gjson.Result

	// if j.config.indexLeaves {
	// 	leafIndices = doIndexLeaves(input)
	// }

	// Go step by step through transformation file and apply the rules, building up the result json

	// Output data has the same structure as input data
	var outData []byte = []byte(j.transformData.Raw)

	it := json.NewIterator(&j.transformData)

	evaluator := transform.NewExpressionEvaluator(input)

	opts := &sjson.Options{
		Optimistic:     true,
		ReplaceInPlace: false,
	}

	for it.Next() {

		cur := it.Value()

		if cur.Value.Type == gjson.String {

			result, err := evaluator.EvaluateExpression(cur.Value.String())

			if err != nil {
				return nil, fmt.Errorf("failed to process instruction '%s': %w", cur.Value.String(), err)
			}

			if outData, err = sjson.SetBytesOptions(outData, strings.Join(cur.Path, "."), result.Value(), opts); err != nil {
				return nil, err
			}
		}

	}

	return outData, nil
}

func (j *JSONTransformer) TransformWithFile(inFile string) ([]byte, error) {

	inBytes, err := ioutil.ReadFile(inFile)

	if err != nil {
		return nil, err
	}

	return j.Transform(inBytes)
}

func doIndexLeaves(data gjson.Result) map[string]*gjson.Result {

	leaves := make(map[string]*gjson.Result)

	it := json.NewIterator(&data)

	for it.Next() {
		cur := it.Value()

		// Element is leaf node if its type is not JSON
		if cur.Value.Type != gjson.JSON {
			leaves[cur.Path[len(cur.Path)-1]] = cur.Value
		}
	}

	return leaves
}

func validateJSON(jsonData []byte) bool {
	return gjson.ValidBytes(jsonData)
}
