package oinstance_test

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) Print() {
	fmt.Println("调用了Print方法")
	fmt.Println("学生的名字是", s.Name)
}

func (s Student) GetSumn(n1, n2 int) int {
	return n1 + n2
}

func (s *Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

func OperStundentStruct(a interface{}) {
	val := reflect.ValueOf(a)
	fmt.Println("val的值的", val)

	//通过val.Value类型操作结构体内部字段
	typ := reflect.TypeOf(a)
	fields := val.Elem().NumField()

	for i := 0; i < fields; i++ {
		fmt.Printf("第%d个字段名为%v,值为 %v;\n", i, typ.Elem().Field(i).Name, val.Elem().Field(i))
	}

	methods := val.Elem().NumMethod()

	for i := 0; i <= methods; i++ {
		fmt.Printf("第%d个方法是 %v;\n", i, typ.Method(i).Name)
		fmt.Printf("第%d个方法是 %v,参数有：%v \n", i, typ.Method(i).Name, typ.Method(i).Type.NumIn())

	}

	var params []reflect.Value

	params = append(params, reflect.ValueOf("Q"))
	params = append(params, reflect.ValueOf(20))
	val.Method(2).Call(params)

	fmt.Println("", val.Interface().(*Student).Name)
	//通过val.Value类型操作结构体内部方法

}

func OperStundentSetStruct(a interface{}) {
	// val := reflect.ValueOf(a)
	// fmt.Println("val的值的", val)

	// //通过val.Value类型操作结构体内部字段
	// typ := reflect.TypeOf(a)
	// fields := val.Elem().NumField()

	// for i := 0; i < fields; i++ {
	// 	fmt.Printf("第%d个字段名为%v,值为 %v;\n", i, typ.Elem().Field(i).Name, val.Elem().Field(i))

	// }
	// val.Elem().Field(0).SetString()
	// methods := val.Elem().NumMethod()

	// for i := 0; i <= methods; i++ {
	// 	fmt.Printf("第%d个方法是 %v;\n", i, typ.Method(i).Name)
	// 	fmt.Printf("第%d个方法是 %v,参数有：%v \n", i, typ.Method(i).Name, typ.Method(i).Type.NumIn())

	// }

	// var params []reflect.Value

	// params = append(params, reflect.ValueOf("Q"))
	// params = append(params, reflect.ValueOf(20))
	// val.Method(2).Call(params)

	// fmt.Println("", val.Interface().(*Student).Name)
	//通过val.Value类型操作结构体内部方法

}
