package main

type Product struct {
	ID    int
	Name  string
	Price float64
}

// Our in-memory store
var products = make(map[int]Product)
