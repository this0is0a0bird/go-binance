package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// FapiGetAccountService get account info
type FapiGetAccountService struct {
	c *Client
}

// Do send request
func (s * FapiGetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Account define account info
type FapiAccount struct {
	MakerCommission  int64     `json:"makerCommission"`
	TakerCommission  int64     `json:"takerCommission"`
	BuyerCommission  int64     `json:"buyerCommission"`
	SellerCommission int64     `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	UpdateTime       uint64    `json:"updateTime"`
	AccountType      string    `json:"accountType"`
	Balances         []Balance `json:"balances"`
	Permissions      []string  `json:"permissions"`
}

// Balance define user balance of your account
type FapiBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// GetAccountSnapshotService all account orders; active, canceled, or filled
type  FapiGetAccountSnapshotService struct {
	c           *Client
	accountType string
	startTime   *int64
	endTime     *int64
	limit       *int
}

// Type set account type ("SPOT", "MARGIN", "FUTURES")
func (s * FapiGetAccountSnapshotService) Type(accountType string) *FapiGetAccountSnapshotService {
	s.accountType = accountType
	return s
}

// StartTime set starttime
func (s * FapiGetAccountSnapshotService) StartTime(startTime int64) *FapiGetAccountSnapshotService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s * FapiGetAccountSnapshotService) EndTime(endTime int64) * FapiGetAccountSnapshotService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s * FapiGetAccountSnapshotService) Limit(limit int) * FapiGetAccountSnapshotService {
	s.limit = &limit
	return s
}

// Do send request
func (s * FapiGetAccountSnapshotService) Do(ctx context.Context, opts ...RequestOption) (res *Snapshot, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/accountSnapshot",
		secType:  secTypeSigned,
	}
	r.setParam("type", s.accountType)

	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Snapshot{}, err
	}
	res = new(Snapshot)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &Snapshot{}, err
	}
	return res, nil
}

// Snapshot define snapshot
type FapiSnapshot struct {
	Code     int            `json:"code"`
	Msg      string         `json:"msg"`
	Snapshot []*SnapshotVos `json:"snapshotVos"`
}

// SnapshotVos define content of a snapshot
type FapiSnapshotVos struct {
	Data       *SnapshotData `json:"data"`
	Type       string        `json:"type"`
	UpdateTime int64         `json:"updateTime"`
}

// SnapshotData define content of a snapshot
type FapiSnapshotData struct {
	MarginLevel         string `json:"marginLevel"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`

	Balances   []*SnapshotBalances   `json:"balances"`
	UserAssets []*SnapshotUserAssets `json:"userAssets"`
	Assets     []*SnapshotAssets     `json:"assets"`
	Positions  []*SnapshotPositions  `json:"position"`
}

// SnapshotBalances define snapshot balances
type FapiSnapshotBalances struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// SnapshotUserAssets define snapshot user assets
type FapiSnapshotUserAssets struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// SnapshotAssets define snapshot assets
type FapiSnapshotAssets struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	WalletBalance string `json:"walletBalance"`
}

// SnapshotPositions define snapshot positions
type FapiSnapshotPositions struct {
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
}

//Account Status  /sapi/v1/account/status
type FapiGetAccountStatusService struct {
	c *Client
}

func (s *FapiGetAccountStatusService) Do(ctx context.Context, opts ...RequestOption) (res *AccountStatus, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/account/status",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &AccountStatus{}, err
	}
	res = new(AccountStatus)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &AccountStatus{}, err
	}
	return res, nil
}

type FapiAccountStatus struct {
	Data string `json:"data"`
}

//Account Api Trade Status  /
type FapiGetAccountApiTradeStatusService struct {
	c *Client
}

func (s *FapiGetAccountApiTradeStatusService) Do(ctx context.Context, opts ...RequestOption) (res *AccountApiTradeStatusResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/account/apiTradingStatus",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &AccountApiTradeStatusResponse{}, err
	}
	res = new(AccountApiTradeStatusResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &AccountApiTradeStatusResponse{}, err
	}
	return res, nil
}

type FapiAccountApiTradeStatusResponse struct {
	Data AccountApiTradeStatus `json:"data"`
}

type FapiAccountApiTradeStatus struct {
	IsLocked           bool `json:"isLocked"`
	PlannedRecoverTime int  `json:"plannedRecoverTime"`
	TriggerCondition   struct {
		GCR  int `json:"GCR"`
		IFER int `json:"IFER"`
		UFR  int `json:"UFR"`
	} `json:"triggerCondition"`
	UpdateTime int64 `json:"updateTime"`
}

// get api enable perms
type FapiGetAccountApiPermListService struct {
	c *Client
}

func (s *FapiGetAccountApiPermListService) Do(ctx context.Context, opts ...RequestOption) (res *GetAccountApiPermResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/account/apiRestrictions",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &GetAccountApiPermResponse{}, err
	}
	res = new(GetAccountApiPermResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &GetAccountApiPermResponse{}, err
	}
	return res, nil
}

type FapiGetAccountApiPermResponse struct {
	IpRestrict                     bool  `json:"ipRestrict"`
	CreateTime                     int64 `json:"createTime"`
	EnableWithdrawals              bool  `json:"enableWithdrawals"`
	EnableInternalTransfer         bool  `json:"enableInternalTransfer"`
	PermitsUniversalTransfer       bool  `json:"permitsUniversalTransfer"`
	EnableVanillaOptions           bool  `json:"enableVanillaOptions"`
	EnableReading                  bool  `json:"enableReading"`
	EnableFutures                  bool  `json:"enableFutures"`
	EnableMargin                   bool  `json:"enableMargin"`
	EnableSpotAndMarginTrading     bool  `json:"enableSpotAndMarginTrading"`
	TradingAuthorityExpirationTime int64 `json:"tradingAuthorityExpirationTime"`
}
