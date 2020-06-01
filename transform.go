package jsont

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/lucku/jsont/json"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// JSONTransformer transforms JSON input data of a file into output data using a set of transformation descriptions
type JSONTransformer struct {
	options       options
	inputFile     string
	transformData gjson.Result
}

// NewJSONTransformer instantiates a new instance of a JSONTransformer struct which is used to perform
// the transformation of input JSON data by following instructions given in a JSON format as well. This JSON
// giving the transformation description has to be provided as bytes as an input to this function.
//
// As further input, this function takes a list of options which allows certain control of the transformation
// behavior. See config.go for a list of configuration options.
//
// This function returns an error if the provided transformation bytes are not valid JSON.
func NewJSONTransformer(transformJSON []byte, opts ...Option) (*JSONTransformer, error) {

	options := newOptions(opts...)

	if options.validate {
		if ok := validateJSON(transformJSON); !ok {
			return nil, errors.New("Transformation file has invalid JSON syntax")
		}
	}

	return &JSONTransformer{options: options, transformData: gjson.ParseBytes(transformJSON)}, nil
}

// NewJSONTransformerWithFile instantiates a new instance of a JSONTransformer struct which is used to perform
// the transformation of input JSON data by following instructions given in a JSON format as well. Unlike the
// instantiation using NewJSONTransformer(), this function takes a filesystem path to a file that contains
// the JSON with transformation instructions. The path can be either relative to the execution context of the program
// or absolute on the system.
//
// As further input, this function takes a list of options which allows certain control of the transformation
// behavior. See config.go for a list of configuration options.
//
// This function returns an error in at least one of the following cases:
//
// (1) the provided filesystem path is not valid
//
// (2) the provided transformation bytes are not valid JSON
func NewJSONTransformerWithFile(transformFile string, opts ...Option) (*JSONTransformer, error) {

	fileData, err := ioutil.ReadFile(transformFile)

	if err != nil {
		return nil, err
	}

	return NewJSONTransformer(fileData, opts...)
}

// Transform takes as an input a JSON as byte array and applies the JSONTransformer's transformation description
// on the provided data. In case of a successful transformation, the output JSON is returned as byte array.
//
// This function returns an error in at least one of the following cases:
//
// (1) the provided input bytes are not valid JSON
//
// (2) an instruction of the transformation description cannot be performed due to invalid syntax of the expression
// or wrong input data types
func (j *JSONTransformer) Transform(inJSON []byte) ([]byte, error) {

	if j.options.validate {
		if ok := validateJSON(inJSON); !ok {
			return nil, errors.New("Input file has invalid JSON syntax")
		}
	}

	input := gjson.ParseBytes(inJSON)

	// var leafIndices map[string]gjson.Result

	// if j.config.indexLeaves {
	// 	leafIndices = doIndexLeaves(input)
	// }

	// Go step by step through transformation file and apply the rules, building up the result json

	// Output data has the same structure as input data
	var outData []byte = []byte(j.transformData.Raw)

	it := json.NewIterator(&j.transformData)

	evaluator := newExpressionEvaluator(input)

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

// TransformWithFile takes JSON data as an input and applies the JSONTransformer's transformation description
// on the provided data. In case of a successful transformation, the output JSON is returned as byte array. Unlike
// Transform(), which accepts raw JSON bytes as input, this function takes a filesystem path to a file that contains
// the JSON to be transformed. The path can be either relative to the execution context of the program or absolute on
// the system.
//
// This function returns an error in at least one of the following cases:
//
// (1) the provided filesystem path is not valid
//
// (2) the provided input bytes are not valid JSON
//
// (3) an instruction of the transformation description cannot be performed due to invalid syntax of the expression
// or wrong input data types
func (j *JSONTransformer) TransformWithFile(inJSONFile string) ([]byte, error) {

	inBytes, err := ioutil.ReadFile(inJSONFile)

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
