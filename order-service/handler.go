package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
    var order Order
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if err := CreateOrder(&order); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    order, err := GetOrder(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(order)
}
