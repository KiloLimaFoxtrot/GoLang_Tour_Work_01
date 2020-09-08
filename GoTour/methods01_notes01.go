package main

/*
type VrtxStrct01 struct {
	X, Y float64
}

func main() {
	fmt.Println()
	fmt.Println("**** Methods 01:")*/
/*
	Go does not have classes,
	where you can create objects/instances of classes that have
	states (field values) and behaviors (
	methods that can be called on those objects).

	Go does allow you to define methods on data types.
	(ie Go allows you to treat variables assigned to data types
	as objects, and define methods on those types)

	A method is a function with a special receiver argument.

	The receiver appears in its own argument list between the func
	keyword and the method name.

	In this example,
	the Abs01 method has a receiver of type Vertex02 named v.

	From geeksforgeeks on GoLang methods
	https://www.geeksforgeeks.org/methods-in-golang/
	Go language support methods.
	Go methods are similar to Go function with one difference,
	i.e, the method contains a receiver argument in it.
	With the help of the receiver argument,
	the method can access the properties of the receiver.
	[more at geeks site...]

	(Again, Go allows one to treat variables assigned to data
	types as objects, and can define methods on them)

	With the below syntax,
	the receiver can be accessed within the method

	syntax:
	func(receiver_name type) method_name(
	method parameter_list / ie profile?) (return type){
	// code...
	}
*/

// Receiver_name initialized and assigned to a data type
// (the struct named VrtxStrct01),
// ie a variable is assigned to a data type, a struct,
// and this variable will have a method defined on it below
/*v1 := VrtxStrct01{
	X: 3,
	Y: 4,
}*/

// Here just below,
// we have a variable v1AbsVal assigned to the method call Abs01()

// The method call Abs01() broken down is,
// 1. A call on the function Abs01(), which is
// 2. defined on the data type v1,
// which is a struct defined as VrtxStrct01 (
// instantiated and initialized above)

// The method Abs01() is instantiated/initialized/defined further
// below

/*v1AbsVal := v1.Abs01()

	fmt.Println()
	fmt.Println("v1AbsVal: ", v1AbsVal)

}*/

// Abs01() method instantiation/initialization/definition etc...

// Below, this is a method, not just a function,
// because it defines a relationship between a data type and a
// function, a relationship similar to an object and its behavior
// /methods that can be called on that object

// In java, only objects have both in java terms...
// 1. 'behavior',
// ie have functions/methods (
// semantically the same in java) that can be called on them,
// and
// 2. 'states',
// ie field values that can be get/retrieved and set/changed...

// in GoLang, data types, of struct or non-struct type,
// can have
// 1. functions defined/called on them,
// ie data types can have 'behavior',
// and
// 2. hold values that can be get and set in the data type itself,
// ie data types can have 'states'

// ie this method defines a function on a data type,
// ie this method defines Abs01() (with a return tpye of float64)
// on the variable v1 (assigned to the struct VrtxStrct01))

/*func (v1 VrtxStrct01) Abs01() float64 {
	absVal := math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
	return absVal
}
*/
