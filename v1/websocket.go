package bitfinex

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/bitfinexcom/bitfinex-api-go/utils"

	"github.com/gorilla/websocket"
)

// Pairs available
const (
	// Pairs
	BTCUSD = "BTCUSD"
	LTCUSD = "LTCUSD"
	LTCBTC = "LTCBTC"
	ETHUSD = "ETHUSD"
	ETHBTC = "ETHBTC"
	ETCUSD = "ETCUSD"
	ETCBTC = "ETCBTC"
	BFXUSD = "BFXUSD"
	BFXBTC = "BFXBTC"
	ZECUSD = "ZECUSD"
	ZECBTC = "ZECBTC"
	XMRUSD = "XMRUSD"
	XMRBTC = "XMRBTC"
	RRTUSD = "RRTUSD"
	RRTBTC = "RRTBTC"
	XRPUSD = "XRPUSD"
	XRPBTC = "XRPBTC"
	EOSETH = "EOSETH"
	EOSUSD = "EOSUSD"
	EOSBTC = "EOSBTC"
	IOTUSD = "IOTUSD"
	IOTBTC = "IOTBTC"
	IOTETH = "IOTETH"
	BCCBTC = "BCCBTC"
	BCUBTC = "BCUBTC"
	BCCUSD = "BCCUSD"
	BCUUSD = "BCUUSD"
	SANBTC = "SANBTC"
	SANETH = "SANETH"
	OMGBTC = "OMGBTC"
	OMGETH = "OMGETH"
	AVTBTC = "AVTBTC"
	AVTETH = "AVTETH"
	EDOBTC = "EDOBTC"
	EDOETH = "EDOETH"
	YYWBTC = "YYWBTC"
	YYWETH = "YYWETH"
	GNTBTC = "GNTBTC"
	GNTETH = "GNTETH"
	SNTBTC = "SNTBTC"
	SNTETH = "SNTETH"
	BATBTC = "BATBTC"
	BATETH = "BATETH"
	FUNBTC = "FUNBTC"
	FUNETH = "FUNETH"
	ZRXBTC = "ZRXBTC"
	ZRXETH = "ZRXETH"
	TNBBTC = "TNBBTC"
	TNBETH = "TNBETH"
	RCNBTC = "RCNBTC"
	RCNETH = "RCNETH"
	AIDBTC = "AIDBTC"
	AIDETH = "AIDETH"
	REPBTC = "REPBTC"
	REPETH = "REPETH"
	ELFBTC = "ELFBTC"
	ELFETH = "ELFETH"
	IOSBTC = "IOSBTC"
	IOSETH = "IOSETH"
	AIOBTC = "AIOBTC"
	AIOETH = "AIOETH"
	REQBTC = "REQBTC"
	REQETH = "REQETH"
	RDNBTC = "RDNBTC"
	RDNETH = "RDNETH"
	LRCBTC = "LRCBTC"
	LRCETH = "LRCETH"
	WAXBTC = "WAXBTC"
	WAXETH = "WAXETH"
	DAIBTC = "DAIBTC"
	DAIETH = "DAIETH"
	CFIBTC = "CFIBTC"
	CFIETH = "CFIETH"
	AGIBTC = "AGIBTC"
	AGIETH = "AGIETH"
	BFTBTC = "BFTBTC"
	BFTETH = "BFTETH"
	MTNBTC = "MTNBTC"
	MTNETH = "MTNETH"
	ODEBTC = "ODEBTC"
	ODEETH = "ODEETH"
	ANTBTC = "ANTBTC"
	ANTETH = "ANTETH"
	DTHBTC = "DTHBTC"
	DTHETH = "DTHETH"
	MITBTC = "MITBTC"
	MITETH = "MITETH"
	MKRBTC = "MKRBTC"
	MKRETH = "MKRETH"
	KNCBTC = "KNCBTC"
	KNCETH = "KNCETH"
	LYMBTC = "LYMBTC"
	LYMETH = "LYMETH"
	UTKBTC = "UTKBTC"
	UTKETH = "UTKETH"
	VEEBTC = "VEEBTC"
	VEEETH = "VEEETH"
	DADBTC = "DADBTC"
	DADETH = "DADETH"
	ORSBTC = "ORSBTC"
	ORSETH = "ORSETH"
	AUCBTC = "AUCBTC"
	AUCETH = "AUCETH"
	POYBTC = "POYBTC"
	POYETH = "POYETH"
	FSNBTC = "FSNBTC"
	FSNETH = "FSNETH"
	CBTBTC = "CBTBTC"
	CBTETH = "CBTETH"
	ZCNBTC = "ZCNBTC"
	ZCNETH = "ZCNETH"
	SENBTC = "SENBTC"
	SENETH = "SENETH"
	NCABTC = "NCABTC"
	NCAETH = "NCAETH"
	CNDBTC = "CNDBTC"
	CNDETH = "CNDETH"
	CTXBTC = "CTXBTC"
	CTXETH = "CTXETH"
	SEEBTC = "SEEBTC"
	SEEETH = "SEEETH"
	ESSBTC = "ESSBTC"
	ESSETH = "ESSETH"
	ATMBTC = "ATMBTC"
	ATMETH = "ATMETH"
	HOTBTC = "HOTBTC"
	HOTETH = "HOTETH"
	DTABTC = "DTABTC"
	DTAETH = "DTAETH"
	WPRBTC = "WPRBTC"
	WPRETH = "WPRETH"
	ZILBTC = "ZILBTC"
	ZILETH = "ZILETH"
	BNTBTC = "BNTBTC"
	BNTETH = "BNTETH"
	MANETH = "MANETH"

	// Channels
	ChanBook   = "book"
	ChanTrade  = "trades"
	ChanTicker = "ticker"
)

