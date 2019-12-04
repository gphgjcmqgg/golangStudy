package main
import (
	"fmt"	
	_"go_code/project01/func/utils"
)

// func getSumAndSub(n1 int, n2 int) (int, int) {
// 	sum := n1 + n2
// 	sub := n1 - n2
// 	return sum, sub
// }
// func main() {
// 		var n1 float64 = 1.25
// 		var n2 float64 = 2.26
// 		var reuslt = utils.CalAdd(n1, n2)
// 		fmt.Println(reuslt)

// 		res1,res2 := getSumAndSub(2, 1)

// 		fmt.Println("sum=",res1,", sub=", res2)

// 		_,res3 := getSumAndSub(5, 1)

// 		fmt.Println(" sub=", res3)
// }
func test01(n1 int) {
	n1 = n1 +10
	fmt.Println("test01 n1=" , n1)
}

func test02(n1 *int) {
	*n1 = *n1 +10
	fmt.Println("test01 n1=" , *n1)
}

func getSum(n1 int,n2 int) int {
	return n1 + n2
}

func myFum(funvar func(int,int) int, n1 int ,n2 int) int {
	return funvar(n1, n2)
}
func main() {
	var n1 int = 20
	// test01(n1)
	test02(&n1)
	fmt.Println("main n1=" , n1)

	res := myFum(getSum, 5, 10)
	fmt.Println("main res=" , res)
}