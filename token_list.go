package birdeye

import (
	"fmt"
	"net/url"
)

type TokenListV1 struct {
	Success bool `json:"success"`
	Data    struct {
		UpdateUnixTime int    `json:"updateUnixTime"`
		UpdateTime     string `json:"updateTime"`
		Tokens         []struct {
			Address           string  `json:"address"`
			Decimals          int     `json:"decimals"`
			Price             float64 `json:"price"`
			LastTradeUnixTime int     `json:"lastTradeUnixTime"`
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
