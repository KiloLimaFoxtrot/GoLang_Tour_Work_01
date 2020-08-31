package main

import (
	"fmt"
)

// // Package point is a basic package for points.
// package point

// for testing..
func main() {
	fmt.Println()
	fmt.Println()

	// p01X := X(p01)
	// p01Y := Y(p01)

	p01 := Point{
		x: 24,
		y: 38,
	}

	p02 := New(p01.X(), p01.Y())

	fmt.Println()
	fmt.Println("p02: ", p02)

}

// Point is a basic point. Although simple, the member variables are kept
// private to ensure that Point remains immutable.
type Point struct {
	x int64
	y int64
}

// New creates a new immutable Point.
func New(x, y int64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// X returns the x-value.
func (p *Point) X() int64 {
	return p.x
}

// Y returns the y-value.
func (p *Point) Y() int64 {
	return p.y
}
