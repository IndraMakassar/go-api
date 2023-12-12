package entities

type Account struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

var Accounts = []Account{
	{101, "kas", "Aset"},
	{102, "piutang", "Aset"},
	{103, "barang", "Aset"},
}
