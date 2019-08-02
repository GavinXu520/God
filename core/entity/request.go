package entity

//type Req struct {
//	Sign string `json:"sign"`
//	Data
//}

type RegisterReq struct {
	Timestamp     string `json:"timestamp"`
	MobileNo      string `json:"mobileNo"`
	LoginPassword string `json:"loginPassword"`
	TradePassword string `json:"tradePassword"` // 支付密码
}

type ReqHeader struct {
	Terminalid int    `json:"terminalid"`
	Devicecode string `json:"devicecode"`
	Version    string `json:"version"`
}
