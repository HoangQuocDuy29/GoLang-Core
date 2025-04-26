package main

import "fmt"

func main() {
	// 1. Toán tử gán cơ bản =
	var a int = 10
	fmt.Printf("1. Gán cơ bản: a = %d\n", a)

	// 2. Các toán tử gán số học
	a += 5 // Tương đương a = a + 5
	fmt.Printf("2. += : a += 5 → a = %d\n", a)

	a -= 3 // Tương đương a = a - 3
	fmt.Printf("3. -= : a -= 3 → a = %d\n", a)

	a *= 2 // Tương đương a = a * 2
	fmt.Printf("4. *= : a *= 2 → a = %d\n", a)

	a /= 4 // Tương đương a = a / 4
	fmt.Printf("5. /= : a /= 4 → a = %d\n", a)

	a %= 3 // Tương đương a = a % 3
	fmt.Printf("6. %%= : a %%= 3 → a = %d\n", a)

	// 3. Các toán tử gán bitwise
	b := 0b1100 // 12 trong hệ nhị phân
	b &= 0b1010 // AND bit: 1100 & 1010 = 1000 (8)
	fmt.Printf("7. &= : b &= 0b1010 → b = %b (%d)\n", b, b)

	b |= 0b0011 // OR bit: 1000 | 0011 = 1011 (11)
	fmt.Printf("8. |= : b |= 0b0011 → b = %b (%d)\n", b, b)

	b ^= 0b1111 // XOR bit: 1011 ^ 1111 = 0100 (4)
	fmt.Printf("9. ^= : b ^= 0b1111 → b = %b (%d)\n", b, b)

	b <<= 1 // Dịch trái 1 bit: 0100 << 1 = 1000 (8)
	fmt.Printf("10. <<= : b <<= 1 → b = %b (%d)\n", b, b)

	b >>= 2 // Dịch phải 2 bit: 1000 >> 2 = 0010 (2)
	fmt.Printf("11. >>= : b >>= 2 → b = %b (%d)\n", b, b)

	b &^= 0b0011 // AND NOT bit: 0010 &^ 0011 = 0010 (2)
	fmt.Printf("12. &^= : b &^= 0b0011 → b = %b (%d)\n", b, b)

	// 4. Gán nhiều giá trị cùng lúc
	x, y := 5, 10
	fmt.Printf("\n13. Trước khi hoán đổi: x = %d, y = %d\n", x, y)
	x, y = y, x // Hoán đổi giá trị
	fmt.Printf("14. Sau khi hoán đổi: x = %d, y = %d\n", x, y)

	// 5. Gán cho con trỏ
	value := 100
	var ptr *int
	ptr = &value
	fmt.Printf("\n15. Giá trị của ptr: %p, Giá trị trỏ tới: %d\n", ptr, *ptr)
}