var successPair = map[string]string{
	BTCUSD: "BTCUSD",
	LTCUSD: "LTCUSD",
	LTCBTC: "LTCBTC",
	ETHUSD: "ETHUSD",
	ETHBTC: "ETHBTC",
	ETCUSD: "ETCUSD",
	ETCBTC: "ETCBTC",
	BFXUSD: "BFXUSD",
	BFXBTC: "BFXBTC",
	ZECUSD: "ZECUSD",
	ZECBTC: "ZECBTC",
	XMRUSD: "XMRUSD",
	XMRBTC: "XMRBTC",
	RRTUSD: "RRTUSD",
	RRTBTC: "RRTBTC",
	XRPUSD: "XRPUSD",
	XRPBTC: "XRPBTC",
	EOSETH: "EOSETH",
	EOSUSD: "EOSUSD",
	EOSBTC: "EOSBTC",
	IOTUSD: "IOTUSD",
	IOTBTC: "IOTBTC",
	IOTETH: "IOTETH",
	BCCBTC: "BCCBTC",
	BCUBTC: "BCUBTC",
	BCCUSD: "BCCUSD",
	BCUUSD: "BCUUSD",
	SANBTC: "SANBTC",
	SANETH: "SANETH",
	OMGBTC: "OMGBTC",
	OMGETH: "OMGETH",
	AVTBTC: "AVTBTC",
	AVTETH: "AVTETH",
	EDOBTC: "EDOBTC",
	EDOETH: "EDOETH",
	YYWBTC: "YYWBTC",
	YYWETH: "YYWETH",
	GNTBTC: "GNTBTC",
	GNTETH: "GNTETH",
	SNTBTC: "SNTBTC",
	SNTETH: "SNTETH",
	BATBTC: "BATBTC",
	BATETH: "BATETH",
	FUNBTC: "FUNBTC",
	FUNETH: "FUNETH",
	ZRXBTC: "ZRXBTC",
	ZRXETH: "ZRXETH",
	TNBBTC: "TNBBTC",
	TNBETH: "TNBETH",
	RCNBTC: "RCNBTC",
	RCNETH: "RCNETH",
	AIDBTC: "AIDBTC",
	AIDETH: "AIDETH",
	REPBTC: "REPBTC",
	REPETH: "REPETH",
	ELFBTC: "ELFBTC",
	ELFETH: "ELFETH",
	IOSBTC: "IOSBTC",
	IOSETH: "IOSETH",
	AIOBTC: "AIOBTC",
	AIOETH: "AIOETH",
	REQBTC: "REQBTC",
	REQETH: "REQETH",
	RDNBTC: "RDNBTC",
	RDNETH: "RDNETH",
	LRCBTC: "LRCBTC",
	LRCETH: "LRCETH",
	WAXBTC: "WAXBTC",
	WAXETH: "WAXETH",
	DAIBTC: "DAIBTC",
	DAIETH: "DAIETH",
	CFIBTC: "CFIBTC",
	CFIETH: "CFIETH",
	AGIBTC: "AGIBTC",
	AGIETH: "AGIETH",
	BFTBTC: "BFTBTC",
	BFTETH: "BFTETH",
	MTNBTC: "MTNBTC",
	MTNETH: "MTNETH",
	ODEBTC: "ODEBTC",
	ODEETH: "ODEETH",
	ANTBTC: "ANTBTC",
	ANTETH: "ANTETH",
	DTHBTC: "DTHBTC",
	DTHETH: "DTHETH",
	MITBTC: "MITBTC",
	MITETH: "MITETH",
	MKRBTC: "MKRBTC",
	MKRETH: "MKRETH",
	KNCBTC: "KNCBTC",
	KNCETH: "KNCETH",
	LYMBTC: "LYMBTC",
	LYMETH: "LYMETH",
	UTKBTC: "UTKBTC",
	UTKETH: "UTKETH",
	VEEBTC: "VEEBTC",
	VEEETH: "VEEETH",
	DADBTC: "DADBTC",
	DADETH: "DADETH",
	ORSBTC: "ORSBTC",
	ORSETH: "ORSETH",
	AUCBTC: "AUCBTC",
	AUCETH: "AUCETH",
	POYBTC: "POYBTC",
	POYETH: "POYETH",
	FSNBTC: "FSNBTC",
	FSNETH: "FSNETH",
	CBTBTC: "CBTBTC",
	CBTETH: "CBTETH",
	ZCNBTC: "ZCNBTC",
	ZCNETH: "ZCNETH",
	SENBTC: "SENBTC",
	SENETH: "SENETH",
	NCABTC: "NCABTC",
	NCAETH: "NCAETH",
	CNDBTC: "CNDBTC",
	CNDETH: "CNDETH",
	CTXBTC: "CTXBTC",
	CTXETH: "CTXETH",
	SEEBTC: "SEEBTC",
	SEEETH: "SEEETH",
	ESSBTC: "ESSBTC",
	ESSETH: "ESSETH",
	ATMBTC: "ATMBTC",
	ATMETH: "ATMETH",
	HOTBTC: "HOTBTC",
	HOTETH: "HOTETH",
	DTABTC: "DTABTC",
	DTAETH: "DTAETH",
	WPRBTC: "WPRBTC",
	WPRETH: "WPRETH",
	ZILBTC: "ZILBTC",
	ZILETH: "ZILETH",
	BNTBTC: "BNTBTC",
	BNTETH: "BNTETH",
	MANETH: "MANETH",
}

