package main

import "os"

func main() {

	panic("a fatal error")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
