package main

import (
	"StatisticsCollectionService/internal/api"
	"StatisticsCollectionService/internal/db"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	http.HandleFunc("/orderbook/get", api.GetOrderBookHandler)
	http.HandleFunc("/orderbook/save", api.SaveOrderBookHandler)
	http.HandleFunc("/orderhistory/get", api.GetOrderHistory)
	http.HandleFunc("/order/save", api.SaveOrder)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
