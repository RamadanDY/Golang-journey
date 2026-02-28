package main

import (
	"fmt"
	"net/http"
)

// the list or arrays and also slices
var thetask = []string{
	"Learn Go",
	"Build a Todo App",
	"Practice HTTP",
}

func main(){
	fmt.Println("##### Welcome to our todo list app")

	// lets write a handler ie a response 
	http.HandleFunc("/",helloUser)
	http.HandleFunc("/show-tasks",showtasks)

	http.ListenAndServe(":8082",nil)	
	

}
// lets create a func that will be use to handle the res and the req 
// from the users info that he will send to the route that has been 
// created and named helloUser 

func helloUser(writer http.ResponseWriter,handler *http.Request) {
	var greeting = "hello user wellcome to our todo list"
	fmt.Fprintf(writer,greeting)
	fmt.Println("this is my first route in gp")
}


func showtasks(writer http.ResponseWriter,handler *http.Request) {
	// var thisis = "this is what i just want to display"

	for index , task := range thetask {
		fmt.Fprintf(writer,"%d  %s\n",index+1,task)
	}
	fmt.Fprintln(writer,thetask)
}


func PrintTask(thetask []string){

	fmt.Println("these are the list of todo app")

	// the for loop in golang to display them

	for index, task := range thetask {
		fmt.Printf("format: %d\n %s\n",index,task)
	}
}

func addingTasks(thetask []string,newTask string) ([]string) {
	var updated = append(thetask,newTask)
	return updated

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
