package main

import "fmt"

func main() {
	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch nana's Golang full Course"
	var rewardDessert = "reward myself with a donut"

	// Slice of strings
	var taskItem = []string{shortGolang, fullGolang, rewardDessert}

	// we pass the parameter into this call back func in the main func
	printTasks(taskItem)

	fmt.Println("::::::::")
	addTask(taskItem, "hiiii")

}

// this is a type of package or global scoped function
// with the parameters side we pass the taskItem in here because we want the function to be able to access the array in the main function
func printTasks(taskItem []string) {
	// Loop through taskItem with the for range tyoe of looping through an array
	for index, value := range taskItem {
		fmt.Println(index+
			1, ":", value)
	}
}

func addTask(taskItem []string, newTask string) {
	// the newTask is a slice
	var updatedTaskItem = append(taskItem, newTask)
	printTasks(updatedTaskItem)

}
