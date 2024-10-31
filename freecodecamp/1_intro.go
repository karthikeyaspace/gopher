package main

import "fmt"

// https://www.youtube.com/watch?v=un6ZyFkqFKo

/*
	- Go has very fast compile time and execution time
	- Go is slower in execution than C and Rust
	- Statically typed
	- go compiles directly to machine code like Rust and C
	- execution is similar to Java, C#
	- Go uses much less memory than Java and C#
*/

/*
	Data types in GO
	- bool
	- string
	- int, int8, int16, int32, int64, uint8, uint16, uint32, uint64, uintptr
	- byte (alias for uint8)
	- rune (alias for int32) - represents a Unicode code point
	- float32, float64
	- complex64, complex128

	- int, uint, float64 are enough for most cases


	- Go does not allow unused variables
*/

/*
	Important go commands
	- go run <filename>.go
	- go build <filename>.go
	- go fmt <filename>.go
	- go install <filename>.go
	- go get <package-name>
*/

func main() {
	fmt.Println("Hello, World!")

	// Declaring variables

	var age int = 40
	var name string = "John Doe"

	// Shorthand declaration - type inference
	age1 := 40
	name1 := "John Doe"

	fmt.Println(name, "is", age, "years old")
	fmt.Println(name1, "is", age1, "years old")

	i := 42
	f := 3.14
	g := 0.867 + 0.5i

	fmt.Printf("%T %T %T\n", i, f, g)

	// multiple variable declaration
	name2, age2 := "Jane Doe", 30
	fmt.Println(name2, "is", age2, "years old")

	// Printf - formatted output
	// %v - value in default format
	// %T - type of the value
	// %d - decimal
	// %f - float
	// %s - string

	// Conditional statements

	// if else
	if age > 30 {
		fmt.Println("You are old")
	} else {
		fmt.Println("You are young")
	}

	// switch case
	switch age {
	case 20:
		fmt.Println("You are 20 years old")

	case 30:
		fmt.Println("You are 30 years old")

	default:
		fmt.Println("You are neither 20 nor 30 years old")
	}

	// Loops

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop

	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// break and continue

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// Arrays
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
}
