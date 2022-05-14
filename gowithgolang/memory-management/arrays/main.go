package main

import "fmt"
func main(){
	fmt.Println("Arrays in Golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"
	fmt.Println(fruitList)
	// [Apple Banana  Mango]
	fmt.Println(len(fruitList))

	var vegList = [5]string{"potato","tomato","Capsicum"}
	fmt.Println(vegList)
	fmt.Println(len(vegList)) // 5

}