package birdeye

import (
	"fmt"
	"net/url"
	"time"
)

type MintBurnTxs struct {
	Success bool `json:"success"`
	Data    struct {
		Items []struct {
			Amount         string    `json:"amount"`
			BlockHumanTime time.Time `json:"block_human_time"`
			BlockTime      int       `json:"block_time"`
			CommonType     string    `json:"common_type"`
			Decimals       int       `json:"decimals"`
			Mint           string    `json:"mint"`
			ProgramId      string    `json:"program_id"`
			Slot           int       `json:"slot"`
			TxHash         string    `json:"tx_hash"`
			UiAmount       int       `json:"ui_amount"`
			UiAmountString string    `json:"ui_amount_string"`
		} `json:"items"`
	} `json:"data"`
}

func (srv *Client) GetMintBurnTxs(address, sortType, sortBy, Type string, afterTime, beforeTime, offset, limit int) (*MintBurnTxs, error) {
	urlStr := fmt.Sprintf("%s/v3/token/mint-burn-txs", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("address", address)
	query.Add("sort_type", sortType)
	query.Add("sort_by", sortBy)
	query.Add("type", Type)
	query.Add("after_time", fmt.Sprint(afterTime))
	query.Add("before_time", fmt.Sprint(beforeTime))
	query.Add("offset", fmt.Sprint(offset))
	query.Add("limit", fmt.Sprint(limit))
	u.RawQuery = query.Encode()
	return newCli[MintBurnTxs](srv.apiKey).makeRequest("GET", u.String(), nil)
}
