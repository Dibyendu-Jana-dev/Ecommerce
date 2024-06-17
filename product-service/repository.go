package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
    var err error
	var port, user, password, dbname, sslmode string
	host := os.Getenv(POSTGRES_HOST)
	port = os.Getenv(POSTGRES_PORT)
	user = os.Getenv(POSTGRES_USER)
	password = os.Getenv(POSTGRES_PASSWORD)
	dbname = os.Getenv(POSTGRES_DB_NAME)
	sslmode = os.Getenv(POSTGRES_SSLMODE)

	databaseURL := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode
    db, err = sql.Open("postgres", databaseURL)
    if err != nil {
        panic(err)
    }
}

func CreateProduct(product *Product) error {
    _, err := db.Exec("INSERT INTO products (name, price, quantity, version) VALUES ($1, $2, $3, $4)",
        product.Name, product.Price, product.Quantity, 1)
    return err
}

func GetProduct(id string) (*Product, error) {
    var product Product
    err := db.QueryRow("SELECT id, name, price, quantity, version FROM products WHERE id=$1", id).Scan(
        &product.ID, &product.Name, &product.Price, &product.Quantity, &product.Version)
    return &product, err
}

func UpdateProduct(product *Product) error {
    result, err := db.Exec(`UPDATE products SET name=$1, price=$2, quantity=$3, version=version+1 WHERE id=$4 AND version=$5`,
        product.Name, product.Price, product.Quantity, product.ID, product.Version)
    if err != nil {
        return err
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return err
    }
    return nil
}

func DeleteProduct(id string) error {
    _, err := db.Exec("DELETE FROM products WHERE id=$1", id)
    return err
}
