package birdeye

import (
	"fmt"
	"net/url"
)

type TokenListV1 struct {
	Success bool `json:"success"`
	Data    struct {
		UpdateUnixTime int64  `json:"updateUnixTime"`
		UpdateTime     string `json:"updateTime"`
		Tokens         []struct {
			Address           string  `json:"address"`
			Decimals          int     `json:"decimals"`
			Price             float64 `json:"price"`
			LastTradeUnixTime int64   `json:"lastTradeUnixTime"`
			Liquidity         float64 `json:"liquidity"`
			LogoURI           string  `json:"logoURI"`
			Mc                float64 `json:"mc"`
			Name              string  `json:"name"`
			Symbol            string  `json:"symbol"`
			V24HChangePercent float64 `json:"v24hChangePercent"`
			V24HUSD           float64 `json:"v24hUSD"`
		} `json:"tokens"`
		Total int `json:"total"`
	} `json:"data"`
}

func (srv *Client) GetTokenListV1(sortBy, sortType string, offset, limit, minLiquidity, maxLiquidity int) (*TokenListV1, error) {
	urlStr := fmt.Sprintf("%s/tokenlist", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("sort_by", sortBy)
	query.Add("sort_type", sortType)
	query.Add("offset", fmt.Sprint(offset))
	query.Add("limit", fmt.Sprint(limit))
	query.Add("min_liquidity", fmt.Sprint(minLiquidity))
	query.Add("max_liquidity", fmt.Sprint(maxLiquidity))
	u.RawQuery = query.Encode()
	return newCli[TokenListV1](srv.apiKey).makeRequest("GET", u.String(), nil)
}

type TokenListV3 struct {
	Data struct {
		Items []struct {
			Address    string `json:"address"`
			LogoUri    string `json:"logo_uri"`
			Name       string `json:"name"`
			Symbol     string `json:"symbol"`
			Decimals   int    `json:"decimals"`
			Extensions struct {
				CoingeckoId string `json:"coingecko_id"`
				SerumV3Usdc string `json:"serum_v3_usdc"`
				SerumV3Usdt string `json:"serum_v3_usdt"`
				Website     string `json:"website"`
				Telegram    any    `json:"telegram"`
				Twitter     string `json:"twitter"`
				Description string `json:"description"`
				Discord     string `json:"discord"`
				Medium      string `json:"medium"`
			} `json:"extensions"`
			MarketCap              float64 `json:"market_cap"`
			Fdv                    float64 `json:"fdv"`
			Liquidity              float64 `json:"liquidity"`
			LastTradeUnixTime      int64   `json:"last_trade_unix_time"`
			Volume1HUsd            float64 `json:"volume_1h_usd"`
			Volume1HChangePercent  float64 `json:"volume_1h_change_percent"`
			Volume2HUsd            float64 `json:"volume_2h_usd"`
			Volume2HChangePercent  float64 `json:"volume_2h_change_percent"`
			Volume4HUsd            float64 `json:"volume_4h_usd"`
			Volume4HChangePercent  float64 `json:"volume_4h_change_percent"`
			Volume8HUsd            float64 `json:"volume_8h_usd"`
			Volume8HChangePercent  float64 `json:"volume_8h_change_percent"`
			Volume24HUsd           float64 `json:"volume_24h_usd"`
			Volume24HChangePercent float64 `json:"volume_24h_change_percent"`
			Trade1HCount           int     `json:"trade_1h_count"`
			Trade2HCount           int     `json:"trade_2h_count"`
			Trade4HCount           int     `json:"trade_4h_count"`
			Trade8HCount           int     `json:"trade_8h_count"`
			Trade24HCount          int     `json:"trade_24h_count"`
			Price                  float64 `json:"price"`
			PriceChange1HPercent   float64 `json:"price_change_1h_percent"`
			PriceChange2HPercent   float64 `json:"price_change_2h_percent"`
			PriceChange4HPercent   float64 `json:"price_change_4h_percent"`
			PriceChange8HPercent   float64 `json:"price_change_8h_percent"`
			PriceChange24HPercent  float64 `json:"price_change_24h_percent"`
			Holder                 int     `json:"holder"`
			RecentListingTime      any     `json:"recent_listing_time"`
		} `json:"items"`
		HasNext bool `json:"has_next"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetTokenListV3(sortBy, sortType string, offset, limit, minLiquidity, maxLiquidity int) (*TokenListV3, error) {
	urlStr := fmt.Sprintf("%s/v3/token/list", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("sort_by", sortBy)
	query.Add("sort_type", sortType)
	query.Add("offset", fmt.Sprint(offset))
	query.Add("limit", fmt.Sprint(limit))
	query.Add("min_liquidity", fmt.Sprint(minLiquidity))
	query.Add("max_liquidity", fmt.Sprint(maxLiquidity))
	u.RawQuery = query.Encode()
	return newCli[TokenListV3](srv.apiKey).makeRequest("GET", u.String(), nil)
}

type TokenListV3Scroll struct {
	Data struct {
		NextScrollId string `json:"next_scroll_id"`
		ScrollTime   string `json:"scroll_time"`
		Items        []struct {
			Address                string  `json:"address"`
			LogoUri                string  `json:"logo_uri"`
			Name                   string  `json:"name"`
			Symbol                 string  `json:"symbol"`
			Decimals               int     `json:"decimals"`
			Extensions             any     `json:"extensions"`
			MarketCap              float64 `json:"market_cap"`
			Fdv                    float64 `json:"fdv"`
			Liquidity              float64 `json:"liquidity"`
			LastTradeUnixTime      int64   `json:"last_trade_unix_time"`
			Volume1HUsd            float64 `json:"volume_1h_usd"`
			Volume1HChangePercent  float64 `json:"volume_1h_change_percent"`
			Volume2HUsd            float64 `json:"volume_2h_usd"`
			Volume2HChangePercent  float64 `json:"volume_2h_change_percent"`
			Volume4HUsd            float64 `json:"volume_4h_usd"`
			Volume4HChangePercent  float64 `json:"volume_4h_change_percent"`
			Volume8HUsd            float64 `json:"volume_8h_usd"`
			Volume8HChangePercent  float64 `json:"volume_8h_change_percent"`
			Volume24HUsd           float64 `json:"volume_24h_usd"`
			Volume24HChangePercent float64 `json:"volume_24h_change_percent"`
			Trade1HCount           int     `json:"trade_1h_count"`
			Trade2HCount           int     `json:"trade_2h_count"`
			Trade4HCount           int     `json:"trade_4h_count"`
			Trade8HCount           int     `json:"trade_8h_count"`
			Trade24HCount          int     `json:"trade_24h_count"`
			Price                  float64 `json:"price"`
			PriceChange1HPercent   float64 `json:"price_change_1h_percent"`
			PriceChange2HPercent   float64 `json:"price_change_2h_percent"`
			PriceChange4HPercent   float64 `json:"price_change_4h_percent"`
			PriceChange8HPercent   float64 `json:"price_change_8h_percent"`
			PriceChange24HPercent  any     `json:"price_change_24h_percent"`
			Holder                 int     `json:"holder"`
			RecentListingTime      int     `json:"recent_listing_time"`
		} `json:"items"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetTokenListV3Scroll(sortBy, sortType string, offset, limit, minLiquidity, maxLiquidity int) (*TokenListV3Scroll, error) {
	urlStr := fmt.Sprintf("%s/v3/token/list/scroll", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("sort_by", sortBy)
	query.Add("sort_type", sortType)
	query.Add("offset", fmt.Sprint(offset))
	query.Add("limit", fmt.Sprint(limit))
	query.Add("min_liquidity", fmt.Sprint(minLiquidity))
	query.Add("max_liquidity", fmt.Sprint(maxLiquidity))
	u.RawQuery = query.Encode()
	return newCli[TokenListV3Scroll](srv.apiKey).makeRequest("GET", u.String(), nil)
}
