package command

type OpenAccountCommand struct {
	AccountHolder  string  `json:"accountHolder"`
	AccountType    int     `json:"accountType"`
	OpeningBalance float64 `json:"openingBalance"`
}

type DepositFundCommand struct {
	ID     string  `json:"ID"`
	Amount float64 `json:"amount"`
}

type WithDrawFundCommand struct {
	ID     string  `json:"ID"`
	Amount float64 `json:"amount"`
}

type CloseAccountCommand struct {
	ID string `json:"ID"`
}
