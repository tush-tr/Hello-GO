- [Variables in Golang](#variables-in-golang)
- [Types in Golang](#types-in-golang)
	- [Types](#types)
- [Comma ok Syntax and packages in golang](#comma-ok-syntax-and-packages-in-golang)
- [Conversions in golang](#conversions-in-golang)
- [Math, crypto and random numbers](#math-crypto-and-random-numbers)
- [Handling time in Golang](#handling-time-in-golang)
- [Build in Go for Mac, Linux and windows](#build-in-go-for-mac-linux-and-windows)
# Variables in Golang
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

# Types in Golang
<li>Case insensitive; almost
<li>Variable type should be known in advance
<li>Everything is a type.

## Types
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


# Comma ok Syntax and packages in golang

<li>Getting inputs in golang

```Golang
reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok // err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
```


# Conversions in golang

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



# Math, crypto and random numbers

Two ways we can generate random numbers in Golang

1. Math package

```golang
import (
	"fmt"
	"math/rand"
)
func main() {
	fmt.Println("Random number generation")
	rand.Seed(30)
	fmt.Println(rand.Intn(5))
}
```

2. Crypto package 

```golang
import (
	"crypto/rand"
	"fmt"
	"math/big"
)
func main() {
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
```


# Handling time in Golang

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	// Formatting the time and date
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// Note: this is only format need to stick with it

	createdDate := time.Date(2020, time.March, 10, 23, 23, 0, 0, time.Local)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
```

# Build in Go for Mac, Linux and windows
Go allows us to build mac, linux and windows build files.

check env variables in Go:

```sh
go env
```

using above command you will see this environment variable:

```sh
GOOS="linux"
```

Now for building your project for a target Operating system and target architecture 

```sh
$ GOOS="linux" GOARCH="amd64" go build
$ GOOS="windows" GOARCH="amd64" go build # build exe file
```
