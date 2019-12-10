package main

import (
	"fmt"
)

func bubbleSort(arr *[6]int) {
	 temp:= 0
	 for i:=0;i< len(*arr) -1; i++ {
		 for j:=0; j < len(*arr) -1 -i; j++ {
			 if (*arr)[j] > (*arr)[j+1] {
				temp= (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp
			 }
		 }
	 }
	 
}

func main() {

	var intArr [6]int = [6]int{23, 8, 56, 4, 81, 5}
	bubbleSort(&intArr)
	fmt.Println(intArr)
}