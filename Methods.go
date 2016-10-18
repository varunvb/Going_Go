package main


import (
	"fmt"
)


type rect struct {
	width, height int
}

func (r *rect) area () int {
	return r.height*r.width
}

func (r rect) perim () int {
	return 2*r.height+2*r.width
}

func main () {
	rectangle := rect{width: 10, height: 30}
	fmt.Println("area is ", rectangle.area())
	fmt.Printf("perimeter is %d\n", rectangle.perim())
	
	rp := &rectangle
	s := fmt.Sprintf("area is %d", rp.area())
	fmt.Println(s)
	fmt.Sprintf("perim is %d", rp.perim())
} 
