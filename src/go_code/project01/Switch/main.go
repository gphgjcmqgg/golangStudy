package main 
import (
	"fmt"
)

func main() {
	switch score:= 90; {
		case score > 90:
			fmt.Println("1")
		case score >= 70 && score <= 90:
			fmt.Println("2")
		default :
			fmt.Println("3")
	}
}