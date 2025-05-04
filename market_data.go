package birdeye

import (
	"fmt"
	"strings"
)

type MarketDataSingle struct {
	Data struct {
		Address           string  `json:"address"`
		Liquidity         float64 `json:"liquidity"`
		Price             float64 `json:"price"`
		TotalSupply       float64 `json:"total_supply"`
		CirculatingSupply float64 `json:"circulating_supply"`
		Fdv               float64 `json:"fdv"`
		MarketCap         float64 `json:"market_cap"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetMarketDataSingle(address string) (*MarketDataSingle, error) {
	url := fmt.Sprintf("%s/v3/token/meta-data/single?address=%s", baseURL, address)
	return newCli[MarketDataSingle](srv.apiKey).makeRequest("GET", url, nil)
}

type MarketDataMultiple struct {
	Data map[string]struct {
		Address           string  `json:"address"`
		Liquidity         float64 `json:"liquidity"`
		Price             float64 `json:"price"`
		TotalSupply       float64 `json:"total_supply"`
		CirculatingSupply float64 `json:"circulating_supply"`
		FDV               float64 `json:"fdv"`
		MarketCap         float64 `json:"market_cap"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetMarketDataMultiple(address ...string) (*MarketDataMultiple, error) {
	url := fmt.Sprintf("%s/v3/token/market-data/multiple?list_address=%s", baseURL, strings.Join(address, ","))
	return newCli[MarketDataMultiple](srv.apiKey).makeRequest("GET", url, nil)
}
