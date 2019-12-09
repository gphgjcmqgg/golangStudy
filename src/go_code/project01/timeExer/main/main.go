package main
import (
	"fmt"
	"time"
	"strconv"
)

func strAdd () {
	var str string
	for i := 0; i < 100000;i++ {
		str += "hello" + strconv.Itoa(i)
	}

}

func main() {
	start := time.Now().Unix()
	strAdd()
	end := time.Now().Unix()
	fmt.Printf("共用时%v秒", end -start)
}