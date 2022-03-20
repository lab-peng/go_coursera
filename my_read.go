package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	f, _ := os.Open("../go_files/names.txt")
	r := bufio.NewScanner(f)
	names := []name{}

	for r.Scan() {
		line := strings.Split(r.Text(), " ")
		n := name{
			fname: line[0],
			lname: line[1],
		}
		names = append(names, n)
	}
	
	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}