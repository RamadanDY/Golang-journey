// functions
package main

import "fmt"

func main() {
	saka(1, 2, 3, "kdd")
}

func saka(a, b, c int, k string) {
	// a = 23
	// b = 2
	// c = 1
	// k = "Ama"
	fmt.Println(a)
}

// func concat(s1 string, s2 string) string {
// 	return s1 + s2
// }

// func main() {
// 	test("name", "is ")
// 	test("rapper", "is ")
// 	test("shiitttt", "is ")
// }
// func test(s1 string, s2 string) {
// 	fmt.Println(concat(s1, s2))

// }

// GO SWITCH STATEMENTS
// package main

// import "fmt"

// func main() {
// 	age := 22
// 	switch age {
// 	case 1:
// 		fmt.Println("yes thats my age ")
// 	case 22:
// 		fmt.Println("okay")
// 	default:
// 		fmt.Println("kk")
// 	}

// }

// package main

// import "fmt"

// func main() {
// 	height := 3
// 	if height > 4 {
// 		fmt.Println("You are tall enough!")
// 	}

// 	email := "ghana@gmail.com"
// 	// /we can do it in a short format ie
// 	if length := len(email); length >
// 		1 {
// 		fmt.Println("the email is correct ")
// 	} else {
// 		fmt.Println("hwllo world")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	name := "Alice"
// 	Lastname := "Yakubu"
// 	age := 28
// 	score := 93.456

// 	msg := fmt.Sprintf("Name: %s | Lastname: %s| Age: %d | Score: %.2f%%", name, Lastname, age, score)
// 	fmt.Println(msg)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	a := 12
// 	b := 34

// 	// Convert to strings
// 	astr := strconv.Itoa(a)
// 	bstr := strconv.Itoa(b)

// 	// Concatenate
// 	result := astr + bstr

// 	fmt.Println("Concatenated:", result) // "1234"
// }
