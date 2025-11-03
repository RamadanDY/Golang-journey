package main

import "fmt"

// fuc to square our positive number and alerting us if the number is a negative number

func square(x int) (int, error) {

	if x < 0 {
		return 0, fmt.Errorf("the value isnt a positive number : %v", x)

	}
	return x * x, nil

}

func main() {

	result, err := square(-3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

// package main

// import "fmt"

// func square(x int) (int, error) {
// 	if x < 0 {
// 		// Why return 0 for the int value?
// 		// When a function needs to return both a normal result and an error, it must return some value for the normal result even if computation failed. 0 is a neutral placeholder since no valid square is computed.
// 		// “Return 0 as the result and an error object describing that the input value (x) was negative.”
// 		// square returns two values to the caller:

// 		// first value: 0 (an int placeholder because the computation didn’t succeed)

// 		// second value: the error object described above
// 		return 0, fmt.Errorf("the value that we want to square if a negative value ie %v", x)
// 	}
// 	// nill means no error
// 	return x * x, nil
// }

// func main() {

// 	result, err := square(-3)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}

// }
