package transform

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type JSONTransformer struct {
	config        *Config
	inputFile     string
	transformData gjson.Result
}

func NewJSONTransformer(transformBytes []byte, config *Config) (*JSONTransformer, error) {

	if config == nil {
		config = NewConfig()
	}

	if config.validate {
		if ok := validateJSON(transformBytes); !ok {
			return nil, errors.New("Transformation file has invalid JSON syntax")
		}
	}

	return &JSONTransformer{config: config, transformData: gjson.ParseBytes(transformBytes)}, nil
}

func NewJSONTransformerWithFile(transformFile string, config *Config) (*JSONTransformer, error) {

	fileData, err := ioutil.ReadFile(transformFile)

	if err != nil {
		return nil, err
	}

	return NewJSONTransformer(fileData, config)
}

func (j *JSONTransformer) Transform(inData []byte) ([]byte, error) {

	if j.config.validate {
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

	it := JSONIterator{Data: &j.transformData}

	opts := &sjson.Options{
		Optimistic:     true,
		ReplaceInPlace: false,
	}

	for it.Next() {

		cur := it.Value()

		if cur.Value.Type == gjson.String {

			//fmt.Println(string(outData))

			result, err := processInstruction(&input, cur.Value.String())

			if err != nil {
				return nil, err
			}

			if outData, err = sjson.SetBytesOptions(outData, strings.Join(cur.Path, "."), result, opts); err != nil {
				return nil, err
			}
		}

	}

	// if outFile == nil {
	// 	// Write to std.out
	// }

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

	it := JSONIterator{Data: &data}

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

func processInstruction(input *gjson.Result, instr string) (interface{}, error) {

	// implement a state machine

	fmt.Println("Process", instr)

	result := input.Get(instr)

	// First step: Just act like there are no operators and simply resolve references

	// make a copy of type of instr

	return result.Value(), nil
}
