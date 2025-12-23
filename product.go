package main

import (
	"fmt"
)

// type Product struct {
// 	id int
// 	title string
// 	price float64
// 	stock int
// 	is_active bool
// }

// type ProductBuilder struct {
// 	*Product
// }

func main() {
	fmt.Println("Enter your name: ")
	var name string
	result := fmt.Scanln(&name)
	fmt.Println(result)
}

// func (product ProductBuilder) createProduct() *ProductBuilder {
// 	product := ProductBuilder(product.)	
// }