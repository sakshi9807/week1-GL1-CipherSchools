package main

import "fmt"

func main() {
	var num int
	num = 10
	fmt.Println(num)
	var getInt func(int) int
	getInt = func(x int) int {

		fmt.Println("In a 1st number")
		var getInt func(int) int
		getInt = func(x int) {
			fmt.Println("In a 1st function")
			return 20 + x
		}
		g(getInt, 8)
		//	getInt=func (x int ) int {
		//		fmt.Println("In a second func")
		//		return 10+x
		//	}
		var y func()
		y = func() {
			fmt.Println(9)
		}
		y()
		j := func(x int) int {
			fmt.Println("In a 1st function")
			return 20 + x
		}(10)
		fmt.Println(j)
		
}
	func g(getInt func (int) int , u int){

			getInt(78)
	}
