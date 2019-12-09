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

// 斐波那契数 ->切片实现
// 1,1,2,3,5,8,13
func feiboSlice(num int) []uint64 {
	fbnSlice := make([]uint64, num)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < num; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	var num int = 7

	result := feibo(num)
	fmt.Println(result)

	num2 := feiboSlice(7)
	fmt.Println(num2)
}