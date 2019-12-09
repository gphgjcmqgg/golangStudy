package main

import (
	"fmt"
	"errors"
)

func readConf(name string) (err error) {
	if name == "conf.ini"{
		return nil
	} else {
		return errors.New("文件错误") // 自定义errors错误
	}
}

func test() {
	defer func () {
    err := recover() // recover()内置函数、可以捕获到异常
    if err != nil {
        fmt.Println("err=",err)
    }
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	test()
	fmt.Println("main exec")
	err := readConf("123.ini")
	if err != nil {
		panic(err)
	}
}