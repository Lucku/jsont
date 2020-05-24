package transform

import (
	"fmt"
	"go/scanner"
	"go/token"
	"strconv"

	"github.com/lucku/jsont/json"
	"github.com/lucku/jsont/operator"
	"github.com/tidwall/gjson"
)

type ExpressionEvaluator interface {
	EvaluateExpression(expr string) (*gjson.Result, error)
}

type expressionEvaluator struct {
	inputData  gjson.Result
	operands   Stack
	operators  Stack
	qualifiers Stack
}

func newExpressionEvaluator(inputData gjson.Result) ExpressionEvaluator {
	operands := newTypeSafeStack(0)
	operators := newTypeSafeStack(0)
	qualifiers := newTypeSafeStack(0)
	return &expressionEvaluator{inputData: inputData, operands: operands, operators: operators, qualifiers: qualifiers}
}

const _qualifierIndicator = "$"

func (e *expressionEvaluator) EvaluateExpression(expr string) (*gjson.Result, error) {

	s := initScanner(expr)

	prevTok := token.ILLEGAL
	prevLit := ""
	readingQualifier := false
	currentQualifier := ""

Scan:
	for {

		_, tok, lit := s.Scan()

		switch {
		case tok == token.EOF:
			break Scan
		case isOperator(tok.String()):
			if readingQualifier {
				e.qualifiers.Push(currentQualifier)
				currentQualifier = ""
				readingQualifier = false
			}
			e.evalAllQualifiers()
			e.evalAllLowerOperators(tok.String())
		case tok == token.PERIOD:
			// we have a select operation here: check if last token was an operand
			if !readingQualifier || prevTok == token.ILLEGAL {
				return nil, fmt.Errorf("not able to parse instruction: path separator (.) without preceding qualifier")
			}
		case tok == token.LPAREN:
			e.operators.Push(tok.String())
		case tok == token.RPAREN:
			if readingQualifier {
				e.qualifiers.Push(currentQualifier)
				currentQualifier = ""
				readingQualifier = false
			}
			for e.operators.Size() > 0 && e.operators.Peek().(string) != "(" {
				if err := e.evalOperation(); err != nil {
					return nil, err
				}
			}

			// pop closing parenthesis
			if e.operators.Pop().(string) != "(" {
				return nil, fmt.Errorf("no matching parenthesis found")
			}
		case tok == token.FLOAT, tok == token.INT:
			rawNum, _ := strconv.ParseFloat(lit, 64)
			jsonNum := gjson.Result{Type: gjson.Number, Num: rawNum}
			e.operands.Push(jsonNum)
		case lit == "true" && !readingQualifier: // if not part of qualifier, this is a 'true' literal
			jsonTrue := gjson.Result{Type: gjson.True}
			e.operands.Push(jsonTrue)
		case lit == "false" && !readingQualifier: // if not part of qualifier, this is a 'false' literal
			jsonFalse := gjson.Result{Type: gjson.False}
			e.operands.Push(jsonFalse)
		case tok == token.SEMICOLON: // semicolon is automatically added at the end by the scanner, can be used to push last qualifier
			if readingQualifier {
				e.qualifiers.Push(currentQualifier)
			}
		case lit == _qualifierIndicator:
			if readingQualifier {
				return nil, fmt.Errorf("illegal \"$\" inside of qualifier")
			}
			readingQualifier = true
		case prevLit == _qualifierIndicator, prevTok == token.PERIOD: // path segment of qualifier
			if currentQualifier != "" {
				currentQualifier += "."
			}
			currentQualifier += lit
		default: // string literal
			jsonString := gjson.Result{Type: gjson.String, Str: lit}
			e.operands.Push(jsonString)
		}

		prevTok = tok
		prevLit = lit
	}

	e.evalAllQualifiers()

	for e.operators.Size() > 0 {

		if err := e.evalOperation(); err != nil {
			return nil, err
		}
	}

	res := e.operands.Pop().(gjson.Result)

	return &res, nil
}

func initScanner(in string) scanner.Scanner {

	src := []byte(in)

	fs := token.NewFileSet()
	file := fs.AddFile("", fs.Base(), len(src))

	var s scanner.Scanner
	s.Init(file, src, nil, 0)

	return s
}

func isOperator(lit string) bool {
	return operator.IsOperator(lit)
}

func (e *expressionEvaluator) evalAllLowerOperators(op string) {

	parsedOp := operator.GetOperator(op)

	for e.operators.Size() > 0 {

		nextOp := e.operators.Peek().(string)

		nextOpParsed := operator.GetOperator(nextOp)

		if nextOpParsed == nil {
			break
		}

		if parsedOp.Precedence >= nextOpParsed.Precedence {
			break
		}

		e.evalOperation()
	}

	e.operators.Push(op)
}

func (e *expressionEvaluator) evalAllQualifiers() {

	for e.qualifiers.Size() > 0 {
		q := e.qualifiers.Pop().(string)
		val := e.inputData.Get(q)

		e.operands.Push(val)
	}
}

func (e *expressionEvaluator) evalOperation() error {

	op := e.operators.Pop().(string)

	parsedOp := operator.GetOperator(op)

	args := make([]gjson.Result, len(parsedOp.ArgTypes))

	for i := len(parsedOp.ArgTypes) - 1; i >= 0; i-- {

		arg := e.operands.Pop()

		if arg == nil {
			return fmt.Errorf("not enough arguments in call to %s (%s) operator", parsedOp.Identifier, parsedOp.Sign)
		}

		val := arg.(gjson.Result)

		if exp := parsedOp.ArgTypes[i]; !json.CheckType(val, exp) {
			return fmt.Errorf("argument '%v' is not of expected type %v", val, exp)
		}

		args[i] = val
	}

	res := parsedOp.Apply(args...)

	e.operands.Push(res)

	return nil
}
