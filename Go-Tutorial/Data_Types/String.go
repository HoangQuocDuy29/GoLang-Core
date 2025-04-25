package main

import (
	"fmt"
)

func main() {
	var txt1 string = "<UNK>"
	var txt2 string
	txt3 := "ABC"

	fmt.Printf("Type : %T, value : %v \n", txt1, txt1)
	fmt.Printf("Type : %T, value : %v \n", txt2, txt2)
	fmt.Printf("Type : %T, value : %v \n", txt3, txt3)
}
