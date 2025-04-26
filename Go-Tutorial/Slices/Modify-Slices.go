package main

import (
	"fmt"
)

func main() {
	slice := []string{"A", "B"}
	fmt.Println("Trước : ", slice, "Len: ", len(slice), "Cap: ", cap(slice))

	// Thêm 1 phần tử
	slice = append(slice, "C")
	fmt.Println("Sau khi thêm 1 append: ", slice)

	//Cập nhật phần tử bằng index và gán giá trị mới
	slice[2] = "D"
	fmt.Println("Cập nhật: ", slice)

	fmt.Println(slice[2])

}
