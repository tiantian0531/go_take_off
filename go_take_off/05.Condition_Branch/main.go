package main

import "fmt"

func main() {
	//condtionIF(30)
	//condtionIF(0)
	//condtionIF(-30)
	conditionSwitch(0)
}

// if else
func condtionIF(param int32) {
	if count := param; count > 0 {
		fmt.Print("大于0")
	} else if count < 0 {
		fmt.Print("小于0")
	} else {
		fmt.Print("等于0")
	}
}

func conditionSwitch(param int32) {
	switch count := param; {
	case count < 0:
		fmt.Print("小于0")
	case count > 0:
		fmt.Print("大于0")

	default:
		fmt.Print("等于0")
	}
}
