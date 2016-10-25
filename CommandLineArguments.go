package main

import (
	"fmt"
	"os"
)

func main() {
	p := fmt.Println
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	p(argsWithProg)
	p(argsWithoutProg)
	p(arg)
}
