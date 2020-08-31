package main

// import (
// 	"fmt"
// )
//
// type pointStrct02 struct {
// 	x, y int
// }
//
// func main() {
//
// 	fmt.Println()
// 	fmt.Println("Go Game Translate SGF Point 01")
//
// 	a := 'a'
// 	aval := int(a)
// 	fmt.Println(aval - 97)
//
// 	c := 'c'
// 	cval := int(c)
// 	fmt.Println(cval - 97)
//
// 	// This for loop use the char's ascii value - 97 to get a
// 	// corresponding board x / y axis int value
// 	fmt.Println()
// 	fmt.Println("For each for loop")
// 	for _, char := range "abcdefghilkmnopqrstuvwxyz" {
// 		fmt.Printf("%q: %d\n", char, int(char)-97)
// 	}
//
// 	fmt.Println()
// 	fmt.Println("Trad for loop: ")
// 	alphabet := "abcdefghilkmnopqrstuvwxyz"
// 	for i := 0; i < len(alphabet); i++ {
// 		fmt.Printf("%q: %d\n", alphabet[i], int(alphabet[i])-97)
// 	}
//
// 	/*
// 		so, looks like you don't need to pass in the boardSize
//
// 		Signatures:
// 		1。 func translateToSGF(pt Point) string { .... }
// 		2。 func translateToPoint02(sgfPt string) Point {}
//
// 		Should be sufficient then
// 	*/
//
// 	// Test of 1。 func translateToSGF(pt Point) string { .... }
// 	intPoint01 := pointStrct02{
// 		x: 16,
// 		y: 23,
// 	}
//
// 	intPoint02 := pointStrct02{
// 		x: 0,
// 		y: 1,
// 	}
// 	fmt.Println()
// 	fmt.Println("ttSGF01 intPoint01", pointToSGF02(intPoint01))
// 	fmt.Println()
// 	fmt.Println("ttSGF01 intPoint01", pointToSGF02(intPoint02))
//
// 	// Test of 2。 func translateToPoint02(sgfPt string) Point {}
// 	sgfPoint01 := "ab"
// 	sgfPoint02 := "qx"
// 	fmt.Println()
// 	fmt.Println("ttp01 sgfPoint01: ", sGFToPoint02(sgfPoint01))
// 	fmt.Println()
// 	fmt.Println("ttp01 sgfPoint02: ", sGFToPoint02(sgfPoint02))
//
// 	// To test the alternate of above that accepts an sgf string point
// 	// with a comma between the characters
// 	// sgfPoint03 := "a,b"
// 	// sgfPoint04 := "q,x"
// 	// fmt.Println()
// 	// fmt.Println("ttp02 sgfPoint01: ", translateToPoint02(sgfPoint03))
// 	// fmt.Println()
// 	// fmt.Println("ttp01 sgfPoint02: ", translateToPoint02(sgfPoint04))
//
// }
// func pointToSGF02(ptIN pointStrct02) string {
// 	sgfPtX := string(rune((ptIN.x) + 97))
// 	sgfPtY := string(rune((ptIN.y) + 97))
// 	return sgfPtX + sgfPtY
// }
//
// func sGFToPoint02(sgfPt string) pointStrct02 {
// 	sgfPtX := sgfPt[0]
// 	sgfPtY := sgfPt[1]
// 	return pointStrct02{
// 		x: int(sgfPtX) - 97,
// 		y: int(sgfPtY) - 97,
// 	}
// }
//
// // Alternate of above that accepts an sgf string point with a comma
// // between the characters
// // func translateToPoint02(sgfPt string) pointStrct02 {
// // 	sgfPtX := sgfPt[0]
// // 	sgfPtY := sgfPt[2]
// // 	return pointStrct02{
// // 		x: int(sgfPtX) - 97,
// // 		y: int(sgfPtY) - 97,
// // 	}
// // }
//
// // ??to split the string on the comma??
// // func translateToPoint02(sgfPt string) pointStrct02 {
// // 	sgfPtSplit := strings.Split(sgfPt, ",")
// // 	return pointStrct02{
// // 		x: int(sgfPtSplit[0]) - 97,
// // 		y: int(sgfPtY) - 97,
// // 	}
// // }
