package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, add string
	fmt.Println("Type in your name")
	fmt.Scan(&name)
	fmt.Println("Type in your address")
	fmt.Scan(&add)

	m := make(map[string]string)
	m["name"] = name
	m["address"] = add

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
