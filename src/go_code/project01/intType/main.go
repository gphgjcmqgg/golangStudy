package main
import (
	"fmt"
	"unsafe"
)

func main() {
	var n = 100

	fmt.Printf("n的数据类型%T 占用字节数是%d", n, unsafe.Sizeof(n))

}