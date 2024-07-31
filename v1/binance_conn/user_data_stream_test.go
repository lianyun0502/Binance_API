package binance_conn_test

import (
	"encoding/json"
	"log"
	"testing"
	"github.com/lianyun0502/exchange_conn/binance_conn"
	"github.com/lianyun0502/exchange_conn/common"
	"github.com/stretchr/testify/assert"
)

func TestCreateListenKey(t *testing.T) {
	assert := assert.New(t)
	client := binance_conn.NewClient(
		apiKey,
		secretKey,
		testURL,
	)

	cak := new(binance_conn.CreateAPIKeyService)
	
	data, err := cak.Do(client)
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