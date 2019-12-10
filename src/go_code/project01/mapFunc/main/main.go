package main 

import (
	"fmt"
)

func main() {
	var a map[string]string
	// 在使用map前，需要先make，make的作用就是给map分配数据空间
	a = make(map[string]string,10)
	a["no1"] = "123"
	a["no2"] = "456"
	fmt.Println(a)

	var b map[string]string = map[string]string {
    "no1":"成都",
	}
	fmt.Println(b)

	// map遍历
	for k,v := range a {
		fmt.Printf("k=%v,v=%v",k,v)
	}
}