package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/register", RegisterHandler).Methods("POST")
    r.HandleFunc("/login", LoginHandler).Methods("POST")
    
    log.Println("Starting authentication service on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
