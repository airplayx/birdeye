package birdeye

import (
	"fmt"
	"net/url"
)

type TopTraders struct {
	Success bool `json:"success"`
	Data    struct {
		Items []struct {
			TokenAddress string        `json:"tokenAddress"`
			Owner        string        `json:"owner"`
			Tags         []interface{} `json:"tags"`
			Type         string        `json:"type"`
			Volume       float64       `json:"volume"`
			Trade        int           `json:"trade"`
			TradeBuy     int           `json:"tradeBuy"`
			TradeSell    int           `json:"tradeSell"`
			VolumeBuy    float64       `json:"volumeBuy"`
			VolumeSell   float64       `json:"volumeSell"`
		} `json:"items"`
	} `json:"data"`
}

func (srv *Client) GetTopTraders(address, timeFrame, sortType, sortBy string, offset, limit int) (*TopTraders, error) {
	urlStr := fmt.Sprintf("%s/v2/tokens/top_traders", baseURL)
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("address", address)
	query.Add("time_frame", timeFrame)
	query.Add("sort_type", sortType)
	query.Add("sort_by", sortBy)
	query.Add("offset", fmt.Sprint(offset))
	query.Add("limit", fmt.Sprint(limit))
	u.RawQuery = query.Encode()
	return newCli[TopTraders](srv.apiKey).makeRequest("GET", u.String(), nil)
}
