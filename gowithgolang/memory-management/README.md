- [Memory Management](#memory-management)
		- [**_new()_**](#new)
		- [**_make()_**](#make)
		- [Garbage collector](#garbage-collector)
- [Pointers in Golang](#pointers-in-golang)
- [Arrays in Golang](#arrays-in-golang)
- [Slices in Golang](#slices-in-golang)
	- [Slice length and capacity](#slice-length-and-capacity)
	- [Declaring slices](#declaring-slices)
	- [append method:](#append-method)
	- [Nil slices](#nil-slices)
	- [Another syntax for declaring Slices: make()](#another-syntax-for-declaring-slices-make)
	- [Remove a value from Slice based on Index](#remove-a-value-from-slice-based-on-index)


# Memory Management

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

- An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

- A slice is formed by specifying two indices, a low and high bound, separated by a colon:  
    ```golang
    a[low : high]
    ```

- Slices are like reference to arrays:
  - A slice does not store any data, it just describes a section of an underlying array.
  - Changing the elements of a slice modifies the corresponding elements of its underlying array.
  - Other slices that share the same underlying array will see those changes.


## Slice length and capacity

- A slice has both a length and a capacity.
- The length of a slice is the number of elements it contains.
- The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
- The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
- You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

	```golang
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	```

## Declaring slices

```golang
var fruitList = []string{"Apple","banana","mango"}
```

here using this above syntax, we always need to initialize the slice. 

- In slices we don't need to specify length of the slice so that we can add as many items we want to the slices.
- It automatically expends the memory for us.
- to put a value to the slice we use a method, called append():

## append method:

```golang
fruitList = append(fruitList, "Grapes","melon")

// for slicing our slice.....(making parts of our slice)
fruitList = append(fruitList[1:])
```

## Nil slices
- The zero value of a slice is nil.
- A nil slice has a length and capacity of 0 and has no underlying array.


## Another syntax for declaring Slices: make()

```golang
highscores := make([]int, 4)
```

- Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

- The make function allocates a zeroed array and returns a slice that refers to that array:

	```golang
	a := make([]int, 5)  // len(a)=5
	```


## Remove a value from Slice based on Index 
```go
var courses = []string{"reactjs","javascript","swift","elasticSearch","docker","Kubernetes"}
index := 2
courses = append(courses[:index],courses[index+1:]...)
```

# Hey

