package main
import (
	"fmt"	
	"go_code/project01/func/utils"
)

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
func main() {
		var n1 float64 = 1.25
		var n2 float64 = 2.26
		var reuslt = utils.CalAdd(n1, n2)
		fmt.Println(reuslt)

		res1,res2 := getSumAndSub(2, 1)

		fmt.Println("sum=",res1,", sub=", res2)

		_,res3 := getSumAndSub(5, 1)

		fmt.Println(" sub=", res3)
}