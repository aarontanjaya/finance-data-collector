package dto

type PasardanaSnapshotResponse struct {
	Id                   int     `json:"id"`
	BloombergQuote       float64 `json:"BloombergQuote"`
	Name                 string  `json:"Name"`
	IMName               string  `json:"InvestmentManagerName"`
	IMFee                string  `json:"InvestmentManagerFee"`
	CustodianBankName    string  `json:"CustodianBankName"`
	Type                 int     `json:"Type"`
	IsSharia             bool    `json:"Sharia"`
	IsETF                bool    `json:"ExchangeTradedFund"`
	IsIndex              bool    `json:"Index"`
	NAV                  float64 `json:"NetAssetValue"`
	DailyReturn          float64 `json:"DailyReturn"`
	WeeklyReturn         float64 `json:"WeeklyReturn"`
	MonthlyReturn        float64 `json:"MonthlyReturn"`
	QuarterlyReturn      float64 `json:"QuarterlyReturn"`
	YTDReturn            float64 `json:"YtdReturn"`
	YearlyReturn         float64 `json:"YearlyReturn"`
	AUM                  float64 `json:"AssetUnderManagement"`
	TotalUnit            float64 `json:"TotalUnit"`
	AUMLastUpdate        string  `json:"AumLastUpdate"`
	LastUpdate           string  `json:"LastUpdate"`
	ConservativeCategory string  `json:"ConservativeCategory"`
}
