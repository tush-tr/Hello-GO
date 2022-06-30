package main

import "fmt"

func main() {
	fmt.Println("________Slices data structures in Golang_______")

	// SYNTAX 1: for declaring slices, this syntax needs to be initialized with declaration.
	var fruitList = []string{"Apple","banana","mango"} // we need to put the values also using this syntax.
	// printing the type of slice
	fmt.Printf("Type of fruitList is:  %T\n", fruitList) // output: []string

	// append new items into our slice:
	fruitList = append(fruitList, "Grapes","melon")
	fmt.Println(fruitList)

	// removing/ slicing our slice(making parts of our slice)
	fmt.Println(fruitList[0:2])
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	// second syntax for declaring the slices
	highscores := make([]int, 4)
	highscores[0] = 232
	highscores[1] = 232313
	highscores[2] = 29319
	highscores[3] = 12312

	// we can increase the length of slice if we append more items in the slice.
	highscores = append(highscores,232,34352,24324,2323)
	fmt.Println(highscores)
	// print length of slice
	fmt.Println(len(highscores))
}
