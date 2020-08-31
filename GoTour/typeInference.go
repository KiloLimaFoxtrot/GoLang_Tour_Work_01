package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println("Data Type Inference: ")
	fmt.Println()

	//When declaring a variable without specifying an explicit type,
	//(either by using the := syntax or var = expression syntax),
	//the variable's data type is inferred from the right hand side.
	var h = 42        // h is an int
	i := 3.142        // i is a float64
	j := 0.867 + 0.5i // complex128

	fmt.Println("Variable declaration without specifying the explicit" +
		" type: \n" +
		"- The variable's type is inferred from the right side" +
		" value of the equation... ")
	fmt.Println()

	fmt.Println("E.g. " +
		"\nvar h = 42        // h is an int" +
		"\ni := 3.142        // i is a float64" +
		"\nj := 0.867 + 0.5i // complex128")
	fmt.Println()

	fmt.Printf("Variable: h \n"+
		"Type: %T \n"+
		"Value: %v\n", h, h)
	fmt.Println()

	fmt.Printf("Variable: i \n"+
		"Type: %T\n"+
		"Value: %v\n", i, i)
	fmt.Println()

	fmt.Printf("Variable: j \n"+
		"Type: %T\n"+
		"Value: %v\n", j, j)
	fmt.Println()

	//When the right hand side of the declaration is typed,
	var k int // k is declared an int
	//the new variable is of that same type:
	m := k // m is assigned to the value of K,
	// hence K is consequently also an int

	fmt.Println("Variable declaration with specifying the " +
		"explicit type: \n" +
		"- New variables assigned to a value equal the original" +
		" variable's value that do not have their explicit type" +
		" declared will take the type of the original variable... ")
	fmt.Println()

	fmt.Println("E.g. " +
		"\n// k is declared an int" +
		"\nvar k int " +
		"\n\n// m is assigned to the value of K," +
		"\n// hence K is consequently also an int:" +
		"\nm := k ")
	fmt.Println()

	fmt.Printf("Variable: m \n"+
		"Type: %T\n"+
		"Value: %v\n", m, m)
	fmt.Println()

}
