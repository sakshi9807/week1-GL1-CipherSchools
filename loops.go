package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(i)
		}
	}

	nums := []int{12, 13, 14, 15}

	for k, v := range nums {
		fmt.Println(k)
		fmt.Println(v)

	}

	g := []int{12, 13, 14, 15}

	for _, value := range nums {
		if value == 13 {
			break
		}
		fmt.Println(value)

	}
}
