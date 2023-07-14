package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Write sequence of integers split by space")

	arr := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)
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

	if len(arr) > 10 {
		panic("This program only accepts up to 10 elements")
	}

	BubbleSort(arr)
	fmt.Println(arr)
}

func BubbleSort(nums []int) {
	for i := range nums {
		for i2 := range nums[:len(nums)-i-1] {
			if nums[i2] > nums[i2+1] {
				Swap(nums, i2)
			}
		}
	}
}

func Swap(arr []int, i1 int) {
	arr[i1], arr[i1+1] = arr[i1+1], arr[i1]
}
