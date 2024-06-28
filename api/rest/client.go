package rest

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dimkus/okex"
	requests "github.com/dimkus/okex/requests/rest/public"
	responses2 "github.com/dimkus/okex/responses"
	responses "github.com/dimkus/okex/responses/public_data"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"strings"
	"time"
)

// ClientRest is the rest api client
type ClientRest struct {
	Account        *Account
	SubAccount     *SubAccount
	Trade          *Trade
	Funding        *Funding
	Market         *Market
	PublicData     *PublicData
	TradeData      *TradeData
	apiKey         string
	secretKey      []byte
	passphrase     string
	destination    okex.Destination
	baseURL        okex.BaseURL
	client         *http.Client
	serverTimeDiff time.Duration
}

// NewClient returns a pointer to a fresh ClientRest
func NewClient(apiKey, secretKey, passphrase string, baseURL okex.BaseURL, destination okex.Destination) *ClientRest {
	c := &ClientRest{
		apiKey:      apiKey,
		secretKey:   []byte(secretKey),
		passphrase:  passphrase,
		baseURL:     baseURL,
		destination: destination,
		client:      http.DefaultClient,
	}
	c.Account = NewAccount(c)
	c.SubAccount = NewSubAccount(c)
	c.Trade = NewTrade(c)
	c.Funding = NewFunding(c)
	c.Market = NewMarket(c)
	c.PublicData = NewPublicData(c)
	c.TradeData = NewTradeData(c)
	return c
}

func (c *ClientRest) WithHTTPClient(httpClient *http.Client) *ClientRest {
	c.client = httpClient
	return c
}

// Do the http request to the server
func (c *ClientRest) Do(ctx context.Context, method, path string, private bool, params ...map[string]string) (*http.Response, error) {
	u := fmt.Sprintf("%s%s", c.baseURL, path)
	var (
		r    *http.Request
		err  error
		j    []byte
		body string
	)
	if method == http.MethodGet {
		r, err = http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
		if err != nil {
			return nil, err
		}
		if len(params) > 0 {
			q := r.URL.Query()
			for k, v := range params[0] {
				q.Add(k, strings.ReplaceAll(v, "\"", ""))
			}
			r.URL.RawQuery = q.Encode()
			if len(params[0]) > 0 {
				path += "?" + r.URL.RawQuery
			}
		}
	} else {
		j, err = json.Marshal(params[0])
		if err != nil {
			return nil, err
		}
		body = string(j)
		if body == "{}" {
			body = ""
		}
		r, err = http.NewRequestWithContext(ctx, method, u, bytes.NewBuffer(j))
		if err != nil {
			return nil, err
		}
		r.Header.Add("Content-Type", "application/json")
	}
	if err != nil {
		return nil, err
	}
	if private {
		timestamp, sign := c.sign(method, path, body)
		r.Header.Add("OK-ACCESS-KEY", c.apiKey)
		r.Header.Add("OK-ACCESS-PASSPHRASE", c.passphrase)
		r.Header.Add("OK-ACCESS-SIGN", sign)
		r.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)
	}
	if c.destination == okex.DemoServer {
		r.Header.Add("x-simulated-trading", "1")
	}
	return c.client.Do(r)
}

// Status
// Get event status of system upgrade
//
// https://www.okex.com/docs-v5/en/#rest-api-status
func (c *ClientRest) Status(ctx context.Context, req requests.Status) (response responses.Status, err error) {
	p := "/api/v5/system/status"
	m := okex.S2M(req)
	res, err := c.Do(ctx, http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

func (c *ClientRest) SetSystemTimeDiff(timeDiff time.Duration) {
	c.serverTimeDiff = timeDiff
}

func (c *ClientRest) sign(method, path, body string) (string, string) {
	format := "2006-01-02T15:04:05.999Z07:00"
	currentTime := time.Now().Add(c.serverTimeDiff).UTC()
	t := currentTime.Format(format)
	ts := fmt.Sprint(t)
	s := ts + method + path + body
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *ClientRest) decode(reader io.Reader, v any) error {
	err := json.NewDecoder(reader).Decode(&v)

	if err != nil {
		return err
	}

	_, ok := v.(responses2.BasicI)
	if !ok {
		return nil
	}

	vBasic := v.(responses2.BasicI)
	if vBasic.GetCode() != 0 {
		return errors.New(fmt.Sprintf("code: %d, msg: %s", vBasic.GetCode(), vBasic.GetMsg()))
	}

	return nil
}
