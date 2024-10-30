package main

import (
	"encoding/json"
	"fmt"
)

type DivisionError struct {
	Divisor int
}

type Person2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func makeMultiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("cannot divide by %d", e.Divisor)
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, &DivisionError{Divisor: y}
	}
	return x / y, nil
}

func main() {
	// Higher-order function - returning functions
	double := makeMultiplier(2)
	fmt.Println(double(3)) // Output: 6

	// Custom errors
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// JSON
	p := Person2{"Karthikeya", 19}

	// Marshal
	b, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(b))

	// Unmarshal
	var p2 Person2
	err = json.Unmarshal(b, &p2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(p2)

}
