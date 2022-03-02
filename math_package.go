package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64 = 12

	var result float64 = math.Sqrt(num)
	// var intResult = math.Round(result)
	// var intResult = math.Ceil(result)
	var intResult = math.Floor(result)

	fmt.Println("result", intResult)
	fmt.Printf("The result is %.2g", result)

	//or
	//fmt.Printf("The result is %.2f", result)
}
