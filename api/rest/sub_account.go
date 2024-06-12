package rest

import (
	"context"
	"github.com/dimkus/okex"
	requests "github.com/dimkus/okex/requests/rest/subaccount"
	responses "github.com/dimkus/okex/responses/sub_account"
	"net/http"
	"strings"
)

// SubAccount
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount
type SubAccount struct {
	client *ClientRest
}

// NewSubAccount returns a pointer to a fresh SubAccount
func NewSubAccount(c *ClientRest) *SubAccount {
	return &SubAccount{c}
}

// ViewList
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-view-sub-account-list
func (c *SubAccount) ViewList(ctx context.Context, req requests.ViewList) (response responses.ViewList, err error) {
	p := "/api/v5/users/subaccount/list"
	m := okex.S2M(req)
	res, err := c.client.Do(ctx, http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = c.client.decode(res.Body, &response)
	return
}

// CreateAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-create-an-apikey-for-a-sub-account
func (c *SubAccount) CreateAPIKey(ctx context.Context, req requests.CreateAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/apikey"
	m := okex.S2M(req)
	if len(req.IP) > 0 {
		m["ip"] = strings.Join(req.IP, ",")
	}
	res, err := c.client.Do(ctx, http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = c.client.decode(res.Body, &response)
	return
}

// QueryAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-query-the-apikey-of-a-sub-account
func (c *SubAccount) QueryAPIKey(ctx context.Context, req requests.QueryAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/apikey"
	m := okex.S2M(req)
	res, err := c.client.Do(ctx, http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = c.client.decode(res.Body, &response)
	return
}

// ResetAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-reset-the-apikey-of-a-sub-account
func (c *SubAccount) ResetAPIKey(ctx context.Context, req requests.CreateAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/modify-apikey"
	m := okex.S2M(req)
	if len(req.IP) > 0 {
		m["ip"] = strings.Join(req.IP, ",")
	}
	res, err := c.client.Do(ctx, http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = c.client.decode(res.Body, &response)
	return
}

// DeleteAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-delete-the-apikey-of-sub-accounts
func (c *SubAccount) DeleteAPIKey(ctx context.Context, req requests.DeleteAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/delete-apikey"
	m := okex.S2M(req)
	res, err := c.client.Do(ctx, http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = c.client.decode(res.Body, &response)
	return
}

// GetBalance
// Query detailed balance info of Trading Account of a sub-account via the master account
// (applies to master accounts only)
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-get-sub-account-balance
func (c *SubAccount) GetBalance(ctx context.Context, req requests.GetBalance) (response responses.GetBalance, err error) {
	p := "/api/v5/account/subaccount/balances"
	m := okex.S2M(req)
	res, err := c.client.Do(ctx, http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = c.client.decode(res.Body, &response)
	return
}

// HistoryTransfer
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-history-of-sub-account-transfer
func (c *SubAccount) HistoryTransfer(ctx context.Context, req requests.HistoryTransfer) (response responses.HistoryTransfer, err error) {
	p := "/api/v5/account/subaccount/bills"
	m := okex.S2M(req)
	res, err := c.client.Do(ctx, http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = c.client.decode(res.Body, &response)
	return
}

// ManageTransfers
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-master-accounts-manage-the-transfers-between-sub-accounts
func (c *SubAccount) ManageTransfers(ctx context.Context, req requests.ManageTransfers) (response responses.ManageTransfer, err error) {
	p := "/api/v5/account/subaccount/transfer"
	m := okex.S2M(req)
	res, err := c.client.Do(ctx, http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = c.client.decode(res.Body, &response)
	return
}
