package helper

import "fmt"

func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("Demoninator can not be zero")
	}

	return num1 / num2, nil
}
