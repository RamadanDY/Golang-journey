package main

import (
	"fmt"
)

// CustomError is a custom error type
type CustomError struct {
	Code    int
	Message string
}

// Error implements the error interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &CustomError{
			Code:    400,
			Message: "division by zero is not allowed",
		}
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}
	fmt.Printf("Result: %.2f\n", result)
}
