package main

import "fmt"

func digui(n int) {
	if(n>2) {
		n--
		digui(n)
	}
	fmt.Println("n=",n)
}

func digui2(n int) {
	if(n>2) {
		n--
		digui2(n)
	} else {
		fmt.Println("n=",n)
	}
}

func main() {
	var n int = 4
	digui2(n)
}