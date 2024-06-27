package repository

import "StatisticsCollectionService/internal/models"

type Repository interface {
	GetOrderBook(exchangeName, pair string) ([]*models.DepthOrder, error)
	SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) error
	GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error)
	SaveOrder(client *models.Client, order *models.HistoryOrder) error
}
