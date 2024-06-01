package ws

import (
	"github.com/dimkus/okex"
	"github.com/dimkus/okex/events"
	"github.com/dimkus/okex/events/private"
	requests "github.com/dimkus/okex/requests/ws/private"
	"github.com/goccy/go-json"
)

// Private
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel
type Private struct {
	*ClientWs
	aCh   chan *private.Account
	pCh   chan *private.Position
	bnpCh chan *private.BalanceAndPosition
	oCh   chan *private.Order
}

// NewPrivate returns a pointer to a fresh Private
func NewPrivate(c *ClientWs) *Private {
	return &Private{ClientWs: c}
}

// Account
// Retrieve account information. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-account-channel
func (p *Private) Account(req requests.Account, ch ...chan *private.Account) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.aCh = ch[0]
	}
	return p.Subscribe(true, []okex.ChannelName{"account"}, m)
}

// UAccount
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-account-channel
func (p *Private) UAccount(req requests.Account, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.aCh = nil
	}
	return p.Unsubscribe(true, []okex.ChannelName{"account"}, m)
}

// Position
// Retrieve position information. Initial snapshot will be pushed according to subscription granularity. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-positions-channel
func (p *Private) Position(req requests.Position, ch ...chan *private.Position) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.pCh = ch[0]
	}
	return p.Subscribe(true, []okex.ChannelName{"positions"}, m)
}

// UPosition
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-positions-channel
func (p *Private) UPosition(req requests.Position, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.pCh = nil
	}
	return p.Unsubscribe(true, []okex.ChannelName{"positions"}, m)
}

// BalanceAndPosition
// Retrieve account balance and position information. Data will be pushed when triggered by events such as filled order, funding transfer.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
func (p *Private) BalanceAndPosition(ch ...chan *private.BalanceAndPosition) error {
	m := make(map[string]string)
	if len(ch) > 0 {
		p.bnpCh = ch[0]
	}
	return p.Subscribe(true, []okex.ChannelName{"balance_and_position"}, m)
}

// UBalanceAndPosition unsubscribes a position channel
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
func (p *Private) UBalanceAndPosition(rCh ...bool) error {
	m := make(map[string]string)
	if len(rCh) > 0 && rCh[0] {
		p.bnpCh = nil
	}
	return p.Unsubscribe(true, []okex.ChannelName{"balance_and_position"}, m)
}

// Order
// Retrieve position information. Initial snapshot will be pushed according to subscription granularity. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-order-channel
func (p *Private) Order(req requests.Order, ch ...chan *private.Order) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.oCh = ch[0]
	}
	return p.Subscribe(true, []okex.ChannelName{"orders"}, m)
}

// UOrder
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-order-channel
func (p *Private) UOrder(req requests.Order, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.oCh = nil
	}
	return p.Unsubscribe(true, []okex.ChannelName{"orders"}, m)
}

func (p *Private) Process(data []byte, e *events.Basic) bool {
	if e.Event == "" && e.Arg != nil && e.Data != nil && len(e.Data) > 0 {
		ch, ok := e.Arg.Get("channel")
		if !ok {
			return false
		}
		switch ch {
		case "account":
			if p.aCh == nil {
				return false
			}
			e := new(private.Account)
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			p.aCh <- e
			return true
		case "positions":
			if p.pCh == nil {
				return false
			}
			e := new(private.Position)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.pCh <- e
			return true
		case "balance_and_position":
			if p.bnpCh == nil {
				return false
			}
			e := new(private.BalanceAndPosition)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.bnpCh <- e
			return true
		case "orders":
			if p.oCh == nil {
				return false
			}
			e := new(private.Order)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.oCh <- e
			return true
		}
	}
	return false
}
