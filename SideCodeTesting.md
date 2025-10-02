package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 12
	b := 34

	// Convert to strings
	astr := strconv.Itoa(a)
	bstr := strconv.Itoa(b)

	// Concatenate
	result := astr + bstr

	fmt.Println("Concatenated:", result) // "1234"
}
