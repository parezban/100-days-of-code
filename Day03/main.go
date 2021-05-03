package main

import "fmt"

type Vertext struct {
	Height int
	Name   string
}

func main() {

	test := Vertext{Height: 54, Name: "Arian"}
	p := &test
	np := test

	p.Height = 31
	np.Height = 32

	fmt.Println(test, *p, p, np)
}
