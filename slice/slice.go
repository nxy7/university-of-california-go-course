package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sl := make([]int, 3)
	for {
		var scanned string
		fmt.Println("Type integer")
		fmt.Scan(&scanned)

		if scanned == "X" {
			return
		}

		var i int
		i, err := strconv.Atoi(scanned)
		if err != nil {
			continue
		}

		sl = append(sl, i)
		sort.Ints(sl)

		fmt.Println(sl)
	}
}
