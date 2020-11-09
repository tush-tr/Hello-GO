// Generate a random number in go with module math/rand

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var a = rand.Intn(100)
	fmt.Println("My fav number is", a)
	var b float64 = 81
	fmt.Println("Square root of number is ", math.Sqrt(b))
}
