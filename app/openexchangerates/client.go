package openexchangerates

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"
)

const BaseUSD string = "USD"

type RatesResponse struct {
	Rates map[string]float64 `json:"rates"`
}

type Client struct {
	appID  string
	url    string
	client *resty.Client
}

func NewClient(url, appID string) *Client {
	return &Client{
		appID:  appID,
		url:    url,
		client: resty.New(),
	}
}

func (c *Client) GetExchangeRates(currencies []string) (*RatesResponse, error) {
	resp, err := c.client.R().
		SetQueryParam("app_id", c.appID).
		SetQueryParam("symbols", strings.Join(currencies, ",")).
		SetQueryParam("base", BaseUSD).
		SetResult(&RatesResponse{}).
		Get(fmt.Sprintf("%s%s", c.url, "/api/latest.json"))

	if err != nil {
		return nil, fmt.Errorf("[1633483348] client error: %v", err)
	}

	if resp.Error() != nil {
		return nil, fmt.Errorf("[1950155015] client error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("[1950245024] %s", resp.Status())
	}

	return resp.Result().(*RatesResponse), nil
}
