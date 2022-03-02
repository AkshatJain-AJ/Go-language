package main

import (
	"fmt"
)

func main() {
	num := 3
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("None")
	}
}