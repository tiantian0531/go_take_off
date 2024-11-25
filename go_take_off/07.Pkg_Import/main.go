package main

import (
	calcs "first/07.Pkg_Import/calcutils"
	cc "first/07.Pkg_Import/dbutils" //包取别名
	"fmt"
)

func main() {

	calcs.Calc()
	cc.GetConn()
	fmt.Println("这是main函数执行")
}
