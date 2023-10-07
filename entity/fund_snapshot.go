package entity

type FundSnapshot struct {
	Source        int
	Id            int
	Name          string
	IMName        string
	Type          string //Protected, Money Market, Fixed Income, etc
	TypeID        int
	IsIndex       bool //Index, conventional etc
	IsETF         bool
	IsSharia      bool
	NAV           float64
	DailyRR       float64
	MonthlyRR     float64
	YTDRR         float64
	YearlyRR      float64
	AUM           float64
	AUMLastUpdate uint64
	Units         float64
	LastUpdate    uint64
}
