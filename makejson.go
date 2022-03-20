package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	addressBook := make(map[string]string)

	fmt.Println("Enter Your Name: ")
	var name string
	fmt.Scanln(&name)
	
	fmt.Println("Enter Your Address: ")
	var address string
	fmt.Scanln(&address)

	addressBook["name"] = name
	addressBook["address"] = address

	data, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(string(data))
	}
}