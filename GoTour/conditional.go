package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
	//fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))

	fmt.Println()
	fmt.Println(posNegValue(2))
	fmt.Println(posNegValue(-2))
	fmt.Println(posNegValue(0))

	//The calls to the powXX functions will return their results
	//before the call to the fmt.Println statement begins
	fmt.Println()
	fmt.Println(
		pow01(3, 2, 10),
		pow01(3, 3, 20),
		pow01(4, 3, 100),
		pow01(4, 4, 100),
	)

	fmt.Println()
	fmt.Println(
		pow02(3, 2, 10),
		pow02(3, 3, 20),
	)

	fmt.Println()
	fmt.Println(
		pow03(3, 2, 10),
		pow03(3, 2, 20),
	)

	fmt.Println()
	fmt.Println(
		pow04(3, 2, 10),
		pow04(3, 3, 20),
		pow04(2, 3, 10),
		pow04(2, 4, 10),
	)
}

func pow04(x, n, lim float64) float64 {
	//The short statement outside and before the if statement,
	//to test if the short statement is visible inside any of the
	//else blocks, until the end of the if/statement..
	v := math.Pow(x, n)

	if v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim

}

func pow03(x, n, lim float64) float64 {
	//The if statement can start with a short statement to execute
	//before the if statement. E.g. [ v := math.Pow(x,n) ] below..
	//Variables declared in the short statement are only visible,
	//ie in the if statement scope,
	//and are visible inside any of the else blocks, until the end of
	//the if/statement..
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim

}

//The if condition checks for negative values being passed in,
//if the value is negative,
//the function will return the positive value of the negative value
//as a String with "i" appended to the end

//If the value passed in is positive, then the if condition fails,
//and the program moves on to the below return statement,
//which returns a string of the value of the square root of the value
//passed into the function

//Sprint method?? Does it convert values to a string version? And in
//this case giving us a string that can be returned??

func pow02(x, n, lim float64) float64 {
	//Short statement before if to test..
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	return lim

}

func pow01(x, n, lim float64) float64 {
	//The if statement can start with a short statement to execute
	//before the if statement. E.g. [ v := math.Pow(x,n) ] below..
	//Variables declared in the short statement are only visible,
	//ie in the scope, until the end of the if statement..

	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim

}

func posNegValue(x float64) string {
	if x > 0 {
		return "The value [" + fmt.Sprint(x) + "] is positive. "
	} else if x == 0 {
		return "The value [" + fmt.Sprint(x) + "] is zero. "
	} else {
		return "The value [" + fmt.Sprint(x) + "] is negative. "
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
