package main

import (
	"fmt"
)

func main() {
	fmt.Println("Testing loops in go")
	fmt.Print("Enter the limit to stop at: ")
	var a int
	fmt.Scan(&a)
	
	// For loop with count 
	for i := 0; i < a; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// For till range is hit
	count := a
	for count > 0 {
		fmt.Print(count)
		count--
	}
	fmt.Println()

	// For loop through a list
	// First add elements to it
	list := []int{}
	for i := 0; i < a; i++ {
		list = append(list, i)
	}
	
	for i, v := range list {
		fmt.Printf("`%v:%v`", i, v)
	}
	fmt.Println()

}