package yichang

import "fmt"

func SafeDivide(a, b int) (result int, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("错误panic:%v", r)
		}
	}()
	result = a / b
	return
}