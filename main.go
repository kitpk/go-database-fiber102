package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

var db *sql.DB

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal()
	}

	app := fiber.New()

	app.Get("/product", getProductsHandler)
	app.Get("/product/:id", getProductHandler)
	app.Post("/product", createProductHandler)
	app.Put("/product/:id", updateProductHandler)
	app.Delete("/product/:id", deleteProductHandler)

	app.Get("/supplier", getSuppliersHandler)
	app.Get("/supplier/:id", getSupplierHandler)
	app.Post("/supplier", createSupplierHandler)
	app.Put("/supplier/:id", updateSupplierHandler)
	app.Delete("/supplier/:id", deleteSupplierHandler)

	app.Get("/productAndSupplier", getProductsAndSuppliersHandler)

	app.Listen(":8080")
}
