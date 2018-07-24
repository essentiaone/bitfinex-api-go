package bitfinex

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"testing"
)

func TestOrderService_CreateSuccess(t *testing.T) {
	httpDo = func(req *http.Request) (*http.Response, error) {
		msg := `{
            "ID": 14647523792,
			"Symbol": "ethbtc",
			"Exchange": "bitfinex",
			"Price": "0.060237",
			"avg_execution_price": "0.0",
			"Side": "sell",
			"Type": "exchange market",
			"Timestamp": "1532343485.784226238",
			"is_live": true,
			"is_cancelled": false,
			"is_hidden": false,
			"was_forced": false,
			"original_amount": "0.25",
			"remaining_amount": "0.25",
			"executed_amount": "0.0"
        }`
		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	or := OrderRequest{"ETHBTC",0.0316, 1,"exchange market", "sell"}
	resp, err := NewClient().Orders.Create(or)

	if err != nil {
		t.Error(err)
	}

	r := Order{
		ID: 14647523792,
		Symbol: "ethbtc",
		Exchange: "bitfinex",
		Price: "0.060237",
		AvgExecutionPrice: "0.0",
		Side: "sell",
		Type: "exchange market",
		Timestamp: "1532343485.784226238",
		IsLive: true,
		IsCanceled: false,
		IsHidden: false,
		WasForced: false,
		OriginalAmount: "0.25",
		RemainingAmount: "0.25",
		ExecutedAmount: "0.0"}

	if *resp != r {
		t.Error("Expected", r)
		t.Error("Actual ", resp)
	}
}

func TestOrderService_CreateFailed(t *testing.T) {
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

	or := OrderRequest{"ETHBTC", 0.0316, 1, "exchange market", "sell"}
	_, err := NewClient().Orders.Create(or)

	if err == nil {
		t.Error("TestOrderService_CreateFailed failed because of err = nil")
		return
	}
	if err.Error() != "Error from func Order Create in func do, error: POST http://: 500 some error" {
		t.Error("Expected","Error from func Order Create in func do, error: POST http://: 500 some error")
		t.Error("Actual ", err.Error())
	}
}