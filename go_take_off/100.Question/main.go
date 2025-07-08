package main

import (
	// "first/100.Question/qassert"
	"first/100.Question/yichang"
	"fmt"
)

func main() {
 r,err:=	yichang.SafeDivide(1,0)
 fmt.Println("结果为",r,err)
}