// PairAction used to determine the purchase or sale transactions
// If Action == true - buy (+)
// If Action == false - sell (-)
type PairAction struct {
	Pair   string
	Action bool
}

// GetSuccessPair return success pait and arithmetic sign
// return value bool = true this is bay
// return value bool = false this is sell
func GetSuccessPair(pair string) (*PairAction, error) {

	pairAction := &PairAction{}

	targetPair, ok := successPair[pair]

	if !ok {
		needlePair :=  pair[3:] + pair[0:3]

		_, ok = successPair[needlePair]

		if !ok {
			return nil, fmt.Errorf("%s unknown pair", needlePair)
		}

		pairAction.Pair = needlePair
		pairAction.Action = false // sell

		return pairAction, nil
	}

	pairAction.Pair = targetPair
	pairAction.Action = true // bye

	return pairAction, nil
}

// WebSocketService allow to connect and receive stream data
// from bitfinex.com ws service.
type WebSocketService struct {
	// http client
	client *Client
	// websocket client
	ws *websocket.Conn
	// special web socket for private messages
	privateWs *websocket.Conn
	// map internal channels to websocket's
	chanMap    map[float64]chan []float64
	subscribes []subscribeToChannel
}

type subscribeMsg struct {
	Event   string  `json:"event"`
	Channel string  `json:"channel"`
	Pair    string  `json:"pair"`
	ChanID  float64 `json:"chanId,omitempty"`
}

