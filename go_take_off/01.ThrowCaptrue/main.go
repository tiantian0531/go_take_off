package main

import "fmt"
 
func main() {
	calculateSum()
	fmt.Println("计算方法执行成功")
	fmt.Println("正常执行下面的逻辑")
}

func calculateSum() {
	//defer+ recover捕获
	defer capture()
	num1 := 10
	num2 := 0
	// defer fmt.Println("延迟执行")
	result := num1 / num2
	fmt.Println(result)
}

func capture() {
	err := recover()
	//如果没有捕获错误，recover不是返回零值：nil
	if err != nil {
		fmt.Println("错误已经捕获")
		fmt.Println("err 是", err)
	}
}
