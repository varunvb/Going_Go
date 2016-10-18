package main


import (
	"fmt"
)


func worker (done chan bool){
	fmt.Print("working....")
	fmt.Println("done")
	
	done <- true
}


func main () {
		
	done := make (chan bool, 1)
	go worker (done)
	
	<-done
	
}
