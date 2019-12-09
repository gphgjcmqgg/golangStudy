package main

import (
	"fmt"
)

func test(arr *[3]int) {
	(*arr)[0] = 55
}
func main() {
	// 数组几种初始化的方式
	var numArr01 [3]int = [3]int{1,2,3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{5,6,7}
	fmt.Println("numArr02=", numArr02)

	var numArr03 = [...]int{7,8,9}
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1:700,0:800,2:900}
	fmt.Println("numArr04=", numArr04)

	numArr05 := [3]int{15,16,17}
	fmt.Println("numArr05=", numArr05)

	for i,v := range numArr03 {
		fmt.Printf("index=%d, value=%v \n", i, v)
	}
	test(&numArr01)
	fmt.Println("numArr01=", numArr01)
}
