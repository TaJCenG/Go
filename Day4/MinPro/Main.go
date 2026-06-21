package main

func main() {
	// Add products
	AddProduct(Product{ID: 1, Name: "Laptop", Price: 50000})
	AddProduct(Product{ID: 2, Name: "Phone", Price: 20000})

	// List all
	ListProducts()

	// Get one
	if p, ok := GetProduct(1); ok {
		println("Fetched:", p.Name)
	}

	// Delete one
	DeleteProduct(2)

	// List again
	ListProducts()
}

// Maps → storing products by ID.

// Structs → modeling product data.

// Functions → encapsulating operations (Add, Get, Delete, List).

// Memory behavior → maps are reference types, so changes persist across functions.
