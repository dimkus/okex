package ws

import (
	"fmt"
	"github.com/dimkus/okex"
	"github.com/dimkus/okex/events"
	"github.com/dimkus/okex/events/public"
	requests "github.com/dimkus/okex/requests/ws/public"
	"github.com/goccy/go-json"
	"strings"
)

// Public
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels
type Public struct {
	*ClientWs
	iCh    chan *public.Instruments
	tCh    chan *public.Tickers
	oiCh   chan *public.OpenInterest
	cCh    chan *public.Candlesticks
	trCh   chan *public.Trades
	edepCh chan *public.EstimatedDeliveryExercisePrice
	mpCh   chan *public.MarkPrice
	mpcCh  chan *public.MarkPriceCandlesticks
	plCh   chan *public.PriceLimit
	obCh   chan *public.OrderBook
	osCh   chan *public.OPTIONSummary
	frCh   chan *public.FundingRate
	icCh   chan *public.IndexCandlesticks
	itCh   chan *public.IndexTickers
}

// NewPublic returns a pointer to a fresh Public
func NewPublic(c *ClientWs) *Public {
	return &Public{ClientWs: c}
}

// Instruments
// The full instrument list will be pushed for the first time after subscription. Subsequently, the instruments will be pushed if there's any change to the instrument’s state (such as delivery of FUTURES, exercise of OPTION, listing of new contracts / trading pairs, trading suspension, etc.).
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-instruments-channel
func (p *Public) Instruments(req requests.Instruments, ch ...chan *public.Instruments) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.iCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"instruments"}, m)
}

// UInstruments
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-instruments-channel
func (p *Public) UInstruments(req requests.Instruments, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.iCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"instruments"}, m)
}

// Tickers
// Retrieve the last traded price, bid price, ask price and 24-hour trading volume of instruments. Data will be pushed every 100 ms.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-tickers-channel
func (p *Public) Tickers(req requests.Tickers, ch ...chan *public.Tickers) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.tCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"tickers"}, m)
}

// UTickers
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-tickers-channel
func (p *Public) UTickers(req requests.Tickers, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.tCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"tickers"}, m)
}

// OpenInterest
// Retrieve the open interest. Data will by pushed every 3 seconds.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-open-interest-channel
func (p *Public) OpenInterest(req requests.OpenInterest, ch ...chan *public.OpenInterest) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.oiCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"open-interest"}, m)
}

// UOpenInterest
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-open-interest-channel
func (p *Public) UOpenInterest(req requests.OpenInterest, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.oiCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"open-interest"}, m)
}

// Candlesticks
// Retrieve the open interest. Data will by pushed every 3 seconds.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (p *Public) Candlesticks(req requests.Candlesticks, ch ...chan *public.Candlesticks) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.cCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{}, m)
}

// UCandlesticks
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (p *Public) UCandlesticks(req requests.Candlesticks, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.cCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{}, m)
}

// Trades
// Retrieve the recent trades data. Data will be pushed whenever there is a trade.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-trades-channel
func (p *Public) Trades(req requests.Trades, ch ...chan *public.Trades) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.trCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"trades"}, m)
}

// UTrades
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-trades-channel
func (p *Public) UTrades(req requests.Trades, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.trCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"trades"}, m)
}

// EstimatedDeliveryExercisePrice
// Retrieve the estimated delivery/exercise price of FUTURES contracts and OPTION.
//
// Only the estimated delivery/exercise price will be pushed an hour before delivery/exercise, and will be pushed if there is any price change.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-estimated-delivery-exercise-price-channel
func (p *Public) EstimatedDeliveryExercisePrice(req requests.EstimatedDeliveryExercisePrice, ch ...chan *public.EstimatedDeliveryExercisePrice) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.edepCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"estimated-price"}, m)
}

// UEstimatedDeliveryExercisePrice
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-estimated-delivery-exercise-price-channel
func (p *Public) UEstimatedDeliveryExercisePrice(req requests.EstimatedDeliveryExercisePrice, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.edepCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"estimated-price"}, m)
}

