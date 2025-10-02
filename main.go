// package main

// import "fmt"

// func main() {
// 	// these are the two ways of  naming variables in go
// 	var nameOf string = "Gaza"
// 	mylast := "RAmadan"
// 	fmt.Println("welcome to our todo list app ")
// 	fmt.Println(nameOf)
// 	fmt.Println(mylast)

// }
package main

import (
	"fmt"
)

func main() {
	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch nana's Golang full Course"
	var rewardDessert = "reward myself with a donut"

	// Slice of strings
	var taskItem = []string{shortGolang, fullGolang, rewardDessert}

	fmt.Println("Here is an example slice of integers:")

	// Slice of integers
	mySlice := []int{1, 2, 3, 4, 5, 65, 3}
	fmt.Println(mySlice)

	fmt.Println("Tasks")

	// Loop through taskItem
	for index, task := range taskItem {
		fmt.Println(index+1, ":", task)
	}
}

/*

func main() {
	var a string
	var b string
	var c int
	var d int
	// var b int
	// var c bool
	a = "Ramadan"
	b = "23"
	c = 45
	d = 20

	fmt.Println(a + "hellooo" + b)
	fmt.Println(c + d)
	// fmt.Println(b)
	// fmt.Println(c)

	// var taskItem = [2]string{a, b}
	// // we can also use the ones without initial value
	// taskName := [9]int{1, 2, 3, 4, 5}
	// fmt.Println(taskItem)
	// fmt.Println(taskName)
}
*/
