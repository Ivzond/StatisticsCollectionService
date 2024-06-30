package services

import (
	"StatisticsCollectionService/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetOrderBook(exchangeName, pair string) ([]*models.DepthOrder, error) {
	args := m.Called(exchangeName, pair)
	return args.Get(0).([]*models.DepthOrder), args.Error(1)
}

func (m *MockRepository) SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) error {
	args := m.Called(exchangeName, pair, orderBook)
	return args.Error(0)
}

func (m *MockRepository) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	args := m.Called(client)
	return args.Get(0).([]*models.HistoryOrder), args.Error(1)
}

func (m *MockRepository) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	args := m.Called(client, order)
	return args.Error(0)
}

func TestService_GetOrderBook(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	exchangeName := "Binance"
	pair := "BTC/USD"
	expectedOrders := []*models.DepthOrder{
		{Price: 10.0, BaseQty: 1.0},
		{Price: 15.0, BaseQty: 2.0},
	}
	mockRepo.On("GetOrderBook", exchangeName, pair).Return(expectedOrders, nil)
	orders, err := service.GetOrderBook(exchangeName, pair)
	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, orders)

	mockRepo.AssertExpectations(t)
}

func TestService_SaveOrderBook(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	exchangeName := "Binance"
	pair := "BTC/USD"
	orderBook := []*models.DepthOrder{
		{Price: 10.0, BaseQty: 1.0},
		{Price: 15.0, BaseQty: 2.0},
	}

	mockRepo.On("SaveOrderBook", exchangeName, pair, orderBook).Return(nil)
	err := service.SaveOrderBook(exchangeName, pair, orderBook)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_GetOrderHistory(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	client := &models.Client{ClientName: "John Doe"}
	expectedHistory := []*models.HistoryOrder{
		{
			ClientName:          "John Doe",
			ExchangeName:        "Binance",
			Label:               "order1",
			Pair:                "BTC/USD",
			Side:                "buy",
			Type:                "limit",
			BaseQty:             1.0,
			Price:               10.0,
			AlgorithmNamePlaced: "alg1",
			LowestSellPrice:     8.0,
			HighestBuyPrice:     15.0,
			CommissionQuoteQty:  1.0,
			TimePlaced:          time.Now(),
		},
	}

	mockRepo.On("GetOrderHistory", client).Return(expectedHistory, nil)
	history, err := service.GetOrderHistory(client)
	assert.NoError(t, err)
	assert.Equal(t, expectedHistory, history)
	mockRepo.AssertExpectations(t)
}

func TestService_SaveOrder(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	order := &models.HistoryOrder{
		ClientName:          "John Doe",
		ExchangeName:        "Binance",
		Label:               "order1",
		Pair:                "BTC/USD",
		Side:                "buy",
		Type:                "limit",
		BaseQty:             1.0,
		Price:               10.0,
		AlgorithmNamePlaced: "alg1",
		LowestSellPrice:     8.0,
		HighestBuyPrice:     15.0,
		CommissionQuoteQty:  1.0,
		TimePlaced:          time.Now(),
	}
	client := &models.Client{
		ClientName:   order.ClientName,
		ExchangeName: order.ExchangeName,
		Label:        order.Label,
		Pair:         order.Pair,
	}

	mockRepo.On("SaveOrder", client, order).Return(nil)
	err := service.SaveOrder(client, order)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
