package exchange_conn

import (
	"net/http"
	"github.com/lianyun0502/exchange_conn/binance_conn"
)
type IExchange interface {
	
	Request(string, string, bool, bool, ...any) (*binance_conn.Request)
	SetRequest(*binance_conn.Request) (*http.Request, error)
	Call(*http.Request) ([]byte, error)
}

type IRequest interface {
	SetQuery(key string, value interface{}) (IRequest)
	SetParam(key string, value interface{}) (IRequest)
}

type Client struct {
	Exchange IExchange
	request  *binance_conn.Request
}

func NewClient(ex IExchange) *Client {
	return &Client{
		Exchange: ex,
	}
}

func (c *Client) Request(method string, endpoint string, key bool, signed bool, args ...any) *Client {
	c.request = c.Exchange.Request(method, endpoint, key, signed, args)
	return c
}

func (c *Client) SetQuery(key string, value interface{}) *Client {
	c.request.SetQuery(key, value)
	return c
}

func (c *Client) SetParam(key string, value interface{}) *Client {
	c.request.SetParam(key, value)
	return c
}

func (c *Client) Send() (data []byte, err error) {
	req, err := c.Exchange.SetRequest(c.request)
	if err != nil {
		return
	}
	data, err = c.Exchange.Call(req)
	if err != nil {
		return
	}
	return
}
