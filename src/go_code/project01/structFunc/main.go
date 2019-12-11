package main
import (
	"fmt"	
	"encoding/json"
)

type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}

type Person struct {
	Name string `json:"name"`
	Age int	`json:"age"`
}
type A struct{
	Num int
}

type B struct{
	Num int
}

func main() {
	// 创建一个cat变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "white"
	cat1.Hobby = "吃鱼"
	fmt.Println(cat1)

	// 方式2
	var cat2 Cat = Cat{"Tom", 20, "Red", "旅游"}
	fmt.Println(cat2)

	// 方式3
	var c3 *Cat = new(Cat)
	(*c3).Name = "Jerry"
	fmt.Println(*c3)

	// 方式4
	var c4 *Cat = &Cat{}
	(*c4).Name = "Tom"
	c4.Name = "Mary"
	fmt.Println(*c4)

	var a A
	var b B
	a = A(b)
	fmt.Println(a)

	var person = Person{"boy",50}
	fmt.Println("person=", person)
	jsonStr,err := json.Marshal(person)
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println("jsonStr=", string(jsonStr))
	}
}