package birdeye

import (
	"fmt"
	"strings"
)

type MetaDataSingle struct {
	Data struct {
		Address    string `json:"address"`
		Symbol     string `json:"symbol"`
		Name       string `json:"name"`
		Decimals   int    `json:"decimals"`
		Extensions struct {
			CoingeckoId string `json:"coingecko_id"`
			Website     string `json:"website"`
			Twitter     string `json:"twitter"`
			Discord     string `json:"discord"`
			Medium      string `json:"medium"`
		} `json:"extensions"`
		LogoUri string `json:"logo_uri"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetMetaDataSingle(address string) (*MetaDataSingle, error) {
	url := fmt.Sprintf("%s/v3/token/meta-data/single?address=%s", baseURL, address)
	return newCli[MetaDataSingle](srv.apiKey).makeRequest("GET", url, nil)
}

type MetaDataMultiple struct {
	Data map[string]struct {
		Address    string `json:"address"`
		Symbol     string `json:"symbol"`
		Name       string `json:"name"`
		Decimals   int    `json:"decimals"`
		Extensions struct {
			CoingeckoID string `json:"coingecko_id"`
			Website     string `json:"website"`
			Twitter     string `json:"twitter"`
			Discord     string `json:"discord"`
			Medium      string `json:"medium"`
		} `json:"extensions"`
		LogoURI string `json:"logo_uri"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetMetaDataMultiple(address ...string) (*MetaDataMultiple, error) {
	url := fmt.Sprintf("%s/v3/token/meta-data/multiple?list_address=%s", baseURL, strings.Join(address, ","))
	return newCli[MetaDataMultiple](srv.apiKey).makeRequest("GET", url, nil)
}
