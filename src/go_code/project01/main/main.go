package main

import "fmt"

// 定义全局变量
var n1 = 100
var n2 = 200
var n3 = 300
var(
	m1 = 10
	m2 = 12.5
	m3 = "jack"
)
func main() {

	// 输出语句
	fmt.Println("Hello")

	// 变量使用方式1
	var age int
	fmt.Println("age=", age)

	// 变量使用方式2
	name := "gaoj"
	fmt.Println("name=", name)

	// 变量使用方式3  类型推导
	var num = 10.10
	fmt.Println("num=", num)

	var n1,n2,n3 int
	var m1,m2,m3 = 10,"name",10.12

	q1,q2,q3 := "a",2,33

	fmt.Println(n1,n2,n3)
	fmt.Println(m1,m2,m3)
	fmt.Println(q1,q2,q3)

}