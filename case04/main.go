package main

import "fmt"

type Human struct {
	Age int
	Height float32
	Sex bool
}

func main() {
	// var c = Human{
	// 	Age: 35,
	// 	Height: 180,
	// 	Sex: true,
	// }
	
	// if c.Age > 10 {
	// 	fmt.Printf("%+v", c)
	// }

	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
}