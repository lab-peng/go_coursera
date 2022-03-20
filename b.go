package main

import "fmt"

func main() {

	var s = []int{1, 2, 3}


	for i, v := range s {
		fmt.Println(i, v)
	}
	// fmt.Println("Enter the floating point number: ")

	// var float float64

	// fmt.Scanln(&float)

	// s = append(s, int(float))


	// fmt.Println(s)
}