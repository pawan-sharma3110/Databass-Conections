package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	connStr := "host=localhost port=8080 dbname=TestDB user=postgres password=Pawan@2003 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	CreateProductTable(db)
	data := []Product{}
	rows, err := db.Query("SELECT name, available,price FROM product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var name string
	var available bool
	var price float64
	for rows.Next() {
		err := rows.Scan(&name, &available, &price)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Product{name, price, available})
	}
	print(data)
}

// product := Product{"Book", 1.55, true}
// 	pk := insertProduct(db, product)

func CreateProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product(
		id SERIAL PRIMARY KEY,
		name VARCHAR (100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		available BOOLEAN,
		created TIMESTAMP DEFAULT NOW ()
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// func insertProduct(db *sql.DB, product Product) int {
// 	query := `INSERT INTO product (name,price,available)
// 	VALUES($1,$2,$3) RETURNING id`

// 	var pk int
// 	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return pk

// }
