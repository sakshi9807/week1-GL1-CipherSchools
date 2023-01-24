package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var input int
	fmt.Scanln(&input)
	if input%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if x := 10; x%2 == 0 {
		fmt.Println("hi")
	}
	m := 100
	switch m {
	case 1:
		m = 1
		fmt.Println(m)
	case 2:
		m = 10
		fmt.Println(m)
	case 100:
		m = 20
		fmt.Println(m)
		fallthrough
	case 20:
		m = 200
		fmt.Println(m)
	}
}
