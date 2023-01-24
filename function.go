package main

import "fmt"

func main() {
	w := new(int)
	name := new(string)
	t := "hjbgg"
	result, x := C(*w, &t)
	fmt.Println(result)
	fmt.Println(x)
	fmt.Println(*name)
	fmt.Println(w)
	//r = B(1, 2, 3, 4, 5, 6, 6)
	//fmt.Print(r)
}

func A() (int, string) {
	return 122, "abc"
}

func B(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func C(w int, name *string) (i int, j string) {
	i = 10
	j = "rahul"
	w = 100
	*name = "code"
	return
}
