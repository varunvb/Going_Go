package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "Processing job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 110)
	result := make(chan int, 110)

	for w := 0; w <= 110; w++ {
		go worker(w, jobs, result)
	}
	for j := 1; j <= 110; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 110; a++ {
		fmt.Println("result from j is", <-result)
		fmt.Println("a is ", a)
	}
}
