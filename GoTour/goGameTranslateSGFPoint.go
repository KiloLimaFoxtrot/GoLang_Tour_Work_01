package main

import (
	"fmt"
)

type intPointStrct02 struct {
	x, y int
}

type sgfPointStrct02 struct {
	x string
	y int
}

// type sgfPointStruct struct {
// 	// y-axis, x-axis
// 	xColVal, yRowVal string
// }

func main() {

	fmt.Println()
	fmt.Println("Go Game Translate SGF Point99 01")
	/*Signatures:
	1。 func translateToSGF(pt Point99) string { .... }
	2。 func translateToPoint02(sgfPt string) Point99 {}
	*/
	fmt.Println()
	// Tester Method
	testMethods02()
}

/*
### Point99 to SGF
*/
func transPointToSGF01(ptIN intPointStrct02) string {
	sgfPtX := string(rune((ptIN.x) + 97))
	sgfPtY := string(rune((ptIN.y) + 97))
	return sgfPtX + sgfPtY
}

func transPointToSGF02(ptIN sgfPointStrct02) string {
	sgfPtY := string(rune((ptIN.y) + 97))
	// SGF format of column x string then row y string
	return ptIN.x + sgfPtY
}

/*
### SGF to Point99
*/
func transSGFToPoint01(sgfPt string) intPointStrct02 {
	sgfPtX := sgfPt[0]
	sgfPtY := sgfPt[2]
	return intPointStrct02{
		x: int(sgfPtX) - 97,
		y: int(sgfPtY) - 97,
	}
}

// Mixed string int x
func transSGFToPoint02(sgfPt string) sgfPointStrct02 {
	sgfPtX := string(sgfPt[0])
	sgfPtY := sgfPt[1]
	return sgfPointStrct02{
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
	sgfPoint01 := sgfPointStrct02{
		x: "q",
		y: 23,
	}
	fmt.Println()
	fmt.Println("trnsPntToSGF02 sgfPoint01",
		transPointToSGF02(sgfPoint01))

	sgfPoint02 := "qx"
	fmt.Println()
	fmt.Println("trnsSGFToPoint02 sgfPoint03",
		transSGFToPoint02(sgfPoint02))

	// **Test of non-mixed int string point structs
	// Test of 1。 func translateToSGF(pt Point99) string { .... }
	intPoint01 := intPointStrct02{
		x: 16,
		y: 23,
	}

	intPoint02 := intPointStrct02{
		x: 0,
		y: 1,
	}
	fmt.Println()
	fmt.Println("ttSGF01 intPoint01", transPointToSGF01(intPoint01))
	fmt.Println("ttSGF01 intPoint01", transPointToSGF01(intPoint02))

	// Test of 2。 func translateToPoint02(sgfPt string) Point99 {}
	sgfPoint03 := "a,b"
	sgfPoint04 := "q,x"
	fmt.Println()
	fmt.Println("ttp01 sgfPoint03: ", transSGFToPoint01(sgfPoint03))
	fmt.Println("ttp01 sgfPoint04: ", transSGFToPoint01(sgfPoint04))

}
