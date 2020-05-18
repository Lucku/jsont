# jsont
[![GoDoc](https://godoc.org/github.com/Lucku/jsont?status.svg)](https://godoc.org/github.com/Lucku/jsont)

Transform JSON data using .. JSON

A high performance Go library to transform JSON data using simple and easily understandable instructions written in JSON.

The transformation is performed by providing a JSON of the desired structure, but having substitution statements as leaf-level values, provided as strings. They include selectors to the input JSON's fields and further allow for operations between various
input fields. The value of an input JSON field could simply be put under a new key in the desired JSON, or even be part of an operation together with other input fields.

## Example

Input:
```json
{
    "version": "1.0",
    "flight": {
        "arrivalAirport": "HAM",
        "departureAirport": "DUS",
        "economy": {
            "seats": 150,
            "wifi": true
        },
        "premiumEconomy": {
            "seats": 100,
            "wifi": true
        },
        "business": {
            "seats": 10,
            "wifi": true
        }
    }
}
```

Transformation description:
```json
{
    "from": "flight.departureAirport",
    "to": "flight.arrivalAirport",
    "passengers": "flight.economy.seats + flight.premiumEconomy.seats + flight.business.seats",
    "allWifi": "flight.economy.wifi & flight.premiumEconomy.wifi & flight.business.wifi"
}
```

Output:
```json
{
    "from": "DUS",
    "to": "HAM",
    "passengers": 260,
    "allWifi": true
}
```

## Installation

### Using go get

Install using `go get`:

```
$ go get github.com/Lucku/jsont/...
```

### Using magefile (recommended)

Install using `mage` (see https://www.github.com/magefile/mage)

```
$ git clone https://github.com/Lucku/jsont
$ go get github.com/magefile/mage
$ mage install
```

## Getting Started

### CLI

jsont offers a Command Line Interface (CLI) to be used handily from the shell.

```
$ jsont transform -t testdata/trans1.json -o testdata testdata/input1.json
```

### Code

```go
jt, err := transform.NewJSONTransformerWithFile("testdata/trans1.json")

if err != nil {
    return err
}

out, err := jt.TransformWithFile("testdata/input1.json")

if err != nil {
    return err
}

fmt.Println(out)

/* 
output: 
{
    "from": "DUS",
    "to": "HAM",
    "passengers": 260,
    "allWifi": true
}
*/
```

## Operators

- Arithmetic *(number, number) -> number*
  - Add (`+`)
  - Subtract (`-`)
  - Multiply (`*`)
  - Divide (`/`)

- Boolean *(bool, bool) -> bool*
  - And (`&`)
  - Or (`|`)

- Strings *(string, string) -> string*
  - Concatenate (`:`)

- Others
  - IfNull (`?`) *(any, any) -> any*: If the first value happens to be null, use the second one
    - Example: `"{ "name": "aircraft.iata?aircraft.name" }"` - if aircraft.iata is null, take aircraft.name instead
  - Compare (`=`) *(any, any) -> bool*: Returns true if the arguments are equal
    - Example: `"{ "equalNames": "family.father.name=family.firstSon.name" }"`

## ToDos

- Provide all GoDoc comments
- Write a full-fledged CLI
- Extend the magefile to include version information in binary
- Add more operators
- Test cases
- Completely own implementation of JSON parsing and queries
- Access of array elements in selector