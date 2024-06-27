package repository

import (
	"StatisticsCollectionService/internal/models"
	"database/sql"
	"encoding/json"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetOrderBook(exchangeName, pair string) ([]*models.DepthOrder, error) {
	query := `SELECT asks, bids FROM order_books WHERE exchange = $1 AND pair = $2`
	row := r.db.QueryRow(query, exchangeName, pair)

	var asksJSON, bidsJSON []byte
	err := row.Scan(&asksJSON, &bidsJSON)
	if err != nil {
		return nil, err
	}

	var asks, bids []*models.DepthOrder
	err = json.Unmarshal(asksJSON, &asks)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bidsJSON, &bids)
	if err != nil {
		return nil, err
	}

	return append(asks, bids...), nil
}

func (r *PostgresRepository) SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) error {
	asksJSON, err := json.Marshal(orderBook[:len(orderBook)/2])
	if err != nil {
		return err
	}
	bidsJSON, err := json.Marshal(orderBook[len(orderBook)/2:])
	if err != nil {
		return err
	}
	query := `INSERT INTO order_books (exchange, pair, asks, bids) VALUES ($1, $2, $3, $4)`
	_, err = r.db.Exec(query, exchangeName, pair, asksJSON, bidsJSON)
	return err
}

func (r *PostgresRepository) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	query := `SELECT client_name, exchange_name, label, pair, side, type, base_qty, price, algorithm_name_placed, lowest_sell_prc, highest_buy_prc, commission_quote_qty, time_placed FROM order_history WHERE client_name = $1`
	rows, err := r.db.Query(query, client.ClientName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.HistoryOrder
	for rows.Next() {
		var order models.HistoryOrder
		err := rows.Scan(
			&order.ClientName,
			&order.ExchangeName,
			&order.Label,
			&order.Pair,
			&order.Side,
			&order.Type,
			&order.BaseQty,
			&order.Price,
			&order.AlgorithmNamePlaced,
			&order.LowestSellPrice,
			&order.HighestBuyPrice,
			&order.CommissionQuoteQty,
			&order.TimePlaced,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (r *PostgresRepository) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	//TODO implement me
	panic("implement me")
}
