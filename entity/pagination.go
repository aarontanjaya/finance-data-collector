package entity

type Pagination struct {
	Start     int
	Length    int
	SortField string
	SortOrder string
	KeyWord   string
}

type PasardanaPagination struct {
	Start     int    `json:"pageBegin"`
	Length    int    `json:"pageLength"`
	SortField string `json:"sortField"`
	SortOrder string `json:"sortOrder"`
	KeyWord   string `json:"Keywords"`
}
