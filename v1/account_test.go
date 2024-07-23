package binance_connect_test

import (
	"encoding/json"
	"log"
	"practice_go/binance_conn"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestAccountInfo(t *testing.T) {
	assert := assert.New(t)
	client := binance_connect.NewClient(
		apiKey,
		secretKey,
		testURL,
	)
	// Test AccountInfo
	ai := new(binance_connect.AccountInfoService)

	data, err := ai.Do(client)
	if err != nil {
		t.Error(err)
		return

	}
	var j interface{}
	err = json.Unmarshal(data, &j)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(binance_connect.PrettyPrint(j))
	assert.NotEqual("{}", string(data))
}

func TestAccountTradeList(t *testing.T) {
	assert := assert.New(t)
	client := binance_connect.NewClient(
		apiKey,
		secretKey,
		testURL,
	)
	// Test AccountTradeList
	atl := new(binance_connect.AccountTradeListService)

	data, err := atl.Symbol("BTCUSDT").Do(client)
	if err != nil {
		t.Error(err)
		return

	}
	var j interface{}
	err = json.Unmarshal(data, &j)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(binance_connect.PrettyPrint(j))
	assert.NotEqual("{}", string(data))
}
