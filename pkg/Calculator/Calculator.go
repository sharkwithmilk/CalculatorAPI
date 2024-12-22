package Calculator

import "fmt"
import "strconv"
import "strings"

func Calc(expression string) (float64, error) {
	expression = strings.Replace(expression, " ", "", -1)
	if expression == "" {
		return 0, ErrEmptyExpression
	}
	res, err := eval(expression)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func eval(str string) (float64, error) {
	for strings.Contains(str, "(") {
		openInd := strings.LastIndex(str, "(")
		closeInd := strings.Index(str[openInd:], ")") + openInd
		if closeInd <= openInd {
			return 0, ErrInvalidParentheses
		}
		innerResult, err := eval(str[openInd+1 : closeInd])
		if err != nil {
			return 0, err
		}
		str = str[:openInd] + fmt.Sprintf("%f", innerResult) + str[closeInd+1:]
	}
	return parseAddSub(str)
}

func parseAddSub(str string) (float64, error) {
	parts := splitByOperators(str, "+-")
	total := 0.0
	sign := 1.0

	for i := 0; i < len(parts); i++ {
		part := parts[i]
		if part == "+" {
			sign = 1.0
		} else if part == "-" {
			sign = -1.0
		} else {
			value, err := parseMulDiv(part)
			if err != nil {
				return 0, err
			}
			total += sign * value
		}
	}
	return total, nil
}

func parseMulDiv(str string) (float64, error) {
	parts := splitByOperators(str, "*/")
	if len(parts) == 0 {
		return 0, ErrEmptyExpression
	}

	chislo, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, ErrNumberParsing
	}

	for i := 1; i < len(parts); i += 2 {
		if i+1 >= len(parts) {
			return 0, ErrMissingNumber
		}

		operator := parts[i]
		nextChislo, err := strconv.ParseFloat(parts[i+1], 64)
		if err != nil {
			return 0, ErrUnrecognizedNumber
		}

		if operator == "*" {
			chislo *= nextChislo
		} else if operator == "/" {
			if nextChislo == 0 {
				return 0, ErrDivisionByZero
			}
			chislo /= nextChislo
		}
	}
	return chislo, nil
}

func splitByOperators(str, operators string) []string {
	var result []string
	current := ""

	for _, char := range str {
		if strings.ContainsRune(operators, char) {
			if current != "" {
				result = append(result, current)
				current = ""
			}
			result = append(result, string(char))
		} else {
			current += string(char)
		}
	}

	if current != "" {
		result = append(result, current)
	}
	return result
}
