package main

import (
	oinstance_test "first/16.Reflect/OInstance"
	"fmt"
	"reflect"
)

func main() {

}

func Reflect01() {
	var i int = 2

	//获取类型
	var t = reflect.TypeOf(i)
	fmt.Printf("%T", t)

	//过去值类型
	v := reflect.ValueOf(i)
	fmt.Printf("%T \n", v)

	p := v.Interface()
	n := p.(int)

	fmt.Println("最终值为", n+5)

	newV := 3
	v1 := reflect.ValueOf(&i)

	v1.Elem().SetInt(int64(newV))

	fmt.Println("...", i)
}

func reflect02() {
	stu := oinstance_test.Student{
		Name: "小米",
		Age:  18,
	}

	fmt.Println("", reflect.TypeOf(stu))
	fmt.Println("", reflect.ValueOf(stu))
	fmt.Println("", reflect.TypeOf(stu).Kind())
	fmt.Println("", reflect.ValueOf(stu).Kind())

	v := reflect.ValueOf(stu).Interface()
	a := v.(oinstance_test.Student)
	fmt.Println("", a)
}

func reflect03() {
	stu := &oinstance_test.Student{
		Name: "小米1",
		Age:  20,
	}

	oinstance_test.OperStundentStruct(stu)
}
