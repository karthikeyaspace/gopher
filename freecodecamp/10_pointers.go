package main

// Pointers in go
// Go has pointers. A pointer holds the memory address of a value.
// * -> pointer operator
// & -> address operator

// pointers are mainly used for increasing performance but is dangerous to use because it can cause memory leaks and nil pointer exceptions

import "fmt"

func pointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	mystr := "karthikeya"
	mystrptr := &mystr

	fmt.Println("Value of mystr: ", mystr)
	fmt.Println("Address of mystr: ", mystrptr)

	// Dereferencing the pointer
	fmt.Println("Value of mystrptr: ", *mystrptr)

}
