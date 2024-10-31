package main

import "fmt"

// in go, arrays have a fixed size, slices are a dynamic array
// slices are like references to arrays
// a slice does not store any data, it just describes a section of an underlying array
// changing the elements of a slice modifies the corresponding elements of its underlying array

func demo() {
	// array
	var a [5]int
	b := [5]int{1, 2, 3, 4, 5}
	a[2] = 7
	fmt.Println(a)
	fmt.Println(b)

	// slice - dynamic array
	c := []int{1, 2, 3, 4, 5}
	var d []int = b[1:4]
	fmt.Println(c)
	fmt.Println(d)

	// slice length and capacity
	fmt.Println(len(d))
	fmt.Println(cap(d))

	// make function - create a slice
	e := make([]int, 5)
	fmt.Println(e)

	// append function - add elements to a slice
	e = append(e, 6)
	fmt.Println(e)

	// copy function - copy elements from one slice to another
	f := make([]int, len(e))
	copy(f, e)
	fmt.Println(f)

	// multi-dimensional slices
	g := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(g)

}
