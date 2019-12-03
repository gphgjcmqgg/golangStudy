package main
import (
	"fmt"	
	"go_code/project01/func/utils"
)

func main() {
		var n1 float64 = 1.25
		var n2 float64 = 2.26
		var reuslt = utils.CalAdd(n1, n2)
		fmt.Println(reuslt)
}