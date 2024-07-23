package binance_connect

import (
	"net/http"
)

type OrderBook struct {
	symbol string
	limit int
}

func (api *OrderBook) Symbol(s string) *OrderBook {
	api.symbol = s
	return api
}

func (api *OrderBook) Limit(l int) *OrderBook {
	api.limit = l
	return api
}

func (api *OrderBook) Do(c *Client) (data []byte, err error){
	req := NewBinanceRequest(
		http.MethodGet,
		"/api/v3/depth",
		None,
	)
	req.SetQuery("symbol", api.symbol)
	if api.limit != 0 {
		req.SetQuery("limit", api.limit)
	}

	b_req, err := c.SetBinanceRequest(req)
	if err != nil {
		return
	}

	data, err = c.Call(b_req)
	if err != nil {
		return
	}

	return
}


