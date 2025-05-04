package birdeye

import (
	"fmt"
	"net/url"
	"strings"
)

type Trades struct {
	Data []struct {
		Address        string  `json:"address"`
		TotalVolume    float64 `json:"total_volume"`
		TotalVolumeUsd float64 `json:"total_volume_usd"`
		VolumeBuyUsd   float64 `json:"volume_buy_usd"`
		VolumeSellUsd  float64 `json:"volume_sell_usd"`
		VolumeBuy      float64 `json:"volume_buy"`
		VolumeSell     float64 `json:"volume_sell"`
		TotalTrade     int     `json:"total_trade"`
		Buy            int     `json:"buy"`
		Sell           int     `json:"sell"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetTradesSingle(address, timeFrame string) (*Trades, error) {
	urlStr := fmt.Sprintf("%s/v3/all-time/trades/single", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("address", address)
	query.Add("time_frame", timeFrame)
	u.RawQuery = query.Encode()
	return newCli[Trades](srv.apiKey).makeRequest("GET", u.String(), nil)
}

func (srv *Client) GetTradesMultiple(timeFrame string, address ...string) (*Trades, error) {
	urlStr := fmt.Sprintf("%s/v3/all-time/trades/multiple", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("address", strings.Join(address, ","))
	query.Add("time_frame", timeFrame)
	u.RawQuery = query.Encode()
	return newCli[Trades](srv.apiKey).makeRequest("POST", u.String(), nil)
}
