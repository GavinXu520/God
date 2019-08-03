package entity

type ReqHeader struct {
	Terminalid int    `json:"terminalid"`
	Devicecode string `json:"devicecode"`
	Version    string `json:"version"`
}

type RegisterReq struct {
	Sign string    `json:"sign"`
	Data *register `json:"data"`
}

type register struct {
	Timestamp     int    `json:"timestamp"`
	MobileNo      string `json:"mobileNo"`
	LoginPassword string `json:"loginPassword"`
	TradePassword string `json:"tradePassword"` // pay pwd
}

type LoginReq struct {
	Sign string `json:"sign"`
	Data *login `json:"data"`
}

type login struct {
	Timestamp     int    `json:"timestamp"`
	MobileNo      string `json:"mobileNo"`
	LoginPassword string `json:"loginPassword"`
}
