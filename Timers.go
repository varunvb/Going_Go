package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer1 Expired")
	timer2 := time.NewTimer(time.Second * 3)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
}
