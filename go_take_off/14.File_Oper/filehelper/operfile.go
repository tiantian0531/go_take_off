package filehelper

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func OperFileDemo1() {

	fmt.Println(os.Getwd())

	file, err := os.Open("E:/SigleProject/go_learn/go_take_off/go_take_off/14.File_Oper/filehelper/Test.txt")
	if err != nil {
		fmt.Println("打开文件出错")
	}
	fmt.Printf("文件= %v ", file)

	data := make([]byte, 1000)
	count, readErr := file.Read(data)

	if readErr != nil {
		fmt.Println("读取文件出错")
	}
	fmt.Printf("读取数据 %d 字节: %q\n", count, data[:count])
	closeErr := file.Close()
	if closeErr != nil {

		fmt.Println("关闭文件错误")
	}
}

// 一次性读取文件
func OperFileDemo2() {

	file, err := os.ReadFile("E:/SigleProject/go_learn/go_take_off/go_take_off/14.File_Oper/filehelper/Test.txt")
	if err != nil {
		fmt.Println("读取文件出错")
	}
	fmt.Printf("文件= %v \n ", file)

	fmt.Printf("读取数据 %d 字节: %q \n", len(file), file[:len(file)])

	fmt.Printf("读取数据为：%v", string(file))
	// closeErr := file.Close()
	// if closeErr != nil {

	// 	fmt.Println("关闭文件错误")
	// }
}

func OperFileDemo3() {

	file, err := os.Open("E:/SigleProject/go_learn/go_take_off/go_take_off/14.File_Oper/filehelper/Test.txt")
	if err != nil {
		fmt.Println("打开文件失败，err = ", err)
		return
	}
	fmt.Printf("文件= %v \n ", file)
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {

			fmt.Println("关闭文件错误")
		}
		fmt.Printf("关闭文件成功")
	}()

	reader := bufio.NewReader(file)

	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF && str == "" {
			break
		}
		fmt.Printf("读取数据为：%v", str)
	}

}
