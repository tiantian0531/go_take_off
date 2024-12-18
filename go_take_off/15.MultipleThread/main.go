package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)
var (
	event sync.WaitGroup
	lock sync.Mutex
	totalNum int =0
	totalNum1 int32 =0
	rwLock sync.RWMutex
)
func main() {
	// MutlpleThreadOperOneData()
	// MutlpleThreadOperOneDataAtomiclock()
	MutlpleThreadOperOneDataRwLock()
}


func WaitGroupTest(){
	event.Add(1)
	go func(){
		defer event.Done()
		fmt.Println("协程")
	}()
	//  time.Sleep(time.Second * 1000)
	event.Wait()
}

func StartMutlple(){
	event.Add(5)
	for i := 0; i < 5; i++ {
		go func(n int){
			defer event.Done()
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Second * 2)
	event.Wait()
}


func MutlpleThreadOperOneDataLock(){
	event.Add(2)
	go func()  {
		defer event.Done()
		lock.Lock()
		for i := 0; i < 10000; i++ {			
			totalNum +=i	
		}
		lock.Unlock()
	}()
	go func()  {
		defer event.Done()
		lock.Lock()
		for i := 0; i < 10000; i++ {
		
			totalNum -=i
			
		}
		lock.Unlock()
	}()
	event.Wait()

	fmt.Println(totalNum)
}

func MutlpleThreadOperOneDataAtomiclock(){
	event.Add(2)
	go func()  {
		defer event.Done()
		for i := 1; i<= 100; i++ {		
			atomic.AddInt32(&totalNum1,int32(i))	
		}
	}()
	go func()  {
		defer event.Done()
		for i := 1; i <= 100; i++ {
		
			atomic.AddInt32(&totalNum1,-int32(i))	
			
		}
	}()
	event.Wait()

	fmt.Println(totalNum1)
}


func MutlpleThreadOperOneDataRwLock(){
	event.Add(200)

		for i := 1; i<= 100; i++ {		
			go func()  {
				defer event.Done()
				rwLock.RLock()
				fmt.Println("读取数据")
				time.Sleep(time.Second)
				fmt.Println("读取数据成功")
				rwLock.RUnlock()
			}()
			
		}

	

		for i := 1; i <= 100; i++ {
			go func()  {
				defer event.Done()
				rwLock.Lock()
				fmt.Println("修改数据")
				time.Sleep(time.Second*1)
				fmt.Println("修改数据成功")
				rwLock.Unlock()
			}()
		}

	event.Wait()

	fmt.Println(totalNum1)
}