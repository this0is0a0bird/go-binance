package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAllCurrencyService get all currency
type GetAllCurrencyService struct {
	c *Client
}

type T struct {
	Storage           string `json:"storage"`
	Trading           bool   `json:"trading"`
	WithdrawAllEnable bool   `json:"withdrawAllEnable"`
	Withdrawing       string `json:"withdrawing"`
}

type GetAllCurrencyServiceItem struct {
	Coin              string                             `json:"coin"`
	DepositAllEnable  bool                             `json:"depositAllEnable"`
	Free              string                             `json:"free"`
	Freeze            string                             `json:"freeze"`
	Ipoing            string                             `json:"ipoing"`
	IsLegalMoney      bool                             `json:"is_legal_money"`
	Locked            string                             `json:"locked"`
	Name              string                             `json:"name"`
	Storage           string                             `json:"storage"`
	Trading           bool                               `json:"trading"`
	WithdrawAllEnable bool                               `json:"withdrawAllEnable"`
	Withdrawing       string                             `json:"withdrawing"`
	NetworkList       []GetAllCurrencyServiceNetworkItem `json:"networkList"`
}

type GetAllCurrencyServiceNetworkItem struct {
	AddressRegex            string `json:"addressRegex"`
	Coin                    string `json:"coin"`
	DepositDesc             string `json:"depositDesc"`
	DepositEnable           bool   `json:"depositEnable"`
	IsDefault               bool   `json:"isDefault"`
	MemoRegex               string `json:"memoRegex"`
	MinConfirm              int    `json:"minConfirm"`
	Name                    string `json:"name"`
	Network                 string `json:"network"`
	ResetAddressStatus      bool   `json:"resetAddressStatus"`
	SpecialTips             string `json:"specialTips"`
	UnLockConfirm           int    `json:"unLockConfirm"`
	WithdrawDesc            string `json:"withdrawDesc"`
	WithdrawEnable          bool   `json:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
	WithdrawMax             string `json:"withdrawMax"`
	WithdrawMin             string `json:"withdrawMin"`
	SameAddress             bool   `json:"sameAddress"`
}

// Do send request
func (s *GetAllCurrencyService) Do(ctx context.Context, opts ...RequestOption) (res []*GetAllCurrencyServiceItem, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/config/getall",
		secType:  secTypeSigned,
	}
	response := make([]*GetAllCurrencyServiceItem, 0)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return response, err
	}
	rawMessages := make([]*json.RawMessage, 0)
	err = json.Unmarshal(data, &rawMessages)
	if err != nil {
		return response, err
	}
	for _, j := range rawMessages {
		o := new(GetAllCurrencyServiceItem)
		if err := json.Unmarshal(*j, o); err != nil {
			return response, err
		}
		response = append(response, o)
	}
	return response, nil
}
