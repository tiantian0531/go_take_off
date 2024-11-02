package main

import "fmt"

func main() {

	//5个学生的成绩
	// score1 := 95
	// score2 := 91
	// score3 := 39
	// score4 := 60
	// score5 := 21

	// sum := score1 + score2 + score3 + score4 + score5
	// //平均数
	// avg := sum / 5
	// fmt.Printf("成绩的总和是：%v,成绩的平均数是：%v", sum, avg)

	// var scores [5]int
	// scores[0] = 95
	// scores[1] = 91
	// scores[2] = 39
	// scores[3] = 60
	// scores[4] = 21

	// var sum = 0
	// for i := 0; i < len(scores); i++ {
	// 	sum += scores[i]
	// }
	// avg := sum / len(scores)
	// fmt.Printf("容量是：%v,长度是%v", cap(scores), len(scores))
	// fmt.Printf("成绩的总和是：%v,成绩的平均数是：%v", sum, avg)

	//引用指针传递
	var arr_ref = [3]int{1, 2, 3}
	assignment_ref(&arr_ref)
	fmt.Println(arr_ref)

	//值传递
	var arr_val = [3]int{7, 8, 9}
	assignment_val(arr_val)
	fmt.Println(arr_val)
}

func assignment_ref(arr *[3]int) {
	arr[0] = 999
}
func assignment_val(arr [3]int) {
	arr[0] = 888
}
