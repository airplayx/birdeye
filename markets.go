package birdeye

import (
	"fmt"
	"net/url"
	"time"
)

type Markets struct {
	Success bool `json:"success"`
	Data    struct {
		Items []struct {
			Address string `json:"address"`
			Base    struct {
				Address  string `json:"address"`
				Decimals int    `json:"decimals"`
				Symbol   string `json:"symbol"`
				Icon     string `json:"icon"`
			} `json:"base"`
			CreatedAt time.Time   `json:"createdAt"`
			Liquidity float64     `json:"liquidity"`
			Name      string      `json:"name"`
			Price     interface{} `json:"price"`
			Quote     struct {
				Address  string `json:"address"`
				Decimals int    `json:"decimals"`
				Icon     string `json:"icon"`
				Symbol   string `json:"symbol"`
			} `json:"quote"`
			Source                       string  `json:"source"`
			Trade24H                     int     `json:"trade24h"`
			Trade24HChangePercent        float64 `json:"trade24hChangePercent"`
			UniqueWallet24H              int     `json:"uniqueWallet24h"`
			UniqueWallet24HChangePercent float64 `json:"uniqueWallet24hChangePercent"`
			Volume24H                    float64 `json:"volume24h"`
		} `json:"items"`
		Total int `json:"total"`
	} `json:"data"`
}

func (srv *Client) GetMarkets(address, timeFrame, sortType, sortBy string, offset, limit int) (*Markets, error) {
	urlStr := fmt.Sprintf("%s/v2/markets", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("address", address)
	query.Add("time_frame", timeFrame)
	query.Add("sort_type", sortType)
	query.Add("sort_by", sortBy)
	query.Add("offset", fmt.Sprint(offset))
	query.Add("limit", fmt.Sprint(limit))
	u.RawQuery = query.Encode()
	return newCli[Markets](srv.apiKey).makeRequest("GET", u.String(), nil)
}
