package main

import (
	"StatisticsCollectionService/internal/api"
	"StatisticsCollectionService/internal/db"
	"StatisticsCollectionService/internal/repository"
	"StatisticsCollectionService/internal/services"
	"log"
	"net/http"

	_ "StatisticsCollectionService/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API Сервиса Сбора Статистики
// @version 1.0
// @description Это микросервис на golang для сбора статистики
// @host localhost:8080
// @BasePath /

// Основная функция для запуска сервера
func main() {
	db.InitDB()
	defer db.DB.Close()

	repo := repository.NewPostgresRepository(db.DB)
	service := services.NewService(repo)

	http.HandleFunc("/orderbook/get", api.GetOrderBookHandler(service))
	http.HandleFunc("/orderbook/save", api.SaveOrderBookHandler(service))
	http.HandleFunc("/orderhistory/get", api.GetOrderHistoryHandler(service))
	http.HandleFunc("/order/save", api.SaveOrderHandler(service))

	// Swagger endpoint
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
