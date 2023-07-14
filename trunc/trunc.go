package main

import "fmt"

func main() {
	var num float64
	fmt.Println("Type some float number:")
	fmt.Scan(&num)
	var truncated int64
	truncated = int64(num)
	fmt.Println(truncated)
}
