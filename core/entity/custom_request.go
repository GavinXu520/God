package entity

type LoanReq struct {
	Sign string `json:"sign"`
	Data *loan  `json:"data"`
}

type loan struct {
	Timestamp int `json:"timestamp"`
}
