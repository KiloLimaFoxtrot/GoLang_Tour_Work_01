package main

import (
	"fmt"
)

/*
// ### ### START For testing...
*/

type Point struct {
	x int64
	y int64
}

func main() {
	fmt.Println()
	fmt.Println("pointTranslateINTSGF: ")

	p01 := Point{
		x: 22,
		y: 44,
	}

	p02 := New(p01.X(), p01.Y())

	fmt.Println()
	fmt.Println("p02: ", p02)
}

// Package point is a basic package for points.

// Point is a basic point. Although simple, the member variables are kept
// private to ensure that Point remains immutable.

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

/*
// ### ### END For testing...
*/

// Translation methods

// WIll use the GoGameSGFPointTranslate01 translate methods ..
