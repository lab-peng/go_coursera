package main

import "fmt"

func main() {
	fmt.Println("Enter the floating point number: ")

	var float float64

	fmt.Scanln(&float)

	var integer int64 = int64(float)

	fmt.Println("The integer value is: ", integer)
}