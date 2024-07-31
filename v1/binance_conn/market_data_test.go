package binance_conn_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/lianyun0502/exchange_conn/common"
	"github.com/lianyun0502/exchange_conn/binance_conn"
)

type OrderBookResponse struct {
	LastUpdateId int64 `json:"lastUpdateId"`
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}
func TestOrderBook(t *testing.T) {
	assert := assert.New(t)
	client := binance_conn.NewClient(
		apiKey,
		secretKey,
		testURL,
	)
	// Test OrderBook
	ob := new(binance_conn.OrderBook)

	limit := 5

	data, err := ob.Symbol("BTCUSDT").Limit(limit).Do(client)
	if err != nil {
		t.Error(err)
	}
	var j OrderBookResponse
	err = json.Unmarshal(data, &j)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(common.PrettyPrint(j))
	assert.Equal(limit, len(j.Bids))
	assert.Equal(limit, len(j.Asks))
}