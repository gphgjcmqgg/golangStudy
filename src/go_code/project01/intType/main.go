package main
import (
	"fmt"	// fmt包中提供格式化、输出、输入的函数
	"unsafe"
)

// 转移字符		
// \t   制表符
// \\		斜杠转义	
// \n		换行符	
// \"		双引号转义
// \r	  回车
func main() {
	fmt.Println("char\tChange")		// \t   制表符
	fmt.Println("char\nChange") 	// \n		换行符	

	fmt.Println("D:\\GolangCode\\src\\go_code\\project01\\main")	// \\		斜杠转义	
	fmt.Println("Tome say \"I love you\"")	// \"		双引号转义

	fmt.Println("我是一个\r超人")	// \r	  回车	从当前行的最前面开始输出，覆盖掉以前内容


	var n = 100

	fmt.Printf("n的数据类型%T 占用字节数是%d", n, unsafe.Sizeof(n))

}