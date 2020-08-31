package main

import (
	"fmt"
)

type gamePointStrct01 struct {
	x, y int
}

type gamePointStrct02 struct {
	x string
	y int
}

type sgfPointStruct struct {
	// y-axis, x-axis
	xColVal, yRowVal string
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
	testMethods02()
}

/*
### Point to SGF
*/
func transPointToSGF01(ptIN gamePointStrct01) string {
	sgfPtX := string(rune((ptIN.x) + 97))
	sgfPtY := string(rune((ptIN.y) + 97))
	return sgfPtX + sgfPtY
}

func transPointToSGF02(ptIN gamePointStrct02) string {
	sgfPtY := string(rune((ptIN.y) + 97))
	// SGF format of column x string then row y string
	return ptIN.x + sgfPtY
}

/*
### SGF to Point
*/
func transSGFToPoint01(sgfPt string) gamePointStrct01 {
	sgfPtX := sgfPt[0]
	sgfPtY := sgfPt[2]
	return gamePointStrct01{
		x: int(sgfPtX) - 97,
		y: int(sgfPtY) - 97,
	}
}

// Mixed string int x
func transSGFToPoint02(sgfPt string) gamePointStrct02 {
	sgfPtX := string(sgfPt[0])
	sgfPtY := sgfPt[1]
	return gamePointStrct02{
		x: sgfPtX,
		y: int(sgfPtY) - 97,
	}
}

/*
Tester Method
*/
func testMethods02() {
	fmt.Println("Results: ")
	
	// **Test of mixed int string point structs
	gamePoint01 := gamePointStrct02{
		x: "q",
		y: 23,
	}
	fmt.Println()
	fmt.Println("trnsPntToSGF02 gamePoint01",
		transPointToSGF02(gamePoint01))
	
	sgfPoint03 := "qx"
	fmt.Println()
	fmt.Println("trnsSGFToPoint02 sgfPoint01",
		transSGFToPoint02(sgfPoint03))
	
	// **Test of non-mixed int string point structs
	// Test of 1。 func translateToSGF(pt Point) string { .... }
	intPoint01 := gamePointStrct01{
		x: 16,
		y: 23,
	}
	
	intPoint02 := gamePointStrct01{
		x: 0,
		y: 1,
	}
	fmt.Println()
	fmt.Println("ttSGF01 intPoint01", transPointToSGF01(intPoint01))
	fmt.Println("ttSGF01 intPoint01", transPointToSGF01(intPoint02))
	
	// Test of 2。 func translateToPoint02(sgfPt string) Point {}
	sgfPoint01 := "a,b"
	sgfPoint02 := "q,x"
	fmt.Println()
	fmt.Println("ttp01 sgfPoint01: ", transSGFToPoint01(sgfPoint01))
	fmt.Println("ttp01 sgfPoint02: ", transSGFToPoint01(sgfPoint02))
	
}
