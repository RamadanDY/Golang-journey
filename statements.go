package main

import "fmt"

// the switch staement in go
func main() {

	day := 30

	switch day {
	case 1:
		fmt.Println("monday")

	case 2:
		fmt.Println("tuesday")

	case 3:
		fmt.Println("wednesday")

	case 4:
		fmt.Println("thursday")

	case 5:
		fmt.Println("friday")

	case 6:
		fmt.Println("saturday")

	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("no day")
	}

}

// The Nested if Statement
// You can have if statements inside if statements, this is called a nested if.

// Syntax
// if condition1 {
//    // code to be executed if condition1 is true
//   if condition2 {
//      // code to be executed if both condition1 and condition2 are true
//   }
// }

// Go else if Statement

// package main
// import ("fmt")

// func main() {
//   time := 22
//   if time < 10 {
//     fmt.Println("Good morning.")
//   } else if time < 20 {
//     fmt.Println("Good day.")
//   } else {
//     fmt.Println("Good evening.")
//   }
// }

// if statement

// Use the else statement to specify a block of code to be executed if the condition is false.
// The brackets in the else statement should be like } else {:

// func main() {
// 	agec := 423
// 	class := 67
// 	if agec <= class {
// 		fmt.Println("wroneg class")
// 	} else {
// 		fmt.Println("correcvt class")
// 	}
// }
