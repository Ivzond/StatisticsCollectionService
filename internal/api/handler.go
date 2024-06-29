package api

import (
	"StatisticsCollectionService/internal/models"
	"StatisticsCollectionService/internal/services"
	"encoding/json"
	"net/http"
)

// @Summary Получить книгу ордеров
// @Description Получить книгу ордеров для указанной биржи и пары валют
// @Param exchange_name query string true "Имя биржи"
// @Param pair query string true "Валютная пара"
// @Success 200 {array} models.DepthOrder
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /orderbook/get [get]
func GetOrderBookHandler(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		exchangeName := r.URL.Query().Get("exchange_name")
		pair := r.URL.Query().Get("pair")

		orderBook, err := service.GetOrderBook(exchangeName, pair)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(orderBook)
	}
}

// @Summary Сохранить книгу ордеров
// @Description Сохранить книгу ордеров для указанной биржи и пары валют
// @Param order body models.OrderBook true "Книга ордеров"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /orderbook/save [post]
func SaveOrderBookHandler(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			ExchangeName string               `json:"exchange_name"`
			Pair         string               `json:"pair"`
			OrderBook    []*models.DepthOrder `json:"order_book"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := service.SaveOrderBook(request.ExchangeName, request.Pair, request.OrderBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// @Summary Получить историю ордеров
// @Description Получить историю ордеров для указанного клиента
// @Param client body models.Client true "Клиент"
// @Success 200 {array} models.HistoryOrder
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /orderhistory/get [get]
func GetOrderHistoryHandler(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var client models.Client
		if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		history, err := service.GetOrderHistory(&client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(history)
	}
}

// @Summary Сохранить ордер
// @Description Сохранить новый ордер для указанного клиента
// @Param order body models.HistoryOrder true "Ордер"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /order/save [post]
func SaveOrderHandler(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order models.HistoryOrder
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		client := &models.Client{
			ClientName:   order.ClientName,
			ExchangeName: order.ExchangeName,
			Label:        order.Label,
			Pair:         order.Pair,
		}
		err := service.SaveOrder(client, &order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
