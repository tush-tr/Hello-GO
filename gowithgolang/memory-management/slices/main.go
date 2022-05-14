package main

import "fmt"

func main() {
	fmt.Println("Slices data structures in Golang")
	// this syntax needs to be initialized
	var fruitList = []string{"Apple","banana","mango"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Grapes","melon")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
}
