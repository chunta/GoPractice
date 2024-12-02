package main

import (
	"fmt"
)

type GetArea interface {
	save(int, string) error
	update(string)
}

type Circle struct {
	Radius float64
}

func (c Circle) GetArea() float64 {
	return c.Radius * c.Radius * 3.1416
}
 

func add[T int|float64](a, b T) T {
	return a + b
}

func main() {
	result := add(10, 11)
	fmt.Print(result)
}
