package calc

import "errors"


func Calculator(a float64, op string, b float64) (float64, error) {
	switch op {
		case "+":
			return a + b, nil
		case "-":
			return a - b, nil
		case "*", "x", "X":
			return a * b, nil
		case "/":
			if b == 0 {
				return 0, errors.New("division by zero") 
			}
			return a / b, nil
		default:
			return 0, errors.New("invalid operator")
	}
}

