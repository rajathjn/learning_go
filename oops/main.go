// Test code for shapes
package main

import (
	"fmt"
	"learning_go/oops/circle"
	"learning_go/oops/rectangle"
)

type shape interface{
	Scale(factor float64)
}

func Scale_shape(s shape, factor float64) {
	s.Scale(factor)
}

func main() {
	c := circle.Circle{
		Radius: 5.0,
	}
	fmt.Println("Circle Area:",c.Area())
	fmt.Printf("Circle Radius: %v\n", c.Radius)
	Scale_shape(&c, 4)
	fmt.Printf("Circle Radius: %v\n", c.Radius)
	fmt.Println("Circle Area:",c.Area())


	r := rectangle.Rectangle{
		Width: 10,
		Height: 24,
	}
	fmt.Println("Rectangle Area:",r.Area())
	
	fmt.Printf("Rectangle width: %v and height: %v\n", r.Width, r.Height)
	Scale_shape(&r, 3)
	fmt.Printf("Rectangle width: %v and height: %v\n", r.Width, r.Height)
	fmt.Println("Area:",r.Area())
}
