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

func CreateUser(username, password string) error {
    _, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
    return err
}

func GetUserPassword(username string) (string, error) {
    var password string
    err := db.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&password)
    return password, err
}
