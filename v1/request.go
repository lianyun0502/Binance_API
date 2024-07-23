package binance_connect

import (
	"encoding/json"
)

type WSApiRequest struct {
	Id int `json:"id"`
	Method string `json:"method"`
	Params interface{} `json:"params"`
}


func (r *WSApiRequest) genText() []byte {
	payload, _ := json.Marshal(r)
	return payload
}