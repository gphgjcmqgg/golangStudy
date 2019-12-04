package main

import "fmt"

// 递归算法
// 斐波那契数
// 1,1,2,3,5,8,13
func feibo(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else 	{
		return feibo(num-2) + feibo(num-1)
	}

}
func main() {
	var num int = 7

	result := feibo(num)
	fmt.Println(result)
}