package autotype

import (
	"github.com/go-vgo/robotgo"
)

func TypeStr(content string) bool {
	robotgo.TypeStr(content)
	if content == "\n" {
		robotgo.KeyTap("enter")
	}
	return true
}

// func CurrentMousePosition() {
// 	for {
// 		fmt.Println(robotgo.GetMousePos())
// 	}
// }
