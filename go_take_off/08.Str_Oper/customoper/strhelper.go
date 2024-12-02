package customoper

import (
	"fmt"
	"strings"
	"unicode"
)

func StrHelperOper() {
	var str string = " 1231 456 你好世界 "
	// fmt.Print("长度是：", len(str))

	//查找指定字符在给定的字符串中出现的次数
	fmt.Println(strings.Count(str, "00"))

	//不区分大小写比较
	fmt.Println(strings.EqualFold("hello", "hELLO"))
	//区分大小写比较
	fmt.Println("hello" == "hELLO")

	//返回指定字符第一次出现的位置，没有就是-1
	fmt.Println(strings.Index(str, "0"))

	fmt.Println(strings.HasSuffix(str, "界"))

	//去掉前后空格
	fmt.Println(strings.TrimSpace(str))

	//去掉所有的空格，按照空格切分然后连接
	fmt.Println(strings.Join(strings.FieldsFunc(str, unicode.IsSpace), ""))

}
