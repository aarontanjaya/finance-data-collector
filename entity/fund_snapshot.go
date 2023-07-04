package entity

import "time"

type FundSnapshot struct {
	Name          string
	IMName        string
	Type          string
	TypeID        int
	Category      string
	CategoryID    int
	NAV           float64
	DailyRR       float64
	MonthlyRR     float64
	YTDRR         float64
	YearlyRR      float64
	AUM           float64
	AUMLastUpdate time.Time
	Units         float64
	LastUpdate    time.Time
}
