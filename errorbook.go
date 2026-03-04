package main

import (
	"fmt"
)

var myerror = fmt.Errorf("hellpoooo  u messed up here ")


//It checks if x is negative.

func square(x int) (int ,error) {
	if x > 0 {
		//lets introduce our custom message 
		return 0, myerror		//return 0, fmt.Errorf("x must be negative: %d", x)
	}

	return x * x, nil

}

func main() {
	result ,err :=  square(2)

	if err != nil {
		 fmt.Println(err) 
	} else { 
		fmt.Println(result)
    }
}









// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// // the error is a return value that's why it's singular

// func processFile(path string) error {
// 	if path == "" {
// 		return errors.New("path is empty")
// 	}

// 	f, err := os.Open(path)
// 	if err != nil {
// 		return err
// 	}

// 	// Always close opened files
// 	defer f.Close()

// 	if f.Name() != "Arrays&Slices.txt" {
// 		return fmt.Errorf("unexpected filename: %s", f.Name())
// 	}

// 	return nil
// }

// func main() {
// 	err := processFile("./helloo.txt")

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("File processed successfully")
// }








// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
// 	if denominator == 0 {
// 		return 0, 0, errors.New("denominator or cannot br  is 0")
// 	}
// 	return numerator / denominator, numerator % denominator, nil
// }

// func main() {
// 	result1, res2, err := calcRemainderAndMod(8, 4)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result1, res2)
// 	}

// }
