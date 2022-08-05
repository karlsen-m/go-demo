package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 1)

func set() {
	for i := 1; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second * 5)
	}
}

func get() {
	for {
		select {
		case val := <-ch:
			if val >= 1 {
				go print(val)
			}
		}
		time.Sleep(time.Second * 1)
	}

}
func print(val int) {
	fmt.Println(val)
}
func main() {
	go get()
	go set()
	time.Sleep(time.Second * 50)
}
