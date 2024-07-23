package binance_connect

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWSApiRequest_genText(t *testing.T) {
	assert := assert.New(t)
	ws := new(WSApiRequest)
	ws.Id = 1
	ws.Method = "test"
	ws.Params = &struct {
		Test string `json:"test"`
		Test2 string `json:"test2"`
	}{"test", "test2"}
	
	data := ws.genText()
	assert.NotEqual("{}", string(data))
	assert.Equal(`{"id":1,"method":"test","params":{"test":"test","test2":"test2"}}`, string(data))
}