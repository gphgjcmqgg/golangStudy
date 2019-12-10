package main 

import (
	"fmt"
	"sort"
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
		fmt.Printf("k=%v,v=%v\n",k,v)
	}

	// map切片
	var slice []map[string]string
	slice = make([]map[string]string,2)
	if slice[0] == nil {
		slice[0]= make(map[string]string,2)
		slice[0]["name"] = "jack"
		slice[0]["age"] = "34"
	}
	
	fmt.Println(slice)
	monster := map[string]string {
		"name" : "zhangsan",
		"age" : "200",
	}
	slice = append(slice,monster)
	fmt.Println(slice)

	// map排序
	map1 := map[int]int {
		10:99,
		1:5,
		4:23,
		7:8,
	}
	// 排序
	var keys []int
	for k,_ := range map1 {
			keys = append(keys,k)
	}
	sort.Ints(keys)

	for _ ,k := range keys {
			fmt.Printf("map[%v]=%v\t", k, map1[k])
	}

}