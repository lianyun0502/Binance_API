package binance_connect


import (
	"encoding/json"
)


func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

