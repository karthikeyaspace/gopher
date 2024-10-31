package main

/*

// Channels and Concurrency

- Synchronous/Sequential Programming - Code is executed line by line, everything si happening in a order - this is good but not efficient to run large programs

- cpu has something called clock speed ins/secs 1.8 GHz - 1.8 billion instructions per second
1 instruction = 1 operation

- Synchronous programming does not use the full potential and advantage of the CPU

- Concurrency - multiple things happening at the same time following the rules of concurrency
rules of concurrency -
1. Share Memory by Communicating
2. Don't communicate by sharing memory

- Concurrency in Go allows the use of full potential of the CPU cores and threads

- Channels - are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine


- Go has a built-in feature called goroutines - lightweight threads managed by the Go runtime

- Goroutines can be thought of as light weight threads. The cost of creating a goroutine is tiny when compared to a thread. Hence its common for Go applications to have thousands of goroutines running concurrently

*/

import (
	"fmt"
	"time"
)

func concurrency() {
	// Create a channel
	c := make(chan int)

	// Launch a goroutine
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// Receive values from the channel
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	// Close the channel
	close(c)

	// Wait for a second
	time.Sleep(time.Second)
}

// Email sending service

func addEmailsToQueue(emails []string) chan string {
	emailQ := make(chan string, len(emails))
	for _, email := range emails {
		emailQ <- email
	}
	close(emailQ)
	return emailQ
}

func sendEmail(batchSize int, emailQ chan string) {
	for i := 0; i < batchSize; i++ {
		email, ok := <-emailQ
		if !ok {
			return
		}
		fmt.Println("Sending email to", email)
	}
}

func emailService(emails ...string) {
	emailQ := addEmailsToQueue(emails)
	go sendEmail(3, emailQ)
	go sendEmail(3, emailQ)
	time.Sleep(time.Second)
}

// Logger service

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				chEmails = nil
			}
			logEmail(email)
		case sms, ok := <-chSms:
			if !ok {
				chSms = nil
			}
			logSms(sms)
		}
	}
}
func logEmail(email string) {
	fmt.Println("Logging email", email)
}

func logSms(sms string) {
	fmt.Println("Logging sms", sms)
}

func loggerService() {
	chEmails := make(chan string)
	chSms := make(chan string)

	go logMessages(chEmails, chSms)

	chEmails <- "email 1"
	chSms <- "sms 1"
	chEmails <- "email 2"
	chSms <- "sms 2"

	close(chEmails)
	close(chSms)

	time.Sleep(time.Second)
}
