package birdeye

import (
	"fmt"
	"net/url"
)

type TokenTrending struct {
	Data struct {
		UpdateUnixTime int    `json:"updateUnixTime"`
		UpdateTime     string `json:"updateTime"`
		Tokens         []struct {
			Address                string  `json:"address"`
			Decimals               int     `json:"decimals"`
			Liquidity              float64 `json:"liquidity"`
			LogoURI                string  `json:"logoURI"`
			Name                   string  `json:"name"`
			Symbol                 string  `json:"symbol"`
			Volume24HUSD           float64 `json:"volume24hUSD"`
			Volume24HChangePercent int     `json:"volume24hChangePercent"`
			Rank                   int     `json:"rank"`
			Price                  float64 `json:"price"`
			Price24HChangePercent  int     `json:"price24hChangePercent"`
			Fdv                    int64   `json:"fdv"`
			Marketcap              int     `json:"marketcap"`
		} `json:"tokens"`
		Total int `json:"total"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetTokenTrending(sortType, sortBy string, offset, limit int) (*TokenTrending, error) {
	urlStr := fmt.Sprintf("%s/token_trending", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("sort_type", sortType)
	query.Add("sort_by", sortBy)
	query.Add("offset", fmt.Sprint(offset))
	query.Add("limit", fmt.Sprint(limit))
	u.RawQuery = query.Encode()
	return newCli[TokenTrending](srv.apiKey).makeRequest("GET", u.String(), nil)
}
