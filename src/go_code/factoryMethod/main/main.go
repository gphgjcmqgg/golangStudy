package main
import (
	"fmt"	
	"go_code/factoryMethod/model"
)

func main() {
	var stu = model.StudentFactory("gj", 15.15)
	fmt.Println("stu=", *stu)
	fmt.Println("name=", (*stu).Name)
	fmt.Println("score=", stu.GetScore())
}