package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Random number generation")
	// var mynumberOne int = 2;
	// var mynumberTwo float64 = 3.44
	// fmt.Println("The sum is: ",mynumberOne+int(mynumberTwo))

	// random number
	// rand.Seed(30)
	// fmt.Println(rand.Intn(5)) // math/rand

	// TODO: Random number generation  using crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
