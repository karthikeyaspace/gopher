package main

import "fmt"

// make function - The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it.

// make(map[string]int) - creates a map with string keys and int values
// make(map[string]int, 100) - initial cap of 100

type vehcile struct {
	vtype string
	price int
}

func maps() {
	// map[keyType]valueType
	// map[keyType]valueType{key: value, key: value}

	customers := map[string]string{
		"c_01": "karthikeya",
		"c_02": "elon",
		"c_03": "jeff",
		"c_04": "zuck",
	}

	printAllUsers(customers)

	// check if key exists
	if value, ok := customers["c_01"]; ok {
		fmt.Println("Customer ID: c_01 exists with name: ", value)
	}

	// make function

	products := make(map[string]int)

	products["p_01"] = 100
	products["p_02"] = 200
	products["p_03"] = 300

	fmt.Println(products)

	// delete key
	delete(products, "p_01")

	// update key
	products["p_02"] = 250

	// using structs to make nested maps

	vehicles := map[string]vehcile{
		"v_01": {vtype: "car",
			price: 10000},
		"v_02": {vtype: "bike",
			price: 2000},
		"v_03": {vtype: "cycle",
			price: 500},
	}

	for key, value := range vehicles {
		fmt.Println("Vehicle ID: ", key, "Vehicle type: ", value.vtype, "Vehicle price: ", value.price)

	}

	// check if key exists
	if value, ok := vehicles["v_01"]; ok {
		fmt.Println("Vehicle ID: v_01 exists with type: ", value.vtype, "price: ", value.price)
	}

}

func printAllUsers(customers map[string]string) {
	for key, value := range customers {
		fmt.Println("Customer ID: ", key, "Customer name: ", value)
	}
}
