package binance_connect

import (
	// "fmt"
	"log"
	"time"

	"github.com/lxzan/gws"
)

type ErrHandler func(err error)
type WsHandler func(message []byte)

type WebSocketEvent struct{
	Err_Handler func(err error)
	Ws_Handler func(message []byte)

	client *WsClient

	pingTimer Timer
}
func (conn *WebSocketEvent) OnOpen(socket *gws.Conn) {
	log.Println("OnOpen")

	conn.pingTimer = Timer{
		Interval: 10*time.Second,
		handle: func() { 
			log.Println("Ping server timeout") 
			socket.NetConn().Close()
		},
	}
	conn.pingTimer.Start(nil)
	socket.WritePing([]byte("ping"))
}
func (conn *WebSocketEvent) OnPing(socket *gws.Conn, message []byte) {
	log.Println("OnPing")
	socket.WritePong(message)
}
func (conn *WebSocketEvent) OnPong(socket *gws.Conn, message []byte) {
	log.Println("OnPong")
	conn.pingTimer.Stop()
	go func () {
		time.Sleep(5*time.Second)
		socket.WritePing([]byte("ping"))
		conn.pingTimer.Start(nil)
	}()
	
}
func (conn *WebSocketEvent) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	log.Println("OnMessage")
	if conn.Ws_Handler == nil { return }
	conn.Ws_Handler(message.Data.Bytes())
}
func (conn *WebSocketEvent) OnClose(socket *gws.Conn, err error) {
	log.Println("OnClose")
	conn.pingTimer.Stop()
	if conn.Err_Handler == nil { return }
	if err != nil {
		conn.Err_Handler(err)
	}
	conn.client.ReconnectSignal <- struct{}{}
}

type WsClient struct {
	clientOption *gws.ClientOption
	wsEvent gws.Event
	conn *gws.Conn
	DoneSignal chan struct{}
 	stopSignal chan struct{}
	ReconnectSignal chan struct{}
}

func (client *WsClient) Reconnect() (err error) {
	client.conn, _, err = gws.NewClient(client.wsEvent, client.clientOption)
	if err == nil{
		client.StartLoop()
	}
	return
}

func (client *WsClient) StartLoop() {
	go func() {
		client.conn.ReadLoop()
	}()
}
func (client *WsClient) Close() {
	client.stopSignal <- struct{}{}
}

func NewWsClient(url string, wsEvent *WebSocketEvent) (client *WsClient, err error){
	clientOption := gws.ClientOption{
		ReadBufferSize: 655350,
		Addr : url,
		HandshakeTimeout: 45*time.Second,
		PermessageDeflate: gws.PermessageDeflate{
			Enabled: true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
	}
	
	conn, _, err := gws.NewClient(
		wsEvent,
		&clientOption,
	)

	if err != nil {
		return nil, err
	}

	client = &WsClient{
		clientOption: &clientOption,
		wsEvent: wsEvent,
		conn: conn,
		DoneSignal: make(chan struct{}),
		stopSignal: make(chan struct{}),
		ReconnectSignal: make(chan struct{}),
	}
	wsEvent.client = client

	go func() {
		<-client.stopSignal
		client.conn.NetConn().Close()
		client.DoneSignal <- struct{}{}
	}()

	return
}
