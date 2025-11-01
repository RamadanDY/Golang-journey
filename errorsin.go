package main

import (
	"errors"
	"fmt"
)

// Basic function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Custom error type
type CustomError struct {
	Code    int
	Message string
}

// Implement the error interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

func main() {
	// Basic error handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)

	// Custom error example
	customErr := &CustomError{
		Code:    500,
		Message: "something went wrong",
	}
	fmt.Println("Custom error:", customErr)

	// Multiple error handling
	err = someOperation()
	if err != nil {
		fmt.Println("Operation failed:", err)
	}
}

func someOperation() error {
	// Wrapping errors with fmt.Errorf
	err := anotherOperation()
	if err != nil {
		return fmt.Errorf("failed to perform operation: %w", err)
	}
	return nil
}

func anotherOperation() error {
	return errors.New("operation failed")
}
