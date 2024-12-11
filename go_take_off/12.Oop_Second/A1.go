package main

import "fmt"

type A struct {
	Name string
	Age  int
}

type B struct {
	A
	NickName string
}

type C struct {
	A
	EnName string
}

func (a A) Hello() {
	fmt.Println("hello")
}
