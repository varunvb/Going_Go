package main

import (
	"fmt"
)


func main () {
	
	messages := make (chan string, 1)
	signals := make (chan bool)

	msg := "test"
	messages <- msg
	
	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	default:
		fmt.Println ("No message received")
	}

	msg1 := "hi"
	select {
	case messages <-msg1:
		fmt.Println ("Sent message", msg1)
	default:
		fmt.Println ("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println ("received message", msg)
	case sig := <-signals:
		fmt.Println ("received signal", sig)
	default:
		fmt.Println ("no activity")
	}
}
