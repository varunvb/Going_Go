package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	f := createFile("/tmp/defer.txt")

	defer closeFile(f)
	done := make(chan bool, 1)
	writeFile(f, done)
	<-done
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, w chan bool) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
	time.Sleep(time.Second)
	w <- true
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
