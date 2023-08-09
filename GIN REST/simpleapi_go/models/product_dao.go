package models

import (
	"database/sql"
	"fmt"
)

type ProductDAO struct {
	db *sql.DB
}

func NewProductDAO(db *sql.DB) *ProductDAO {
	return &ProductDAO{db}
}

func (dao *ProductDAO) GetProducts() []Product {
	results, err := dao.db.Query("SELECT * FROM product")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer results.Close()

	products := []Product{}
	for results.Next() {
		var prod Product
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		if err != nil {
			fmt.Println("Err", err.Error())
			continue
		}
		products = append(products, prod)
	}

	return products
}

func (dao *ProductDAO) GetProduct(code string) *Product {
	prod := &Product{}
	results, err := dao.db.Query("SELECT * FROM product WHERE code=?", code)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer results.Close()

	if results.Next() {
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		if err != nil {
			fmt.Println("Err", err.Error())
			return nil
		}
	} else {
		return nil
	}

	return prod
}

func (dao *ProductDAO) AddProduct(product Product) {
	insert, err := dao.db.Exec(
		"INSERT INTO product (code, name, qty, last_updated) VALUES (?, ?, ?, NOW())",
		product.Code, product.Name, product.Qty)
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}

	affectedRows, _ := insert.RowsAffected()
	if affectedRows == 0 {
		fmt.Println("No rows affected")
	}
}
