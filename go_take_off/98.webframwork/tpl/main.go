/*
 * @Author: yu_xianglong yu_xianglong@qq.com
 * @Date: 2025-07-10 10:02:56
 * @LastEditors: yu_xianglong yu_xianglong@qq.com
 * @LastEditTime: 2025-07-10 14:33:46
 * @FilePath: \go_take_off.git\go_take_off\98.webframwork\tpl\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if tmpl !=nil{
		fmt.Println("解析模版失败err:",err)
	}

	name:="小猴子"
	tmpl.Execute(w,name)

}

func main() {
	http.HandleFunc("/sayhello",sayHello)

	err:=http.ListenAndServe(":9000",nil)
	if err !=nil{
		fmt.Println("启动服务失败, err:", err)
		return
	}
}