module practice_go/binance_api

go 1.22.5

require (
	github.com/binance/binance-connector-go v0.6.0
	practice_go/binance_conn v0.0.0-00010101000000-000000000000
)

require (
	github.com/bitly/go-simplejson v0.5.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
)

replace practice_go/binance_conn => ./v1
