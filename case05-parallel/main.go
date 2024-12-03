package main

import (
	"fmt"
	"sync"
)

var (
	n int32
	lock = sync.Mutex{}
)

func foo() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		n++
		lock.Unlock()
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		foo()
	}()
		go func() {
		defer wg.Done()
		foo()
	}()
	wg.Wait()
	fmt.Println(n);
}