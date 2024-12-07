package main

import (
	"fmt"
	"os"
)

func echo1Tets() {

	//for i := 0; i < len(os.Args); i++ {
	//	s += semp + os.Args[i]
	//	semp = " "
	//}

	for index, value := range os.Args {
		fmt.Println("索引是", index, "值是", value)
	}

}
