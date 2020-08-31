package main

// ** This work is continued in the file translateIntLtrCrdnts01.go

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
// type gameStonLtrCrdnt struct {
// 	xAxsLtr string
// 	yAxsLtr string
// }
//
// type gameStonIntCrdnt struct {
// 	xAxsInt int
// 	yAxsInt int
// }
//
// func main() {
// 	fmt.Println()
// 	fmt.Println("**** Methods03:")
//
// 	// array of alphabet/int coordinates
// 	bordLttrIntCrdntArry01 := []string{"a", "b", "c", "d", "e", "f",
// 		"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r",
// 		"s", "t", "u", "v", "w", "x", "y", "z"}
//
// 	/*
// 		**** Arbitrary game run to build material to work with,
// 		more detail above the gameRun01 function
// 	*/
// 	// gameRunSlceLen := len(bordLttrIntCrdntArry01) + 2 // ??
// 	gameRunSlceLen := 26
//
// 	gameRunSlce01 := gameRun01(gameRunSlceLen, bordLttrIntCrdntArry01)
//
// 	// iterate to print the array for my test
// 	fmt.Println()
// 	printGameSlce(gameRunSlce01)
//
// 	/*
// 	**** Getters
// 	 */
// 	// Get game run letter coordinates
// 	gameLtrCrdntsSlce := getLtrCrdnts01(gameRunSlceLen, gameRunSlce01)
//
// 	// iterate to print the array for my test
// 	fmt.Println()
// 	printGameLtrCrdnt(gameLtrCrdntsSlce)
//
// 	// Get game run int coordinates
// 	gameIntCrdntsSlce := getIntCrdnts01(gameRunSlceLen, gameRunSlce01)
//
// 	// iterate to print the array for my test
// 	fmt.Println()
// 	printGameIntCrdnt(gameIntCrdntsSlce)
//
// 	/*
// 	**** Translators
// 	 */
// 	// Translate game run letter coordinates to int coordinates
// 	gameLtrCrdntsTrnsSlce := trnsLtrCrdnts01(gameRunSlceLen,
// 		gameLtrCrdntsSlce, bordLttrIntCrdntArry01)
//
// 	// iterate to print the array for my test
// 	fmt.Println()
// 	printGameIntCrdnt(gameLtrCrdntsTrnsSlce)
//
// 	// Translate game run int coordinates to letter coordinates
// 	gameIntCrdntsTrnsSlce := trnsIntCrdnts01(gameRunSlceLen,
// 		gameIntCrdntsSlce, bordLttrIntCrdntArry01)
//
// 	// iterate to print the array for my test
// 	fmt.Println()
// 	printGameLtrCrdnt(gameIntCrdntsTrnsSlce)
//
// }
//
// /*
// **** Translators
//  */
//
// func trnsIntCrdnts01(gameRunSlceLenIN int,
// 	gameIntCrdntsSlceIN []gameStonIntCrdnt,
// 	bordLttrIntCrdntArryIN []string) []gameStonLtrCrdnt {
//
// 	// array/slice of
// 	gameLtrCrdntsSlceOut := make([]gameStonLtrCrdnt, gameRunSlceLenIN)
//
// 	for i := 0; i < len(gameLtrCrdntsSlceOut); i++ {
//
// 		gameStonObj := gameIntCrdntsSlceIN[i]
// 		xAxsInt01 := gameStonObj.xAxsInt
// 		yAxsInt01 := gameStonObj.yAxsInt
//
// 		xAxsLtr01 := bordLttrIntCrdntArryIN[xAxsInt01]
// 		yAxsLtr01 := bordLttrIntCrdntArryIN[yAxsInt01]
//
// 		// create a temp gameStonIntCrdnt to fill the slice/array of
// 		// gameStonIntCrdnt objects..
// 		gameLtrCrdntsSlceOut[i] = gameStonLtrCrdnt{
// 			xAxsLtr: xAxsLtr01,
// 			yAxsLtr: yAxsLtr01,
// 		}
// 	}
// 	return gameLtrCrdntsSlceOut
//
// }
//
// func trnsLtrCrdnts01(gameRunSlceLenIN int,
// 	gameLtrCrdntsSlceIN []gameStonLtrCrdnt,
// 	bordLttrIntCrdntArryIN []string) []gameStonIntCrdnt {
//
// 	// array/slice of
// 	gameIntCrdntsSlceOut := make([]gameStonIntCrdnt, gameRunSlceLenIN)
//
// 	for i := 0; i < len(gameIntCrdntsSlceOut); i++ {
//
// 		gameStonObj := gameLtrCrdntsSlceIN[i]
// 		xAxsLtr01 := gameStonObj.xAxsLtr
// 		yAxsLtr01 := gameStonObj.yAxsLtr
//
// 		xAxsInt01 := sort.SearchStrings(bordLttrIntCrdntArryIN,
// 			xAxsLtr01)
// 		yAxsInt01 := sort.SearchStrings(bordLttrIntCrdntArryIN,
// 			yAxsLtr01)
//
// 		// create a temp gameStonIntCrdnt to fill the slice/array of
// 		// gameStonIntCrdnt objects..
// 		gameIntCrdntsSlceOut[i] = gameStonIntCrdnt{
// 			xAxsInt: xAxsInt01,
// 			yAxsInt: yAxsInt01,
// 		}
// 	}
// 	return gameIntCrdntsSlceOut
//
// }
//
// /*
// **** Getters
//  */
//
// /*
// Get the Stone struct objects' letter coordinates
// */
// func getIntCrdnts01(gameRunSlceLenIN int,
// 	gameRunSlceIn []gameStonObj01) []gameStonIntCrdnt {
//
// 	// array/slice of
// 	gameIntCrdntsSlceOut := make([]gameStonIntCrdnt, gameRunSlceLenIN)
//
// 	for i := 0; i < len(gameIntCrdntsSlceOut); i++ {
// 		// create a temp gameStonObj,
// 		// to extract that struct's x and y access letters
// 		gameStonObj := gameRunSlceIn[i]
// 		xAxsInt01 := gameStonObj.xAxsInt
// 		yAxsInt01 := gameStonObj.yAxsInt
//
// 		// create a temp gameStonIntCrdnt to fill the slice/array of
// 		// gameStonIntCrdnt objects..
// 		gameIntCrdntsSlceOut[i] = gameStonIntCrdnt{
// 			xAxsInt: xAxsInt01,
// 			yAxsInt: yAxsInt01,
// 		}
// 	}
// 	return gameIntCrdntsSlceOut
//
// }
//
// /*
// Get the Stone struct objects' letter coordinates
// */
// func getLtrCrdnts01(gameRunSlceLenIN int,
// 	gameRunSlceIn []gameStonObj01) []gameStonLtrCrdnt {
//
// 	// array/slice of
// 	gameLtrCrdntsSlceOut := make([]gameStonLtrCrdnt, gameRunSlceLenIN)
//
// 	for i := 0; i < len(gameLtrCrdntsSlceOut); i++ {
// 		// create a temp gameStonObj,
// 		// to extract that struct's x and y access letters
// 		gameStonObj := gameRunSlceIn[i]
// 		xAxsLtr01 := gameStonObj.xAxsLtr
// 		yAxsLtr01 := gameStonObj.yAxsLtr
//
// 		// create a temp gameStonLtrCrdnt to fill the slice/array of
// 		// gameStonLtrCrdnt objects..
// 		gameLtrCrdntsSlceOut[i] = gameStonLtrCrdnt{
// 			xAxsLtr: xAxsLtr01,
// 			yAxsLtr: yAxsLtr01,
// 		}
// 	}
// 	return gameLtrCrdntsSlceOut
//
// }
//
// /*
// **** Arbitrary game run to fill a slice of board stone structs
// /objects that have both letter and int coordinates that correlate
// with each other by slice index position
// */
//
// func gameRun01(gameRunSlceLenIN int, bordLttrIntCrdntArryIN []string,
// ) []gameStonObj01 {
//
// 	// array/slice of boardstone01 types/objects,
// 	// ie a coordinates history,
// 	// with an arbitrary size to cycle the alphabet
// 	gameRunSlce01 := make([]gameStonObj01, gameRunSlceLenIN)
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
// func printGameIntCrdnt(gameIntCrdntsSlceIN []gameStonIntCrdnt) {
// 	for i := 0; i < len(gameIntCrdntsSlceIN); i++ {
// 		fmt.Printf("INT val gameRunSlceIN[%v]:  %v \n", i,
// 			gameIntCrdntsSlceIN[i])
// 	}
// }
//
// func printGameLtrCrdnt(gameLtrCrdntsSlceIN []gameStonLtrCrdnt) {
// 	for i := 0; i < len(gameLtrCrdntsSlceIN); i++ {
// 		fmt.Printf("LTR val gameRunSlceIN[%v]:  %v \n", i,
// 			gameLtrCrdntsSlceIN[i])
// 	}
// }
//
// func printGameSlce(gameRunSlceIN []gameStonObj01) {
// 	for i := 0; i < len(gameRunSlceIN); i++ {
// 		fmt.Printf("ALL val gameRunSlceIN[%v]:  %v \n", i,
// 			gameRunSlceIN[i])
// 	}
// }
