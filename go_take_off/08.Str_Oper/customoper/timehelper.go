package customoper

import (
	"fmt"
	"time"
)

func TimeHelperOper() {

	//2024-12-01 22:53:14.057792 +0800 CST m=+0.000000001
	var now = time.Now()
	fmt.Printf("%v 类型为：%T", now, now)

	// t := &time.Time{}

	var utcnow = now.UTC()
	fmt.Println(utcnow)

	fmt.Println(utcnow.Local())
}
