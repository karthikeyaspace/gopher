package main

import (
	"errors"
	"fmt"
)

// Advanced Functions
// Functions can be passed as arguments to other functions.
// Functions can return other functions.
// Functions as data.

// Higher Order Functions
// Functions that take other functions as arguments or return functions are called higher order functions.

// create a custom logger

func getLogger(formatter func(string ,string) string) func (string, string) {
	return func (msg string, level string) {
		fmt.Println(formatter(msg, level))
	}
}

func advfunctions() {

	dbErrors := []error{
		errors.New("connection error"),
		errors.New("query error"),
		errors.New("timeout error"),
		errors.New("out of memory"),
	}

	serverErrors := []error{
		errors.New("404 not found"),
		errors.New("500 internal server error"),
		errors.New("503 service unavailable"),
	}

	// create a logger that logs errors to the console

	consoleLogger := getLogger(func(msg string, level string) string {
		return fmt.Sprintf("[%s] %s", level, msg)
	})



	// log database errors

	for _, err := range dbErrors {
		consoleLogger(err.Error(), "ERROR")
	}

	// log server errors

	for _, err := range serverErrors {
		consoleLogger(err.Error(), "ERROR")
	}



	//Defer Keyword
	// The defer keyword is used to delay the execution of a function until the surrounding function returns.
	// Defer is commonly used to ensure that resources are released after a function has completed.




}