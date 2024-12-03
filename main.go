package main

import "fmt"

type produce struct {
	title string
	id    string
	price float64
}

func main() {
	prices := [100]float64{11.0}
	fmt.Println(prices)
}
