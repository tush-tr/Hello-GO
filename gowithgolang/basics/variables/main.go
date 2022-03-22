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
