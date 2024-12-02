package main

import (
	"fmt"
	"time"
)


func foo() {
	fmt.Println("extra goroutine")
}

// func f2() {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("f2")
// }

func f1() {
	// fmt.Println("f1")
	// go f2()
	// It's similar to try catch 
	defer func () {
		err := recover()

		if err != nil {
			fmt.Println("error: ", err)
		}
	}()
	a,b := 3,0
	fmt.Println(a, b)
	// will happen panic -> exit process immediately
	_ = a / b
	fmt.Println("f1 finished")
}

func main() {
	go f1()
	time.Sleep(1 * time.Second)
	fmt.Println("main finished")
}

// func main() {
// 	// go foo()
// 	wg := sync.WaitGroup{}
// 	// Add function includes the number of goroutines
// 	wg.Add(2)
// 	// Anonymous functions
// 	go func() {
// 		// before the function ends, it will call Done function
// 		defer wg.Done()
// 		fmt.Println("test1")
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("test2")
// 	}()
// 	wg.Wait()
// }