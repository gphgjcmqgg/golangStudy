package main
import (
	"fmt"
)

func main() {
		var c1 byte = 'a'
		var c2 byte = '0'

		fmt.Println("c1=", c1)
		fmt.Println("c2=", c2,"\n")

		fmt.Printf(" c1=%c ,c2=%c ", c1, c2)

		//var c3 byte = '北'
		var c4 int = '北'

		// fmt.Printf(" c3=%c" , c3)
		fmt.Printf("c4=%c" , c4)
}