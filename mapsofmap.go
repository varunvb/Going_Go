package main

import "fmt"

func main() {

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "helium",
			"state": "gas",
		},
	}

	if el, ok := elements["He"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
