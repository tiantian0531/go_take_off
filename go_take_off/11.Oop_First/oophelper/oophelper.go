package oophelper

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Person struct {
	Name string
	Age  int
}

func OopOper() {

	var s = new(Student)
	var p = new(Person)
	s.Age = 28
	*p = Person(*s)
	s.Age = 26
	fmt.Println(*s)
	fmt.Println(*p)
}

// 按值传递
func updatePersonByValue(p Person) {
	p.Age = 30
	fmt.Println("按值传递:", p)
}

// 按引用传递
func updatePersonByReference(p *Person) {
	p.Age = 30
	fmt.Println("指针传递:", *p)
}

func TransmitInstance() {

	person := Person{Name: "Alice", Age: 25}

	// 调用按值传递的函数
	updatePersonByValue(person)
	fmt.Println("Outside 按值传递之后:", person) // 输出：Age仍为25

	// 调用按引用传递的函数
	updatePersonByReference(&person)
	fmt.Println("指针传递之后:", person) // 输出：Age已改为30
}
