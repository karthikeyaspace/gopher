package main

type car struct {
	make string
	model string
	year int
	price float64
}
// car1 := car{"Toyota", "Corolla", 2020, 15000.00}
// car2 := car{"Ford", "F150", 2020, 25000.00} 



// struct key can hold any type 

// we can also nest structs
type person struct {
	name string
	age int
	car car
}
// person1 := person{"John", 25, car1}

// think of structs in go as typescript in steroids

type messageToSend struct {
	message string
	to user
	from user
}

type user struct {
	name string
	mail string
	alias string
	number int
}




// nesting structs - anonymouse structs

type messageToSend2 struct {
	message string
	to struct {
		name string
		mail string
		alias string
		number int
	}
	from struct {
		name string
		mail string
		alias string
		number int
	}
}
 

// go is not an object oriented language but it has some object oriented features

type truck struct {
	car
	wheels int
}

// all the fields of car are now available in truck as top level fields

// eg
// truck1 := truck{
// 	car: car{
// 		make: "Toyota",
// 		model: "Corolla",
// 		year: 2020,
// 		price: 15000.00,
// 	},
// 	wheels: 4,
// }

// truck1.make // Toyota

 
