package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string
	fmt.Println("Type some string")
	fmt.Scan(&userInput)
	userInput = strings.ToLower(userInput)
	if strings.Index(userInput, "i") != 0 || strings.LastIndex(userInput, "n") != len(userInput)-1 || strings.Index(userInput, "a") == -1 {
		fmt.Println("Not Found!")
		return
	}
	fmt.Println("Found!")
}
