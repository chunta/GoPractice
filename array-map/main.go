package main

import "fmt"

type produce struct {
	title string
	id    string
	price float64
}

func main() {
	productNameList := [4]string{"sarah", "rex", "kenny", "gogoro"}
	fmt.Println(productNameList)
	fmt.Println(productNameList[1])

	prices := [100]float64{11.0}
	fmt.Println(prices)
	
	// Sliced array - this is not a copy of origin array. its ref type.
	slicedPrices := productNameList[1:3]
	fmt.Println(slicedPrices)
	slicedToEndPrices := productNameList[1:]
	fmt.Println(slicedToEndPrices)
	slicedToEndPrices[0] = "stan"
	fmt.Println(productNameList)
	fmt.Println(len(slicedToEndPrices), cap(slicedToEndPrices))

	slicedToEndPrices2 := productNameList[:1]
	fmt.Println(slicedToEndPrices2)

	// Slice and Copy array
	original := []int{1, 2, 3, 4, 5}
	copySlice := make([]int, len(original)) 
	copy(copySlice, original)
	fmt.Println("copySlices:", copySlice)
	for index, value := range copySlice {
		fmt.Println(index, value)
	}

	// Map
	websites := map[string]string{
		"Google":"http://google.com",
		"Amazon":"http://aws.com", 
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	fmt.Println(websites["Amazon"])
	fmt.Println(websites["Amaz"])
	websites["Twitter"] = "http://www.twitter.com"
	delete(websites, "Amazon")
	fmt.Println(websites)

	for key, value := range websites {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}
}
