package main

import (
	"fmt"
)

type pointIntStrct01 struct {
	x, y int
}


func main() {
	
	fmt.Println()
	fmt.Println("Go Game Translate SGF Point 01")
	/*Signatures:
	1。 func translateToSGF(pt Point) string { .... }
	2。 func translateToPoint02(sgfPt string) Point {}
	*/
	fmt.Println()
	// Tester Method
	testMethods()
}

func translateToSGF01(ptIN pointIntStrct01) string {
	sgfPtX := string(rune((ptIN.x) + 97))
	sgfPtY := string(rune((ptIN.y) + 97))
	return sgfPtX + sgfPtY
}

func translateToPoint01(sgfPt string) pointIntStrct01 {
	sgfPtX := sgfPt[0]
	sgfPtY := sgfPt[2]
	return pointIntStrct01{
		x: int(sgfPtX) - 97,
		y: int(sgfPtY) - 97,
	}
}

/*
Tester Method
*/
func testMethods() {
	fmt.Println("Results: ")
	
	// Test of 1。 func translateToSGF(pt Point) string { .... }
	intPoint01 := pointIntStrct01{
		x: 16,
		y: 23,
	}
	
	intPoint02 := pointIntStrct01{
		x: 0,
		y: 1,
	}
	fmt.Println()
	fmt.Println("ttSGF01 intPoint01", translateToSGF01(intPoint01))
	fmt.Println("ttSGF01 intPoint01", translateToSGF01(intPoint02))
	
	// Test of 2。 func translateToPoint02(sgfPt string) Point {}
	sgfPoint01 := "a,b"
	sgfPoint02 := "q,x"
	fmt.Println()
	fmt.Println("ttp01 sgfPoint01: ", translateToPoint01(sgfPoint01))
	fmt.Println("ttp01 sgfPoint02: ", translateToPoint01(sgfPoint02))
	
}
