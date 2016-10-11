package main

import "fmt"


func vals () (int, int) {
	return 3, 7
}
func main () {
	
	fmt.Println ("print return values")
	a, b := vals ()
	fmt.Println ("a is=", a)
	fmt.Println ("b is=", b)
	
	_, c := vals()
	fmt.Println (c)

}	