type subscribeToChannel struct {
	Channel string
	Pair    string
	Chan    chan []float64
}

// NewWebSocketService returns a WebSocketService using the given client.
func NewWebSocketService(c *Client) *WebSocketService {
	return &WebSocketService{
		client:     c,
		chanMap:    make(map[float64]chan []float64),
		subscribes: make([]subscribeToChannel, 0),
	}
}

// Connect create new bitfinex websocket connection
func (w *WebSocketService) Connect() error {
	var d = websocket.Dialer{
		Subprotocols:    []string{"p1", "p2"},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		Proxy:           http.ProxyFromEnvironment,
	}

	if w.client.WebSocketTLSSkipVerify {
		d.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	ws, _, err := d.Dial(w.client.WebSocketURL, nil)
	if err != nil {
		return err
	}
	w.ws = ws
	return nil
}

// Close web socket connection
func (w *WebSocketService) Close() {
	w.ws.Close()
}

func (w *WebSocketService) AddSubscribe(channel string, pair string, c chan []float64) {
	s := subscribeToChannel{
		Channel: channel,
		Pair:    pair,
		Chan:    c,
	}
	w.subscribes = append(w.subscribes, s)
}

func (w *WebSocketService) ClearSubscriptions() {
	w.subscribes = make([]subscribeToChannel, 0)
}

func (w *WebSocketService) sendSubscribeMessages() error {
	for _, s := range w.subscribes {
		msg, _ := json.Marshal(subscribeMsg{
			Event:   "subscribe",
			Channel: s.Channel,
			Pair:    s.Pair,
		})

		err := w.ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

// Subscribe allows to subsribe to channels and watch for new updates.
// This method supports next channels: book, trade, ticker.
func (w *WebSocketService) Subscribe() error {
	// Subscribe to each channel
	if err := w.sendSubscribeMessages(); err != nil {
		return err
	}

	for {
		_, p, err := w.ws.ReadMessage()
		if err != nil {
			return err
		}

		if bytes.Contains(p, []byte("event")) {
			w.handleEventMessage(p)
		} else {
			w.handleDataMessage(p)
		}
	}

	return nil
}

func (w *WebSocketService) handleEventMessage(msg []byte) {
	// Check for first message(event:subscribed)
	event := &subscribeMsg{}
	err := json.Unmarshal(msg, event)

	// Received "subscribed" resposne. Link channels.
	if err == nil {
		for _, k := range w.subscribes {
			if event.Event == "subscribed" && event.Pair == k.Pair && event.Channel == k.Channel {
				w.chanMap[event.ChanID] = k.Chan
			}
		}
	}
}

func (w *WebSocketService) handleDataMessage(msg []byte) {
	// Received payload or data update
	var dataUpdate []float64
	err := json.Unmarshal(msg, &dataUpdate)
	if err == nil {
		chanID := dataUpdate[0]
		// Remove chanID from data update
		// and send message to internal chan
		w.chanMap[chanID] <- dataUpdate[1:]
	}

	// Payload received
	var fullPayload []interface{}
	err = json.Unmarshal(msg, &fullPayload)

	if err != nil {
		log.Println("Error decoding fullPayload", err)
	} else {
		if len(fullPayload) > 3 {
			itemsSlice := fullPayload[3:]
			i, _ := json.Marshal(itemsSlice)
			var item []float64
			err = json.Unmarshal(i, &item)
			if err == nil {
				chanID := fullPayload[0].(float64)
				w.chanMap[chanID] <- item
			}
		} else {
			itemsSlice := fullPayload[1]
			i, _ := json.Marshal(itemsSlice)
			var items [][]float64
			err = json.Unmarshal(i, &items)
			if err == nil {
				chanID := fullPayload[0].(float64)
				for _, v := range items {
					w.chanMap[chanID] <- v
				}
			}
		}
	}
}

/////////////////////////////
// Private websocket messages
/////////////////////////////

type privateConnect struct {
	Event       string `json:"event"`
	APIKey      string `json:"apiKey"`
	AuthSig     string `json:"authSig"`
	AuthPayload string `json:"authPayload"`
}

// Private channel auth response
type privateResponse struct {
	Event  string  `json:"event"`
	Status string  `json:"status"`
	ChanID float64 `json:"chanId,omitempty"`
	UserID float64 `json:"userId"`
}

type TermData struct {
	// Data term. E.g: ps, ws, ou, etc... See official documentation for more details.
	Term string
	// Data will contain different number of elements for each term.
	// Examples:
	// Term: ws, Data: ["exchange","BTC",0.01410829,0]
	// Term: oc, Data: [0,"BTCUSD",0,-0.01,"","CANCELED",270,0,"2015-10-15T11:26:13Z",0]
	Data  []interface{}
	Error string
}

func (c *TermData) HasError() bool {
	return len(c.Error) > 0
}

func (w *WebSocketService) ConnectPrivate(ch chan TermData) {

	var d = websocket.Dialer{
		Subprotocols:    []string{"p1", "p2"},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		Proxy:           http.ProxyFromEnvironment,
	}

	if w.client.WebSocketTLSSkipVerify {
		d.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	ws, _, err := d.Dial(w.client.WebSocketURL, nil)
	if err != nil {
		ch <- TermData{
			Error: err.Error(),
		}
		return
	}

	nonce := utils.GetNonce()
	payload := "AUTH" + nonce
	connectMsg, _ := json.Marshal(&privateConnect{
		Event:       "auth",
		APIKey:      w.client.APIKey,
		AuthSig:     w.client.signPayload(payload),
		AuthPayload: payload,
	})

	// Send auth message
	err = ws.WriteMessage(websocket.TextMessage, connectMsg)
	if err != nil {
		ch <- TermData{
			Error: err.Error(),
		}
		ws.Close()
		return
	}

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			ch <- TermData{
				Error: err.Error(),
			}
			ws.Close()
			return
		}

		event := &privateResponse{}
		err = json.Unmarshal(p, &event)
		if err != nil {
			// received data update
			var data []interface{}
			err = json.Unmarshal(p, &data)
			if err == nil {
				if len(data) == 2 { // Heartbeat
					// XXX: Consider adding a switch to enable/disable passing these along.
					ch <- TermData{Term: data[1].(string)}
					return
				}

				dataTerm := data[1].(string)
				dataList := data[2].([]interface{})

				// check for empty data
				if len(dataList) > 0 {
					if reflect.TypeOf(dataList[0]) == reflect.TypeOf([]interface{}{}) {
						// received list of lists
						for _, v := range dataList {
							ch <- TermData{
								Term: dataTerm,
								Data: v.([]interface{}),
							}
						}
					} else {
						// received flat list
						ch <- TermData{
							Term: dataTerm,
							Data: dataList,
						}
					}
				}
			}
		} else {
			// received auth response
			if event.Event == "auth" && event.Status != "OK" {
				ch <- TermData{
					Error: "Error connecting to private web socket channel.",
				}
				ws.Close()
			}
		}
	}
}
