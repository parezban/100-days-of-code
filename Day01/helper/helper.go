package helper

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello,", name)
}

func AddItForMe(x, y int) int {
	return x + y
}

func DoSomethingOnIt(a, b int) (x, y int) {
	x = a + b
	y = a * b

	return
}
