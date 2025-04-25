package main

// khai báo nhiều biến cùng lúc
import "fmt"

func main() {
	//var a,b,c int = 1,2,3;
	//
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	var a, b int = 1, 2
	c, d := 3, "Duy"

	// hoán đổi giá trị
	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
