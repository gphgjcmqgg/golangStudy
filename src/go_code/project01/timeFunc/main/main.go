package main 
import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v new type=%T\n", now, now)

	fmt.Printf("当前年月日 时分秒 %d-%d-%d %d:%d:%d\n",now.Year(), now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())

	dateStr := fmt.Sprintf("当前年月日 时分秒 %d-%d-%d %d:%d:%d\n",now.Year(), now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Printf("dateStr=%v", dateStr)

	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println("")
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println("")
	fmt.Printf(now.Format("15:04:05"))
}