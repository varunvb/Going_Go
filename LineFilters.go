package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		result := strings.ToUpper(scanner.Text())
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}
