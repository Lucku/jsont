package main

import (
	"fmt"
	"log"

	"github.com/valyala/fastjson"
)

func ExampleScanner() {
	var sc fastjson.Scanner

	json := `{
		"test": 1,
		"test2": "string",
		"test3:" true,
		"test4", nil	
	}`

	sc.Init(json)

	for sc.Next() {
		fmt.Printf("%s\n", sc.Value())
	}
	if err := sc.Error(); err != nil {
		log.Fatalf("unexpected error: %s", err)
	}

	// Output:
	// {"foo":"bar"}
	// []
	// 12345
	// "xyz"
	// true
	// false
	// null
}
