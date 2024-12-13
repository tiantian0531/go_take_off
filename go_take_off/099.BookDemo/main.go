package main

import (
	"fmt"
)

type integer int

func main() {

	var i integer = 20
	i.print()
	fmt.Println("\n =================")
	fmt.Println(i)

	var c integer
	fmt.Println(i)
	c.print1()
	fmt.Println("\n =================")
	fmt.Println(i)
}

func (i *integer) print() {
	*i = 21
	fmt.Println(*i)
}

func (i integer) print1() {
	i = 22
	fmt.Println(i)
}

// func (i integer) String() string {
// 	fmt.Println("12306")
// 	return "2566"
// }
