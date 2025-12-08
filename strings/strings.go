package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "こんにちは世界!" // String containing Unicode characters ( Hello World Japanese)
	fmt.Println("String:", s)
	fmt.Println("Length of string (in bytes):", len(s)) // len() returns byte count

	// Counting runes in a string
	runeCount := utf8.RuneCountInString(s)
	fmt.Println("Number of runes:", runeCount)

	// Iterating over runes in a string
	fmt.Println("Runes:")
	for i, r := range s {
		fmt.Printf("  Index %d: %v %c (Unicode code point: %U)\n", i, r, r, r)
	}

	// Declaring a rune variable
	myRune := '😎'
	fmt.Printf("My Rune: %c %v (Unicode code point: %U)\n", myRune, myRune, myRune)
}
