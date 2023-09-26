package entity

type DataPoint struct {
	Date  uint64
	Value float64
	Rate  float64
}

type FundBenchmark struct {
	Id   int
	Name string
	Data []DataPoint
}
