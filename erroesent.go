package main

import (
	"errors"
	"fmt"
)

// the sentinal Error

var ErrNoBread = errors.New("No Bread Left In Stock")

func BreadInStock(stock int) (string, error) {
	if stock == 0 {
		return "", ErrNoBread

	}
	return "Here’s your bread!", nil

}

func main() {
	result, err := BreadInStock(10)
	if err != nil {
		// we want to check if the err that we get is the same as the sentinel one created above

		if errors.Is(err, ErrNoBread) {
			fmt.Println("sorry no bread")

		} else {
			fmt.Println("smt else went wrong", err)
		}

	}

	fmt.Println(result)

}
