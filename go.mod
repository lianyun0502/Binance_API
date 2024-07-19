module practice_go/binance_api

go 1.22.5

require practice_go/binance_conn v0.0.0-00010101000000-000000000000

require (
	github.com/dolthub/maphash v0.1.0 // indirect
	github.com/klauspost/compress v1.17.5 // indirect
	github.com/lxzan/gws v1.8.5 // indirect
)

replace practice_go/binance_conn => ./v1
