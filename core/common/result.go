package common

type Result struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

type OK_Result struct {
	StaStatus int         `json:"status"`
	Data      interface{} `json:"data"`
}

type Failed_Result struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
