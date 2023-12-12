package entities

type Finance struct {
	ID          string  `json:"id"`
	Nominal     float64 `json:"nominal"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Account     int     `json:"account"`
}

var Finances = []Finance{
	{"1", 10000, "Selotip", "2023-12-11", 201},
	{"2", 15000, "Penggaris", "2023-12-11", 201},
	{"3", 5000, "Penghapus", "2023-12-11", 201},
}
