package main

import "fmt"

// Constants are declared like variables,
//but with the const keyword.
const Pi = 3.14
const Truth = true

func main() {
	fmt.Println()
	fmt.Println("Constants: ")
	// Constants legal data types are character, string, boolean,
	//or numeric value.
	const World = "世界"
	fmt.Println("Hello ", World)
	//fmt.Println("Happy ", Pi, "Day ")
	fmt.Printf("Happy %v Day \n", Pi)
	fmt.Println("Go rules? ", Truth)
	fmt.Println()

	fmt.Println("Constants: ")
	fmt.Printf("Variable Pi: %T\n", Pi)
	fmt.Printf("Variable Truth: %T\n", Truth)
	fmt.Println()

	fmt.Println("Non-Constants: ")
	fmt.Printf("Variable World: %T\n", World)
	fmt.Println()

	//Constants cannot be declared using the := syntax

}
