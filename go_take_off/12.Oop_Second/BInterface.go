package main

import "fmt"

type Animal interface {
	SayHello()
}

type Cat struct {
	Name string
}
type Dog struct {
	Name string
}

func (c Cat) SayHello() {

	fmt.Println("喵喵喵")
}
func (d Dog) SayHello() {
	fmt.Println("wangwangwang")
}
