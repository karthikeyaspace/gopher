package main

import (
	"fmt"
	"time"
	"errors"
)

type Person struct {
	fn string
	ln string
	age int
}

func main() {
	//go from chatgpt


	fmt.Println("Hello world")
	
	// variables

	var a int = 10
	b := 20
	const PI = 3.14


	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(PI)
	

	// control flow

	age := 19
	if age>=18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("child")
	}
	
	// switch

	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Monday")
	case "Sunday":
		fmt.Println("Funday")
	default:
		fmt.Println("Regualr boring day!")

	}
	

	// function
	fmt.Println(add(4, 5))
	
	// variadic function
	fmt.Println(sum(1, 2, 3, 4))



	// pointers

	x := 420
	p := &x
	fmt.Println("x: ", x)
	fmt.Println("p: ", p)
	fmt.Println("P: ", *p)

	// structs

	person := Person{"Karthikeya", "Veruturi", 19}
	fmt.Println(person)
	fmt.Println(person.age)



	// go routines
	go printNums()
	fmt.Println("Done!")
	time.Sleep(3 * time.Second)



	// error handling
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Resumt: ", res)
	}
}


func add(x, y int) int {
	return x+y
}

func sum(nums ...int) int {
	total := 0
	for _,num := range nums {
		total += num
	}
	return total
}

func printNums() {
	for i := 1; i<=5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}

}

func divide(a, b int) (int, error) {
	if b==0 {
		return 0, errors.New("Cannont divide with zero")
	}
	return a/b, nil
}

