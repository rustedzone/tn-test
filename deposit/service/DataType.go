package depositservice

type Deposit struct {
	DepositID     int64   `json:"deposit_id" map:"deposit_id"`
	AccountNumber string  `json:"account_number" map:"account_number"`
	Date          string  `json:"date" map:"date"`
	Deposit       float64 `json:"deposit" map:"deposit"`
	Sum           float64 `json:"sum" map:"sum"`
}

type AccountDeposit struct {
	AccountNumber string  `json:"account_number" map:"account_number"`
	Sum           float64 `json:"sum" map:"sum"`
	Detail        []struct {
		DepositID int64   `json:"deposit_id" map:"deposit_id"`
		Deposit   float64 `json:"deposit" map:"deposit"`
	} `json:"detail" map:"detail"`
}
