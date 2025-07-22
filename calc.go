package main

import (
	"fmt"
	"errors"
	"strconv"
)


func Calculator(a float64, op string, b float64) (float64, error) {
	switch op {
		case "+":
			return a + b, nil
		case "-":
			return a - b, nil
		case "*":
			return a * b, nil
		case "/":
			if (b == 0) {
				return 0, errors.New("division by zero") 
			}
			return a / b, nil
		default:
			return 0, errors.New("invalid operator")
	}
}


func convertToNumber(num string) (float64, error) {
	res, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot parse value %q: %w", num, err)
	}

	return res, err
}
