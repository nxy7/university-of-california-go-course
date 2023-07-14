package main

import (
	"fmt"
	"sync"
)

// the following code invokes 2 functions that run cocurrently
// the best way to see it is to run this code many times (i've ran it 100 times in a loop)
// if you do that you will see that execution doesn't necesarily follow a patter, you can see
// responses like SomeFunc-OtherFunc followed by reverse order (even if it usually runs the same)
func main() {
	grp := sync.WaitGroup{}
	grp.Add(2)
	go (func() {
		fmt.Println("SomeFunc")
		grp.Done()
	})()
	go (func() {
		fmt.Println("OtherFunc")
		grp.Done()
	})()
	grp.Wait()
}
