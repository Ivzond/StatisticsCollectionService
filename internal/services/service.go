package services

import (
	"StatisticsCollectionService/internal/models"
	"StatisticsCollectionService/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func (s *Service) GetOrderBook(exchangeName, pair string) ([]*models.DepthOrder, error) {
	return s.repo.GetOrderBook(exchangeName, pair)
}

func (s *Service) SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) ([]*models.DepthOrder, error) {
	return s.repo.SaveOrderBook(exchangeName, pair, orderBook)
}

func (s *Service) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	return s.repo.GetOrderHistory(client)
}

func (s *Service) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	return s.repo.SaveOrder(client, order)
}
