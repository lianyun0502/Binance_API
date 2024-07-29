package binance_connect_test

import (
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
	"practice_go/binance_conn"
)

var errHandler = func(err error) {
	if err != nil {
		log.Println(err)
	}
}

// type WsApiPingResponse struct {
// 	ID string `json:"id"`
// 	Method string `json:"method"`
// }


func TestWsApiPing(t *testing.T) {
	assert := assert.New(t)
	client, err := binance_connect.NewWebSocketAPI(
		apiKey,
		secretKey,
		"wss://testnet.binance.vision/ws-api/v3",
		errHandler,
	)
	if err != nil {
		t.Error(err)
		return
	}
	client.StartLoop()
	for i := 0; i < 20; i++ {
		resp, err := client.PingServer()
		if err != nil {
			t.Error(err)
			return
		}
		data :=  binance_connect.PrettyPrint(resp)
		log.Printf("response:\n%s", data)
		assert.NotEqual("{}", data)
	}
	client.StopLoop()
}