package main

import "fmt"

func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1)) // 0 (slice rỗng)
	fmt.Println(cap(myslice1)) // 0 (chưa cấp phát)
	fmt.Println(myslice1)      // []

	myslices2 := [4]string{"A", "B", "C", "D"}
	fmt.Println(len(myslices2)) // 4 (4 phần tử)
	fmt.Println(cap(myslices2)) // 4 (bằng len khi khởi tạo trực tiếp)
	fmt.Println(myslices2)      // [A B C D]
}
