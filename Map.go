package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 8
	fmt.Println("map ", m)

	v1 := m["k1"]
	fmt.Println("len: ", len(m))
	fmt.Println("value of k1 (v1) is", v1)

	delete(m, "k2")
	fmt.Println("map after deletion ", m)
	e, prs := m["k2"]
	fmt.Println("prs", prs)
	fmt.Println("empty ", e)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n ", n)
}
