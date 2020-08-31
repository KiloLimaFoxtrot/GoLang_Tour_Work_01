package main

import (
	"fmt"
)

func main() {
	
	fmt.Println()
	fmt.Println("Go Game Translate SGF Point Notes 01")
	
	// Josh notes from Discord:
	/*
		1.
		FWIW, I'm looking for something a bit simpler.
		Essentially a method with the signature:
	
		Signatures:
		1. func translateToSGF(pt Point) string { .... }
		2. func translateToPoint02(sgfPt string) Point { .... }
	
		2.
		Another recommendation: You're testing via running a main
		binary -- often the right approach!
		Another way to go is to create tests
		https://golang.org/pkg/testing/
	
		There's one trick that'll make it vastly more simple:
		https://www.socketloop.
		com/tutorials/golang-how-to-convert-character-to-ascii-and
		-back#:~:text=Solution%3A,Here%20you%20go!
	
		basically you can just take a char a and convert it to its
		integer value 97 Then you don't need to store the index w/ an
		array
	
		https://play.golang.org/p/AQ0szjJBk26
	*/
	
	// fmt.Println("Hello, playground")
	// ? So to get the board position int value,
	// we can use the char's ascii value - 97 to get a corresponding
	// board x / y axis int value, see below
	// From the website
	// a := 'a'
	// aval := int(a)
	// fmt.Println(aval - 97)
	//
	// c := 'c'
	// cval := int(c)
	// fmt.Println(cval - 97)
	//
	// // This for loop use the char's ascii value - 97 to get a
	// // corresponding board x / y axis int value
	// fmt.Println()
	// fmt.Println("For each for loop")
	// for _, char := range "abcdefghilkmnopqrstuvwxyz" {
	// 	fmt.Printf("%q: %d\n", char, int(char)-97)
	// }
	//
	// fmt.Println()
	// fmt.Println("Trad for loop: ")
	// alphabet := "abcdefghilkmnopqrstuvwxyz"
	// for i := 0; i < len(alphabet); i++ {
	// 	fmt.Printf("%q: %d\n", alphabet[i], int(alphabet[i])-97)
	// }
	
	/*
	
		I was saying that you can use the ascii value to get you most
		of the way there. You might need to do some translation based
		on where the x / y axis are,
		but I think you only need to pass in the board size for that
		so you'd have the signatures
	
		Signatures:
		1。 func translateToSGF(pt Point, int boardSize) string { .... }
		2。 func translateToPoint02(sgfPt string, int boardSize) Point {}
	
		Or maybe not.
		I  hink SGF points might be in the correct orientation
		-- from the top-left
	
		Here's the sort of method I was thinking of,
		but in JavaScript:
		https://github.com/Kashomon/glift-core/blob/master/src/util
		/point.js#L167
	
		* Return a string representation of the coordinate.  I.e., "12,3".
		* @return {!glift.PtStr}
	
		// JavaScript
		// toString: function() {
		// return glift.util.coordToString(this.x(), this.y());
		// }
	
		* @param {number} x
		* @param {number} y
		* @return {!glift.Point} a new point that's a translation from this one.
	
		// JavaScript, example function
		// translate: function(x, y) {
		// 	return glift.util.point(this.x() + x, this.y() + y);
		// }
	
		Core logic and rules for Glift, for handling SGFs and such.
		https://discordapp.com/channels/@me/705599069300719626/746816145985896519
	
		so, looks like you don't need to pass in the boardSize
	
		Signatures:
		1。 func translateToSGF(pt Point) string { .... }
		2。 func translateToPoint02(sgfPt string) Point {}
	
		Should be sufficient then
	
	
		3。
		func translateToSGF(pt Point) string { .... }
		func translateToPoint02(sgfPt string) Point {}
	*/
}
