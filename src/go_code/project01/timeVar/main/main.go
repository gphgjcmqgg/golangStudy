package main 
import (
	"fmt"
	"time"
)

func main() {
	i:=0
	for {
			fmt.Println(i)
			i++
			time.Sleep(time.Millisecond * 100) // 休眠0.1秒 100毫秒
			if i == 100 {
					break
			}
	}
}