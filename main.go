package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"practice_go/binance_conn"

)

func wsHandler(message []byte) {
	fmt.Println(string(message))
}
func errorHandler(err error) {
	fmt.Println(err)
}

func testAPI(){
	client := binance_connect.NewClient(
		"PkuSskgddMZuLdPV35rBdN1kqdtffKERnG0plDdTDqLIdqH1tfD4xavK1AjxNo1F", 
		"1vs3rVyqEG64wDuWAQqlxY3ZclMAB3WjvVUTSdguz7IrWO7hyC9xeIjL8PtFSxPz", 
		"https://api1.binance.com",
	)

	req := binance_connect.Request{
		Method: http.MethodGet,
		Endpoint: "/api/v3/depth",
		SercType: binance_connect.None,
	}
	req.Query = url.Values{}
	req.Query.Add("symbol", "BTCUSDT")
	req.Query.Add("limit", "100")

	b_req, err := client.SetBinanceRequest(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := client.Call(b_req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("data: %s", data)

}
func main() {
	// url := "wss://stream.binance.com:9443/ws/btcusdt@depth@100ms"
	// url := "wss://ws-api.binance.com:443/ws-api/v3"
	// url := "wss://stream.binance.com:9443/stream?streams=btcusdt@trade/btcusdt@aggTrade"

	// doneCh, stopCh, err := binance_connect.WsClient(url, wsHandler, errorHandler)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// go func() {
	// 	time.Sleep(30*time.Second)
	// 	stopCh <- struct{}{}
	// }()

	// <- doneCh
	testAPI()	
}