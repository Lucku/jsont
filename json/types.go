package json

import (
	"github.com/tidwall/gjson"
)

// Type represents a JSON type
type Type int

const (
	// Null is a JSON null value and default value if a value is not set
	Null Type = iota
	// Number is JSON number
	Number
	// Bool is a JSON boolean
	Bool
	// String is a JSON string
	String
	// Object is a JSON object
	Object
	// Array is a JSON array
	Array
	// Any represents any JSON data type (including null) and is only used as a "all-matcher" for
	// the checkType function
	Any
)

func (t Type) String() string {

	switch t {
	case Null:
		return "null"
	case Number:
		return "number"
	case Bool:
		return "bool"
	case String:
		return "string"
	case Object:
		return "object"
	case Array:
		return "array"
	case Any:
		return "any"
	}

	return ""
}

// GetJSONType returns the type of a gjson.Result, mapping its gjson.Type to the internally used one
func GetJSONType(r gjson.Result) Type {

	switch r.Type {
	case gjson.False, gjson.True:
		return Bool
	case gjson.Number:
		return Number
	case gjson.String:
		return String
	case gjson.JSON:
		if r.IsArray() {
			return Array
		}
		return Object
	default:
		return Null
	}
}

// CheckType takes a gjson.Result and an internal type and returns true, if the JSON value satisfies the type
//
// This function *must* be used whenever a gjson.Result has to be checked against an internal type and acts as
// a single source of truth.
func CheckType(r gjson.Result, t Type) bool {

	if t == Any || t == GetJSONType(r) {
		return true
	}

	return false
}
