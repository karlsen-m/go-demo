package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	x, y := robotgo.GetMousePos()
	fmt.Println("Current mouse position:", x, y)

	// 输出鼠标移动是否成功
	success := robotgo.MoveSmooth(100, 100, 1.0, 100.0)
	fmt.Println("Mouse move success:", success)

	time.Sleep(2 * time.Second)

	success = robotgo.MoveSmooth(x, y, 1.0, 100.0)
	fmt.Println("Mouse move back success:", success)
}
