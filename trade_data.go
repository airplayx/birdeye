package birdeye

import (
	"fmt"
	"strings"
)

type TradeData struct {
	Address                      string  `json:"address"`
	Holder                       int     `json:"holder"`
	Market                       int     `json:"market"`
	LastTradeUnixTime            int64   `json:"last_trade_unix_time"`
	LastTradeHumanTime           string  `json:"last_trade_human_time"`
	Price                        float64 `json:"price"`
	History30MPrice              float64 `json:"history_30m_price"`
	PriceChange30MPercent        float64 `json:"price_change_30m_percent"`
	History1HPrice               float64 `json:"history_1h_price"`
	PriceChange1HPercent         float64 `json:"price_change_1h_percent"`
	History2HPrice               float64 `json:"history_2h_price"`
	PriceChange2HPercent         float64 `json:"price_change_2h_percent"`
	History4HPrice               float64 `json:"history_4h_price"`
	PriceChange4HPercent         float64 `json:"price_change_4h_percent"`
	History6HPrice               float64 `json:"history_6h_price"`
	PriceChange6HPercent         float64 `json:"price_change_6h_percent"`
	History8HPrice               float64 `json:"history_8h_price"`
	PriceChange8HPercent         float64 `json:"price_change_8h_percent"`
	History12HPrice              float64 `json:"history_12h_price"`
	PriceChange12HPercent        float64 `json:"price_change_12h_percent"`
	History24HPrice              float64 `json:"history_24h_price"`
	PriceChange24HPercent        float64 `json:"price_change_24h_percent"`
	UniqueWallet30M              int     `json:"unique_wallet_30m"`
	UniqueWalletHistory30M       int     `json:"unique_wallet_history_30m"`
	UniqueWallet30MChangePercent any     `json:"unique_wallet_30m_change_percent"`
	UniqueWallet1H               int     `json:"unique_wallet_1h"`
	UniqueWalletHistory1H        int     `json:"unique_wallet_history_1h"`
	UniqueWallet1HChangePercent  any     `json:"unique_wallet_1h_change_percent"`
	UniqueWallet2H               int     `json:"unique_wallet_2h"`
	UniqueWalletHistory2H        int     `json:"unique_wallet_history_2h"`
	UniqueWallet2HChangePercent  any     `json:"unique_wallet_2h_change_percent"`
	UniqueWallet4H               int     `json:"unique_wallet_4h"`
	UniqueWalletHistory4H        int     `json:"unique_wallet_history_4h"`
	UniqueWallet4HChangePercent  any     `json:"unique_wallet_4h_change_percent"`
	UniqueWallet8H               int     `json:"unique_wallet_8h"`
	UniqueWalletHistory8H        int     `json:"unique_wallet_history_8h"`
	UniqueWallet8HChangePercent  any     `json:"unique_wallet_8h_change_percent"`
	UniqueWallet24H              int     `json:"unique_wallet_24h"`
	UniqueWalletHistory24H       int     `json:"unique_wallet_history_24h"`
	UniqueWallet24HChangePercent any     `json:"unique_wallet_24h_change_percent"`
	Trade30M                     int     `json:"trade_30m"`
	TradeHistory30M              int     `json:"trade_history_30m"`
	Trade30MChangePercent        float64 `json:"trade_30m_change_percent"`
	Sell30M                      int     `json:"sell_30m"`
	SellHistory30M               int     `json:"sell_history_30m"`
	Sell30MChangePercent         float64 `json:"sell_30m_change_percent"`
	Buy30M                       int     `json:"buy_30m"`
	BuyHistory30M                int     `json:"buy_history_30m"`
	Buy30MChangePercent          float64 `json:"buy_30m_change_percent"`
	Volume30M                    float64 `json:"volume_30m"`
	Volume30MUsd                 float64 `json:"volume_30m_usd"`
	VolumeHistory30M             float64 `json:"volume_history_30m"`
	VolumeHistory30MUsd          float64 `json:"volume_history_30m_usd"`
	Volume30MChangePercent       float64 `json:"volume_30m_change_percent"`
	VolumeBuy30M                 float64 `json:"volume_buy_30m"`
	VolumeBuy30MUsd              float64 `json:"volume_buy_30m_usd"`
	VolumeBuyHistory30M          float64 `json:"volume_buy_history_30m"`
	VolumeBuyHistory30MUsd       float64 `json:"volume_buy_history_30m_usd"`
	VolumeBuy30MChangePercent    float64 `json:"volume_buy_30m_change_percent"`
	VolumeSell30M                float64 `json:"volume_sell_30m"`
	VolumeSell30MUsd             float64 `json:"volume_sell_30m_usd"`
	VolumeSellHistory30M         float64 `json:"volume_sell_history_30m"`
	VolumeSellHistory30MUsd      float64 `json:"volume_sell_history_30m_usd"`
	VolumeSell30MChangePercent   float64 `json:"volume_sell_30m_change_percent"`
	Trade1H                      int     `json:"trade_1h"`
	TradeHistory1H               int     `json:"trade_history_1h"`
	Trade1HChangePercent         float64 `json:"trade_1h_change_percent"`
	Sell1H                       int     `json:"sell_1h"`
	SellHistory1H                int     `json:"sell_history_1h"`
	Sell1HChangePercent          float64 `json:"sell_1h_change_percent"`
	Buy1H                        int     `json:"buy_1h"`
	BuyHistory1H                 int     `json:"buy_history_1h"`
	Buy1HChangePercent           float64 `json:"buy_1h_change_percent"`
	Volume1H                     float64 `json:"volume_1h"`
	Volume1HUsd                  float64 `json:"volume_1h_usd"`
	VolumeHistory1H              float64 `json:"volume_history_1h"`
	VolumeHistory1HUsd           float64 `json:"volume_history_1h_usd"`
	Volume1HChangePercent        float64 `json:"volume_1h_change_percent"`
	VolumeBuy1H                  float64 `json:"volume_buy_1h"`
	VolumeBuy1HUsd               float64 `json:"volume_buy_1h_usd"`
	VolumeBuyHistory1H           float64 `json:"volume_buy_history_1h"`
	VolumeBuyHistory1HUsd        float64 `json:"volume_buy_history_1h_usd"`
	VolumeBuy1HChangePercent     float64 `json:"volume_buy_1h_change_percent"`
	VolumeSell1H                 float64 `json:"volume_sell_1h"`
	VolumeSell1HUsd              float64 `json:"volume_sell_1h_usd"`
	VolumeSellHistory1H          float64 `json:"volume_sell_history_1h"`
	VolumeSellHistory1HUsd       float64 `json:"volume_sell_history_1h_usd"`
	VolumeSell1HChangePercent    float64 `json:"volume_sell_1h_change_percent"`
	Trade2H                      int     `json:"trade_2h"`
	TradeHistory2H               int     `json:"trade_history_2h"`
	Trade2HChangePercent         float64 `json:"trade_2h_change_percent"`
	Sell2H                       int     `json:"sell_2h"`
	SellHistory2H                int     `json:"sell_history_2h"`
	Sell2HChangePercent          float64 `json:"sell_2h_change_percent"`
	Buy2H                        int     `json:"buy_2h"`
	BuyHistory2H                 int     `json:"buy_history_2h"`
	Buy2HChangePercent           float64 `json:"buy_2h_change_percent"`
	Volume2H                     float64 `json:"volume_2h"`
	Volume2HUsd                  float64 `json:"volume_2h_usd"`
	VolumeHistory2H              float64 `json:"volume_history_2h"`
	VolumeHistory2HUsd           float64 `json:"volume_history_2h_usd"`
	Volume2HChangePercent        float64 `json:"volume_2h_change_percent"`
	VolumeBuy2H                  float64 `json:"volume_buy_2h"`
	VolumeBuy2HUsd               float64 `json:"volume_buy_2h_usd"`
	VolumeBuyHistory2H           float64 `json:"volume_buy_history_2h"`
	VolumeBuyHistory2HUsd        float64 `json:"volume_buy_history_2h_usd"`
	VolumeBuy2HChangePercent     float64 `json:"volume_buy_2h_change_percent"`
	VolumeSell2H                 float64 `json:"volume_sell_2h"`
	VolumeSell2HUsd              float64 `json:"volume_sell_2h_usd"`
	VolumeSellHistory2H          float64 `json:"volume_sell_history_2h"`
	VolumeSellHistory2HUsd       float64 `json:"volume_sell_history_2h_usd"`
	VolumeSell2HChangePercent    float64 `json:"volume_sell_2h_change_percent"`
	Trade4H                      int     `json:"trade_4h"`
	TradeHistory4H               int     `json:"trade_history_4h"`
	Trade4HChangePercent         float64 `json:"trade_4h_change_percent"`
	Sell4H                       int     `json:"sell_4h"`
	SellHistory4H                int     `json:"sell_history_4h"`
	Sell4HChangePercent          float64 `json:"sell_4h_change_percent"`
	Buy4H                        int     `json:"buy_4h"`
	BuyHistory4H                 int     `json:"buy_history_4h"`
	Buy4HChangePercent           float64 `json:"buy_4h_change_percent"`
	Volume4H                     float64 `json:"volume_4h"`
	Volume4HUsd                  float64 `json:"volume_4h_usd"`
	VolumeHistory4H              float64 `json:"volume_history_4h"`
	VolumeHistory4HUsd           float64 `json:"volume_history_4h_usd"`
	Volume4HChangePercent        float64 `json:"volume_4h_change_percent"`
	VolumeBuy4H                  float64 `json:"volume_buy_4h"`
	VolumeBuy4HUsd               float64 `json:"volume_buy_4h_usd"`
	VolumeBuyHistory4H           float64 `json:"volume_buy_history_4h"`
	VolumeBuyHistory4HUsd        float64 `json:"volume_buy_history_4h_usd"`
	VolumeBuy4HChangePercent     float64 `json:"volume_buy_4h_change_percent"`
	VolumeSell4H                 float64 `json:"volume_sell_4h"`
	VolumeSell4HUsd              float64 `json:"volume_sell_4h_usd"`
	VolumeSellHistory4H          float64 `json:"volume_sell_history_4h"`
	VolumeSellHistory4HUsd       float64 `json:"volume_sell_history_4h_usd"`
	VolumeSell4HChangePercent    float64 `json:"volume_sell_4h_change_percent"`
	Trade8H                      int     `json:"trade_8h"`
	TradeHistory8H               int     `json:"trade_history_8h"`
	Trade8HChangePercent         float64 `json:"trade_8h_change_percent"`
	Sell8H                       int     `json:"sell_8h"`
	SellHistory8H                int     `json:"sell_history_8h"`
	Sell8HChangePercent          float64 `json:"sell_8h_change_percent"`
	Buy8H                        int     `json:"buy_8h"`
	BuyHistory8H                 int     `json:"buy_history_8h"`
	Buy8HChangePercent           float64 `json:"buy_8h_change_percent"`
	Volume8H                     float64 `json:"volume_8h"`
	Volume8HUsd                  float64 `json:"volume_8h_usd"`
	VolumeHistory8H              float64 `json:"volume_history_8h"`
	VolumeHistory8HUsd           float64 `json:"volume_history_8h_usd"`
	Volume8HChangePercent        float64 `json:"volume_8h_change_percent"`
	VolumeBuy8H                  float64 `json:"volume_buy_8h"`
	VolumeBuy8HUsd               float64 `json:"volume_buy_8h_usd"`
	VolumeBuyHistory8H           float64 `json:"volume_buy_history_8h"`
	VolumeBuyHistory8HUsd        float64 `json:"volume_buy_history_8h_usd"`
	VolumeBuy8HChangePercent     float64 `json:"volume_buy_8h_change_percent"`
	VolumeSell8H                 float64 `json:"volume_sell_8h"`
	VolumeSell8HUsd              float64 `json:"volume_sell_8h_usd"`
	VolumeSellHistory8H          float64 `json:"volume_sell_history_8h"`
	VolumeSellHistory8HUsd       float64 `json:"volume_sell_history_8h_usd"`
	VolumeSell8HChangePercent    float64 `json:"volume_sell_8h_change_percent"`
	Trade24H                     int     `json:"trade_24h"`
	TradeHistory24H              int     `json:"trade_history_24h"`
	Trade24HChangePercent        float64 `json:"trade_24h_change_percent"`
	Sell24H                      int     `json:"sell_24h"`
	SellHistory24H               int     `json:"sell_history_24h"`
	Sell24HChangePercent         float64 `json:"sell_24h_change_percent"`
	Buy24H                       int     `json:"buy_24h"`
	BuyHistory24H                int     `json:"buy_history_24h"`
	Buy24HChangePercent          float64 `json:"buy_24h_change_percent"`
	Volume24H                    float64 `json:"volume_24h"`
	Volume24HUsd                 float64 `json:"volume_24h_usd"`
	VolumeHistory24H             float64 `json:"volume_history_24h"`
	VolumeHistory24HUsd          float64 `json:"volume_history_24h_usd"`
	Volume24HChangePercent       float64 `json:"volume_24h_change_percent"`
	VolumeBuy24H                 float64 `json:"volume_buy_24h"`
	VolumeBuy24HUsd              float64 `json:"volume_buy_24h_usd"`
	VolumeBuyHistory24H          float64 `json:"volume_buy_history_24h"`
	VolumeBuyHistory24HUsd       float64 `json:"volume_buy_history_24h_usd"`
	VolumeBuy24HChangePercent    float64 `json:"volume_buy_24h_change_percent"`
	VolumeSell24H                float64 `json:"volume_sell_24h"`
	VolumeSell24HUsd             float64 `json:"volume_sell_24h_usd"`
	VolumeSellHistory24H         float64 `json:"volume_sell_history_24h"`
	VolumeSellHistory24HUsd      float64 `json:"volume_sell_history_24h_usd"`
	VolumeSell24HChangePercent   float64 `json:"volume_sell_24h_change_percent"`
}

// TradeDataSingle represents the response structure for trade data queries
type TradeDataSingle struct {
	Data    TradeData `json:"data"`
	Success bool      `json:"success"`
}

// GetTradeDataSingle retrieves trade history for a token with pagination support
func (srv *Client) GetTradeDataSingle(address string) (*TradeDataSingle, error) {
	url := fmt.Sprintf("%s/v3/token/trade-data/single?address=%s", baseURL, address)
	return newCli[TradeDataSingle](srv.apiKey).makeRequest("GET", url, nil)
}

type TradeDataMultiple struct {
	Data    map[string]TradeData `json:"data"`
	Success bool                 `json:"success"`
}

func (srv *Client) GetTradeDataMultiple(address ...string) (*TradeDataMultiple, error) {
	url := fmt.Sprintf("%s/v3/token/trade-data/multiple?list_address=%s", baseURL, strings.Join(address, ","))
	return newCli[TradeDataMultiple](srv.apiKey).makeRequest("GET", url, nil)
}
