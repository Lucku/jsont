package transform

import (
	"fmt"
	"go/scanner"
	"go/token"

	"github.com/lucku/jsont/operation"
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

func newExpressionEvaluator(inputData gjson.Result) *expressionEvaluator {
	operands := newTypeSafeStack(0)
	operators := newTypeSafeStack(0)
	qualifiers := newTypeSafeStack(0)
	return &expressionEvaluator{inputData: inputData, operands: operands, operators: operators, qualifiers: qualifiers}
}

func (e *expressionEvaluator) EvaluateExpression(expr string) (*gjson.Result, error) {

	s := initScanner(expr)

	prev := token.ILLEGAL

Scan:
	for {

		_, tok, lit := s.Scan()

		switch {
		case tok == token.EOF:
			break Scan
		case isOperator(tok.String()):
			e.evalAllQualifiers()
			e.evalAllLowerOperators(tok.String())
		case tok == token.PERIOD:
			// we have a select operation here: check if last token was an operand
			if prev.IsOperator() || prev == token.ILLEGAL {
				return nil, fmt.Errorf("not able to parse instruction: select (.) without preceding operand")
			}
		case tok == token.LPAREN:
			e.operators.Push(tok.String())
		case tok == token.RPAREN:

			for e.operators.Size() > 0 && e.operators.Peek().(string) != "(" {
				if err := e.evalOperation(); err != nil {
					return nil, err
				}
			}

			// pop closing parenthesis
			if e.operators.Pop().(string) != "(" {
				return nil, fmt.Errorf("no matching parenthesis found")
			}

		case tok == token.SEMICOLON:
		default: //operand

			if prev == token.PERIOD {
				lastQualifier := e.qualifiers.Pop().(string)
				lastQualifier += "."
				lastQualifier += lit
				e.qualifiers.Push(lastQualifier)
			} else {
				e.qualifiers.Push(lit)
			}
		}

		prev = tok
	}

	for e.operators.Size() > 0 {

		e.evalAllQualifiers()

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
	return operation.IsOperator(lit)
}

func (e *expressionEvaluator) evalAllLowerOperators(operator string) {

	parsedOp := operation.GetOperator(operator)

	for e.operators.Size() > 0 {

		nextOp := e.operators.Peek().(string)

		nextOpParsed := operation.GetOperator(nextOp)

		if nextOpParsed == nil {
			break
		}

		if parsedOp.Precedence >= nextOpParsed.Precedence {
			break
		}

		e.evalOperation()
	}

	e.operators.Push(operator)
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

	parsedOp := operation.GetOperator(op)

	args := make([]gjson.Result, len(parsedOp.ArgTypes))

	for i := len(parsedOp.ArgTypes) - 1; i >= 0; i-- {

		arg := e.operands.Pop()

		if arg == nil {
			return fmt.Errorf("not enough arguments in call to %s (%s) operator", parsedOp.Identifier, parsedOp.Sign)
		}

		val := arg.(gjson.Result)

		// check data types of operands
		if exp := parsedOp.ArgTypes[i]; val.Type != exp && exp != operation.Any {
			return fmt.Errorf("argument %v is not of expected type %v", val, exp)
		}

		args[i] = val
	}

	res := parsedOp.Apply(args...)

	e.operands.Push(res)

	return nil
}
