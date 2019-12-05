package main

import (
	"fmt"
)

func sum(n1 int,n2 int) int {
	// 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈中）
	// 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("n1",n1)
	defer fmt.Println("n2",n2)
	res := n1 + n2
	fmt.Println("res=", res)
	return res
}

func main() {
	sum(10, 20)
}