// MarkPrice
// Retrieve the mark price. Data will be pushed every 200 ms when the mark price changes, and will be pushed every 10 seconds when the mark price does not change.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-mark-price-channel
func (p *Public) MarkPrice(req requests.MarkPrice, ch ...chan *public.MarkPrice) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.mpCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"mark-price"}, m)
}

// UMarkPrice
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-mark-price-channel
func (p *Public) UMarkPrice(req requests.MarkPrice, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.mpCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"mark-price"}, m)
}

// MarkPriceCandlesticks
// Retrieve the candlesticks data of the mark price. Data will be pushed every 500 ms.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-mark-price-candlesticks-channel
func (p *Public) MarkPriceCandlesticks(req requests.MarkPriceCandlesticks, ch ...chan *public.MarkPriceCandlesticks) error {
	m := okex.S2M(req)
	m["channel"] = "mark-price-" + m["channel"]
	if len(ch) > 0 {
		p.mpcCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{}, m)
}

// UMarkPriceCandlesticks
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-mark-price-candlesticks-channel
func (p *Public) UMarkPriceCandlesticks(req requests.MarkPriceCandlesticks, rCh ...bool) error {
	m := okex.S2M(req)
	m["channel"] = "mark-price-" + m["channel"]
	if len(rCh) > 0 && rCh[0] {
		p.mpcCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{}, m)
}

// PriceLimit
// Retrieve the maximum buy price and minimum sell price of the instrument. Data will be pushed every 5 seconds when there are changes in limits, and will not be pushed when there is no changes on limit.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-price-limit-channel
func (p *Public) PriceLimit(req requests.PriceLimit, ch ...chan *public.PriceLimit) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.plCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"price-limit"}, m)
}

// UPriceLimit
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-price-limit-channel
func (p *Public) UPriceLimit(req requests.PriceLimit, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.plCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"price-limit"}, m)
}

// OrderBook
// Retrieve order book data for multiple instruments.
//
// Use books for 400 depth levels, books5 for 5 depth levels, books50-l2-tbt tick-by-tick 50 depth levels, and books-l2-tbt for tick-by-tick 400 depth levels.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-order-book-channel
func (p *Public) OrderBook(reqs []requests.OrderBook, ch ...chan *public.OrderBook) error {
	if len(ch) > 0 {
		p.obCh = ch[0]
	}
	var subscriptions []map[string]string
	for _, req := range reqs {
		m := okex.S2M(req)
		subscriptions = append(subscriptions, m)
	}
	return p.Subscribe(false, []okex.ChannelName{}, subscriptions...)
}

// UOrderBook
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-order-book-channel
func (p *Public) UOrderBook(req requests.OrderBook, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.obCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{okex.ChannelName(req.Channel)}, m)
}

// OPTIONSummary
// Retrieve detailed pricing information of all OPTION contracts. Data will be pushed at once.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-option-summary-channel
func (p *Public) OPTIONSummary(req requests.OPTIONSummary, ch ...chan *public.OPTIONSummary) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.osCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"opt-summary"}, m)
}

// UOPTIONSummary
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-option-summary-channel
func (p *Public) UOPTIONSummary(req requests.OPTIONSummary, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.osCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"opt-summary"}, m)
}

// FundingRate
// Retrieve funding rate. Data will be pushed every minute.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-funding-rate-channel
func (p *Public) FundingRate(req requests.FundingRate, ch ...chan *public.FundingRate) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.frCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"funding-rate"}, m)
}

// UFundingRate
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-funding-rate-channel
func (p *Public) UFundingRate(req requests.FundingRate, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.frCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"funding-rate"}, m)
}

// IndexCandlesticks
// Retrieve the candlesticks data of the index. Data will be pushed every 500 ms.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-index-candlesticks-channel
func (p *Public) IndexCandlesticks(req requests.IndexCandlesticks, ch ...chan *public.IndexCandlesticks) error {
	m := okex.S2M(req)
	m["channel"] = req.Channel
	if len(ch) > 0 {
		p.icCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{}, m)
}

