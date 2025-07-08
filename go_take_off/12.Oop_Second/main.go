/*
 * @Author: yu_xianglong yu_xianglong@qq.com
 * @Date: 2024-12-17 10:08:30
 * @LastEditors: yu_xianglong yu_xianglong@qq.com
 * @LastEditTime: 2025-07-08 09:19:10
 * @FilePath: \go_take_off.git\go_take_off\12.Oop_Second\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"encoding/json"
	"first/12.Oop_Second/model"
	"fmt"
	"reflect"
)

func main() {

 user:=model.NewUser("sa",18)

 user.GetInfo()
fmt.Printf("信息是 %v",user)

	t1 := B{
		A:A{Name: "xz", Age: 12},
		NickName: "hha",
	}

	 js,_:= json.MarshalIndent(t1, "", "  ")

 	fmt.Println("json",js)
	// // 输出 JSON 字符串
 	fmt.Println(string(js))
	//多态
	 var dog *Dog = new(Dog)
	 Great(dog)

	 var iValue integer = 10
	 var i Animal = &iValue
	 i.SayHello()

	 T3ReflectValue3()

	var arr [3]Animal
	arr[0] = Dog{Name: "天天"}
	arr[1] = Cat{Name: "杰克"}
	arr[2] = Dog{Name: "阿奇"}

	fmt.Println(arr)
	for _, v := range arr {
		if v != nil {
			v.SayHello()
		}
	}
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

	// b.Name = "管理员"
	// b.Age = 19
	// b.NickName = "黄光头"
	// b.Hello()
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

func Great(i Animal) {
	i.SayHello()
}
