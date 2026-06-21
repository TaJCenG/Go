package main

import "fmt"

// Add Product
func AddProduct(p Product) {
	products[p.ID] = p
	fmt.Println("Added:", p)
}

// Get Product
func GetProduct(id int) (Product, bool) {
	p, exists := products[id]
	return p, exists
}

// Delete Product
func DeleteProduct(id int) {
	delete(products, id)
	fmt.Println("Deleted product with ID:", id)
}

// List Products
func ListProducts() {
	fmt.Println("Product List:")
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}
