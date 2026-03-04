package main

import "fmt"
func firstFunction() (int error) {
	return fmt.Errorf("original error or smt went wrong with the firstFunction")
	
}

func secondFunction() (int error)  {
	///lets bring in the first one 
	err :=  firstFunction()
	if err != nil {
		fmt.Println(err)
	} 
	return nil
}
func main() {

	err :=  secondFunction()
	fmt.Println(err)


}