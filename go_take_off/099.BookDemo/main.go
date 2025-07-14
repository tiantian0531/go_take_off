package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	//Spt()
	//TimeOper()
	//TimestampOper()
	//LoggerOper()
	//FileOper()
	FileOper()
}

func Fpf() {

	fmt.Fprintln(os.Stdout, "hello world")
	file, err := os.OpenFile("./099.BookDemo/t.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错", err)
	}
	name := "yuxl"
	fmt.Fprintf(file, "往文件中写如信息：%s", name)
}

func Spt() {
	s1 := fmt.Sprint("卡布达")
	name := "卡布达"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("卡布达")
	fmt.Println(s1, s2, s3)
}

// 时间基本操作
func TimeBaseOper() {
	now := time.Now()
	fmt.Printf("当前时间为:%v \n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("当前时间为:%d:%02d:%02d-%02d:%02d:%02d: \n", year, month, day, hour, minute, second)
}

// 时间戳操作
func TimestampOper() {
	now := time.Now()
	timestamp := now.Unix()
	timestampNano := now.UnixNano()
	fmt.Printf("时间戳是:%v ,纳秒时间戳是:%v \n", timestamp, timestampNano)

	rNow := time.Unix(timestamp, 0)

	fmt.Printf("转换回来的时间是:%v \n", rNow)

	add := now.Add(time.Hour * 1)
	sub := add.Sub(now)

	duration := time.Since(now)
	fmt.Printf("添加之后的时间是:%v \n", add)
	fmt.Printf("时间差值是:%v \n", sub)
	fmt.Printf("经过的时间是:%v \n", duration)

	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

// 日志
func LoggerOper() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[给一个前缀.]")
	log.Printf("输出一个日志:%s \n", time.Now().Format("2006-01-02 15:04:05"))
}

// 文件操作
func FileOper() {

	file, err := os.Open("./099.BookDemo/t.txt")
	if err != nil {
		fmt.Println("打开文件错误", err)
		return
	}
	// 关闭文件
	defer file.Close()
	var buf [128]byte
	var content []byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取文件错误", err)
			return
		}

		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}
