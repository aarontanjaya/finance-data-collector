package entity

type FundDetail struct {
	Source            int
	Id                int
	Name              string
	Type              int
	Currency          int
	IsSharia          bool
	IsETF             bool
	IsIndex           bool
	HasDividend       bool
	IMId              int
	IMName            string
	IMFee             string
	CustodianBankId   int
	CustodianBankName string
	Benchmarks        []FundBenchmark
	NAVTimeSeries     []DataPoint
	AUM               []DataPoint
	TotalUnits        []DataPoint
}
