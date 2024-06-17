package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/products", CreateProductHandler).Methods("POST")
    r.HandleFunc("/products/{id}", GetProductHandler).Methods("GET")
    r.HandleFunc("/products/{id}", UpdateProductHandler).Methods("PUT")
    r.HandleFunc("/products/{id}", DeleteProductHandler).Methods("DELETE")
    
    log.Println("Starting product management service on port 8001")
    log.Fatal(http.ListenAndServe(":8001", r))
}
