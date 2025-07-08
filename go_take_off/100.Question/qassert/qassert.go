/*
 * @Author: yu_xianglong yu_xianglong@qq.com
 * @Date: 2025-07-08 09:31:55
 * @LastEditors: yu_xianglong yu_xianglong@qq.com
 * @LastEditTime: 2025-07-08 10:10:56
 * @FilePath: \go_take_off.git\go_take_off\100.Question\qassert\qassert.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package qassert

import "fmt"

type MyStruct struct{
	Field int
}

func ValAssert(){
	var x interface{} =10
	y,ok:=x.(int)

	if ok{
		fmt.Println("x断言值类型为",x)
		fmt.Println("y断言值类型为",y)
		y=20
		fmt.Println("修改y之后,x为",x)
	}

	if !ok{
		fmt.Println("断言失败")
	}
}

func RefAssert(){

	var x interface{} = &MyStruct{18}

	
if y,ok:= x.(*MyStruct);ok{
	fmt.Printf("Address of y: %p\n", &y)
	fmt.Printf("Address of x: %p\n", &x)
		fmt.Println("x断言值类型为",x)
		fmt.Println("y断言值类型为",y)
		y.Field =89
		fmt.Println("修改y之后,x为",x)
			fmt.Println("修改y之后,y",y)
	}else{
		fmt.Println("断言失败")
	}

}