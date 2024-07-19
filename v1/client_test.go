package binance_connect

import (
	"encoding/json"
	"net/http"
	"testing"
)

var (
	apiKey = "PkuSskgddMZuLdPV35rBdN1kqdtffKERnG0plDdTDqLIdqH1tfD4xavK1AjxNo1F"
	secretKey = "1vs3rVyqEG64wDuWAQqlxY3ZclMAB3WjvVUTSdguz7IrWO7hyC9xeIjL8PtFSxPz"
	testURL = "https://api1.binance.com"
)


func TestPingAPIServer(t *testing.T) {
	client := NewClient(
		apiKey, 
		secretKey, 
		testURL,
	)

	req := Request{
		Method: http.MethodGet,
		Endpoint: "/api/v3/ping",
		SercType: None,
	}

	b_req, err := client.SetBinanceRequest(&req)
	if err != nil {
		t.Error(err)
		return
	}

	data, err := client.Call(b_req)
	if err != nil {
		t.Error(err)
		return
	}
	if string(data) != "{}" {
		t.Errorf("data: %s", data)
	}
}

type CheckServerTimeResponce struct {
	ServerTime int64 `json:"serverTime"`
}

func TestCheckServerTime(t *testing.T) {
	client := NewClient(
		apiKey, 
		secretKey, 
		testURL,
	)

	req := Request{
		Method: http.MethodGet,
		Endpoint: "/api/v3/time",
		SercType: None,
	}

	b_req, err := client.SetBinanceRequest(&req)
	if err != nil {
		t.Error(err)
		return
	}

	data, err := client.Call(b_req)
	if err != nil {
		t.Error(err)
		return
	}

	j := new(CheckServerTimeResponce)
	err = json.Unmarshal(data, j)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("\ndata: %#v", j)
}

type ExchangeInfo struct {
	TimeZone string `json:"timezone"`
	ServerTime int64 `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval string `json:"interval"`
		Limit int `json:"limit"`
		IntervalNum int `json:"intervalNum"`
	} `json:"rateLimits"`
	ExchangeFilters []struct {
		FilterType string `json:"filterType"`
		MinPrice string `json:"minPrice"`
		MaxPrice string `json:"maxPrice"`
		TickSize string `json:"tickSize"`
	} `json:"exchangeFilters"`
	Symbols []struct {
		Symbol string `json:"symbol"`
		Status string `json:"status"`
		BaseAsset string `json:"baseAsset"`
		BaseAssetPrecision int `json:"baseAssetPrecision"`
		QuoteAsset string `json:"quoteAsset"`
		QuotePrecision int `json:"quotePrecision"`
		QuoteAssetPrecision int `json:"quoteAssetPrecision"`
		BaseCommissionPrecision int `json:"baseCommissionPrecision"`

		OrderTypes []string `json:"orderTypes"`
		IceBergAllowed bool `json:"iceBergAllowed"`
		OcoAllowed bool `json:"ocoAllowed"`
		QuoteOrderQtyMarketAllowed bool `json:"quoteOrderQtyMarketAllowed"`
		AllowTradingStop bool `json:"allowTradingStop"`
		CancelReplaceService bool `json:"cancelReplaceService"`
		IsStopTradingAllowed bool `json:"isSpotTradingAllowed"`
		IsMarginTradingAllowed bool `json:"isMarginTradingAllowed"`
		Filters []struct {
			FilterType string `json:"filterType"`
			MinPrice string `json:"minPrice"`
			MaxPrice string `json:"maxPrice"`
			TickSize string `json:"tickSize"`
		} `json:"filters"`
		Permissions []string `json:"permissions"`
		PermissionSets [][]string `json:"permissionSets"`
		DefaultSelfTradePreventionMode string `json:"defaultSelfTradePreventionMode"`
		AllowedSelfTradePreventionModes []string `json:"allowedSelfTradePreventionModes"`
	} `json:"symbols"`
	Sors []struct {
		BaseAsset string `json:"baseAsset"`
		Symbols []string `json:"symbols"`
	} `json:"sors"`
}

func TestGetExchangeInfo(t *testing.T) {
	client := NewClient(
		apiKey, 
		secretKey, 
		testURL,
	)

	req := NewBinanceRequest(
		http.MethodGet,
		"/api/v3/exchangeInfo",
		None,
	)

	req.Query.Set("symbol", "BTCUSDT")

	b_req, err := client.SetBinanceRequest(req)
	if err != nil {
		t.Error(err)
		return
	}

	data, err := client.Call(b_req)
	if err != nil {
		t.Error(err)
		return
	}

	j := new(ExchangeInfo)
	err = json.Unmarshal(data, j)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %#v", j)
}