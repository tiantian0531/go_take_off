package main

import (
	"fmt"
	"reflect"
)

func main() {
	T3ReflectValue3()
}
func T1() {
	var c *C = new(C)
	c.Name = "管理员"
	c.Age = 18
	c.EnName = "sa"
	c.Hello()
	fmt.Println(*c)
	fmt.Println("=================")
	var b *B = new(B)

	b.Name = "管理员"
	b.Age = 19
	b.NickName = "黄光头"
	b.Hello()
	fmt.Println(*b)
}

func T2jk() {
	var a *Cat = new(Cat)
	a.Name = "jeck"
	a.SayHello()

	var c *Dog = new(Dog)
	c.Name = "天天"
	c.SayHello()
}

func T3ReflectValue() {
	b := 3.4
	var t1 reflect.Type = reflect.TypeOf(b)
	fmt.Println(t1)        //打印t1的类型
	fmt.Println(t1.Kind()) //打印t1的类型

	fmt.Println("==============")

	var t2 reflect.Value = reflect.ValueOf(&b)
	if t2.IsValid() {
		fmt.Println(t2.Kind()) //打印t2的类型
		fmt.Println(t2)        //打印t2的值
		fmt.Println(t2.Elem()) //打印t2指针对应的值
		t2.Elem().SetFloat(1.9)
		fmt.Println(b)
	}
}

func T3ReflectValue2() {
	i := &A{"SA", 18}
	fmt.Println(i)

	v := reflect.ValueOf(i)
	fmt.Println(v)
	v = v.Elem()
	fmt.Println(v)

	f := v.FieldByName("Name")
	f1 := v.FieldByName("Age")
	fmt.Println(f)

	if f.Kind() == reflect.String {
		f.SetString("奥利奥")
	}
	if f1.Kind() == reflect.Int {
		f1.SetInt(1)
	}

	fmt.Println(v)
}

func T3ReflectValue3() {

	i := &A{"SA", 18}
	//fmt.Println(i)
	//
	v := reflect.ValueOf(i)
	//fmt.Println(v)
	//v = v.Elem()
	//fmt.Println(v)
	//
	//fields := v.NumField()
	//
	//for i := 0; i < fields; i++ {
	//	f := v.Field(i)
	//	if f.CanSet() {
	//		switch f.Kind() {
	//		case reflect.String:
	//			f.SetString("奥利奥++")
	//		case reflect.Int:
	//			f.SetInt(1)
	//		}
	//	}
	//}

	methods := v.NumMethod()
	F1 := v.MethodByName("Hello")
	for i := 0; i < methods; i++ {
		m := v.Method(i)
		if m.IsValid() {
			m.Call([]reflect.Value{})
		}
	}

	fmt.Println(*i)
	fmt.Println(F1)
}
