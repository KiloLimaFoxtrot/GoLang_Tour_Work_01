package main

import (
	"fmt"
	"math"
)

type VrtxStrct01 struct {
	X, Y float64
}

func main() {
	fmt.Println()
	fmt.Println("**** Methods 01:")
	
	v1 := VrtxStrct01{
		X: 3,
		Y: 4,
	}
	
	v1AbsVal := v1.Abs()
	
	fmt.Println()
	fmt.Println("v1AbsVal: ", v1AbsVal)
	
}

func (v1 VrtxStrct01) Abs() float64 {
	absVal := math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
	return absVal
}
