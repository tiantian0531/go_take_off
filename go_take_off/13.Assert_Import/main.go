package main

import (
	"fmt"
)

type SayHello interface {
	sayHello()
}

type Chinese struct {
	name string
}

func (person *Chinese) sayHello() {
	fmt.Printf("中国人 %s 说你好 \n", person.name)
}

func (person *Chinese) PlayMajiang() {
	fmt.Printf("中国人 %s 打麻将 \n", person.name)
}

type American struct {
	name string
}

func (person *American) sayHello() {
	fmt.Printf("美国人 %s 说你好 \n", person.name)
}

func Great(s SayHello) {

	//断言
	// var ch Chinese = s.(Chinese) //判断是否能转为Chinese
	if ch1, ok := s.(*Chinese); ok {
		ch1.PlayMajiang()
	} else {
		s.sayHello()
	}
}
func main() {

	c := Chinese{name: "小丽"}
	a := American{name: "rose"}

	Great(&c)
	Great(&a)

}
