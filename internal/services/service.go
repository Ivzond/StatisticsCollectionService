package services

import (
	"StatisticsCollectionService/internal/models"
	"StatisticsCollectionService/internal/repository"
)

// Структура сервиса, предоставляющая бизнес-логику
type Service struct {
	Repo repository.Repository
}

// Конструктор для создания нового сервиса
func NewService(repo repository.Repository) *Service {
	return &Service{Repo: repo}
}

// Метод для получения книги ордеров
func (s *Service) GetOrderBook(exchangeName, pair string) ([]*models.DepthOrder, error) {
	return s.Repo.GetOrderBook(exchangeName, pair)
}

// Метод для сохранения книги ордеров
func (s *Service) SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) error {
	return s.Repo.SaveOrderBook(exchangeName, pair, orderBook)
}

// Метод для получения истории ордеров
func (s *Service) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	return s.Repo.GetOrderHistory(client)
}

// Метод для сохранения ордера
func (s *Service) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	return s.Repo.SaveOrder(client, order)
}
