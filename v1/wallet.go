package bitfinex

import (
	"strconv"
)

const (
	WALLET_DEPOSIT  = "deposit"
)

type WalletService struct {
	client *Client
}

type TransferStatus struct {
	Status  string
	Message string
}

type TransferRequest struct {
	amount   float64
	currency string
	from     string
	to       string
}
// Transfer funds between wallets
func (c *WalletService) Transfer(request TransferRequest) ([]TransferStatus, error) {

	payload := map[string]interface{}{
		"amount":     strconv.FormatFloat(request.amount, 'f', -1, 32),
		"currency":   request.currency,
		"walletfrom": request.from,
		"walletto":   request.to,
	}

	req, err := c.client.newAuthenticatedRequest("GET", "transfer", payload)

	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Transfer", FuncWhat:"newAuthenticatedRequest", FuncError: err.Error()}
	}

	status := make([]TransferStatus, 0)

	_, err = c.client.do(req, &status)
	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Transfer", FuncWhat:"do", FuncError: err.Error()}
	}

	return status, nil
}

type WithdrawStatus struct {
	Status       string
	Message      string
	WithdrawalID int `json:"withdrawal_id"`
}

type WithdrawRequest struct {
	amount             float64
	currency           string
	wallet             string
	destinationAddress string
}
// Withdraw a cryptocurrency to a digital wallet
func (c *WalletService) WithdrawCrypto(request WithdrawRequest) ([]WithdrawStatus, error) {

	payload := map[string]interface{}{
		"amount":         strconv.FormatFloat(request.amount, 'f', -1, 32),
		"walletselected": request.wallet,
		"withdraw_type":  request.currency,
		"address":        request.destinationAddress,
	}

	req, err := c.client.newAuthenticatedRequest("GET", "withdraw", payload)

	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Withdraw", FuncWhat:"newAuthenticatedRequest", FuncError: err.Error()}
	}

	status := make([]WithdrawStatus, 0)

	_, err = c.client.do(req, &status)
	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Withdraw", FuncWhat:"do", FuncError: err.(*ErrorResponse).Error()}
	}
	for _, withdraw := range status {
		if withdraw.Status == "error" {
			return nil, &ErrorHandler{FuncWhere: "Withdraw", FuncWhat:"check status", FuncError: "withddraw status error"}
		}
	}

	return status, nil

}

