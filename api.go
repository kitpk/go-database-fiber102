package main

import (
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

// Products
func getProductsHandler(c *fiber.Ctx) error {
	products, err := getProducts()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(products)
}

func getProductHandler(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	product, err := getProduct(productID)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(product)
}

func createProductHandler(c *fiber.Ctx) error {
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	product, err := createProduct(p)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(product)
}

func updateProductHandler(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	product, err := updateProduct(productID, p)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(product)
}

func deleteProductHandler(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	err = deleteProduct(productID)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Suppliers
func getSuppliersHandler(c *fiber.Ctx) error {
	suppliers, err := getSuppliers()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(suppliers)
}

func getSupplierHandler(c *fiber.Ctx) error {
	supplierID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	supplier, err := getSupplier(supplierID)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(supplier)
}

func createSupplierHandler(c *fiber.Ctx) error {
	s := new(Supplier)
	if err := c.BodyParser(s); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	supplier, err := createSupplier(s)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(supplier)
}

func updateSupplierHandler(c *fiber.Ctx) error {
	supplierID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	s := new(Supplier)
	if err := c.BodyParser(s); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	supplier, err := updateSupplier(supplierID, s)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(supplier)
}

func deleteSupplierHandler(c *fiber.Ctx) error {
	supplierID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	err = deleteSupplier(supplierID)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// ProductWithSupplier
func getProductsAndSuppliersHandler(c *fiber.Ctx) error {
	productAndSuppliers, err := getProductsAndSuppliers()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(productAndSuppliers)
}
