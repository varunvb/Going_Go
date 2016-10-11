package main


import "fmt"


func main () {
	if 7%2 == 0 {
		fmt.Println ("7 is even")
	} else {
		fmt.Println ("7 is Odd")
	}


	if 8%4 == 0 {
		fmt.Println ("8 is divisible by 4")
	} else {
		fmt.Println ("8 isn't divisible by 4")
	}



	if num := 9; num < 0 {
		fmt.Println ("num is negative")
	} else if num < 10 {
		fmt.Println ("has one digit")
	} else {
		fmt.Println (num, "has multiple digits")
	}
}
