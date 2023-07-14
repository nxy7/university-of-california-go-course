package main

import "fmt"

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64
	fmt.Println("Type in acceleration")
	fmt.Scan(&acceleration)
	fmt.Println("Type in initial velocity")
	fmt.Scan(&initialVelocity)
	fmt.Println("Type in initial displacement")
	fmt.Scan(&initialDisplacement)

	f := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println("Type in time value")
	fmt.Scan(&time)
	fmt.Println(f(time))
}

func GenDisplaceFn(acc float64, iVel float64, iDis float64) func(float64) float64 {
	return func(time float64) float64 {
		return (acc * 0.5 * (time * time)) + iVel*time + iDis
	}
}
