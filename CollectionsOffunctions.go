package main

import (
	"fmt"
	"strings"
)

func Exist(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func BoolExist(vs []string, t string) bool {
	return Exist(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {

	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {

	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	var strs = []string{"peach", "apple", "banana", "plum"}
	fmt.Println(Exist(strs, "pear"), "this is from exist function")
	fmt.Println(BoolExist(strs, "plum"), "this is from the bool function")
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}), "this is from Any function using function predicate")
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}), "this is from All function using function predicate")
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}), "this is from Filter function using function predicate")
	fmt.Println(Map(strs, strings.ToUpper), "this is from Map function using function predicate")
}
