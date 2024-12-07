package slicehelper

import "fmt"

func SliceOper() {

	//定义数组
	var intarr [6]int = [6]int{3, 5, 6, 7, 88, 7}
	//切片建立在数组之上，从索引1开始切到3，不包含索引3
	var slice []int = intarr[1:3]

	fmt.Println("slice的长度是：", len(slice))
	fmt.Println("slice的容量是：", cap(slice))
	fmt.Println("\n ===============================================================================================================================================================================")
	fmt.Printf("切片的地址是：%p", &slice[0])
	fmt.Printf("数组的地址是：%p", &intarr[1])

	fmt.Println("\n ===============================================================================================================================================================================")
	fmt.Println("intarr:", intarr)
	fmt.Println("slice:", slice)
	slice[1] = 99

	fmt.Println("intarr:", intarr)
	fmt.Println("slice:", slice)
}

func SliceDefine() {
	//数组创建切片
	var intarr [6]int = [6]int{7, 8, 9, 3, 2, 1}
	slice1 := intarr[1:3]
	fmt.Println("slice1:", slice1)
	fmt.Println("\n ===============================================================================================================================================================================")
	//使用make创建切片
	slice2 := make([]int, 4, 20)
	slice2[2] = 66
	fmt.Println("slice2:", slice2)
	fmt.Println("slice2", len(slice2))
	fmt.Println("slice2", cap(slice2))
	fmt.Println("\n ===============================================================================================================================================================================")
	//定义
	slice3 := []int{2, 4, 6}
	fmt.Println("slice3:", slice3)
	fmt.Println("slice3", len(slice3))
	fmt.Println("slice3", cap(slice3))
}

func SliceIteRate() {

	//使用make创建切片
	slice2 := make([]int, 4, 20)
	slice2[0] = 11
	slice2[1] = 22
	slice2[2] = 66
	slice2[3] = 88

	fmt.Println("slice2:", slice2)
	fmt.Println("slice2", len(slice2))
	fmt.Println("slice2", cap(slice2))
	fmt.Println("\n ===============================================================================================================================================================================")

	for i := 0; i < len(slice2); i++ {
		fmt.Printf("slice2[%v] = %v \t", i, slice2[i])
	}
	fmt.Println("\n ===============================================================================================================================================================================")

	for index, v := range slice2 {
		fmt.Printf("下标：%v ,元素：%v\t", index, v)
	}
}

func SlicePrecautions() {

	//没有意义
	var slice []int
	fmt.Println(slice)
	fmt.Println("\n ===============================================================================================================================================================================")

	//越界
	var intarr [6]int = [6]int{3, 5, 6, 7, 88, 7}
	var slice1 []int = intarr[1:4]
	fmt.Println(slice1)
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])
	fmt.Println(slice1[2])
	// fmt.Println(slice1[3]) out range
	fmt.Println("\n ===============================================================================================================================================================================")

	//切片再切片
	slice2 := slice1[1:2]
	fmt.Println(slice2)
	fmt.Println(slice2[0])
	fmt.Println("\n ===============================================================================================================================================================================")

	//动态增长
	slice3 := intarr[1:4] //5.6.7
	fmt.Println(len(slice3))

	slice3 = append(slice3, 12, 11)
	fmt.Println(slice3)
	fmt.Println(slice3)
	fmt.Println("\n ===============================================================================================================================================================================")

	slice4 := append(slice3, slice2...)
	fmt.Println(slice4)
	fmt.Println(slice4)
}

func SliceCopy() {
	var slice []int = []int{1, 4, 7, 3, 6, 9}
	var b []int = make([]int, 10)

	//拷贝
	copy(b, slice) //将slice赋值到b
	fmt.Println(b)
	fmt.Println("\n ===============================================================================================================================================================================")
	//深拷贝
	slice[1] = 998
	fmt.Println(b)
	fmt.Println(slice)
}
