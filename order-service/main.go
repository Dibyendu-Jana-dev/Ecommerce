package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/orders", CreateOrderHandler).Methods("POST")
    r.HandleFunc("/orders/{id}", GetOrderHandler).Methods("GET")
    
    log.Println("Starting order processing service on port 8002")
    log.Fatal(http.ListenAndServe(":8002", r))
}
