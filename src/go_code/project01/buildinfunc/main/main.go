package main
import (
	"fmt"
)

func main() {
	num1 := 100
	fmt.Printf("num1的类型%T, num1的值是%v, num1的地址%v\n", num1 , num1 , &num1)

	num2 := new(int)
	// num2的类型%T == *int
	// num2的值是== 地址 0xc04200e0e0	（这个地址系统分配）
	// num2的地址 == 地址 0xc042004030	（这个地址系统分配）
	// num2指向的值 == 0
	fmt.Printf("num2的类型%T, num2的值是%v, num2的地址%v, num2指向的值%v\n", num2 , num2 , &num2, *num2)
	*num2 = 100
	fmt.Printf("num2的类型%T, num2的值是%v, num2的地址%v, num2指向的值%v\n", num2 , num2 , &num2, *num2)
}