package main

import (
	"fmt"
)

type (
	Vertex02 struct {
		Lat, Long float64
	}
	
	// my test
	Point01 struct {
		xVal, yVal int
	}
)

var (
	// my notes: the below syntax works as follows?
	// map01 : map variable name
	// map[Key_Type]Value_Type{key1: value1, ... keyN: valueN}
	// below the Value_Type is the struct VrtxStrct01,
	// which has two fields (
	// not sure of their GoLang name) and are initialized in the
	// struct and instantiated in the function mapBldg01 below
	
	// this, which is using the var keyword, so is essentially
	// var map01 map[string]VrtxStrct01
	// creates and initializes an empty map
	map01 map[string]Vertex02
	
	// my test
	map02 map[string]Point01
)

func main() {
	fmt.Println()
	fmt.Println("Maps: ")
	
	// A map maps keys to values
	
	// The zero value of a map is nil.
	// A nil map has no keys, nor can keys be added
	
	// The make function returns a map of the given type,
	// initialized and ready for use
	
	fmt.Println()
	fmt.Println("****Map Basic:")
	
	fmt.Println()
	fmt.Println("  **FUNCTION mapBld01(): ")
	mapBld01()
	
	fmt.Println()
	fmt.Println("  **FUNCTION mapBld02(): ")
	mapBld02()
	
	fmt.Println()
	fmt.Println("****Map literals:")
	// Map literals are like struct literals, but the keys are required
	
	fmt.Println()
	fmt.Println("  **FUNCTION mapBld03(): ")
	mapBld03()
	
	fmt.Println()
	fmt.Println("****Map literals continued:")
	
	// If the top-level type is just a type name,
	// you can omit it from the elements of the literal
	
	fmt.Println()
	fmt.Println("  **FUNCTION mapBld04(): ")
	mapBld04()
}

func mapBld04() {
	
	// **This function is basically the same as mapBldg03,
	// but with the top-level type name omitted
	
	fmt.Println()
	// using the var keyword to create and instantiate the map
	// We are using a string as the key, and
	// the struct VrtxStrct01, which has two 'fields', as the value
	
	//              key:    value:
	var map03 = map[string]Vertex02{
		// key:      value:
		"Bell Labs": {
			Lat:  40.68433, // the values
			Long: -74.39967,
		},
		"Google": {
			Lat:  37.42202,
			Long: -122.08408,
		},
	}
	fmt.Printf("map04: %v \n", map03)
	
}

func mapBld03() {
	fmt.Println()
	
	// using the var keyword to create and instantiate the map
	// We are using a string as the key, and
	// the struct VrtxStrct01, which has two 'fields', as the value
	
	//              key:    value:
	var map03 = map[string]Vertex02{
		// key:      value:
		"Bell Labs": Vertex02{
			Lat:  40.68433, // the values
			Long: -74.39967,
		},
		"Google": Vertex02{
			Lat:  37.42202,
			Long: -122.08408,
		},
	}
	fmt.Printf("map03: %v \n", map03)
	
}

func mapBld02() {
	fmt.Println()
	// make and initialize the map
	map02 = make(map[string]Point01)
	
	// set values in the map
	map02["Point01"] = Point01{
		10, 15,
	}
	
	fmt.Println()
	fmt.Printf("map02: %v \n", map02["Point01"])
	
}

func mapBld01() {
	fmt.Println()
	
	// creating and initialize map01, making it non-empty
	map01 = make(map[string]Vertex02)
	
	// Set the key_Value: "Bell Labs",
	// the value_Value: VrtxStrct01 which is a struct with two
	// 'field' values, Lat and Long)
	map01["Bell Labs"] = Vertex02{
		Lat:  40.68433,
		Long: -74.39967,
	}
	
	// use the key_Value to retrieve the value assigned to that key
	fmt.Printf("map01: %v \n", map01["Bell Labs"])
	
}
