package main

import "fmt"

func main() {
	ch := make(chan int, 2) // khởi tạo có chanel có kích thước là 2
	ch <- 1                 // gửi 1 vào chanel
	ch <- 2                 // gửi 2 vào chanel
	fmt.Println(<-ch)       // nhận giá trị 1
	fmt.Println(<-ch)       // nhận giá trị 2
}
