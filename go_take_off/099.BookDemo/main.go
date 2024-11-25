package main

import (
	"fmt"
)

func main() {
	fmt.Println("2121")
	// echo1Tets()
	a, b := returnTest(2, 3)
	fmt.Println(a)
	fmt.Println(b)
}
func returnTest(r1 int, r2 int) (num int, num2 int) {

	num = r1 + r2
	num2 = num + 1

	fmt.Println(num)
	fmt.Println(num2)
	return
}
