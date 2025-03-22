package main

import "fmt"

func main() {
	age := 23
	fmt.Println("Age: ", age)

	var agePointer *int
	agePointer = &age

	fmt.Println(agePointer)
	fmt.Println(*agePointer)

}
