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

func CreateOrder(order *Order) error {
    _, err := db.Exec("INSERT INTO orders (product_id, quantity, status) VALUES ($1, $2, $3)",
        order.ProductID, order.Quantity, "pending")
    return err
}

func GetOrder(id string) (*Order, error) {
    var order Order
    err := db.QueryRow("SELECT id, product_id, quantity, status FROM orders WHERE id=$1", id).Scan(
        &order.ID, &order.ProductID, &order.Quantity, &order.Status)
    return &order, err
}
