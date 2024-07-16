package v1

import (
	"fmt"
	"time"
	"net/http"

	"github.com/gorilla/websocket"
)

type WsHandler func(message []byte)
type ErrHandler func(err error)

var WsServer = func(url string, handler WsHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	Dialer := websocket.Dialer{
		Proxy:   http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		EnableCompression: false,
	}
	headers := http.Header{}
	headers.Add("User-Agent", fmt.Sprintf("1231$%s", "zsdfs"))
	conn, _, err := Dialer.Dial(url, headers)
	if err != nil{
		return nil, nil, err
	}

	conn.SetReadLimit(655350)
	doneCh = make(chan struct{})
	stopCh = make(chan struct{})

	go func () {
		silent := false
		go func() {
			select {
			case <-stopCh:
				silent = true
			case <-doneCh:
			}
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
			return
			}
			handler(message)
		}

	}()
	return 
}