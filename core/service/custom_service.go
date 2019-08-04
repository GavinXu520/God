package service

import (
	"God/core/entity"
	"God/utils"
	"God/utils/comutil"
	"time"

	"github.com/spf13/viper"
)

type CustomerService struct{}

func (self *CustomerService) Loan(header *entity.ReqHeader, req *entity.LoanReq) (*entity.LoanResp, error) {

	// todo  need finished

	tokenLimit := viper.GetInt("common.tokenDuration")
	sessionLimit := viper.GetInt("common.sessionDuration")
	// generate the token and sessionId
	util.ExpireKV(comutil.TOKEN, header.Token, tokenLimit)
	util.ExpireKV(comutil.SESSION, header.SessionId, sessionLimit)

	return &entity.LoanResp{
		TimeStamp:   int(time.Now().Unix()),
		AccountId:   0,
		Token:       header.Token,
		SessionId:   header.SessionId,
		ForwardPage: 0,
	}, nil
}
