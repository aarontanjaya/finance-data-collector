package dto

type PasardanaDataPoint struct {
	Date  string  `json:"Date"`
	Value float64 `json:"Value"`
	Rate  float64 `json:"DailyReturn"`
}

type PasardanaFundBenchmark struct {
	Id   int                  `json:"Id"`
	Name string               `json:"Name"`
	Data []PasardanaDataPoint `json:"Data"`
}
