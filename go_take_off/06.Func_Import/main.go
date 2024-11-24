package main

import "fmt"

func main() {
	//var str string = "12345是"
	//
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("当前值是：%d \n", str[i])
	//}
	//
	//for index, value := range str {
	//	fmt.Printf("索引为%d,值是%v\n", index, value)
	//}
	//	var _, v2 = calc(1, 3)

	//可变参数
	//test(1)
	//fmt.Println("++++++++++++++++++")
	//test(1, 2, 3)
	//fmt.Println("++++++++++++++++++")
	//test(4, 5, 6)
	//fmt.Println("++++++++++++++++++")
	//test(7, 8, 9, 999)
	//fmt.Println("Hello World")

	var num int = 10
	fmt.Println(num)
	testArr(num)
	fmt.Println(num)

	testPoint(&num)
	fmt.Println(num)
}
