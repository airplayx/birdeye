package birdeye

import (
	"fmt"
	"time"
)

type TokenCreationInfo struct {
	Success bool `json:"success"`
	Data    struct {
		TxHash         string    `json:"txHash"`
		Slot           int       `json:"slot"`
		TokenAddress   string    `json:"tokenAddress"`
		Decimals       int       `json:"decimals"`
		Owner          string    `json:"owner"`
		BlockUnixTime  int64     `json:"blockUnixTime"`
		BlockHumanTime time.Time `json:"blockHumanTime"`
	} `json:"data"`
}

func (srv *Client) GetTokenCreationInfo(address string) (*TokenCreationInfo, error) {
	url := fmt.Sprintf("%s/token_creation_info?address=%s", baseURL, address)
	return newCli[TokenCreationInfo](srv.apiKey).makeRequest("GET", url, nil)
}
