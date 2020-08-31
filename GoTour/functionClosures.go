package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()
	fmt.Println("****Function Closures 01:")
	
	// for the GoTour example
	/*posNum := adder01()
	negNum := adder01()*/
	
	// initialize the struct,
	// note: remember a struct is like a point object in java that
	// can hold two data points, but a struct can hold many data points
	/*type cmputStrct01 struct {
		x int
		y int
	}*/
	
	// my test
	type cmputStrct02 struct {
		xAxsLtr01 string
		yAxsLtr01 string
		xAxsInt01 int
		yAxsInt01 int
	}
	
	// array of alphabet
	abcArry := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z"}
	
	// create an array/slice? of the type of the struct initialized
	// above using structs with arrays is another way to create a multi
	// -dimensional array,
	// i.e. each index position of a 1D array holds a struct of as
	// many value points as a struct can hold
	/*cmptStrctArry01 := make([]cmputStrct01, 10)*/
	
	// my test,this is length 26,
	// to make room for the below creation of A B x y position values
	// going in reverse directions along the abcArry..
	cmptStrctArry02 := make([]cmputStrct02, 26)
	
	// iterate and fill the array with instantiated instances of the
	// struct cmputStrct01
	/*for i := 0; i < 10; i++ {
		x := (i)
		y := (-2 * i)
		cmptStrctArry01[i] = cmputStrct01{posNum(x), negNum(y)}
	
	}*/
	
	// my test iterate and fill the array with instantiated instances
	// of the struct cmputStrct02
	for i := 0; i < len(cmptStrctArry02); i++ {
		// randomly decided to create A and B values by taking first
		// and last index positions
		posA01 := abcArry[i]
		posB01 := abcArry[(len(abcArry)-1)-i]
		
		// to assign the xAxsInt and yAxsInt int value board position to the
		// xAxsLtr01 and yAxsLtr index position values in their arrays,
		// which are the x y position values
		// sort.SearchStrings returns the index value in the given
		// slice/array of the given string
		posX01 := sort.SearchStrings(abcArry, posA01)
		posY01 := sort.SearchStrings(abcArry, posB01)
		
		// fill the array with the struct instances for each
		// iteration of this for loop
		cmptStrctArry02[i] = cmputStrct02{
			posA01, posB01, posX01, posY01,
		}
		
	}
	
	// iterate to print the array
	/*fmt.Println()
	for i := 0; i < len(cmptStrctArry01); i++ {
		fmt.Printf("cmptStrctArry01[%v]: %v \n", i, cmptStrctArry01[i])
	}*/
	
	// iterate to print the array for my test
	fmt.Println()
	for i := 0; i < len(cmptStrctArry02); i++ {
		fmt.Printf("cmptStrctArry02[%v]: %v \n", i, cmptStrctArry02[i])
	}
	
}

func adder01() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
	
}
