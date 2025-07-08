package bibao
import (
	// "first/100.Question/qassert"
	
	"fmt"
)

func Add() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}

	//call 
	// pos:= bibao.Add()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(pos(i))
	// }
}



type Caclulate func(x int,y int) int

func DoublePrint(fn Caclulate) Caclulate{
	return func(x, y int) int {
		fmt.Printf("Calling function with arguments: %d, %d\n", x, y)
		result := fn(x, y)
		result+=1
		fmt.Printf("Function returned: %d\n", result)
		return result
	}
}

// 	decoratedAdd := bibao.DoublePrint(bibao.DoublePrint(Add))
// 	result :=decoratedAdd(1,5)
// 	fmt.Println("r的值是", result)
// 	func Add(a, b int) int {
// 	return a + b
// }