package binance_conn_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/lianyun0502/exchange_conn/binance_conn"
	"github.com/lianyun0502/exchange_conn/common"
	"github.com/lianyun0502/exchange_conn/binance_conn/enums"
	
	"github.com/stretchr/testify/assert"
)


func TestNewOrder(t *testing.T) {
	assert := assert.New(t)
	client := binance_conn.NewClient(
		apiKey,
		secretKey,
		testURL,
	)
	// Test NewOrder
	no := new(binance_conn.NewOrder)
	
	data, err := no.Symbol("BTCUSDT").Side(enums.Buy).Type(enums.Limit).TimeInForce(enums.FOK).Quantity(0.1).Price(10000).Do(client)
	if err != nil {
		t.Error(err)
	}
	var j interface{}
	err = json.Unmarshal(data, &j)
	log.Println(common.PrettyPrint(j))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual("{}", string(data))
}