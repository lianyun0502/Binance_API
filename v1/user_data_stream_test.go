package binance_connect_test

import (
	"encoding/json"
	"log"
	"testing"
	"practice_go/binance_conn"
	"github.com/stretchr/testify/assert"
)

func TestCreateListenKey(t *testing.T) {
	assert := assert.New(t)
	client := binance_connect.NewClient(
		apiKey,
		secretKey,
		testURL,
	)

	cak := new(binance_connect.CreateAPIKeyService)
	
	data, err := cak.Do(client)
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