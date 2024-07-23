package binance_connect_test

import (
	"log"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"practice_go/binance_conn"
)

var errHandler = func(err error) {
	if err != nil {
		log.Println(err)
	}
}

type WsApiPingResponse struct {
	ID string `json:"id"`
	Method string `json:"method"`
}


func TestWsApiPing(t *testing.T) {
	assert := assert.New(t)
	client, _:= binance_connect.NewWebSocketAPI(
		apiKey,
		secretKey,
		"wss://testnet.binance.vision/ws-api/v3",
		errHandler,
	)
	client.StartLoop()

	id := binance_connect.GetUUID()
	res := &WsApiPingResponse{ID: id, Method: "ping"}
	respCh := make(chan []byte)

	data, _ := json.Marshal(res)

	client.SendMessage(id, respCh, data)

	resp := <- respCh

	var j interface{}
	err := json.Unmarshal(resp, &j)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(binance_connect.PrettyPrint(j))
	assert.NotEqual("{}", string(resp))
}