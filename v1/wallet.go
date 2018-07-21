package bitfinex

import "strconv"

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

// Transfer funds between wallets
func (c *WalletService) Transfer(amount float64, currency, from, to string) ([]TransferStatus, error) {

	payload := map[string]interface{}{
		"amount":     strconv.FormatFloat(amount, 'f', -1, 32),
		"currency":   currency,
		"walletfrom": from,
		"walletto":   to,
	}

	req, err := c.client.newAuthenticatedRequest("GET", "transfer", payload)

	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Transfer", FuncWhat:"newAuthenticatedRequest", FuncError: err}
	}

	status := make([]TransferStatus, 0)

	_, err = c.client.do(req, &status)
	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Transfer", FuncWhat:"do", FuncError: err}
	}

	return status, nil
}

type WithdrawStatus struct {
	Status       string
	Message      string
	WithdrawalID int `json:"withdrawal_id"`
}

// Withdraw a cryptocurrency to a digital wallet
func (c *WalletService) WithdrawCrypto(amount float64, currency, wallet, destinationAddress string) ([]WithdrawStatus, error) {

	payload := map[string]interface{}{
		"amount":         strconv.FormatFloat(amount, 'f', -1, 32),
		"walletselected": wallet,
		"withdraw_type":  currency,
		"address":        destinationAddress,
	}

	req, err := c.client.newAuthenticatedRequest("GET", "withdraw", payload)

	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Withdraw", FuncWhat:"newAuthenticatedRequest", FuncError: err}
	}

	status := make([]WithdrawStatus, 0)

	_, err = c.client.do(req, &status)
	if err != nil {
		return nil, &ErrorHandler{FuncWhere: "Withdraw", FuncWhat:"do", FuncError: err}
	}
	for _, withdraw := range status {
		if withdraw.Status == "error" {
			return nil, &ErrorHandler{FuncWhere: "Withdraw", FuncWhat:"check status", FuncError: err}
		}
	}

	return status, nil

}

