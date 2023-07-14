package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	arr := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write sequence of integers bigger than 4 elements")
	l, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	l = strings.TrimSpace(l)

	lSplit := strings.Split(l, " ")
	for _, s := range lSplit {
		asInt, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		arr = append(arr, asInt)
	}

	arrLen := len(arr)
	waitGrp := sync.WaitGroup{}
	var p1, p2, p3, p4 []int
	waitGrp.Add(4)
	go func() {
		defer waitGrp.Done()
		p1 = Sorted(arr[:arrLen/4])
	}()
	go func() {
		defer waitGrp.Done()
		p2 = Sorted(arr[arrLen/4 : arrLen/2])
	}()
	go func() {
		defer waitGrp.Done()
		p3 = Sorted(arr[arrLen/2 : arrLen/4*3])
	}()
	go func() {
		defer waitGrp.Done()
		p4 = Sorted(arr[arrLen/4*3:])
	}()

	waitGrp.Wait()
	merged := make([]int, 0)
	for _, slice := range [][]int{p1, p2, p3, p4} {
		for _, i := range slice {
			merged = append(merged, i)
		}
	}

	merged = Sorted(merged)

}

func Sorted(arr []int) []int {
	sort.Ints(arr)
	fmt.Println(arr)
	return arr
}
