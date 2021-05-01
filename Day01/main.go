package main

import (
	"fmt"

	"github.com/parezban/100-days-of-code/Day01/helper"
)

func main() {
	helper.SayHello("Darling")

	fmt.Println("From the main package ;)")

	fmt.Println(helper.AddItForMe(2, 5))

	fmt.Println(helper.DoSomethingOnIt(4, 3))
}
