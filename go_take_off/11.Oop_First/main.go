package main

import "fmt"

type User struct {
	Name   string
	Age    int
	School string
}

func main() {

	// var t1 User
	t1 := &User{Name: "sa"}
	fmt.Println(t1)

	t1.Name = "张三"
	fmt.Println(t1.Name)
}
