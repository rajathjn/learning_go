// Simple Calculator using go
package main

import (
	"errors"
	"fmt"
)

func main() {
	printHello()

	fmt.Println("Enter two numbers for calculation")
	var a int
	var b int
	fmt.Print("Enter the value for a: ")
	fmt.Scan(&a)
	fmt.Print("Enter the value for b: ")
	fmt.Scan(&b)

	var err error
	var c int 

	c, err = add_num( a, b)
	switch err {
	case nil:
		fmt.Printf("Result of adding %v and %v is %v\n", a, b, c)
	default:
		fmt.Println(err.Error())
	}

	c, err = sub_num(a, b)
	switch err {
	case nil:
		fmt.Printf("Result of subtracting %v and %v is %v\n", a, b, c)
	default:
		fmt.Println(err.Error())
	}

	c, err = mul_num(a, b)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result of multiplying",a,"and",b,"is",c)
	}

	var d int
	c, d, err = div_num(a, b)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result of dividing",a,"and",b,"is",c,"with remainder",d)
	}

}

func add_num( a int, b int) (int,error) {
	var err error
	return a+b, err
}

func sub_num( a int, b int) (int,error) {
	var err error
	return a-b, err
}

func mul_num( a int, b int) (int,error) {
	var err error
	return a*b, err
}

func div_num( a int, b int) (int, int, error) {
	var err error
	if (b==0) {
		err = errors.New("can't divide by zero")
		return 0,0,err
	}
	return a/b, a%b, err
}