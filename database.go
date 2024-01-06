package main

import (
	_ "github.com/lib/pq"
)

// Products
func getProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow(
		"SELECT id, name, price FROM products WHERE id=$1",
		id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func createProduct(product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"INSERT INTO products(name, price) VALUES ($1, $2) RETURNING id, name, price;",
		product.Name,
		product.Price,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"UPDATE products SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price;",
		product.Name,
		product.Price,
		id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func deleteProduct(id int) error {
	_, err := db.Exec(
		"DELETE FROM products WHERE id=$1;",
		id,
	)

	return err
}

// Suppliers
func getSuppliers() ([]Supplier, error) {
	rows, err := db.Query("SELECT id, name FROM suppliers")
	if err != nil {
		return nil, err
	}

	var suppliers []Supplier
	for rows.Next() {
		var p Supplier
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return suppliers, nil
}

func getSupplier(id int) (Supplier, error) {
	var p Supplier
	row := db.QueryRow(
		"SELECT id, name FROM suppliers WHERE id=$1",
		id,
	)

	err := row.Scan(&p.ID, &p.Name)
	if err != nil {
		return Supplier{}, err
	}

	return p, nil
}

func createSupplier(supplier *Supplier) (Supplier, error) {
	var s Supplier
	row := db.QueryRow(
		"INSERT INTO suppliers(name) VALUES ($1) RETURNING id, name;",
		supplier.Name,
	)
	err := row.Scan(&s.ID, &s.Name)
	if err != nil {
		return Supplier{}, err
	}

	return s, nil
}

func updateSupplier(id int, supplier *Supplier) (Supplier, error) {
	var p Supplier
	row := db.QueryRow(
		"UPDATE suppliers SET name=$1 WHERE id=$2 RETURNING id, name;",
		supplier.Name,
		id,
	)

	err := row.Scan(&p.ID, &p.Name)
	if err != nil {
		return Supplier{}, err
	}

	return p, nil
}

func deleteSupplier(id int) error {
	_, err := db.Exec(
		"DELETE FROM suppliers WHERE id=$1;",
		id,
	)

	return err
}

// ProductWithSupplier
func getProductsAndSuppliers() ([]ProductWithSupplier, error) {
	query := `
      SELECT
          p.id AS product_id,
          p.name AS product_name,
          p.price AS product_price,
          s.name AS supplier_name
      FROM
          products p
      INNER JOIN suppliers s
          ON p.supplier_id = s.id;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []ProductWithSupplier
	for rows.Next() {
		var p ProductWithSupplier
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductPrice, &p.SupplierName)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
