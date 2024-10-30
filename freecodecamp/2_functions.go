package main


func sub(x int, y int) int {
	return x - y
}

// here "func sub(x int, y int) int" is the signature of the function


// functions can also return multiple values
func swap(x, y string) (string, string) {
	return y, x
}


func getValues() (x, y int) {
	// x and y are already declared in the signature with value 0
	// reassigning the values of x and y
	x = 5
	y = 7 
	
	return // automatically returns x and y since we specified them in the signature
}

// better wat
func getValues2() (int, int) {
	x := 5
	y := 7
	return x, y
}

// functions can also return named values
func getValues3() (x int, y int) {
	x = 5
	y = 7
	return
}


// functions can take a variable number of arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}


// functions can also return a function - higher order functions
func getAdder() func(int) int {
	return func(x int) int {
		return x + 5
	}
}
// sum := getAdder()
// sum(3) // 8



// functions can also be passed as arguments to other functions
func apply(f func(int, int) int, a, b int) int {
	return f(a, b)
}
// apply(sub, 5, 3) // 2


// functions can also be used as a type
type mathOp func(int, int) int

func apply2(f mathOp, a, b int) int {
	return f(a, b)
}

