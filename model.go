package main

import (
	_ "github.com/lib/pq"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Supplier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductWithSupplier struct {
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"product_price"`
	SupplierName string `json:"supplier_name"`
}
