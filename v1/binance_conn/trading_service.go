package binance_conn

import (
	"net/http"
	"github.com/lianyun0502/exchange_conn/binance_conn/enums"
	"errors"
)

type NewOrder struct {
	symbol        string
	side          enums.OrderSide
	type_         enums.OrderType
	timeInForce   *enums.TimeInForce
	quantity      *float64
	quoteOrderQty *float64
	price         *float64

	// A unique id among open orders. Automatically generated if not sent.
	// Orders with the same newClientOrderID can be accepted only when the previous one is filled,
	// otherwise the order will be rejected.
	newClientOrderId *string
	strategyId       *int32
	strategyType     *int32   //The value cannot be less than 1000000
	stopPrice        *float64 // Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.

	// the percentage of movement in the opposite direction you are willing to tolerateOrderSide
	// The range you can set is 0.1% to 20%.
	// 	Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
	trailingDelta *int64

	// 拆分訂單的最小單位 ex 10BTC iceBergQty = 10 ,每次最少就交易 1BTC 所需的金額 ,無法整除會拆成商+餘數
	// 	Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order.
	iceBergQty *float64

	newOrderRespType        *enums.NewOrderRespType
	selfTradePreventionMode *enums.SelfTradePreventionMode
	receiveWindow           *int64
	// timeStamp               *int64
}

func (api *NewOrder) Do (c *Client) (data []byte, err error) {
	req := NewBinanceRequest(
		http.MethodPost,
		// "/api/v3/order",
		"/api/v3/order/test",
		Trade,
	)
	if api.symbol == "" {
		return nil, errors.New("symbol is required")
	}
	req.SetQuery("symbol", api.symbol)
	if api.side == "" {
		return nil, errors.New("side is required")
	}
	req.SetQuery("side", string(api.side))
	if api.type_ == "" {
		return nil, errors.New("type is required")
	}
	req.SetQuery("type", string(api.type_))
	switch api.type_ {
	case enums.Limit:
		if api.timeInForce == nil || api.quantity == nil || api.price == nil {
			return nil, errors.New("in type LIMIT, timeInForce, quantity, price is required")
		}
		if api.iceBergQty != nil {
			req.SetParam("iceBergQty", *api.iceBergQty)
			api.TimeInForce(enums.GTC)
		}
		req.SetParam("timeInForce", *api.timeInForce)
		req.SetParam("quantity", *api.quantity)
		req.SetParam("price", *api.price)
	case enums.Market:
		if api.quantity == nil {
			return nil, errors.New("in type MARKET, quantity is required")
		}
		req.SetParam("quantity", *api.quantity)
	case enums.StopLoss:
		if api.quantity == nil || api.stopPrice == nil || api.trailingDelta == nil {
			return nil, errors.New("in type STOP_LOSS, quantity, stopPrice, trailingDelta is required")
		}
		req.SetParam("quantity", *api.quantity)
		req.SetParam("stopPrice", *api.stopPrice)
		req.SetParam("trailingDelta", *api.trailingDelta)
	case enums.StopLossLimit:
		if api.timeInForce == nil || api.quantity == nil || api.price == nil || 
		api.stopPrice == nil || api.trailingDelta == nil {
			return nil, errors.New("in type STOP_LOSS_LIMIT, timeInForce, quantity, price, stopPrice, trailingDelta is required")
		}
		req.SetParam("timeInForce", *api.timeInForce)
		req.SetParam("quantity", *api.quantity)
		req.SetParam("price", *api.price)
		req.SetParam("stopPrice", *api.stopPrice)
		req.SetParam("trailingDelta", *api.trailingDelta)
	case enums.TakeProfit:
		if api.quantity == nil || api.stopPrice == nil || api.trailingDelta == nil {
			return nil, errors.New("in type TAKE_PROFIT, quantity, stopPrice, trailingDelta is required")
		}
		req.SetParam("quantity", *api.quantity)
		req.SetParam("stopPrice", *api.stopPrice)
		req.SetParam("trailingDelta", *api.trailingDelta)
	case enums.TakeProfitLimit:
		if api.timeInForce == nil || api.quantity == nil || api.price == nil ||
		api.stopPrice == nil || api.trailingDelta == nil {
			return nil, errors.New("in type TAKE_PROFIT_LIMIT, timeInForce, quantity, price, stopPrice, trailingDelta is required")
		}
		req.SetParam("timeInForce", *api.timeInForce)
		req.SetParam("quantity", *api.quantity)
		req.SetParam("price", *api.price)
		req.SetParam("stopPrice", *api.stopPrice)
		req.SetParam("trailingDelta", *api.trailingDelta)
	case enums.LimitMaker:
		if api.quantity == nil || api.price == nil {
			return nil, errors.New("in type LIMIT_MAKER, quantity, price is required")
		}
		if api.iceBergQty != nil {
			req.SetParam("iceBergQty", *api.iceBergQty)
			api.TimeInForce(enums.GTC)
		}
		req.SetParam("quantity", *api.quantity)
		req.SetParam("price", *api.price)
	}
	if api.newClientOrderId != nil {req.SetParam("quoteOrderQty", api.quoteOrderQty)}
	if api.newClientOrderId != nil {req.SetParam("newClientOrderId", *api.newClientOrderId)}
	if api.strategyId != nil {req.SetParam("strategyId", *api.strategyId)}
	if api.strategyType != nil {req.SetParam("strategyType", *api.strategyType)}
	if api.newOrderRespType != nil {req.SetParam("newOrderRespType", *api.newOrderRespType)}
	if api.selfTradePreventionMode != nil {req.SetParam("selfTradePreventionMode", *api.selfTradePreventionMode)}
	if api.receiveWindow != nil {req.SetParam("receiveWindow", *api.receiveWindow)}
	// if api.timeStamp != nil {req.SetParam("timeStamp", *api.timeStamp)}


	b_req, err := c.SetBinanceRequest(req)
	if err != nil {
		return
	}

	data, err = c.Call(b_req)
	if err != nil {
		return
	}

	return
}

