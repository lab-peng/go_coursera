package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string of letters: ")
	text, _ := reader.ReadString('\n')

	clean := strings.ToLower(strings.TrimSpace(text))
	fmt.Println(clean)

    if (strings.Contains(clean, "a") && strings.HasPrefix(clean, "i") && strings.HasSuffix(clean, "n")) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}