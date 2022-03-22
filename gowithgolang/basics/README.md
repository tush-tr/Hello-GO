## Variables in Golang
to create a variable we use a keyword "var". There're other ways also like valorus operator(:) but we can't declare a variable using volarus operator outside a function(global scope)

- we can use constants using const keyword.



```go
package main

import "fmt"

var jwtToken int = 300000
const LoginToken string = "tushso3409403fn9fncfd8f8390fh304hf938340" // public

func main() {
	var username string = "tushar"
	fmt.Println(username)
	// tushar
	fmt.Printf("Variable is of type: %T \n", username)
	// Variable is of type: string

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 256.34434
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variale is of type: %T \n", anotherVariable)

	// implicit way of declaring variables
	var website = "tush-tr.github.io"
	fmt.Println(website)

	// no var style
	myName := "Tushar"
	fmt.Println(myName)

	fmt.Println(jwtToken)
	fmt.Println(LoginToken)
	fmt.Printf("Variale is of type: %T \n", LoginToken)
}

```

## Types in Golang
<li>Case insensitive; almost
<li>Variable type should be known in advance
<li>Everything is a type.

### Types
  - String
  - Bool
  - Integer
    - uint8,uint64,int8,int64,uintptr
  - Floating
    - float32,float64
  - Complex

  - Array
  - Slices
  - Maps
  - Structs
  - pointers

  - functions
  - channels
  - almost everything

## Comma ok Syntax and packages in golang

<li>Getting inputs in golang

```Golang
reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok // err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
```

## Conversions in golang

```Golang
func main() {
	fmt.Println("Welcome to our app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRating+1)

	}

}
```

## Math, crypto and random numbers

