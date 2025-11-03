package main

import (
	"errors"
)

var (
	// ErrNotFound is used when a requested resource cannot be found
	ErrNotFound = errors.New("resource not found")

	// ErrUnauthorized is used when authentication fails
	ErrUnauthorized = errors.New("unauthorized access")

	// ErrInvalidInput is used when the input parameters are invalid
	ErrInvalidInput = errors.New("invalid input parameters")

	// ErrDatabaseConnection is used when database connection fails
	ErrDatabaseConnection = errors.New("database connection error")

	// ErrTimeout is used when an operation exceeds its time limit
	ErrTimeout = errors.New("operation timed out")

	// ErrAlreadyExists is used when trying to create a resource that already exists
	ErrAlreadyExists = errors.New("resource already exists")
)

// Example usage function
func processResource(id string) error {
	// Simulate a not found condition
	if id == "" {
		return ErrNotFound
	}

	// Simulate an unauthorized access
	if id == "restricted" {
		return ErrUnauthorized
	}

	return nil
}
