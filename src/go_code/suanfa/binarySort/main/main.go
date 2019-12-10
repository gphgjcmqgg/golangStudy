package main

import (
	"fmt"
)

func binarySort(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	 if leftIndex > rightIndex {
		fmt.Println("没有找到")
		return
	 }

	 var middle = (leftIndex + rightIndex) / 2

	 if (*arr)[middle] > findVal {
			binarySort(arr, leftIndex, middle -1, findVal)
	 } else if (*arr)[middle] < findVal {
			binarySort(arr, middle + 1, rightIndex, findVal)
	 } else {
		 fmt.Println("数字找到了,索引是:", middle)
	 }
	 
}

func main() {

	var intArr [6]int = [6]int{1, 8, 10, 89, 1000, 1234}
	binarySort(&intArr, 0, len(intArr) - 1, 1234)
}