package dbutils

import "fmt"

//首字母大写可以被其他包访问
func GetConn() {
	fmt.Println("执行了包下面的GetConn函数")
}
