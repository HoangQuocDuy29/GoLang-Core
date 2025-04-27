package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	ch := make(chan int)

	go func() {
		ch <- 1 // Goroutine gửi giá trị vào channel
	}()
	fmt.Println(<-ch) // nhận giá trị từ channel
}
