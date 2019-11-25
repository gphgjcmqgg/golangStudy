package main

import "fmt"

// 指针类型
func main() {
	var i int = 10

	fmt.Println("i的地址值是：", &i)

	// 1. ptr 是一个指针变量
	// 2. ptr 的类型 *int
	// 3. ptr 本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)

}