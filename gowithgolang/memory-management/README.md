# Memory management and datatypes in Golang

## Memory Management

<li>Memory allocation and deallocation happens automatically.

### **_new()_**

<ul>
<li>Allocate memory but no INIT
<li>You will get memory address
<li>zeroed storage
</ul>

### **_make()_**

<ul>
<li>Allocate memory and INIT
<li>You will get memory address
<li>non - zeroed storage
</ul>

### Garbage collector

<li>Garbage collection happens automatically.
<li>Out of Scope or nil

https://pkg.go.dev/runtime

The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely. The runtime/debug package's SetGCPercent function allows changing this percentage at run time. See https://golang.org/pkg/runtime/debug/#SetGCPercent.

# Pointers in Golang

```sh
$ mkdir pointers
$ touch main.go
$ go mod init pointers
```

When you initialize a pointer with no value in it, it has the value nil

```go
// this is a pointer holding integer into that
	var ptr *int
	fmt.Println("Value of pointer is: ",ptr)
    // output: nil
```


refrencing a variable with a pointer

```go
    myNumber := 23
	// create a pointer which is refrencing to some variable's memory address
	// & - address of operator
	var ptr = &myNumber
	fmt.Println(ptr) // memory address 
	fmt.Println(*ptr) // actual value of the variable
```

# Arrays in Golang

Array is very less used in Golang. But there're other datatypes we can use that utitlize arrays.

```go
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
```

# Slices in Golang

Slices are much more powerful and much used in Golang. Slices are under the hood arrays.

