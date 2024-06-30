package api

import (
	"StatisticsCollectionService/internal/models"
	"StatisticsCollectionService/internal/services"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetOrderBook(exchangeName, pair string) ([]*models.DepthOrder, error) {
	args := m.Called(exchangeName, pair)
	return args.Get(0).([]*models.DepthOrder), args.Error(1)
}

func (m *MockService) SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) error {
	args := m.Called(exchangeName, pair, orderBook)
	return args.Error(0)
}

func (m *MockService) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	args := m.Called(client)
	return args.Get(0).([]*models.HistoryOrder), args.Error(1)
}

func (m *MockService) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	args := m.Called(client, order)
	return args.Error(0)
}

func TestGetOrderBookHandler(t *testing.T) {
	mockService := new(MockService)
	service := &services.Service{Repo: mockService}

	exchangeName := "binance"
	pair := "BTC/USDT"
	expectedOrderBook := []*models.DepthOrder{
		{Price: 50000, BaseQty: 0.1},
		{Price: 50500, BaseQty: 0.2},
	}

	mockService.On("GetOrderBook", exchangeName, pair).Return(expectedOrderBook, nil)

	req, err := http.NewRequest("GET", "/orderbook/get?exchange_name="+exchangeName+"&pair="+pair, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := GetOrderBookHandler(service)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var result []*models.DepthOrder
	err = json.NewDecoder(rr.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, expectedOrderBook, result)

	mockService.AssertExpectations(t)
}

func TestSaveOrderBookHandler(t *testing.T) {
	mockService := new(MockService)
	service := &services.Service{Repo: mockService}

	exchangeName := "binance"
	pair := "BTC/USDT"
	orderBook := []*models.DepthOrder{
		{Price: 50000, BaseQty: 0.1},
		{Price: 49000, BaseQty: 0.2},
	}

	mockService.On("SaveOrderBook", exchangeName, pair, orderBook).Return(nil)

	requestBody, err := json.Marshal(map[string]interface{}{
		"exchange_name": exchangeName,
		"pair":          pair,
		"order_book":    orderBook,
	})
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/orderbook/save", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := SaveOrderBookHandler(service)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	mockService.AssertExpectations(t)
}

func TestGetOrderHistoryHandler(t *testing.T) {
	mockService := new(MockService)
	service := &services.Service{Repo: mockService}

	client := &models.Client{ClientName: "test_client"}
	expectedHistory := []*models.HistoryOrder{
		{ClientName: "test_client", ExchangeName: "binance", Pair: "BTC/USDT"},
	}

	mockService.On("GetOrderHistory", client).Return(expectedHistory, nil)

	requestBody, err := json.Marshal(client)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/orderhistory/get", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := GetOrderHistoryHandler(service)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var result []*models.HistoryOrder
	err = json.NewDecoder(rr.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, expectedHistory, result)

	mockService.AssertExpectations(t)
}

func TestSaveOrderHandler(t *testing.T) {
	mockService := new(MockService)
	service := &services.Service{Repo: mockService}

	order := &models.HistoryOrder{
		ClientName: "test_client", ExchangeName: "binance", Pair: "BTC/USDT",
	}

	client := &models.Client{
		ClientName:   order.ClientName,
		ExchangeName: order.ExchangeName,
		Label:        order.Label,
		Pair:         order.Pair,
	}

	mockService.On("SaveOrder", client, order).Return(nil)

	requestBody, err := json.Marshal(order)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/order/save", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := SaveOrderHandler(service)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	mockService.AssertExpectations(t)
}
