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
	Amount   float64
	Currency string
	From     string
	To       string
}
// Transfer funds between wallets
func (c *WalletService) Transfer(request TransferRequest) ([]TransferStatus, error) {

	payload := map[string]interface{}{
		"amount":     strconv.FormatFloat(request.Amount, 'f', -1, 32),
		"currency":   request.Currency,
		"walletfrom": request.From,
		"walletto":   request.To,
	}

	req, err := c.client.newAuthenticatedRequest("POST", "transfer", payload)

	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Transfer", FuncWhat:"newAuthenticatedRequest", FuncError: err.Error()}
	}

	status := make([]TransferStatus, 0)

	_, err = c.client.do(req, &status)
	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Transfer", FuncWhat:"do", FuncError: err.Error()}
	}
	for _,transfer := range status {
		if transfer.Status == "error" {
			return nil, &ErrorHandler{FuncWhere: "Transfer", FuncWhat:"check status", FuncError: "transfer status error " + transfer.Message}
		}
	}
	return status, nil
}

type WithdrawStatus struct {
	Status       string
	Message      string
	WithdrawalID int `json:"withdrawal_id"`
}

type WithdrawRequest struct {
	Amount             float64
	Currency           string
	Wallet             string
	DestinationAddress string
}
// Withdraw a cryptocurrency to a digital wallet
func (c *WalletService) WithdrawCrypto(request WithdrawRequest) ([]WithdrawStatus, error) {

	payload := map[string]interface{}{
		"amount":         strconv.FormatFloat(request.Amount, 'f', -1, 32),
		"walletselected": request.Wallet,
		"withdraw_type":  request.Currency,
		"address":        request.DestinationAddress,
	}

	req, err := c.client.newAuthenticatedRequest("POST", "withdraw", payload)

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
			return nil, &ErrorHandler{FuncWhere: "Withdraw", FuncWhat:"check status", FuncError: "withdraw status error " + withdraw.Message}
		}
	}

	return status, nil

}

