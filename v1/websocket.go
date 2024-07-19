package binance_connect

import (
	// "fmt"
	"log"
	"time"

	"github.com/lxzan/gws"
)

type ErrHandler func(err error)
type WsHandler func(message []byte)

type WebSocket struct{
	Err_Handler ErrHandler
	Ws_Handler WsHandler
}

func (conn *WebSocket) OnOpen(socket *gws.Conn) {
	log.Println("OnOpen")
}
func (conn *WebSocket) OnPing(socket *gws.Conn, message []byte) {
	log.Println("OnPing")
	socket.WritePong(message)
}
func (conn *WebSocket) OnPong(socket *gws.Conn, message []byte) {
	log.Println("OnPong")
}
func (conn *WebSocket) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	log.Println("OnMessage")
	// fmt.Printf("recv: %s\n", message.Data.String())
	conn.Ws_Handler(message.Data.Bytes())
}
func (conn *WebSocket) OnClose(socket *gws.Conn, err error) {
	log.Println("OnClose")
	if err != nil {
		conn.Err_Handler(err)
	}
}


func WsClient(url string, wsHandler WsHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error){
	conn, _, err := gws.NewClient(
		&WebSocket{
			Err_Handler: errHandler,
			Ws_Handler: wsHandler,
		}, 
		&gws.ClientOption{
			ReadBufferSize: 655350,
			Addr : url,
			HandshakeTimeout: 45*time.Second,
			PermessageDeflate: gws.PermessageDeflate{
				Enabled: true,
				ServerContextTakeover: true,
				ClientContextTakeover: true,
			},
	})

	if err != nil {
		return nil, nil, err
	}

	go conn.ReadLoop()
	doneCh = make(chan struct{})
	stopCh = make(chan struct{})
	go func() {
		select {
		case <-stopCh:
		case <-doneCh:
		}
	}()
	return
}
