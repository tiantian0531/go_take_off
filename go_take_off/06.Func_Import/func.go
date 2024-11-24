package main

import "fmt"

func calc(num int, num2 int) (int, int) {
	return num + 1, num + num2
}

func test(args ...int) {
	for index, value := range args {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
}

func testArr(val int) {
	val = 1
}

func testPoint(val *int) {
	*val = 2
}
