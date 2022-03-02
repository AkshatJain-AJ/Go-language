package main

import "fmt"

func add(a, b int) (int, int) {
	var out = a + b // or out := a+b
	var out1 = a - b
	return out, out1
}

func main() {
	num1 := 5
	num2 := 4
	result1, result2 := add(num1, num2)
	fmt.Print(result1, result2)
}
