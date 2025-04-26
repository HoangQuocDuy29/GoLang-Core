package main

import (
	"fmt"
)

func main() {
	x := 22
	if x < 10 {
		fmt.Println("Good morning")
	} else if x < 20 {
		fmt.Println("Good day")
	} else {
		fmt.Println("Good evening")
	}
}
