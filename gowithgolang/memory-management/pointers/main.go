package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers explanation")
	// this is a pointer holding integer into that
	// var ptr *int
	// fmt.Println("Value of pointer is: ",ptr)
	myNumber := 23
	// create a pointer which is refrencing to some variable's memory address
	// & - address of operator
	var ptr = &myNumber
	fmt.Println(ptr) // memory address 
	fmt.Println(*ptr) // actual value of the variable

	*ptr = *ptr *2
	fmt.Println(myNumber)
}
