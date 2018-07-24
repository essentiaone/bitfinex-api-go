package bitfinex

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAccountInfo(t *testing.T) {
	httpDo = func(req *http.Request) (*http.Response, error) {
		msg := `[{
           "maker_fees":"0.1",
           "taker_fees":"0.2",
           "fees":[{
               "pairs":"BTC",
               "maker_fees":"0.1",
               "taker_fees":"0.2"
            },{
               "pairs":"LTC",
               "maker_fees":"0.1",
               "taker_fees":"0.2"
            },{
               "pairs":"ETH",
               "maker_fees":"0.1",
               "taker_fees":"0.2"
            }]
        }]`

		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	info, err := NewClient().Account.Info()
	//info2 := *info

	if err != nil {
		t.Error(err)
	}

	if len((*info)[0].Fees) != 3 {
		t.Error("Expected", 3)
		t.Error("Actual ", len((*info)[0].Fees))
	}
}

func TestAccountFeesSuccess(t *testing.T) {
	httpDo = func(req *http.Request) (*http.Response, error) {
		msg := `{"withdraw":{"BTC":"0.0004","LTC":"0.001","ETH":"0.0027","ETC":"0.01","ZEC":"0.001","XMR":"0.04","DSH":"0.01","XRP":"0.02","IOT":"0.5","EOS":0,"SAN":"1.2858","OMG":"0.17701","BCH":"0.0001","NEO":0,"ETP":"0.01","QTM":"0.01","AVT":"1.3861","EDO":"1.3211","BTG":0,"DAT":"18.126","QSH":"4.4763","YYW":"20.794","GNT":"4.0836","SNT":"17.915","BAT":"3.6035","MNA":"12.345","FUN":"46.949","ZRX":"1.0804","TNB":"52.735","SPK":"14.073","TRX":"34.395","RCN":"24.022","RLC":"1.6026","AID":"10.816","SNG":"27.003","REP":"0.042028","ELF":"2.0032","NEC":"3.5527","IOS":"55.82","AIO":"1.2415","REQ":"18.05","RDN":"1.5841","LRC":"2.939","WAX":"9.9654","DAI":"1.213","CFI":"33.789","AGI":"9.9694","BFT":"14.951","MTN":"15.528","ODE":"3.2154","ANT":"0.68002","DTH":"33.508","MIT":"1.9425","STJ":"2.3495","XLM":0,"XVG":0,"BCI":0,"MKR":"0.0020843","VEN":"0.61794","KNC":"1.3259","POA":"5.812","EVT":"480.99","LYM":"29.316","UTK":"13.484","VEE":"49.551","DAD":"10.046","ORS":"22.534","AUC":"7.5026","POY":"3.5383","FSN":"0.52736","CBT":"41.545","ZCN":"2.6471","SEN":"74.034","NCA":"74.197","CND":"31.769","CTX":"2.0613","PAI":0,"SEE":"180.0","ESS":"58.798","ATD":0,"ADD":0,"MTO":"12155.0","ATM":"36.053","HOT":"41.635","DTA":"33.755","IQX":0}}`

		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	info, err := NewClient().Account.AccountFees()

	if err != nil {
	t.Error(err)
	}

	if info.Withdraw["BTC"] != "0.0004" {
	t.Error("Expected", "0.0004")
	t.Error("Actual ", info.Withdraw["BTC"])
	}
}

func TestAccountFeesFailed(t *testing.T) {
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

	_, err := NewClient().Account.AccountFees()

	if err == nil {
		t.Error("TestAccountFeesFailed failed because of err = nil")
		return
	}
	if err.Error() != "Error from func AccountFees in func do, error: POST http://: 500 some error" {
		t.Error("TestAccountFeesFailed failed because of err =", err)
	}
}
