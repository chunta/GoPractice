package main

import (
	"fmt"
	"math"
	"case02/fileops"
	 "github.com/Pallinder/go-randomdata"
)

const inflationRate = 2.5

func main() {
	
	var investmentAmount float64 = 1000

	fmt.Print("input:")
	fmt.Scan(&investmentAmount)

	years, expectedReturnRate := 10.0, 5.5
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / inflationRate
	fmt.Println(futureValue)
	fmt.Printf("Future value: %.0f.00\n", futureValue)
	fmt.Println(futureRealValue)

	var choice int64 = 0
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("MAKE 1M")
	case 2:
		fmt.Println("MAKE 2M")
	default:
		fmt.Println("make 10m")
	}
	if choice == 2 {
		panic("dont select 2")
	}
	fileops.WriteFileOp("test.txt", fmt.Sprintf("%d", choice))
	stringContent, error := fileops.ReadFileOp("test.txt")
	if error != nil {
		panic("error while read file")
	}
	fmt.Printf("content: %s\n", stringContent)
	fmt.Println(randomdata.SillyName())
}

