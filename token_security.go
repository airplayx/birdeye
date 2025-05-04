package birdeye

import "fmt"

type TokenSecurity struct {
	Data struct {
		CreatorAddress                 string        `json:"creatorAddress"`
		CreatorOwnerAddress            interface{}   `json:"creatorOwnerAddress"`
		OwnerAddress                   interface{}   `json:"ownerAddress"`
		OwnerOfOwnerAddress            interface{}   `json:"ownerOfOwnerAddress"`
		CreationTx                     string        `json:"creationTx"`
		CreationTime                   int           `json:"creationTime"`
		CreationSlot                   int           `json:"creationSlot"`
		MintTx                         string        `json:"mintTx"`
		MintTime                       int           `json:"mintTime"`
		MintSlot                       int           `json:"mintSlot"`
		CreatorBalance                 float64       `json:"creatorBalance"`
		OwnerBalance                   interface{}   `json:"ownerBalance"`
		OwnerPercentage                interface{}   `json:"ownerPercentage"`
		CreatorPercentage              float64       `json:"creatorPercentage"`
		MetaplexUpdateAuthority        string        `json:"metaplexUpdateAuthority"`
		MetaplexOwnerUpdateAuthority   interface{}   `json:"metaplexOwnerUpdateAuthority"`
		MetaplexUpdateAuthorityBalance float64       `json:"metaplexUpdateAuthorityBalance"`
		MetaplexUpdateAuthorityPercent float64       `json:"metaplexUpdateAuthorityPercent"`
		MutableMetadata                bool          `json:"mutableMetadata"`
		Top10HolderBalance             float64       `json:"top10HolderBalance"`
		Top10HolderPercent             float64       `json:"top10HolderPercent"`
		Top10UserBalance               float64       `json:"top10UserBalance"`
		Top10UserPercent               float64       `json:"top10UserPercent"`
		IsTrueToken                    interface{}   `json:"isTrueToken"`
		FakeToken                      interface{}   `json:"fakeToken"`
		TotalSupply                    float64       `json:"totalSupply"`
		PreMarketHolder                []interface{} `json:"preMarketHolder"`
		LockInfo                       interface{}   `json:"lockInfo"`
		Freezeable                     interface{}   `json:"freezeable"`
		FreezeAuthority                interface{}   `json:"freezeAuthority"`
		TransferFeeEnable              interface{}   `json:"transferFeeEnable"`
		TransferFeeData                interface{}   `json:"transferFeeData"`
		IsToken2022                    bool          `json:"isToken2022"`
		NonTransferable                interface{}   `json:"nonTransferable"`
		JupStrictList                  bool          `json:"jupStrictList"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetSecurity(address string) (*TokenSecurity, error) {
	url := fmt.Sprintf("%s/token_security?address=%s", baseURL, address)
	return newCli[TokenSecurity](srv.apiKey).makeRequest("GET", url, nil)
}
