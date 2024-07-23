package binance_connect_test

import (
	"encoding/json"
	"log"
	"practice_go/binance_conn"
	enums "practice_go/binance_conn/Enums"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestNewOrder(t *testing.T) {
	assert := assert.New(t)
	client := binance_connect.NewClient(
		apiKey,
		secretKey,
		testURL,
	)
	// Test NewOrder
	no := new(binance_connect.NewOrder)
	
	data, err := no.Symbol("BTCUSDT").Side(enums.Buy).Type(enums.Limit).TimeInForce(enums.FOK).Quantity(0.1).Price(10000).Do(client)
	if err != nil {
		t.Error(err)
	}
	var j interface{}
	err = json.Unmarshal(data, &j)
	log.Println(binance_connect.PrettyPrint(j))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual("{}", string(data))
}