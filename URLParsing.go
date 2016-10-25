package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "http://www.example.com:4080/auth=?login.php"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
}
