package entity

type PasardanaPagination struct {
	Start     int    `json:"pageBegin"`
	Length    int    `json:"pageLength"`
	SortField string `json:"sortField"`
	SortOrder string `json:"sortOrder"`
	KeyWord   string `json:"Keywords"`
}
