package main

import "fmt"

type CInterface interface {
	c()
}

type BInterface interface {
	b()
}

type AInterface interface {
	BInterface
	CInterface
	a()
}

type Student struct {
}

func (s *Student) a() {
	fmt.Println("实现a")
}
func (s *Student) b() {
	fmt.Println("实现b")
}
func (s *Student) c() {
	fmt.Println("实现c")
}
