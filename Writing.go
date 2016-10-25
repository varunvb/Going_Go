package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	p := fmt.Printf
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	d2 := []byte{112, 111, 189, 119, 20}
	n2, err := f.Write(d2)
	check(err)
	p("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	p("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	p("wrote %d bytes\n", n4)

	w.Flush()

}
