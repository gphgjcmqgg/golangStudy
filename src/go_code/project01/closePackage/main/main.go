package main 
import (
	"fmt"
	"strings"
)

// func AddUpper() func(int) int {
// 	var n int = 10
// 	return func (x int) int {
// 		n = n + x
// 		return n
// 	}
// }

// strings.HasSuffix(s string， suffix string) bool 判断一个文件是否含有后缀
func makeSuffix(suffix string) func(string) string{
	// var ext string = suffix
	return func(fileName string) string {
		if strings.HasSuffix(fileName, suffix) {
			return fileName
		} else {
			return fileName + suffix
		}
	}
}
func main() {
	//  f := AddUpper()
	//  fmt.Println(f(1)) // 11
	//  fmt.Println(f(2)) // 13
	//  fmt.Println(f(3)) // 16
	makeSuf := makeSuffix(".jpg")
	filename := makeSuf("abc")
	fmt.Println(filename)

	filename = makeSuf("ccc.jpg")
	fmt.Println(filename)
}