package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	err := fmt.Sprintf("can not divide %v by zero",d.dividend) 
	return  err
	}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
