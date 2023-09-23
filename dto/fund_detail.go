package dto

type PasardanaFundDetailResponse struct {
	Id                int                  `json:"id"`
	Name              string               `json:"Name"`
	Type              int                  `json:"Type"`
	Currency          int                  `json:"Currency"`
	IsSharia          bool                 `json:"Sharia"`
	IsETF             bool                 `json:"ExchangeTradedFund"`
	IsIndex           bool                 `json:"Index"`
	HasDividend       bool                 `json:"Dividend"`
	IMId              int                  `json:"InvestmentManagerId"`
	IMName            string               `json:"InvestmentManagerName"`
	IMFee             string               `json:"InvestmentManagerFee"`
	CustodianBankId   int                  `json:"CustodianBankId"`
	CustodianBankName string               `json:"CustodianBankName"`
	NAVTimeSeries     []PasardanaDataPoint `json:"NetAssetValues"`
	Benchmarks        []FundBenchmark      `json:"Benchmarks"`
	AUM               []PasardanaDataPoint `json:"AssetUnderManagements"`
	TotalUnits        []PasardanaDataPoint `json:"TotalUnits"`
}
