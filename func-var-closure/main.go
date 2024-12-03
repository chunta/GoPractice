package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	inPlaceTransformNums(&numbers, double)
	fmt.Println(numbers)

	inPlaceTransformNums(&numbers, triple)
	fmt.Println(numbers)

	inPlaceTransformNums(&numbers, func(number int) int {
		return number * 4
	})
	fmt.Println(numbers)
}

func inPlaceTransformNums(numbers *[]int, transform transformFn) {
	for i, value := range *numbers {
		(*numbers)[i] = transform(value)
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
