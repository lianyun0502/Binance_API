package binance_conn

import (
	"net/http"
)

type CreateAPIKeyService struct {}

func (api *CreateAPIKeyService) Do(c *Client) (data []byte, err error){
	res := NewBinanceRequest(
		http.MethodPost,
		"/api/v3/userDataStream",
		UserStream,
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