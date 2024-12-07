package main

import "fmt"

//返回值
func calc(num int, num2 int) (result int, result1 int) {
	result = num + 1
	result1 = num + num2
	return
}

//形参列表
func test(args ...int) {
	for index, value := range args {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
}

//值类型传递
func testArr(val int) {
	val = 1
}

//指针传递
func testPoint(val *int) {
	*val = 2
}

func returnTest(r1 int, r2 int) (num int, num2 int) {

	num = r1 + r2
	num2 = num + 1

	fmt.Println(num)
	fmt.Println(num2)
	return
}
