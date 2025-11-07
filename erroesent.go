package main

import (
	"errors"
	"fmt"
)

var ErrorBreads = errors.New("there is no bread available")

func createbread(x int) (string, error) {
	if x == 0 {
		return "", ErrorBreads

	}
	return "Here is your bread", nil
}

func main() {
	result ,err:= createbread(0){
		if err != nil {
			if errors.Is(err,ErrorBreads)
		}else{
			fmt.Println("this is a diff error",err)
		}
	}
	fmt.Println(result)
}

/*
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
	return "Here is your bread!", nil

}

func main() {
	result, err := BreadInStock(10)
	if err != nil {
		// we want to check if the err that we get is the same as the sentinel one created above

		if errors.Is(err, ErrNoBread) {
			fmt.Println("d")

		} else {
			fmt.Println("smt else  ", err)
		}

	}

	fmt.Println(result)

}

*/
