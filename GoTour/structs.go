package main

import "fmt"

// A struct is a collection of fields (my notes:
// operates somewhat like a class constructor, see below, one
// can call the struct and pass a compatible parameter profile to
// assign values to the fields and create an instance of the struct by
// assigning that call to a variable.
// Note: The parameter values won't stick as assigned values to the
// fields unless that call with the parameter profile passed is
// assigned to a variable to store that call/instance of the struct.. (
// the variable the instance has a data type syntax here is the method in
// which the call is made, dot ".", and the name of the Struct,
// i.e. type = main.Vertex02 )

type Vertex01 struct {
	// fields X and Y
	X int
	Y int
}

type Struct01 struct {
	// fields P and Q
	P int
	Q int
}

var (
	// A struct literal denotes a newly allocated struct by listing
	// the values of its fields..
	
	// You can just list a subset of fields by using the "Name
	// :" syntax. (And the order of named fields is irrelevant.)
	// as in "v2 = Vertex02{X: 1} " below, Y:0 is implicit (default?)
	
	// The special prefix "&" returns a pointer to the struct value
	
	v1 = Vertex01{1, 2}
	v2 = Vertex01{X: 1} // this uses the "name:" syntax to list a subset
	// of fields, the other fields default to take the value 0
	v3 = Vertex01{}      // X:0  and Y:0
	p  = &Vertex01{1, 2} // this has the type *Vertex02
)

func main() {
	
	fmt.Println()
	fmt.Println("Structs: ")
	
	fmt.Println()
	fmt.Println("Passes Temp field values ")
	// Passes temp field values:
	// call the struct/(constructor) and pass a parameter profile that
	// matches the field profile
	fmt.Println("fmt.Println(Vertex02{1, 2}) returns...")
	fmt.Println(Vertex01{1, 2})
	
	fmt.Println("However, \"fmt.Println(" +
		"Vertex02{})\" returns default values 0, 0 ... ")
	fmt.Println(Vertex01{})
	
	var s1 = Vertex01{2, 4}
	fmt.Println("Assigning struct call to a variable, " +
		"making the struct behave like an object/instance of a class" +
		" with field values...")
	fmt.Println()
	fmt.Println("var s1 = Vertex02{2, 4} ")
	fmt.Printf("s1 type: %T, s1 value %v \n", s1, s1)
	
	fmt.Println()
	// Struct fields are accessed (
	// getting and setting it seems) with a dot.
	v := Vertex01{1, 2}
	fmt.Println("v := Vertex02{1, 2}")
	fmt.Printf("v.X type: %T, v.X value: %v \n", v.X, v.X)
	
	fmt.Println()
	v.X = 4
	fmt.Println("v.X = 4 ")
	fmt.Printf("v.X type: %T, v.X value: %v \n", v.X, v.X)
	
	fmt.Println()
	fmt.Printf("v type: %T, v value %v \n", v, v)
	
	// Pointers to structs:
	// Struct fields can be accessed through a struct pointer
	
	// To access the field X of a struct when we have the struct
	// pointer p we could write (*p).X, but that is cumbersome,
	// so the language permits us instead to write just
	//	p.X
	// without explicit deference
	
	// My variables etc, with the GoTour's lesson
	// assigning calls that passes compatible parameter profiles to
	// variables, ie sort of making "instances" of the structs /"class
	// constructors"
	strct1 := Struct01{1, 2}
	// or
	var strct2 = Struct01{2, 4}
	
	// point at strct1, strct2
	pntr1 := &strct1
	pntr2 := &strct2
	
	// access and set a field value
	pntr1.P = 2
	pntr2.Q = 8
	
	fmt.Println()
	fmt.Printf("strct1 type: %T, strct1 value: %v \n", strct1, strct1)
	fmt.Printf("strct2 type: %T, strct1 value: %v \n", strct2, strct2)
	fmt.Println()
	
	// Struct Literals:
	fmt.Println()
	fmt.Println("Struct Literals: ")
	fmt.Printf("v1 type: %T, v1 value: %v \n", v1, v1)
	fmt.Printf("v1 type: %T, v1 value: %v \n", p, p)
	fmt.Printf("v2 type: %T, v2 value: %v \n", v2, v2)
	fmt.Printf("v3 type: %T, v3 value: %v \n", v3, v3)
	// fmt.Printf("p type: %T, p value: %v \n", p, p)
	
}
