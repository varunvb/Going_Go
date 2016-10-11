package main 


import "fmt"


func zeroval (i int) int {
	i = 0
	return i
}

func zeroptr (iptr *int) int {
	*iptr = 0
	return *iptr
}

func main () {
	i := 1
	fmt.Println ("Intial = ", i)
	zeroval (i)
	fmt.Println ("Zeroval = ", i)
	
	zeroptr (&i)
	fmt.Println ("Zero pointer = ", i)
	fmt.Println ("Pointer = ", &i)
}
