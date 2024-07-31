package binance_conn

import (
	"net/http"
	"errors"
)

type AccountInfoService struct {
	omitZeroBalances	bool
}
func (api *AccountInfoService) Do(c *Client) (data []byte, err error){
	res := NewBinanceRequest(
		http.MethodGet,
		"/api/v3/account",
		UserData,
	)
	b_res, err := c.SetBinanceRequest(res)
	if err != nil {
		return
	}
	data, err = c.Call(b_res)
	if err != nil {
		return
	}
	return 
}

func (api *AccountInfoService) OmitZeroBalances(b bool) *AccountInfoService {
	api.omitZeroBalances = b
	return api
}


type AccountTradeListService struct {
	symbol string
	orderID int64
	startTime int64
	endTime int64
	fromID int64
	limit int
}

func (api *AccountTradeListService) Do(c *Client) (data []byte, err error){
	res := NewBinanceRequest(
		http.MethodGet,
		"/api/v3/myTrades",
		UserData,
	)
	if api.symbol == "" {
		return nil, errors.New("symbol is required")
	}
	res.SetQuery("symbol", api.symbol)
	if api.orderID != 0 {
		res.SetQuery("orderId", api.orderID)
	}
	if api.startTime != 0 {
		res.SetQuery("startTime", api.startTime)
	}
	if api.endTime != 0 {
		res.SetQuery("endTime", api.endTime)
	}
	if api.fromID != 0 {
		res.SetQuery("fromId", api.fromID)
	}
	if api.limit != 0 {
		res.SetQuery("limit", api.limit)
	}
	b_res, err := c.SetBinanceRequest(res)
	if err != nil {
		return
	}
	data, err = c.Call(b_res)
	if err != nil {
		return
	}
	return 
}

func (api *AccountTradeListService) Symbol(s string) *AccountTradeListService {
	api.symbol = s
	return api
}