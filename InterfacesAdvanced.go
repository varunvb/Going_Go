package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h *Human) SayHi() {
	fmt.Println("Hi I'm %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human) Sing(lyrics string) {
	fmt.Println("Lalalalalalala", lyrics)
}

//func (h *Human) Guzzle(beerStein string) {
//	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
//}

func (s *Student) SayHi() {
	name := s.name
	school := s.school
	phone := s.phone
	fmt.Printf("Hi I'm %s, I study at %s. Call me on %s\n", name, school, phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi I'm %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	//Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
	//Guzzle(beerStein string)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {

	mike := &Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := &Student{Human{"Paul", 26, "222-222-YYY"}, "Harvard", 100}
	sam := &Employee{Human{"Sam", 40, "444-222-XXX"}, "Golang Inc", 1000}
	tom := &Employee{Human{"Thomas", 37, "666-555-XXX"}, "Things Ltd", 5000}

	var i Men

	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November Rain")

	i = tom
	fmt.Println("This is Tom, en Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}
