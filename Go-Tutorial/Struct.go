package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

func main() {
	// Tạo các instance của Car
	car1 := Car{"Toyota", "Camry", 2020}
	car2 := Car{
		Brand: "Honda",
		Model: "Accord",
		Year:  2021,
	}

	// Mỗi instance độc lập
	car1.Year = 2022
	fmt.Println(car1) // {Toyota Camry 2022}
	fmt.Println(car2) // {Honda Accord 2021}

	// Instance từ con trỏ
	car3 := new(Car)
	car3.Brand = "Ford"
	fmt.Println(*car3) // {Ford  0}
}
