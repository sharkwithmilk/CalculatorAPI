package Calculator

import "errors"

var (
	ErrInvalidExpression  = errors.New("invalid expression")
	ErrDivisionByZero     = errors.New("division by zero")
	ErrEmptyExpression    = errors.New("empty expression")
	ErrNumberParsing      = errors.New("error parsing number")
	ErrMissingNumber      = errors.New("missing number after operator")
	ErrUnrecognizedNumber = errors.New("failed to recognize number")
	ErrInvalidParentheses = errors.New("invalid parentheses")
)
