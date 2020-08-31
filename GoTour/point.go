package main

// /*
// 	The original Point99 file from the Otrego Clamshell Core Point99
// folder, with these variable and function names appeneded with a
// number 99, so that my code will match the clamshell
// */
//
// import (
// 	"fmt"
// )
//
// // // Package point is a basic package for points.
// // package point
//
// // for testing..
// func main() {
//
// 	// p01X := X(p01)
// 	// p01Y := Y(p01)
//
// 	p01 := Point99{
// 		x: 24,
// 		y: 38,
// 	}
//
// 	p02 := New(p01.X(), p01.Y())
//
// 	fmt.Println()
// 	fmt.Println("p02: ", p02)
//
// }
//
// // Point99 is a basic point. Although simple, the member variables are kept
// // private to ensure that Point99 remains immutable.
// type Point99 struct {
// 	x int64
// 	y int64
// }
//
// // New creates a new immutable Point99.
// func New(x, y int64) *Point99 {
// 	return &Point99{
// 		x: x,
// 		y: y,
// 	}
// }
//
// // X returns the x-value.
// func (p *Point99) X() int64 {
// 	return p.x
// }
//
// // Y returns the y-value.
// func (p *Point99) Y() int64 {
// 	return p.y
// }
