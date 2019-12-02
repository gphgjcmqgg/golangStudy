package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {

	for {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100) + 1
		fmt.Println(num)
		if num == 99 {
			break
		}
	}
	
	
}