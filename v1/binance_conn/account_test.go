package binance_conn_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/lianyun0502/exchange_conn/common"
	"github.com/lianyun0502/exchange_conn/binance_conn"
	
	
)


func TestAccountInfo(t *testing.T) {
	assert := assert.New(t)
	client := binance_conn.NewClient(
		apiKey,
		secretKey,
		testURL,
	)
	// Test AccountInfo
	ai := new(binance_conn.AccountInfoService)

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
	log.Println(common.PrettyPrint(j))
	assert.NotEqual("{}", string(data))
}

func TestAccountTradeList(t *testing.T) {
	assert := assert.New(t)
	client := binance_conn.NewClient(
		apiKey,
		secretKey,
		testURL,
	)
	// Test AccountTradeList
	atl := new(binance_conn.AccountTradeListService)

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
	log.Println(common.PrettyPrint(j))
	assert.NotEqual("{}", string(data))
}
