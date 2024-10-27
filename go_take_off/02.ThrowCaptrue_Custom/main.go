package main

import (
	"errors"
	"fmt"
)

func main() {
	err := calculateSum()
	if err != nil {
		fmt.Println("自定义错误：", err)
		panic(err)
	}
	fmt.Println("计算方法执行成功")
	fmt.Println("正常执行下面的逻辑")
}

func calculateSum() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		//抛出自定义错误类型
		return errors.New("除数不能为0")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}
