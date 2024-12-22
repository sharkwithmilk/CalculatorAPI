package Calculator

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      error
	}{
		{"1+2", 3, nil},
		{"10-5", 5, nil},
		{"2*3", 6, nil},
		{"8/4", 2, nil},
		{"1+2*3", 7, nil},
		{"(1+2)*3", 9, nil},
		{"10/(5-5)", 0, ErrDivisionByZero},
		{"", 0, ErrEmptyExpression},
		{"2*(3+4)-5", 9, nil},
		{"3.5+2.5", 6, nil},
		{"(2+3)*(4-1)", 15, nil},
		{"2*((3+4)-5)", 4, nil},
		{"2+3/", 0, ErrMissingNumber},
		{"(2+3", 0, ErrInvalidParentheses},
		{"2++3", 0, ErrInvalidExpression},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := Calc(test.input)

			if err != test.err {
				t.Errorf("unexpected error: got %v, want %v", err, test.err)
			}

			if result != test.expected {
				t.Errorf("unexpected result: got %v, want %v", result, test.expected)
			}
		})
	}
}
