package binance_connect_test

import (
	"encoding/json"
	"testing"
	// "github.com/stretchr/testify/assert"
	"fmt"
	"log"
	"practice_go/binance_conn"
	"time"
)


func wsHandler(message []byte) {
	log.Println(string(message))
	j := make(map[string]interface{})
	json.Unmarshal(message, &j)
	log.Printf("%v", j["E"])
	log.Printf("%v", time.Now().UnixNano()/int64(time.Millisecond))
}
func errorHandler(err error) {
	log.Println(err)
}


func TestWsClient(t *testing.T) {
	url := "wss://stream.binance.com:9443/ws/btcusdt@depth@100ms"
	// url := "wss://stream.binance.com:9443/ws/btcusdt@aggTrade"
	// url := "wss://ws-api.binance.com:443/ws-api/v3"
	// url := "wss://stream.binance.com:9443/stream?streams=btcusdt@trade/btcusdt@aggTrade"

	client, err := binance_connect.NewWsClient(
		url, 
		&binance_connect.WebSocketEvent{
			Err_Handler: errorHandler,
			Ws_Handler: wsHandler,
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.StartLoop()

	go func() {
		time.Sleep(60*time.Second)
		client.Close()
	}()
	
	for {
		select{
			case <- client.DoneSignal:
				fmt.Printf("end")
				return
			case <- client.ReconnectSignal:
				for i :=0; i<10; i++{
					log.Printf("retry connect {%d}", i)
					err = client.Reconnect()
					if err != nil{
						break
					}
					log.Println(err)
				}
				
		}
		
	}
	
}