package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("****Exercise03: Maps:")
	
	// Implement WordCount.
	//
	// It should return a map of the counts of each “word” in the
	// string s.
	//
	// The wc.Test function runs a test suite against the provided
	// function and prints success or failure.
	//
	// You might find strings.Fields helpful.
	
	fmt.Println()
	fmt.Println("  **FUNCTION WordCount(): ")
	
	stringToCnt := "I will will be glad if if if this works."
	
	wordCntMapRslt := WordCount(stringToCnt)
	
	// to print the map
	fmt.Println()
	fmt.Println("The word count map key (string) - value (count): " +
		"count) pairs... ")
	iteratePrintMap(wordCntMapRslt)
	
}

// Short version without notes
func WordCount(s string) map[string]int {
	
	wordCntMap := make(map[string]int)
	wordSpltArry01 := strings.Split(s, " ")
	
	for i := 0; i < len(wordSpltArry01); i++ {
		subStrXX := wordSpltArry01[i]
		subStrXXCnt := strings.Count(s, subStrXX)
		wordCntMap[subStrXX] = subStrXXCnt
	}
	
	return wordCntMap
}

// Long version with notes etc...
/*func WordCount(s string) map[string]int {
	// var wordSpltArry01 []string

	// to separate the words/substrings of the s,
	// so we can map a key-value pair of their string-value as a key,
	// and their count of occurrences as a value
	wordSpltArry01 := strings.Split(s, " ")

	// the empty map to store the key value; i.e.
	// unique words/substrings,
	// to the value; ie their count of occurrence
	var wordCntMap = make(map[string]int)

	for i := 0; i < len(wordSpltArry01); i++ {
		// each word/substring in the wordSpltArry01
		// subStrXX := wordSpltArry01[i]
		var subStrXX = wordSpltArry01[i]

		// the count of the word/substring in the s, we can use the
		subStrXXCnt := strings.Count(s, subStrXX)
		// print test of the counts
		// fmt.Println("subStrXXCnt: ", subStrXXCnt)

		// enter the key: word/substring and value: word/substring
		// count
		wordCntMap[subStrXX] = subStrXXCnt

		// test print
		// fmt.Printf("Key: %v   \tVal: %v \n\n", subStrXX,
		// 	wordCntMap[subStrXX])
	}

	return wordCntMap
}*/

func iteratePrintMap(wordCntMapRslt map[string]int) {
	// Iterating over the map using the for-range loop
	for key, val := range wordCntMapRslt {
		fmt.Printf("Key (string): %v   \tVal (count) : %v \n", key, val)
	}
	
}
