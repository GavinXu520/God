package service

import (
	"God/core/common"
	"God/core/common/comerr"
	"God/core/dao"
	"God/core/entity"
	"God/core/module"
	util "God/utils"

	"God/utils/comutil"
	"fmt"
	"time"

	"God/core/common/enum"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type UserService struct {
}

var userDao = &dao.UserDao{}

func (self *UserService) Register(header *entity.ReqHeader, req *entity.RegisterReq) (*entity.LoginResp, error) {

	tx := common.DB.Begin()
	if nil != tx.Error {
		return nil, tx.Error
	}
	defer tx.Rollback()

	hasRegister, err := userDao.HasAccountByMobile(tx, req.Data.MobileNo)
	if nil != err {
		common.Logger.Error("Failed to Register: Check Account has register is failed, err", err)
		return nil, err
	}
	if hasRegister {
		common.Logger.Error("Failed to Register: The Account had been register !!")
		return nil, comerr.BizErrorf("this mobile had been register, mobileNo: %s", req.Data.MobileNo)
	}

	// generate a nick
	randomStr := util.RandString(int(6))
	subStr := util.SubString(req.Data.MobileNo, int(8), int(11))
	nick_name := "god" + randomStr + subStr

	account := &module.UserAccount{Mobile: req.Data.MobileNo, Nick: nick_name,
		LoginPwd: req.Data.LoginPassword, PayPwd: req.Data.TradePassword}

	accountId, err := userDao.AddAccount(tx, account)
	if nil != err {
		common.Logger.Error("Failed to Register: Create account is failed, err", err)
		return nil, err
	}

	register := &module.UserRegisterInfo{
		UserID:     accountId,
		TerminalID: header.Terminalid,
		Devicecode: header.Devicecode,
		Version:    header.Version,
	}
	registerId, err := userDao.AddRegisterInfo(tx, register)
	if nil != err {
		common.Logger.Error("Failed to Register: Create registerInfo is failed, err", err)
		return nil, err
	}
	tx.Commit()
	common.Logger.Info("Register Success", "accountId", accountId, "registerId", registerId)

	// default is three days
	tokenLimit := viper.GetInt("common.tokenDuration")
	// default is a hour
	sessionLimit := viper.GetInt("common.sessionDuration")
	// generate the token and sessionId
	token := util.RandSeq(19)
	sessionId := util.RandStr()

	// store the token and sessionId
	util.SetKV(comutil.TOKEN, token, sessionId, tokenLimit)
	util.SetKV(comutil.SESSION, sessionId, fmt.Sprint(accountId), sessionLimit)

	return &entity.LoginResp{
		TimeStamp: int(time.Now().Unix()),
		AccountId: accountId,
		Token:     token,
		SessionId: sessionId,
	}, nil
}

func (self *UserService) LoginByPwd(header *entity.ReqHeader, req *entity.LoginReq) (*entity.LoginResp, error) {
	tx := common.DB.Begin()
	defer tx.Rollback()

	account, err := userDao.GetAccountByMobileAndPwd(req.Data.MobileNo, req.Data.LoginPassword)
	if nil != err {
		common.Logger.Error(fmt.Sprintf("Failed to Login by mobile and loginPwd, mobileNo: %s, err: %v", req.Data.MobileNo, err))
		return nil, err
	}
	if nil == account {
		common.Logger.Warn(fmt.Sprintf("The account is empty by mobile and loginPwd, mobileNo: %s", req.Data.MobileNo))
		return nil, nil
	}

	// add login history
	history := &module.UserLoginHistory{
		UserID:     account.ID,
		LoginType:  enum.Pwd,
		TerminalID: header.Terminalid,
		Devicecode: header.Devicecode,
		Version:    header.Version,
		// cityId ?
	}
	if _, err := userDao.AddUserLoginHistory(tx, history); nil != err {
		common.Logger.Error(fmt.Sprintf("Failed to Login by mobile and loginPwd, add login history is failed, mobile: %s, err: %v", req.Data.MobileNo, err))
		return nil, err
	}

	tokenLimit := viper.GetInt("common.tokenDuration")
	sessionLimit := viper.GetInt("common.sessionDuration")
	// generate the token and sessionId
	token := util.RandSeq(19)
	sessionId := util.RandStr()

	// store the token and sessionId
	util.SetKV(comutil.TOKEN, token, sessionId, tokenLimit)
	util.SetKV(comutil.SESSION, sessionId, fmt.Sprint(account.ID), sessionLimit)

	return &entity.LoginResp{
		TimeStamp: int(time.Now().Unix()),
		AccountId: account.ID,
		Token:     token,
		SessionId: sessionId,
	}, nil
}

func (self *UserService) GetUserBase(id uint32) (*module.UserBase, error) {
	user, err := userDao.GetUserBase(id)
	if nil != err && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if nil != err && err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return user, nil
	}
}
