package main

// import (
// 	"fmt"
// 	"sort"
// )
//
// // instantiate a data type, a struct defined as below,
// // that will be the data type that our method is defined on
// type gameStonObj01 struct {
// 	xAxsLtr string
// 	yAxsLtr string
// 	xAxsInt int
// 	yAxsInt int
// }
//
// func main() {
// 	fmt.Println()
// 	fmt.Println("**** Methods 02:")
//
// 	// array of alphabet/int coordinates
// 	bordLttrIntCrdntArry01 := []string{"a", "b", "c", "d", "e", "f",
// 		"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r",
// 		"s", "t", "u", "v", "w", "x", "y", "z"}
//
// 	// gameRunSlceLen := len(bordLttrIntCrdntArry01) + 2
// 	gameRunSlceLen := 26
//
// 	/*// array/slice of boardstone01 types/objects,
// 	// ie a coordinates history,
// 	// with an arbitrary size to cycle the alphabet
// 	gameRunBlanksSlce01 := make([]gameStonObj01, 26)*/
//
// 	/*gameRunSlce01 := gameRun01(bordLttrIntCrdntArry01,
// 	gameRunBlanksSlce01)*/
// 	gameRunSlce01 := gameRun01(gameRunSlceLen, bordLttrIntCrdntArry01)
//
// 	// iterate to print the array for my test
// 	fmt.Println()
// 	printGameSlce(gameRunSlce01)
//
// }
//
// func gameRun01(gameRunSlceLenIN int, bordLttrIntCrdntArryIN []string,
// ) []gameStonObj01 {
//
// 	// array/slice of boardstone01 types/objects,
// 	// ie a coordinates history,
// 	// with an arbitrary size to cycle the alphabet
// 	gameRunSlce01 := make([]gameStonObj01, gameRunSlceLenIN)
//
// 	/*	// a copy of gameRunBlanksSlceIN, that will be returned
// 		var gameRunBlanksSlceCOPY []gameStonObj01
// 		gameRunBlanksSlceOUT := copy(gameRunBlanksSlceCOPY,
// 			gameRunBlanksSlceIN)*/
//
// 	for i := 0; i < len(gameRunSlce01); i++ {
// 		// randomly decided to create A and B values by taking first
// 		// and last index positions
// 		xAxsLtr01 := bordLttrIntCrdntArryIN[i]
// 		yAxsLtr01 := bordLttrIntCrdntArryIN[(len(
// 			bordLttrIntCrdntArryIN)-1)-i]
//
// 		// to assign the xAxsInt and yAxsInt int value board position to the
// 		// xAxsLtr01 and yAxsLtr index position values in their arrays,
// 		// which are the x y position values
// 		// sort.SearchStrings returns the index value in the given
// 		// slice/array of the given string
// 		xAxsInt01 := sort.SearchStrings(bordLttrIntCrdntArryIN,
// 			xAxsLtr01)
// 		yAxsInt01 := sort.SearchStrings(bordLttrIntCrdntArryIN,
// 			yAxsLtr01)
//
// 		gameRunSlce01[i] = gameStonObj01{
// 			xAxsLtr: xAxsLtr01,
// 			yAxsLtr: yAxsLtr01,
// 			xAxsInt: xAxsInt01,
// 			yAxsInt: yAxsInt01,
// 		}
// 	}
// 	return gameRunSlce01
//
// }
//
// func printGameSlce(gameRunSlceIN []gameStonObj01) {
// 	for i := 0; i < len(gameRunSlceIN); i++ {
// 		fmt.Printf("gameRunSlceIN[%v]: %v \n", i,
// 			gameRunSlceIN[i])
// 	}
// }
