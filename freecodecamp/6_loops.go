package main

import (
	"fmt"
)

// Loops in go

func loop() {
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
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}

	// loop through array
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// loop through map
	m := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// loop through string
	str := "hello"
	for index, value := range str {
		fmt.Println(index, string(value))
	}

	// break

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// continue

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// nested loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}
	}

}
