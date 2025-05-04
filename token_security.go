package birdeye

import "fmt"

type TokenSecurity struct {
	Data struct {
		CreatorAddress                 string  `json:"creatorAddress"`
		CreatorOwnerAddress            any     `json:"creatorOwnerAddress"`
		OwnerAddress                   any     `json:"ownerAddress"`
		OwnerOfOwnerAddress            any     `json:"ownerOfOwnerAddress"`
		CreationTx                     string  `json:"creationTx"`
		CreationTime                   int     `json:"creationTime"`
		CreationSlot                   int     `json:"creationSlot"`
		MintTx                         string  `json:"mintTx"`
		MintTime                       int     `json:"mintTime"`
		MintSlot                       int     `json:"mintSlot"`
		CreatorBalance                 float64 `json:"creatorBalance"`
		OwnerBalance                   any     `json:"ownerBalance"`
		OwnerPercentage                any     `json:"ownerPercentage"`
		CreatorPercentage              float64 `json:"creatorPercentage"`
		MetaplexUpdateAuthority        string  `json:"metaplexUpdateAuthority"`
		MetaplexOwnerUpdateAuthority   any     `json:"metaplexOwnerUpdateAuthority"`
		MetaplexUpdateAuthorityBalance float64 `json:"metaplexUpdateAuthorityBalance"`
		MetaplexUpdateAuthorityPercent float64 `json:"metaplexUpdateAuthorityPercent"`
		MutableMetadata                bool    `json:"mutableMetadata"`
		Top10HolderBalance             float64 `json:"top10HolderBalance"`
		Top10HolderPercent             float64 `json:"top10HolderPercent"`
		Top10UserBalance               float64 `json:"top10UserBalance"`
		Top10UserPercent               float64 `json:"top10UserPercent"`
		IsTrueToken                    any     `json:"isTrueToken"`
		FakeToken                      any     `json:"fakeToken"`
		TotalSupply                    float64 `json:"totalSupply"`
		PreMarketHolder                []any   `json:"preMarketHolder"`
		LockInfo                       any     `json:"lockInfo"`
		Freezeable                     any     `json:"freezeable"`
		FreezeAuthority                any     `json:"freezeAuthority"`
		TransferFeeEnable              any     `json:"transferFeeEnable"`
		TransferFeeData                any     `json:"transferFeeData"`
		IsToken2022                    bool    `json:"isToken2022"`
		NonTransferable                any     `json:"nonTransferable"`
		JupStrictList                  bool    `json:"jupStrictList"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (srv *Client) GetSecurity(address string) (*TokenSecurity, error) {
	url := fmt.Sprintf("%s/token_security?address=%s", baseURL, address)
	return newCli[TokenSecurity](srv.apiKey).makeRequest("GET", url, nil)
}
