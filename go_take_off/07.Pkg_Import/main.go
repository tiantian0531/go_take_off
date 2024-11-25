package main

import (
	calcs "first/07.Pkg_Import/calcutils"
	"first/07.Pkg_Import/dbutils"
	"fmt"
)

func main() {

	calcs.Calc()
	dbutils.GetConn()
	fmt.Println("这是main函数执行")
}
