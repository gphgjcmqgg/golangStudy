package main 
import (
	"fmt"
)

var (
	// fun1就是一个全局匿名函数
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)
func main() {
	res := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=",res)

	// a函数类型
	a := func (n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(20, 10)
	fmt.Println("res2=",res2)
}