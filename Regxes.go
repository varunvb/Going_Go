package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, b := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	fmt.Println("this is b ", b)

	r, c := regexp.Compile("p(a-z)+)ch")

	fmt.Println(c)

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("beach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
