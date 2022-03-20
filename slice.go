package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)


func main() {
	s := make([]int, 3)

	for true {
		fmt.Println("Enter an integer: ")
		var input string
		fmt.Scanln(&input)
		if input == "X" {
			os.Exit(0)
		}

		integer, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}


		s = append(s, integer)
		sort.Ints(s)
		fmt.Println(s)

	}
}