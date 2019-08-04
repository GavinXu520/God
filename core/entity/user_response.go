package entity

type LoginResp struct {
	TimeStamp int    `json:"timestamp"`
	AccountId int    `json:"accountId"`
	Token     string `json:"token"`
	SessionId string `json:"sessionId"`
}
