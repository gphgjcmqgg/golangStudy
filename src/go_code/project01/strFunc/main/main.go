package main 
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1.len       统计字符串长度,按字节len(str)
	// golang的编码统一为utf-8（ascii的字符（字母和数字）占一个字节，汉字占用3个字节）
	var str string = "hello"
	fmt.Println(len(str)) // 5
	str = "hello北"
	fmt.Println(len(str)) // 8
	str = "hello北京"
	for i:= 0; i< len(str); i++ {
		fmt.Printf("字符=%c\n", str[i])
	}
	fmt.Println("")
	// 2.字符串遍历，同时处理有中文的问题 r := []rune(str)
	newStr := []rune(str)
	for i:= 0; i< len(newStr); i++ {
		fmt.Printf("字符=%c\n", newStr[i])
	}

	// 3. 字符串转整数 n,err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("n=", n)
	}
	// 整数转字符串
	str = strconv.Itoa(123)
	fmt.Printf("str=%v,str=%T", str, str)
	fmt.Println("")
	// []byte 转字符串
	str = string([]byte{97,98,99})
	fmt.Println("str=", str)
	fmt.Println("")
	// 查找子串是否在指定的字符串中
	fmt.Printf("是否包含%v\n", strings.Contains("seafood", "foo"))

	// 统计一个字符串有几个指定的子串
	fmt.Printf("有几个sea子串：%v\n", strings.Count("seafood sea", "sea"))

	// 区分大小写的字符串比较
	fmt.Printf("是否相等：%v\n", "sea" == "SEA")
	// 不区分大小写的字符串比较
	fmt.Printf("是否相等：%v\n", strings.EqualFold("sea", "SEA"))

	// 返回子串在字符串第一次出现的index值
	fmt.Printf("字符串出现的索引：%v\n", strings.Index("NLT_abc","abc"))

	// 返回子串在字符串最后一次出现的index值如果没有，返回-1
	fmt.Printf("字符串最后一次出现出现的索引：%v\n", strings.LastIndex("go golang","go")) // 3
	
	// 将指定的子串替换成另外一个子串 -1 表示全部替换
	str = strings.Replace("go go hello", "go", "go语言", -1)		// go语言 go语言 hello
	fmt.Printf("str=%v\n", str)
	str = strings.Replace("go go hello", "go", "go语言", 1)  	 // // go语言 go hello
	fmt.Printf("str=%v\n", str)

	// 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split("hello,world,ok",",")
	for i:= 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	// 将字符串左右两边指定的字符串去掉
	str = strings.Trim("! he!lo! "," !") 		// 去掉左右两边的"!"和" "
	fmt.Printf("str=\"%v\"\n", str)
}