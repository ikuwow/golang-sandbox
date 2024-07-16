package main

import (
	"fmt"
	"math"
	"runtime"
)

func Memo() {
	forLoop()
	while()
	shortIf()

	fmt.Printf("OS name: %s\n", getOSName())

	useStruct()

	arrayAndSlice()

	runFibonacci()
}

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func while() {
	sum := 0
	// for ; sum < 10; {
	for sum < 10 {
		sum += 1
	}
	fmt.Println(sum)
}

func shortIf() {
	const limit = 1000
	var x, n float64 = 3, 5 // change here

	if v := math.Pow(x, n); v < limit {
		fmt.Println(v)
	} else {
		fmt.Println("Warn: Limit exceeded!")
	}
	fmt.Println(limit)
}

func CompareSqrt() {
	const value int = 2

	fmt.Printf("Guessed Sqrt(%d): %f\n", value, sqrt(float64(value)))
	fmt.Printf("math.Sqrt(%d): %f\n", value, math.Sqrt(float64(value)))
}

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func getOSName() (os_name string) {

	switch os := runtime.GOOS; os {
	case "darwin":
		os_name = "macOS"
	case "linux":
		os_name = "Linux"
	default:
		os_name = os
	}

	return
}

func useStruct() {
	type Vertex struct {
		X int
		Y int
	}

	v1 := Vertex{1, 2}
	v2 := Vertex{X: 3, Y: 4}

	p1 := &v1

	p1.X = 333

	fmt.Println(v1.X, v2.X)
}

func arrayAndSlice() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var previous int
	var current int
	return func() int {

		if current == 0 {
			previous = 0
			current = 1
			return 0
		}

		next := previous + current

		previous = current
		current = next

		return next
	}
}

func runFibonacci() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
