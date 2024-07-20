package main

import (
	"fmt"
	"math"
)

func MethodsAndInterfaces() {
	fmt.Println("")
	fmt.Println("Methods and Interfaces:")

	somewhere := Vertex{3, 4}

	fmt.Println(somewhere.Abs())
	fmt.Println(Abs(somewhere))

	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

	stringerInterface()
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func stringerInterface() {
	a := Person{"ikuwow", 100}
	fmt.Println(a)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
