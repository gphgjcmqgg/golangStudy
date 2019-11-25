package main 
import (
	"fmt"
)

func main() {
	var address = "我是一个码农 120 Help me"

	fmt.Println(address)

	var code = `
	package main 
	import (
		"fmt"
	)

	func main() {
		var address = "我是一个码农 120 Help me"

		fmt.Println(address)
		
	}
	`
	fmt.Println(code)
}