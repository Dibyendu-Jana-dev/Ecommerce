package main

type Product struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Price    int    `json:"price"`
    Quantity int    `json:"quantity"`
    Version  int    `json:"version"`
}
