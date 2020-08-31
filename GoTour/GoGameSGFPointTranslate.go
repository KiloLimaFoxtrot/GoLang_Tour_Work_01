package main

import (
	"fmt"
)

// Game point format is x-axis string, y-axis int
type pointStrct01 struct {
	x string
	y int
}

// Or, if the game point format is x-axis int, y-axis int
type pointStrct02 struct {
	x, y int
}

// SGF Format is two lowercase letters
// type sgfPointStruct02 struct {
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
	methodTest01()
}

/*
### Point99 to SGF Method
*/
// If game point format is x-axis string, y-axis int
func pointToSGF01(ptIN pointStrct01) string {
	sgfPtY := string(rune((ptIN.y) + 97))
	// SGF format of column x string then row y string
	return ptIN.x + sgfPtY
}

// If game point format is x-axis int, y-axis int
func pointToSGF02(ptIN pointStrct02) string {
	sgfPtX := string(rune((ptIN.x) + 97))
	sgfPtY := string(rune((ptIN.y) + 97))
	return sgfPtX + sgfPtY
}

/*
### SGF to Point99
*/
// If game point format is x-axis string, y-axis int
func sGFToPoint01(sgfPt string) pointStrct01 {
	sgfPtX := string(sgfPt[0])
	sgfPtY := sgfPt[1]
	return pointStrct01{
		x: sgfPtX,
		y: int(sgfPtY) - 97,
	}
}

// If game point format is x-axis int, y-axis int
func sGFToPoint02(sgfPt string) pointStrct02 {
	sgfPtX := sgfPt[0]
	sgfPtY := sgfPt[1]
	return pointStrct02{
		x: int(sgfPtX) - 97,
		y: int(sgfPtY) - 97,
	}
}

/*
### Tester Method
*/
func methodTest01() {
	fmt.Println("Results: ")

	// Test if game point format is x-axis string, y-axis int
	gamePoint01 := pointStrct01{
		x: "q",
		y: 23,
	}
	fmt.Println()
	fmt.Printf("trnsPntToSGF02 gamePoint01: %v, to SGF: %v\n",
		gamePoint01, pointToSGF01(gamePoint01))

	// Test if game point format is x-axis int, y-axis int
	gamePoint02 := pointStrct02{
		x: 16,
		y: 23,
	}
	fmt.Println()
	fmt.Printf("trnsPntToSGF04 gamePoint02: %v, to SGF: %v\n",
		gamePoint02, pointToSGF02(gamePoint02))

	// Test if game point format is x-axis string, y-axis int
	sgfPoint03 := "qx"
	fmt.Println()
	fmt.Printf("trnsSGFToPoint03 sgfPoint03: %v, to Point99: %v\n",
		sgfPoint03, sGFToPoint01(sgfPoint03))

	// Test if game point format is x-axis int, y-axis int
	fmt.Println()
	fmt.Printf("trnsSGFToPoint04 sgfPoint03: %v, to Point99: %v\n",
		sgfPoint03, sGFToPoint02(sgfPoint03))

	fmt.Println()
	fmt.Println("Otrego Go! ")

	fmt.Println()
	fmt.Println("GitHub Go! ")

	fmt.Println()
	fmt.Println("IntelliJ Go! ")

}
