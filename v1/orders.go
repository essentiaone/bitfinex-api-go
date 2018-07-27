package bitfinex

import (
	"strconv"
)

// Order types that the API can return.
const (
	OrderTypeExchangeLimit        = "exchange limit"
	)

// OrderService manages the Order endpoint.
type OrderService struct {
	client *Client
}

// Order represents one order on the bitfinex platform.
type Order struct {
	ID                int64
	Symbol            string
	Exchange          string
	Price             string
	AvgExecutionPrice string `json:"avg_execution_price"`
	Side              string
	Type              string
	Timestamp         string
	IsLive            bool   `json:"is_live"`
	IsCanceled        bool   `json:"is_cancelled"`
	IsHidden          bool   `json:"is_hidden"`
	WasForced         bool   `json:"was_forced"`
	OriginalAmount    string `json:"original_amount"`
	RemainingAmount   string `json:"remaining_amount"`
	ExecutedAmount    string `json:"executed_amount"`
}

type OrderRequest struct {
	Symbol    string
	Amount    float64
	Price     float64
	OrderType string
	Side      string
}

// Create a new order.
func (s *OrderService) Create(request OrderRequest) (*Order, error) {

	payload := map[string]interface{}{
		"symbol":   request.Symbol,
		"amount":   strconv.FormatFloat(request.Amount, 'f', -1, 32),
		"price":    strconv.FormatFloat(request.Price, 'f', -1, 32),
		"side":     request.Side,
		"type":     request.OrderType,
		"exchange": "bitfinex",
	}

	req, err := s.client.newAuthenticatedRequest("POST", "order/new", payload)
	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Order Create", FuncWhat:"newAuthenticatedRequest", FuncError: err.Error()}
	}

	order := &Order{}
	_, err = s.client.do(req, order)
	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Order Create", FuncWhat:"do", FuncError: err.Error()}
	}

	return order, nil
}
