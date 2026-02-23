package main

import "fmt"

func main(){

	shortGolang := "Watch our crash course"
	fullGolang := "watch Nanas golang full course"
	rewardDessert := "eward myself with a sheesecake"

	fmt.Println("##### Welcome to our todo list app")
	
	thetask := []string {shortGolang,fullGolang,rewardDessert}
	// the list or arrays and also slices
	fmt.Println()


	fmt.Println("these are the list of todo app")
	fmt.Println(thetask)

	// the for loop in golang to display them

	for index, task := range thetask  {
		fmt.Println(index+1,".",task)
		fmt.Printf()
	}




}





// package main

// import "fmt"

// func main() {
// 	var shortGolang = "Watch Go crash course"
// 	var fullGolang = "Watch nana's Golang full Course"
// 	var rewardDessert = "reward myself with a donut"
// 	var what = "what is thr"

// 	// Slice of strings
// 	var taskItem = []string{shortGolang, fullGolang, rewardDessert, what}

// 	// we pass the parameter into this call back func in the main func
// 	printTasks(taskItem)

// 	fmt.Println("::::::::")
// 	addTask(taskItem, "hiiii")

// }

// // this is a type of package or global scoped function
// // with the parameters side we pass the taskItem in here because we want the function to be able to access the array in the main function
// func printTasks(taskItem []string) {
// 	// Loop through taskItem with the for range tyoe of looping through an array
// 	for index, value := range taskItem {
// 		fmt.Println(index+
// 			1, ":", value)
// 	}
// }

// func addTask(taskItem []string, newTask string) {
// 	// the newTask is a slice
// 	var updatedTaskItem = append(taskItem, newTask)
// 	printTasks(updatedTaskItem)

// }
