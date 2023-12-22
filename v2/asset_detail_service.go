package binance

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// GetAssetDetailService fetches all asset detail.
//
// See https://binance-docs.github.io/apidocs/spot/en/#asset-detail-user_data
type GetAssetDetailService struct {
	c     *Client
	asset *string
}

// Asset sets the asset parameter.
func (s *GetAssetDetailService) Asset(asset string) *GetAssetDetailService {
	s.asset = &asset
	return s
}

// Do sends the request.
func (s *GetAssetDetailService) Do(ctx context.Context) (res map[string]AssetDetail, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/assetDetail",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make(map[string]AssetDetail)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// AssetDetail represents the detail of an asset
type AssetDetail struct {
	MinWithdrawAmount string `json:"minWithdrawAmount"`
	DepositStatus     bool   `json:"depositStatus"`
	WithdrawFee       string `json:"withdrawFee"`
	WithdrawStatus    bool   `json:"withdrawStatus"`
	DepositTip        string `json:"depositTip"`
}



type GetFundingAssetService struct {
	c     *Client
	asset *string
	needBtcValuation *string
	recvWindow *int64
	timestamp *int64
}

type GetFundingAsset struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	BtcValuation string `json:"btcValuation"`
}

func (s *GetFundingAssetService) Asset(asset string) *GetFundingAssetService{
		s.asset = &asset
		return s
}

func (s *GetFundingAssetService) NeedBtcValuation(needBtcValuation string) *GetFundingAssetService{
	s.needBtcValuation = &needBtcValuation
	return s
}

func (s *GetFundingAssetService) RecvWindow(recvWindow int64) *GetFundingAssetService{
	s.recvWindow = &recvWindow
	return s
}

func (s *GetFundingAssetService) Timestamp(timestamp int64) *GetFundingAssetService{
	s.timestamp = &timestamp
	return s
}

func (s *GetFundingAssetService) Do(ctx context.Context) (res []*GetFundingAsset, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/asset/get-funding-asset",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.needBtcValuation != nil{
		r.setParam("needBtcValuation", *s.needBtcValuation)
	}
	if s.recvWindow != nil{
		r.setParam("needBtcValuation", *s.recvWindow)
	}
	if s.timestamp != nil{
		r.setParam("timestamp", *s.timestamp)
	}else{
		timestamp := time.Now().UnixNano()/1e6
		r.setParam("timestamp", timestamp)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make([]*GetFundingAsset, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}