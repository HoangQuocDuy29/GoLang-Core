package main

import "fmt"

func main() {
	// Khởi tạo map lưu điểm số sinh viên
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}

	// Thêm phần tử mới
	scores["Charlie"] = 88

	// Truy cập giá trị
	fmt.Println("Điểm của Alice:", scores["Alice"])

	// Kiểm tra key có tồn tại
	if score, ok := scores["David"]; ok {
		fmt.Println("Điểm của David:", score)
	} else {
		fmt.Println("Không tìm thấy David")
	}

	// Xóa phần tử
	delete(scores, "Bob")
}
