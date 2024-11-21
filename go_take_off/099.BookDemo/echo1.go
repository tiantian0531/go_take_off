package main

import (
	"fmt"
	"os"
)

func echo1Tets() {
	var s, semp string
	for i := 0; i < len(os.Args); i++ {
		s += semp + os.Args[i]
		semp = " "
	}
	fmt.Println(s)
}
