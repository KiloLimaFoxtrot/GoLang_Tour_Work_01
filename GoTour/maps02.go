package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Maps02: ")
	
	fmt.Println()
	fmt.Println("****Mutating Maps:")
	
	//  Insert or update an element in map m:
	//
	// m[key] = elem
	//
	// Retrieve an element:
	//
	// elem = m[key]
	//
	// Delete an element:
	//
	// delete(m, key)
	//
	// Test that a key is present with a two-value assignment:
	//
	// elem, ok = m[key]
	//
	// If key is in m, ok is true. If not, ok is false.
	//
	// If key is not in the map,
	// then elem is the zero value for the map's element type.
	//
	// Note: If elem or ok have not yet been declared you could use a
	// short declaration form:
	//
	// elem, ok := m[key]
	
	fmt.Println()
	fmt.Println("  **FUNCTION mapMutate01(): ")
	mapMutate01()
	
}

func mapMutate01() {
	fmt.Println()
	mapMtt01 := make(map[string]int)
	
	mapMtt01["Answer"] = 42
	fmt.Printf("The value: %v \n", mapMtt01["Answer"])
	
	mapMtt01["Answer"] = 48
	fmt.Printf("The value: %v \n", mapMtt01["Answer"])
	
	delete(mapMtt01, "Answer")
	fmt.Printf("The value: %v \n", mapMtt01["Answer"])
	
	value, yes := mapMtt01["Answer"]
	fmt.Printf("The Answer value: %v, Previous value present? %v \n",
		value, yes)
	
}
