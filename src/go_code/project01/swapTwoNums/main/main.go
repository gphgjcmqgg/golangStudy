package main
import (
	"fmt"	
)

// 交换两个数字
func swap(n1 *int , n2 *int) {
	fmt.Println(" swap n1地址=", &n1, "n1所存地址=", n1, "n1所存值=", *n1)
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func main() {
	var n1 int = 10
	var n2 int = 20
	fmt.Println("main n1地址=", &n1)
	swap(&n1, &n2)
	fmt.Println("n1:", n1, "-n2:", n2)
}