package main

import (
	"fmt"
)

func main() {
	fmt.Println("Testing arrays, slice and maps")

	// array
	var Arr [5]int = [5]int{1,2,3,4,5}
	// Works as normal array
	fmt.Println("Value at index 0 is ", Arr[0])
	fmt.Println("Value at range 1-4 is ", Arr[1:4])
	fmt.Println("Lenght of the array is ",len(Arr))

	// Slice
	// Similar to Arrays but get remade when adding elements
	var Slice []int = []int{1,2,3,4,5}
	// Works as array
	fmt.Println("Value at index 0 is ", Slice[0])
	fmt.Println("Value at range 1-4 is ", Slice[1:4])
	fmt.Println("Lenght of the array is ", len(Slice))
	fmt.Println("Slice was first at address ", &Slice[0])
	// Can add value like
	Slice = append(Slice, 6, 7, 8, 9, 10)
	Slice = append(Slice, Slice...)
	fmt.Println("Now check the first address of Slice: ", &Slice[0])
	clear(Slice)

	// Remake slice with better buffer
	var pSlice []int = make([]int, 5, 20)
	pSlice = append(pSlice, 1,2,3,4,5)
	fmt.Println("Slice was first at address ", &pSlice[0])
	pSlice = append(pSlice, 6,7,8,9,10)
	fmt.Println("Slice was first at address ", &pSlice[0])

	// Maps, similar to DefaultDict
	var dict map[string]int = make(map[string]int)
	dict["Rajj"] = 25
	dict["Naidu"] = 30
	dict["Jaiprakash"] = 35
	fmt.Println("Dict value for Rajj ", dict["Rajj"])
	// Doesn't throw error for unfound values
	fmt.Println("Dict value for Rajath ", dict["Rajath"])

	// Check flag for if value is there or not
	value, flag := dict["Rajath"]
	fmt.Printf("flag: %v for unknown key: Rajath with default value: %v", flag, value)
}
