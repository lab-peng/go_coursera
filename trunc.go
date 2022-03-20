package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a Decimal Number: ")
	text, _ := reader.ReadString('\n')
	f, _ := strconv.ParseFloat(strings.TrimSpace(text), 64)
	i := int(f)
	fmt.Println(i)
}