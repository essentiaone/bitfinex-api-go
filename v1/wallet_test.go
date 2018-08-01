package bitfinex

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestWalletTransfer(t *testing.T) {
	httpDo = func(req *http.Request) (*http.Response, error) {
		msg := `[{
          "status":"success",
          "message":"1.0 USD transfered from Exchange to Deposit"
        }]`
		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	tr := TransferRequest{10.0, "BTC", "1WalletA", "1WalletB"}
	response, err := NewClient().Wallet.Transfer(tr)

	if err != nil {
		t.Error(err)
	}

	if response[0].Status != "success" {
		t.Error("Expected", "success")
		t.Error("Actual ", response[0].Status)
	}
}

func TestWithdrawCryptoSuccess(t *testing.T) {
	httpDo = func(req *http.Request) (*http.Response, error) {
		msg := `[{
          "status":"success",
          "message":"Your withdrawal request has been successfully submitted.",
          "withdrawal_id":586829
        }]`
		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	wd := WithdrawRequest{10.0, "bitcoin", WALLET_DEPOSIT, "1WalletABC"}
	response, err := NewClient().Wallet.WithdrawCrypto(wd)

	if err != nil {
		t.Error(err)
	}

	if response[0].Status != "success" {
		t.Error("Expected", "success")
		t.Error("Actual ", response[0].Status)
	}
	if response[0].WithdrawalID != 586829 {
		t.Error("Expected", 586829)
		t.Error("Actual ", response[0].WithdrawalID)
	}

}

func TestWithdrawCryptoError(t *testing.T) {
	httpDo = func(req *http.Request) (*http.Response, error) {
		msg := `[{
          "status":"error",
          "message":"Your withdrawal request has errors.",
			"withdrawal_id":0
        }]`
		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	wd := WithdrawRequest{10.0, "bitcoin", WALLET_DEPOSIT, "1WalletABC"}
	_, err := NewClient().Wallet.WithdrawCrypto(wd)

	if err == nil {
		t.Error("TestWithdrawCryptoError failed because of err = nil")
		return
	}
	if err.Error() != "Error from func Withdraw in func check status, error: withdraw status error Your withdrawal request has errors." {
		t.Error("Expected", "Error from func Withdraw in func check status, error: withdraw status error")
		t.Error("Actual ", err.Error())
	}
}

func TestWithdrawCryptoErrorFailed(t *testing.T) {
	httpDo = func(req *http.Request) (*http.Response, error) {
		msg := `{"response":
					{"response":
						{"request":
							{"Status"  :"500 err",
							"StatusCode": 500, 
							"Method":"POST",
							"URL":{"scheme": "http"} 
							}
						}
					},	
					"message":"some error"}`

		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 500,
		}
		return &resp, nil
	}

	wd := WithdrawRequest{10.0, "bitcoin", WALLET_DEPOSIT, "1WalletABC"}
	_, err := NewClient().Wallet.WithdrawCrypto(wd)

	if err == nil {
		t.Error("TestWithdrawCryptoError failed because of err = nil")
		return
	}
	if err.Error() != "Error from func Withdraw in func do, error: POST http://: 500 some error" {
		t.Error("Expected","Error from func Withdraw in func do, error: POST http://: 500 some error")
		t.Error("Actual ", err.Error())
	}
}

