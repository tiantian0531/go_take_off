/*
 * @Author: yu_xianglong yu_xianglong@qq.com
 * @Date: 2024-12-09 09:52:17
 * @LastEditors: yu_xianglong yu_xianglong@qq.com
 * @LastEditTime: 2025-07-07 17:11:07
 * @FilePath: \go_take_off.git\go_take_off\11.Oop_First\oophelper\oophelper.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package oophelper

import (
	"fmt"
    "reflect"
)

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

func CreateInstance(){

	var p Person
	p.Name ="第一种方式"
	p.Age =12

	fmt.Println("第一种方式",p.Name)
	fmt.Println("第一种方式",p.Age)
    fmt.Println("第一种方式类型为", reflect.TypeOf(p))

	p2:=Person{"sa",20}
	fmt.Println("第二种方式",p2)
    fmt.Println("第二种方式类型为", reflect.TypeOf(p2))


	p3:=Person{Name: "第三种",Age: 90}
	fmt.Println("第三种方式",p3)
    fmt.Println("第三种方式类型为", reflect.TypeOf(p3))

	p4:=new(Person)
	p4.Name ="sa"
	fmt.Println("第四种方式",p4)
    fmt.Println("第四种方式类型为", reflect.TypeOf(p4))

	p5:=&Person{Name: "Frank", Age: 40}
	fmt.Println("第五种方式",p5)
		    fmt.Println("第五种方式类型为", reflect.TypeOf(p5))
}