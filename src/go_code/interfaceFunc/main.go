package main
import (
	"fmt"
	"sort"
	"math/rand"
)

// 1. 声明Student结构体
type Student struct {
	Name string
	Age int
}
// 2. 定义Student结构体切片类型
type StuSlice []Student 

// 3.实现Interface接口
func(ss StuSlice) Len() int {
		return len(ss)
}

// 年龄从大到小
func(ss StuSlice) Less(i, j int) bool {
	return ss[i].Age > ss[j].Age
}

func(ss StuSlice) Swap(i, j int) {
	// temp := ss[i]
	// ss[i] = ss[j]
	// ss[j] = temp
	// 两数交换
	ss[i],ss[j] = ss[j], ss[i]

}
func main() {

	var stuSlice StuSlice
	var stu Student
	for i:=0; i< 10; i++ {
		stu = Student{
			Name : fmt.Sprintf("heroes~%d",rand.Intn(100)),
			Age : rand.Intn(100),
		}
		stuSlice = append(stuSlice, stu)
	}
	for _,v := range stuSlice {
		fmt.Println(v)
	}
	fmt.Println("--------------------")
	// 调用排序
	sort.Sort(stuSlice)
	for _,v := range stuSlice {
		fmt.Println(v)
	}
}