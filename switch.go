package main


import "fmt"
import "time"


func main () {
	i := 2
	fmt.Print("write  ", i, " as  ")
	switch i {
	         case 1:
		    fmt.Println ("one")
		 case 2:
		    fmt.Println ("two")	
		 case 3:
		    fmt.Println ("three")
	}
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's a weekend")
		default:
			fmt.Println ("It's a weekday")
	}
	t := time.Now()	
	fmt.Println (t, "is the time")
	switch {
		case t.Hour() < 12:
		    fmt.Println("It's before Noon")
		default:
		    fmt.Println ("It's afternoon")
	}
}	
