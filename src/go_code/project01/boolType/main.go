package main 
import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Printf("b的占用空间 = %d", unsafe.Sizeof(b))
}