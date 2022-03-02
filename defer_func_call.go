package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Byee.....")
}

// func main() {
// 	a()
// }

// func a() {
// 	fmt.Println("a begins")
// 	defer b()
// 	fmt.Println("a ends")
// }

// func b() {
// 	fmt.Println("in b")
// }
