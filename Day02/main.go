package main

import "fmt"

func main() {
	i := 0
	for i < 300 {
		fmt.Println("Hey")
		i++
	}

	cousyCounter := 0
	for {
		fmt.Println("Hey, This one is infinite loop!")
		if cousyCounter == 2000 {
			fmt.Println("No, It wasn't")
			break
		}

		cousyCounter++
	}

	defer fmt.Println(noReasonMethod(400))
	fmt.Println(noReasonMethod(21))

	switch cousyCounter {
	case 2000:
		fmt.Println("2000, Great!")

	default:
		fmt.Println("This one not print at all")

	}

	switch {
	case cousyCounter > 1900:
		fmt.Println("2000, Great! This will work")

	case cousyCounter > 1500:
		fmt.Println("1500, But not work sad For me:(")

	case cousyCounter > 1000:
		fmt.Println("1000, But not work too:(")
	}
}

func noReasonMethod(x int) int {
	if faky := 23; x == 400 {
		return faky
	} else {
		return faky + 2
	}
}
