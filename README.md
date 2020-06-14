# jsont
[![Go Report Card](https://goreportcard.com/badge/github.com/lucku/jsont)](https://goreportcard.com/report/github.com/lucku/jsont)
[![GoDoc](https://godoc.org/github.com/lucku/jsont?status.svg)](https://godoc.org/github.com/lucku/jsont)
[![CircleCI](https://circleci.com/gh/lucku/jsont.svg?style=shield)](https://circleci.com/gh/lucku/jsont)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/lucku/jsont)](https://www.tickgit.com/browse?repo=github.com/lucku/jsont)
[![License MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://img.shields.io/badge/License-MIT-brightgreen.svg)

Transform JSON data using .. JSON

A high performance Go library to transform JSON data using simple and easily understandable instructions written in JSON.

The transformation is performed by providing a JSON of the desired structure, but having substitution statements as leaf-level values, provided as strings. They include selectors to the input JSON document's fields and further allow for operations between various
input fields. The value of an input JSON field could simply be put under a new key in the desired JSON, or be part of an operation together with other input fields.

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
    "from": "$flight.departureAirport",
    "to": "$flight.arrivalAirport",
    "passengers": "$flight.economy.seats + $flight.premiumEconomy.seats + $flight.business.seats",
    "allWifi": "$flight.economy.wifi & $flight.premiumEconomy.wifi & $flight.business.wifi"
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

## Why use jsont?

The transformation of JSON data is a typical task inside of ETL (Extract, transform, load) processes. Independent of the programing language, these transformations usually result in a lot of boilerplate code: unmarshalling the input JSON on one side to set attributes of a target object and marshalling the latter into JSON again. The code for these transformations can be made much more maintainable, less error-prone and the transformation itself more extendable by giving the surrounding concerns of (un)marshalling into the hands of a library and using a declarative rather than an imperative approach to describe the transformation process. This library tries to define an easily understandable declaration language based on JSON itself, therefore completely staying in the domain of data already working with.

## Installation

### Using go get

Install using `go get`:

```bash
$ go get github.com/lucku/jsont
```

### Using magefile (recommended)

Install using `mage` (see https://www.github.com/magefile/mage)

```bash
$ git clone https://github.com/lucku/jsont
$ go get github.com/magefile/mage
$ mage install
```

## Getting Started

### CLI

jsont offers a Command Line Interface (CLI) to be used handily from the shell.

```bash
$ jsont transform -t testdata/trans1.json -o testdata testdata/input1.json
```

### Code

```go
jt, err := jsont.NewJSONTransformerWithFile("testdata/trans1.json")

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
  - Mod (`%`)
  - Power (`^`)
  - Greater (`>`)
  - Greater Equal (`>=`)
  - Less (`<`)
  - Less Equal (`<=`)

- Boolean *(bool, bool) -> bool*
  - And (`&`)
  - Or (`|`)

- Strings *(string, string) -> string*
  - Concatenate (`:`)

- Others
  - IfNull (`?`) *(any, any) -> any*: If the first value happens to be `null`, use the second one (as other operations, can be arbitrarily chained)
    - Example: `"{ "name": "$aircraft.iata ? $aircraft.name" }"` - if `aircraft.iata` is `null`, take `aircraft.name` instead
  - Equal (`==`) *(any, any) -> bool*: Returns true if the arguments are equal
    - Example: `"{ "equalNames": "family.father.name == Tom" }"`
  - NotEqual (`!=`) *(any, any) -> bool*: Returns true if the arguments are not equal
    - Example: `"{ "equalNames": "$family.mother.numChildren != 3 }"`

## ToDos

- [x] Extend the magefile to include version information in binary
- [x] Provide all GoDoc comments
- [ ] Write a full-fledged CLI
- [ ] Support for unary operators like negation
- [ ] Support for same operator on different data types
- [ ] Add more operators
- [ ] Test cases
- [ ] Provide support for functions and implement standard function
- [ ] Completely own implementation of JSON parsing and queries
- [ ] Access of array elements in selector