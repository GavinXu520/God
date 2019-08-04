package entity

type LoanResp struct {
	TimeStamp   int    `json:"timestamp"`
	AccountId   int    `json:"accountId"`
	Token       string `json:"token"`
	SessionId   string `json:"sessionId"`
	ForwardPage int    `json:"forwardPage"` // Jump to page what need forward
	Balance     int    `json:"balance"`     // the amount need to loan
}
