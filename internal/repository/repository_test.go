package repository

import (
	"StatisticsCollectionService/internal/models"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	connStr := "postgres://postgres:12345678@localhost/stats-collection-test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Error connecting to test database: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS order_books (
		id SERIAL PRIMARY KEY,
		exchange VARCHAR(255) NOT NULL,
		pair VARCHAR(255) NOT NULL,
		asks JSONB NOT NULL,
		bids JSONB NOT NULL,
		UNIQUE (exchange, pair)
	)`)
	if err != nil {
		t.Fatalf("Error creating order_books table: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS order_history (
		id SERIAL PRIMARY KEY,
		client_name VARCHAR(255) NOT NULL,
		exchange_name VARCHAR(255) NOT NULL,
		label VARCHAR(255) NOT NULL,
		pair VARCHAR(255) NOT NULL,
		side VARCHAR(50) NOT NULL,
		type VARCHAR(50) NOT NULL,
		base_qty DOUBLE PRECISION NOT NULL,
		price DOUBLE PRECISION NOT NULL,
		algorithm_name_placed VARCHAR(255) NOT NULL,
		lowest_sell_prc DOUBLE PRECISION NOT NULL,
		highest_buy_prc DOUBLE PRECISION NOT NULL,
		commission_quote_qty DOUBLE PRECISION NOT NULL,
		time_placed TIMESTAMP NOT NULL
	)`)
	if err != nil {
		t.Fatalf("Error creating order_history table: %v", err)
	}

	return db
}

func teardownTestDB(t *testing.T, db *sql.DB) {
	_, err := db.Exec(`DROP TABLE IF EXISTS order_books`)
	if err != nil {
		t.Fatalf("Error dropping order_books table: %v", err)
	}
	_, err = db.Exec(`DROP TABLE IF EXISTS order_history`)
	if err != nil {
		t.Fatalf("Error dropping order_history table: %v", err)
	}
	db.Close()
}

func TestPostgresRepository_GetOrderBook(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewPostgresRepository(db)

	exchangeName := "Binance"
	pair := "BTC/USD"
	asks := []*models.DepthOrder{{Price: 10000.0, BaseQty: 1.0}}
	bids := []*models.DepthOrder{{Price: 10500.0, BaseQty: 2.0}}

	asksJSON, _ := json.Marshal(asks)
	bidsJSON, _ := json.Marshal(bids)

	_, err := db.Exec(`INSERT INTO order_books (exchange, pair, asks, bids) VALUES ($1, $2, $3, $4)`, exchangeName, pair, asksJSON, bidsJSON)
	assert.NoError(t, err)

	orderBook, err := repo.GetOrderBook(exchangeName, pair)
	assert.NoError(t, err)
	assert.Len(t, orderBook, 2)
	assert.Equal(t, asks[0], orderBook[0])
	assert.Equal(t, bids[0], orderBook[1])
}

func TestPostgresRepository_SaveOrderBook(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewPostgresRepository(db)

	exchangeName := "Binance"
	pair := "BTC/USD"
	orderBook := []*models.DepthOrder{
		{Price: 10000.0, BaseQty: 1.0},
		{Price: 10500.0, BaseQty: 2.0},
	}

	err := repo.SaveOrderBook(exchangeName, pair, orderBook)
	assert.NoError(t, err)

	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM order_books WHERE exchange = $1 AND pair = $2`, exchangeName, pair).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestPostgresRepository_GetOrderHistory(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewPostgresRepository(db)

	client := &models.Client{ClientName: "John Doe"}
	order := &models.HistoryOrder{
		ClientName:          "John Doe",
		ExchangeName:        "Binance",
		Label:               "order1",
		Pair:                "BTC/USD",
		Side:                "buy",
		Type:                "limit",
		BaseQty:             1.0,
		Price:               10000.0,
		AlgorithmNamePlaced: "alg1",
		LowestSellPrice:     9900.0,
		HighestBuyPrice:     10050.0,
		CommissionQuoteQty:  10.0,
		TimePlaced:          time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	_, err := db.Exec(`INSERT INTO order_history (client_name, exchange_name, label, pair, side, type, base_qty, price, algorithm_name_placed, lowest_sell_prc, highest_buy_prc, commission_quote_qty, time_placed) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		order.ClientName,
		order.ExchangeName,
		order.Label,
		order.Pair,
		order.Side,
		order.Type,
		order.BaseQty,
		order.Price,
		order.AlgorithmNamePlaced,
		order.LowestSellPrice,
		order.HighestBuyPrice,
		order.CommissionQuoteQty,
		order.TimePlaced,
	)
	assert.NoError(t, err)

	history, err := repo.GetOrderHistory(client)
	assert.NoError(t, err)
	assert.Len(t, history, 1)

	order.TimePlaced = order.TimePlaced.UTC()
	history[0].TimePlaced = history[0].TimePlaced.UTC()

	assert.Equal(t, order, history[0])
}

func TestPostgresRepository_SaveOrder(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewPostgresRepository(db)

	client := &models.Client{ClientName: "John Doe"}
	order := &models.HistoryOrder{
		ClientName:          "John Doe",
		ExchangeName:        "Binance",
		Label:               "order1",
		Pair:                "BTC/USD",
		Side:                "buy",
		Type:                "limit",
		BaseQty:             1.0,
		Price:               10000.0,
		AlgorithmNamePlaced: "alg1",
		LowestSellPrice:     9900.0,
		HighestBuyPrice:     10050.0,
		CommissionQuoteQty:  10.0,
		TimePlaced:          time.Now(),
	}

	err := repo.SaveOrder(client, order)
	assert.NoError(t, err)

	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM order_history WHERE client_name = $1`, client.ClientName).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}
