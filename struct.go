package main

import (
	"fmt"
)

type Student struct {
	roll_no int
	name    string
	marks   int
}

func main() {

	var s1 Student = Student{101, "Akshat", 99}
	fmt.Println(s1)
	fmt.Println(s1.name)

	var s2 Student = Student{roll_no: 102, marks: 59}
	fmt.Println(s2)
	fmt.Println(s2.name) // it will not print any value

}
