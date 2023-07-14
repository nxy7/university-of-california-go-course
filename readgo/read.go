package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Println("What's the file name?")
	fmt.Scan(&filename)

	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	content := string(f)
	byLine := strings.Split(content, "\n")

	people := make([]Person, 0)
	for _, c := range byLine {
		fname := strings.Split(c, " ")[0]
		lname := strings.Split(c, " ")[1]
		people = append(people, Person{fname: fname, lname: lname})
	}

	for _, p := range people {
		fmt.Println("First name: ", p.fname, " | Last name: ", p.lname)
	}
}
