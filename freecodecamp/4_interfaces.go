package main

import (
	"io"
	"os"
)

// Interfaces
// An interface is a collection of method signatures that a type can implement.


type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// in go, we don't need to explicitly say that a type implements an interface
// as long as the type has the methods defined in the interface, it is said to implement the interface


// Eg
type employee interface {
	getName() string
	getSalary() float64
}

type contractor struct {
	name string
	hourlyRate float64
	totalHours float64
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() float64 {
	return c.hourlyRate * c.totalHours
}

// emp := contractor{name: "John", hourlyRate: 10, totalHours: 70}
// fmt.Println(emp.getName())
// fmt.Println(emp.getSalary())




// Multiple interfaces

type expense interface {
	cost() float64
}

type sender interface {
	send()
}

type email struct {
	isSub bool
	body string
}

func (e email) cost() float64 {
	if !e.isSub {
		return 0.05 * float64(len(e.body))
	}
	return 0.10 * float64(len(e.body))
}

func (e email) send() {
	// fmt.Println("Sending email, body: ", e.body)
}



// naming the arguments in method in interface - for readability

type Copier interface {
	Copy(srcFl string, destFl string) (bytesWritten int, err error)
}



type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}