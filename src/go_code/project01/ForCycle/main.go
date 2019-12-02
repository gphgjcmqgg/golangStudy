package main 
import (
	"fmt"
)

func main() {
	for i:=1; i<=10; i++ {
		fmt.Println(i)
	}
	j := 1
	for j <= 10 {
			fmt.Println(j)
			j++
	}

	var str string = "abc~ok"
	for index, val := range str {
			fmt.Printf("index=%d, val=%c \n", index, val)
	}

	for i:= 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}


	// 如何字符串含有中文，传统遍历会出现乱码，传统遍历是按照字节遍历，而一个汉字在utf8编码是对应3个字节
	var str1 string = "hello,world!上海"
	str2 := []rune(str1)
	for i:= 0; i < len(str2); i++ {
			fmt.Printf("%c \n", str2[i])
	}
}

