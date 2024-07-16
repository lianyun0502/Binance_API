package main

import (
	"fmt"
	"time"
	"practice_go/binance_conn"
)

func wsHandler(message []byte) {
	fmt.Println(string(message))
}
func errorHandler(err error) {
	fmt.Println(err)
}
func main() {
	// url := "wss://stream.binance.com:9443/ws/btcusdt@trade"
	url_combined := "wss://stream.binance.com:9443/stream?streams=btcusdt@trade/btcusdt@aggTrade"
	doneCh, stopCh, err := v1.WsServer(url_combined, wsHandler, errorHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		time.Sleep(30*time.Second)
		stopCh <- struct{}{}
	}()

	<- doneCh
}