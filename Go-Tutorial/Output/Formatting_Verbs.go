package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. Các verbs cơ bản
	fmt.Println("=== CÁC VERBS CƠ BẢN ===")
	fmt.Printf("Giá trị mặc định: %v\n", 42)              // %v: giá trị mặc định
	fmt.Printf("Kiểu dữ liệu: %T\n", 42)                  // %T: kiểu dữ liệu
	fmt.Printf("Giá trị theo cú pháp Go: %#v\n", "hello") // %#v: cú pháp Go
	fmt.Printf("Dấu phần trăm: %%\n")                     // %%: in dấu %

	// 2. Định dạng số
	fmt.Println("\n=== ĐỊNH DẠNG SỐ ===")
	fmt.Printf("Số nguyên (cơ số 10): %d\n", 42)         // %d: số nguyên
	fmt.Printf("Số nguyên (cơ số 8): %o\n", 42)          // %o: bát phân
	fmt.Printf("Số nguyên (cơ số 16): %x\n", 42)         // %x: thập lục phân
	fmt.Printf("Số nguyên (cơ số 16 hoa): %X\n", 42)     // %X: thập lục phân hoa
	fmt.Printf("Số thực: %f\n", 3.1415926535)            // %f: số thực
	fmt.Printf("Số thực (khoa học): %e\n", 3.1415926535) // %e: ký hiệu khoa học
	fmt.Printf("Số thực (ngắn gọn): %g\n", 3.1415926535) // %g: %e hoặc %f tùy trường hợp

	// 3. Định dạng boolean và con trỏ
	fmt.Println("\n=== BOOLEAN VÀ CON TRỎ ===")
	fmt.Printf("Giá trị boolean: %t\n", true) // %t: boolean
	ptr := 42
	fmt.Printf("Địa chỉ con trỏ: %p\n", &ptr) // %p: con trỏ

	// 4. Định dạng chuỗi và rune
	fmt.Println("\n=== CHUỖI VÀ RUNE ===")
	fmt.Printf("Chuỗi: %s\n", "Xin chào")          // %s: chuỗi
	fmt.Printf("Chuỗi (quoted): %q\n", "Xin chào") // %q: chuỗi trong dấu nháy
	fmt.Printf("Rune (Unicode): %U\n", 'A')        // %U: mã Unicode
	fmt.Printf("Ký tự: %c\n", 65)                  // %c: ký tự từ mã Unicode

	// 5. Định dạng độ rộng và căn lề
	fmt.Println("\n=== ĐỘ RỘNG VÀ CĂN LỀ ===")
	fmt.Printf("Căn phải (10 ký tự): |%10d|\n", 42)                 // %10d: độ rộng 10, căn phải
	fmt.Printf("Căn trái (10 ký tự): |%-10d|\n", 42)                // %-10d: độ rộng 10, căn trái
	fmt.Printf("Số thực (2 số sau dấu phẩy): %.2f\n", 3.1415926535) // %.2f: 2 số thập phân
	fmt.Printf("Đệm số 0: %05d\n", 42)                              // %05d: đệm số 0

	// 6. Định dạng struct và slice
	fmt.Println("\n=== STRUCT VÀ SLICE ===")
	p := Person{"Alice", 30}
	fmt.Printf("Struct (default): %v\n", p)       // %v: struct mặc định
	fmt.Printf("Struct (+field names): %+v\n", p) // %+v: struct với tên trường
	fmt.Printf("Slice: %v\n", []int{1, 2, 3})     // %v: slice

	// 7. Định dạng thời gian
	fmt.Println("\n=== ĐỊNH DẠNG THỜI GIAN ===")
	now := time.Now()
	fmt.Printf("Thời gian mặc định: %v\n", now) // %v: thời gian mặc định
	fmt.Printf("Thời gian định dạng: %02d/%02d/%d\n", now.Day(), now.Month(), now.Year())
}
