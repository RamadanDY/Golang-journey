package main

import "errors"

// the sentinal Error

var ErrNoBread = errors.New("No Bread Left In Stock")

func BreadInStock(stock int) (string, error) {
	if stock == 0 {
		return "", ErrNoBread

	}
	return "Here’s your bread!", nil

}

func main() {
	result, err := BreadInStock(0)
	if err != nil {

	}

}