func (api *NewOrder) Symbol(s string) *NewOrder {
	api.symbol = s
	return api
}

func (api *NewOrder) Side(s enums.OrderSide) *NewOrder {
	api.side = s
	return api
}

func (api *NewOrder) Type(t enums.OrderType) *NewOrder {
	api.type_ = t
	return api
}

func (api *NewOrder) TimeInForce(t enums.TimeInForce) *NewOrder {
	api.timeInForce = &t
	return api
}

func (api *NewOrder) Quantity(q float64) *NewOrder {
	api.quantity = &q
	return api
}

func (api *NewOrder) QuoteOrderQty(q float64) *NewOrder {
	api.quoteOrderQty = &q
	return api
}

func (api *NewOrder) Price(p float64) *NewOrder {
	api.price = &p
	return api
}

func (api *NewOrder) NewClientOrderId(id string) *NewOrder {
	api.newClientOrderId = &id
	return api
}

func (api *NewOrder) StrategyId(id int32) *NewOrder {
	api.strategyId = &id
	return api
}

func (api *NewOrder) StrategyType(t int32) *NewOrder {
	api.strategyType = &t
	return api
}

func (api *NewOrder) StopPrice(p float64) *NewOrder {
	api.stopPrice = &p
	return api
}

func (api *NewOrder) TrailingDelta(d int64) *NewOrder {
	api.trailingDelta = &d
	return api
}

func (api *NewOrder) IceBergQty(q float64) *NewOrder {
	api.iceBergQty = &q
	return api
}

func (api *NewOrder) NewOrderRespType(t enums.NewOrderRespType) *NewOrder {
	api.newOrderRespType = &t
	return api
}

func (api *NewOrder) SelfTradePreventionMode(m enums.SelfTradePreventionMode) *NewOrder {
	api.selfTradePreventionMode = &m
	return api
}

func (api *NewOrder) ReceiveWindow(w int64) *NewOrder {
	api.receiveWindow = &w
	return api
}

