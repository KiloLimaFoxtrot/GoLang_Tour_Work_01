package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println("Pointers: ")

	/*
		Go has pointers. A Pointer holds the memory address of a value.

		(Is a pointer like a Python "shallow copy",
		which is a copy variable that simply points to the memory
		location of another variable. In contrast,
		a Python "deep copy" would be a variable that has its own memory
		address.

		To me, it seems pointers may be something like a data type
		specific placeholder for variables that represent specific data
		types like int, string, float64 etc..)

		The type *T is a pointer to a T value. Its zero value is nil.

			var p *int //Here p is a pointer to an int val

		The & operator generates a pointer to its operand

			i := 42 //this is will default to type int

			//again, the & operator generates a pointer to its operand
			p = &i

		The * operator denotes the pointer's underlying value.

			fmt.Println(*p) //this is read, i through the pointer p
			*p = 21			//this is read, set i through the pointer p

		This is known as "de-referencing" or "in-directing".

		Unlike C, Go has no pointer arithmetic

	*/

	pointer01()

}

func pointer01() {
	fmt.Println()

	i, j := 42, 2701

	// the & operator generates a pointer to its operand,
	// here i's memory address (?),
	// and we are assigning that operator to the variable p

	// i.e. &i is read as, point to i
	// (and then we assign that point-ing to p?)
	p := &i

	// my test
	q := &j

	// Test prints
	fmt.Printf("Variable i type: %T, value: %v \n", i, i)
	fmt.Printf("Variable j type: %T, value: %v \n", j, j)

	fmt.Println()
	// the below prints out p's type: a pointer to a int value,
	// and p's current value: a memory address (
	// presumably for the int value)
	fmt.Printf("Pointer p type: %T, value: %v \n", p, p)

	// the below prints out q's type: a pointer to a int value,
	// and q's current value: a memory address (
	// presumably for the int value)
	fmt.Printf("Pointer q type: %T, value: %v \n", q, q)

	// the * operator denotes the pointer's underlying value (
	// stored in the memory address?)
	// *p is read as, read i through the pointer p
	fmt.Println()
	fmt.Printf("*p value: %v \n", *p)
	fmt.Println(" (*p is read as, read i through the pointer p)")
	fmt.Printf("*p type: %T, *p value %v \n", *p, *p)

	// from above, q := &j, &j is read as, point to j,
	// ie q is set to point to j...
	// *q is read as, read j through the pointer q
	fmt.Printf("*q value: %v \n", *q)
	fmt.Println(" (*q is read as, read j through the pointer q)")
	fmt.Printf("*q type: %T, *q value %v \n", *q, *q)

	fmt.Println()
	*q = *q / 37
	fmt.Println("divide j through (*) the pointer q")
	fmt.Println("*q = *q / 37 ")
	fmt.Printf("j type: %T, j value: %v \n", j, j)

	// my test...
	// read as, set the value of i through the pointer p
	fmt.Println()
	*p = 42 / 2
	fmt.Println("*p = 42 / 2 ")
	fmt.Println(" (set the value of i through the pointer p)")
	fmt.Printf("i type: %T, i value: %v \n", i, i)

	// read as, set the value of j through the pointer q
	fmt.Println()
	*q = 2701 / 2
	fmt.Println("*q = 2701 / 2 ")
	fmt.Println(" (set the value of i through the pointer p)")
	fmt.Printf("j type: %T, j value: %v \n", j, j)

}
