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
