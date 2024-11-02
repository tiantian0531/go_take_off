package main

import "fmt"

func main() {
	var scores [5]int
	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21
	//声明未赋值，所有值为0
	for i := 0; i < len(scores); i++ {
		fmt.Printf("%v", scores[i])
	}
	//数组变量的地址和第一个下标的内存地址一致
	fmt.Println()
	fmt.Printf("数组的地址是：%p", &scores)
	fmt.Println()
	fmt.Printf("数组第一个空间的地址是：%p", &scores[0])

	//数组第二个下标的地址 = 当前地址 + 下标*字节数（int8为1）（int64为8）
	fmt.Println()
	fmt.Printf("数组第二个空间的地址是：%p", &scores[1])
	fmt.Println()
	fmt.Printf("数组第三个空间的地址是：%p", &scores[2])

	fmt.Printf("数组第四个空间的地址是：%p", &scores[3])
	// 	00000
	// 数组的地址是：0xc000010390
	// 数组第一个空间的地址是：0xc000010390
	// 数组第二个空间的地址是：0xc000010398
	// 数组第三个空间的地址是：0xc0000103a0

}