// UIndexCandlesticks
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-index-candlesticks-channel
func (p *Public) UIndexCandlesticks(req requests.IndexCandlesticks, rCh ...bool) error {
	m := okex.S2M(req)
	m["channel"] = req.Channel
	if len(rCh) > 0 && rCh[0] {
		p.icCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{}, m)
}

// IndexTickers
// Retrieve index tickers data
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-index-tickers-channel
func (p *Public) IndexTickers(req requests.IndexTickers, ch ...chan *public.IndexTickers) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		p.itCh = ch[0]
	}
	return p.Subscribe(false, []okex.ChannelName{"index-tickers"}, m)
}

// UIndexTickers
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-index-tickers-channel
func (p *Public) UIndexTickers(req requests.IndexTickers, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		p.itCh = nil
	}
	return p.Unsubscribe(false, []okex.ChannelName{"index-tickers"}, m)
}

func (p *Public) Process(data []byte, e *events.Basic) bool {
	if e.Event == "" && e.Arg != nil && e.Data != nil && len(e.Data) > 0 {
		ch, ok := e.Arg.Get("channel")
		if !ok {
			return false
		}
		switch ch {
		case "instruments":
			if p.iCh == nil {
				return false
			}
			e := new(public.Instruments)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.iCh <- e
			return true
		case "tickers":
			if p.tCh == nil {
				return false
			}
			e := new(public.Tickers)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.tCh <- e
			return true
		case "open-interest":
			if p.oiCh == nil {
				return false
			}
			e := new(public.OpenInterest)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.oiCh <- e
			return true
		case "trades":
			if p.trCh == nil {
				return false
			}
			e := new(public.Trades)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.trCh <- e
			return true
		case "estimated-price":
			if p.edepCh == nil {
				return false
			}
			e := new(public.EstimatedDeliveryExercisePrice)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.edepCh <- e
			return true
		case "mark-price":
			if p.mpCh == nil {
				return false
			}
			e := new(public.MarkPrice)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.mpCh <- e
			return true
		case "price-limit":
			if p.plCh == nil {
				return false
			}
			e := new(public.PriceLimit)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.plCh <- e
			return true
		case "opt-summary":
			if p.osCh == nil {
				return false
			}
			e := new(public.OPTIONSummary)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.osCh <- e
			return true
		case "funding-rate":
			if p.osCh == nil {
				return false
			}
			e := new(public.OPTIONSummary)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.osCh <- e
			return true
		case "index-tickers":
			if p.itCh == nil {
				return false
			}
			e := new(public.IndexTickers)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.itCh <- e
			return true
		case "bbo-tbt":
			// order book
			// 1 depth level snapshot will be pushed every time. Snapshot data will be pushed every 10 ms when there are changes in the 1 depth level snapshot.
			if p.obCh == nil {
				return false
			}
			e := new(public.OrderBook)
			err := json.Unmarshal(data, e)
			if err != nil {
				return false
			}
			p.obCh <- e
			return true
		default:
			// special cases
			// market price candlestick channel
			chName := fmt.Sprint(ch)
			// market price channels
			if strings.Contains(chName, "mark-price-candle") {
				if p.mpcCh == nil {
					return false
				}
				e := new(public.MarkPriceCandlesticks)
				err := json.Unmarshal(data, e)
				if err != nil {
					return false
				}
				p.mpcCh <- e
				return true
			} else if strings.Contains(chName, "index-candle") {
				// index chandlestick channels
				if p.icCh == nil {
					return false
				}
				e := new(public.IndexCandlesticks)
				err := json.Unmarshal(data, e)
				if err != nil {
					return false
				}
				p.icCh <- e
				return true
			} else if strings.Contains(chName, "candle") {
				// candlestick channels
				if p.cCh == nil {
					return false
				}
				e := new(public.Candlesticks)
				err := json.Unmarshal(data, &e)
				if err != nil {
					return false
				}
				p.cCh <- e
				return true
			} else if strings.Contains(chName, "books") {
				// order book channels
				if p.obCh == nil {
					return false
				}
				e := new(public.OrderBook)
				err := json.Unmarshal(data, e)
				if err != nil {
					return false
				}
				p.obCh <- e
				return true
			}
		}
	}
	return false
}
