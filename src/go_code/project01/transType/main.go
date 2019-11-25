package main 
import (
	"fmt"
	"strconv"
)

func main() {

	var i int32 = 100
	// int => float
	var f float32 = float32(i)
	fmt.Println(f)

	var n1 int32 = 12
	var n2 int8
	// var n3 int8
	n2 = int8(n1) + 127				// 编译通过，但结果不是127+12，会溢出处理
	// n3 = int8(n1) + 128		编译不过 128 超过int8值范围
	fmt.Println(n2)

	//------------------------------------------------------------------
	// 基本数据类型转string
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string // 空的

	// 方式1 fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n\n", str, str)

	// 方式2 使用strconv包的函数
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	// strconv.FormatFloat(num2, 'f', 10 , 64)
	// 说明： 'f'格式  10：表示小数位保留10位		64表示这个小数是64位的
	str = strconv.FormatFloat(num2, 'f', 10 , 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type %T str=%q\n", str, str)

	// 使用strconv包中有一个函数Itoa
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str=%q\n", str, str)
	//------------------------------------------------------------------
	// string转基本数据类型
	var str1 string = "true"
	var b3 bool
	// 1.strconv.ParseBool(str) 函数会返回两个值 （vaule bool, err error）
	// 2.只想获取到value bool，不需要获取err error，所以使用 _来忽略
	b3 , _ = strconv.ParseBool(str1)
	fmt.Printf("b3 type %T b3=%v\n", b3, b3)

	var str2 string = "1234590"
	var i3 int64
	i3 , _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("i3 type %T i3=%v\n", i3, i3)

	var str3 string = "123.456"
	var f4 float64
	f4 , _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f4 type %T f4=%v\n", f4, f4)
}