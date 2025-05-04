package birdeye

import (
	"fmt"
	"net/url"
)

type Holder struct {
	Success bool `json:"success"`
	Data    struct {
		Items []struct {
			Amount       string `json:"amount"`
			Decimals     int    `json:"decimals"`
			Mint         string `json:"mint"`
			Owner        string `json:"owner"`
			TokenAccount string `json:"token_account"`
			UiAmount     int    `json:"ui_amount"`
		} `json:"items"`
	} `json:"data"`
}

func (srv *Client) GetHolder(address string, offset, limit int) (*Holder, error) {
	urlStr := fmt.Sprintf("%s/v3/token/holder", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("address", address)
	query.Add("offset", fmt.Sprint(offset))
	query.Add("limit", fmt.Sprint(limit))
	u.RawQuery = query.Encode()
	return newCli[Holder](srv.apiKey).makeRequest("GET", u.String(), nil)
}
