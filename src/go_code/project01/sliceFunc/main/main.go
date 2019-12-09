package main

import (
	"fmt"
)

func main() {

	var intArr [5]int = [5]int{1,22,33,66,99}

	// slice为切片
	// intArr[1:3]表示slice引用到intArr这个数组的第二个元素到第三个元素
	// 引用intArr数组的起始下标为1，最后的下标为3（但是不包含3）
	slice := intArr[1:3]
	fmt.Println(slice)

	// make
	var sliceMake []int = make([]int, 5 , 10)
	fmt.Println(sliceMake)

	var slice1 []int = []int{100,200,300}
	slice1 = append(slice1, 400,500,600)
	fmt.Println(slice1)

	slice1 = append(slice1, slice1...)
	fmt.Println(slice1)

	var slice4 []int = []int{1,2,3,4,5}
	var slice5 = make([]int,10)
	copy(slice5, slice4)
	fmt.Println(slice4)
	fmt.Println(slice5)

}