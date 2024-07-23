package binance_connect_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"practice_go/binance_conn"
)

type OrderBookResponse struct {
	LastUpdateId int64 `json:"lastUpdateId"`
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}
func TestOrderBook(t *testing.T) {
	assert := assert.New(t)
	client := binance_connect.NewClient(
		apiKey,
		secretKey,
		testURL,
	)
	// Test OrderBook
	ob := new(binance_connect.OrderBook)

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
	log.Println(binance_connect.PrettyPrint(j))
	assert.Equal(limit, len(j.Bids))
	assert.Equal(limit, len(j.Asks))
}