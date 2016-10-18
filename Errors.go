package main


import (
	"fmt"
	"errors"
)




func f1 (arg int) (int, error) {
	if arg == 43 {
			return -1, errors.New("cannot work with 43")
	}
	return arg + 3, nil 
}

type argError struct {
	arg int
	prob string
}

func (f *argError) Error() string {
	return fmt.Sprintf ("%d--%s", f.arg, f.prob)
}

func f2(arg int) (int, error) {
	if arg == 43 {
			return -1, &argError{arg, "Can't work with it!"}
	}
	return arg + 3, nil
}

func main () {

	for _, i := range [] int {7, 43} {
	   if r, e := f1 (i); e != nil {
		fmt.Println ("f1 failed: ", e)
	   } else {
		fmt.Println ("f1 worked: ", r)
	   }
	}
	for _, i := range [] int {7, 43} {
		if r, e := f2(i); e  != nil {
			fmt.Println ("f2 failed", e)
		} else {
			fmt.Println ("f2 worked", r)
		}
        }
	_, e := f2(43) 
	fmt.Println ("this is e: ", e)
	fmt.Println("this is function", e.(*argError))
	if ae, ok := e.(*argError); ok {
	       fmt.Println ("this is ae: ", ae)
	       fmt.Println (ae.arg)
	       fmt.Println (ae.prob)
	}
}
 

