package exchange_conn_test

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lianyun0502/exchange_conn/binance_conn"
	"github.com/lianyun0502/exchange_conn"
	"github.com/lianyun0502/exchange_conn/common"
)

func TestExchangePing(t *testing.T) {
	client := exchange_conn.NewClient(
		binance_conn.NewClient(
			"",
			"",
			"https://api.binance.com",
		),
		
	)

	data, err := client.Request(http.MethodGet, "/api/v3/ping", false, false).Send()
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, string(data), "{}")
}

func TestExchangeGetInfo(t *testing.T) {
	client := exchange_conn.NewClient(
		binance_conn.NewClient(
			"",
			"",
			"https://api.binance.com",
		),
	)

	data, err := client.Request(http.MethodGet, "/api/v3/exchangeInfo", false, false).Send()
	if err != nil {
		t.Error(err)
		return
	}

	assert.NotEqual(t, string(data), "{}")
	j := new(interface{})
	json.Unmarshal(data, &j)
	log.Println(common.PrettyPrint(j))
}
