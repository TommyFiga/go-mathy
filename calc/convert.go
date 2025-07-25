package calc

import (
	"fmt"
	"strconv"
)


func ConvertToNumber(num string) (float64, error) {
	res, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot parse value %q: %w", num, err)
	}

	return res, err
}