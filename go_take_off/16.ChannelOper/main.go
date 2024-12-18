package main

import (
	"fmt"
	"sync"
	"time"
)

var syncEvent sync.WaitGroup

func main() {

	ChannelSelectTest()

	// ch := make(chan int, 5)
	// go WriteOnlyChannel(ch)

	// for v := range ch {
	// 	fmt.Println(v)
	// }
}

// 管道基础
func ChannelBasicTest() {

	var intChan chan int = make(chan int, 3)
	fmt.Printf("intChan的值：%v \n", intChan)

	//写入数据到管道中,不能超过管道声明时的容量，否则会报错 all goroutines are asleep - deadlock!
	intChan <- 10
	intChan <- 20
	intChan <- 30

	//管道长度
	fmt.Println(len(intChan))
	//容量
	fmt.Println(cap(intChan))

	//从管道中读取值
	num := <-intChan
	num1 := <-intChan
	num2 := <-intChan
	fmt.Printf("读取管道的值是%v,%v,%v", num, num1, num2)

	//如果取数超过管道数量会报错 all goroutines are asleep - deadlock!
	//必须关闭	close(intChan) 然后读取就是0 ，但是不会报错
	close(intChan)
	num3 := <-intChan
	fmt.Printf("读取管道的值是%v", num3)

	////关闭后只能读不能写，如果往已经关闭的管道写入值也会报错panic: send on closed channel
	intChan <- 50
}

// 遍历管道
func ChannelIterate() {

	var intChan chan int = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i
	}

	close(intChan)
	for v := range intChan {
		fmt.Println("Value:", v)
	}
}

// /分别不同的协程写入和读取管道
func ChannelDemandTest() {

	var intChan chan int = make(chan int, 50)
	syncEvent.Add(2)
	go func() {
		defer func() {
			syncEvent.Done()
			close(intChan)
		}()
		for i := 0; i < 50; i++ {
			intChan <- i
		}
	}()

	go func() {
		defer syncEvent.Done()
		for v := range intChan {
			fmt.Println("read value:", v)
		}
	}()

	syncEvent.Wait()
}

// 声明一个只写或者只读的管道
func DefineChannel() {

	//只写管道只写
	var intChan chan<- int = make(chan int, 3)
	intChan <- 2
	intChan <- 3
	// num := <-intChan
	// fmt.Println("value:", num)

	//声明只读
	var intChanRead <-chan int = ReadOnlyChannel(200)
	for v := range intChanRead {
		fmt.Println("只读管道读取:", v)
	}

}

// 只读
func ReadOnlyChannel(length int) <-chan int {
	var intChanRead chan int = make(chan int, length)
	go func() {
		for i := 0; i < length; i++ {
			intChanRead <- i
		}
		close(intChanRead)
	}()
	return intChanRead
}

// 只写
func WriteOnlyChannel(content chan<- int) {
	for i := 0; i < 5; i++ {
		content <- i
	}
	close(content)
}

// /阻塞
func ChannelBlockTest() {

	var intChan chan int = make(chan int, 50)
	syncEvent.Add(1)
	go func() {
		defer func() {
			syncEvent.Done()
			close(intChan)
		}()
		for i := 0; i < 50; i++ {
			intChan <- i
		}
	}()

	// go func() {
	// 	defer syncEvent.Done()
	// 	for v := range intChan {
	// 		fmt.Println("read value:", v)
	// 	}
	// }()

	syncEvent.Wait()
}

// 管道的select功能
func ChannelSelectTest() {

	var intChan chan int = make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		intChan <- 12
	}()

	var strChan2 chan string = make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		strChan2 <- "测试一下"
	}()

	select {
	case v := <-intChan:
		fmt.Println(v)
	case v := <-strChan2:
		fmt.Println(v)
	default:
		fmt.Println("防止所有的select被阻塞")
	}

}
