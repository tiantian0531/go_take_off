package main

import "fmt"

type integer int

func (i *integer) SayHello() {
	fmt.Println("自定义数据类型实现接口", *i)
}
