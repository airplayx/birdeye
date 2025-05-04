package birdeye

import (
	"fmt"
	"net/url"
	"strconv"
)

type NewListing struct {
	Success bool `json:"success"`
	Data    struct {
		Items []struct {
			Address          string  `json:"address"`
			Symbol           string  `json:"symbol"`
			Name             string  `json:"name"`
			Decimals         int     `json:"decimals"`
			Source           string  `json:"source"`
			LiquidityAddedAt string  `json:"liquidityAddedAt"`
			LogoURI          *string `json:"logoURI"`
			Liquidity        float64 `json:"liquidity"`
		} `json:"items"`
	} `json:"data"`
}

func (srv *Client) GetNewListing(timeTo int64, limit int, memePlatformEnabled bool) (*NewListing, error) {
	urlStr := fmt.Sprintf("%s/v2/tokens/new_listing", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	query := u.Query()
	query.Add("time_to", fmt.Sprint(timeTo))
	query.Add("limit", fmt.Sprint(limit))
	query.Add("meme_platform_enabled", strconv.FormatBool(memePlatformEnabled))
	u.RawQuery = query.Encode()
	return newCli[NewListing](srv.apiKey).makeRequest("GET", u.String(), nil)
}
