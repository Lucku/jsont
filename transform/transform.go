package transform

import (
	"io/ioutil"

	operator "github.com/lucku/jsont/operator"
	"github.com/tidwall/gjson"
	json "github.com/tidwall/gjson"
	"github.com/valyala/fastjson"
)

type JSONTransformer struct {
	config        *Config
	inputFile     string
	transformData json.Result
	resultData    []byte
	operators     []operator.Operator
}

func NewJSONTransformerWithFile(filePath string, config *Config) (*JSONTransformer, error) {

	if config == nil {
		config = NewConfig()
	}

	fileData, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	return &JSONTransformer{config: config, transformData: json.ParseBytes(fileData)}, nil
}

func NewJSONTransformer(transformBytes []byte, config *Config) *JSONTransformer {

	if config == nil {
		config = NewConfig()
	}

	return &JSONTransformer{config: config, transformData: json.ParseBytes(transformBytes)}
}

func (j *JSONTransformer) Transform(inFile string, outFile *string) error {

	// fmt.Printf("Transform %s with to %s", inFile, outFile)

	// var outJSON fastjson.Arena

	// inBytes, err := ioutil.ReadFile(inFile)

	// if err != nil {
	// 	return err
	// }

	// var p fastjson.Parser

	// jval, err := p.ParseBytes(inBytes)

	// if err != nil {
	// 	return err
	// }

	// var leafIndeces map[string]*fastjson.Value

	// if j.config.indexLeaves {
	// 	leafIndeces = doIndexLeaves()
	// }

	// // Go step by step through transformation file and apply the rules, building up the result json

	// transformResult := gjson.ParseBytes(j.transformData)

	// transformResult.ForEach(processInstruction)

	// if outFile == nil {
	// 	// Write to std.out
	// }

	return nil
}

func doIndexLeaves() map[string]*fastjson.Value {
	// BFS over JSON to extract all leaf nodes and their values into the map

	return nil
}

func processInstruction(key gjson.Result, value gjson.Result) bool {
	return false
}
