package bitfinex


type AccountService struct {
	client *Client
}

type AccountPairFee struct {
	Pair      string
	MakerFees float64 `json:"maker_fees,string"`
	TakerFees float64 `json:"taker_fees,string"`
}

type AccountInfo struct {
	MakerFees float64 `json:"maker_fees,string"`
	TakerFees float64 `json:"taker_fees,string"`
	Fees      []AccountPairFee
}

// GET account_infos
func (a *AccountService) Info() (AccountInfo, error) {
	req, err := a.client.newAuthenticatedRequest("GET", "account_infos", nil)

	if err != nil {
		return AccountInfo{}, err
	}

	var v []AccountInfo
	_, err = a.client.do(req, &v)

	if err != nil {
		return AccountInfo{}, err
	}

	return v[0], nil
}

type Fees struct {
	Withdraw  map[string]interface{} `json:"withdraw"`
}

// AccountFees return withdraw and deposit(are equal for 20.07.2018) fee
func (a *AccountService) AccountFees() (Fees, error) {
	req, err := a.client.newAuthenticatedRequest("POST", "account_fees", nil)

	if err != nil {
		return Fees{}, err
	}
	var f Fees
	_, err = a.client.do(req, &f)
	if err != nil {
		return Fees{}, err
	}

	return f, nil
}
