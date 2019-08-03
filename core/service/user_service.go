package service

import (
	"God/core/common"
	"God/core/common/comerr"
	"God/core/dao"
	"God/core/entity"
	"God/core/module"

	"github.com/jinzhu/gorm"
)

type UserService struct {
}

var userDao = &dao.UserDao{}

func (self *UserService) Register(header *entity.ReqHeader, req *entity.RegisterReq) (int, error) {

	tx := common.DB.Begin()
	if nil != tx.Error {
		return 0, tx.Error
	}
	defer tx.Rollback()

	hasRegister, err := userDao.HasAccountByMobile(tx, req.MobileNo)
	if nil != err {
		common.Logger.Error("Failed to Register: Check Account has register is failed, err", err)
		return 0, err
	}
	if hasRegister {
		common.Logger.Error("Failed to Register: The Account had been register !!")
		return 0, comerr.BizErrorf("this mobile had been register, mobileNo: %s", req.MobileNo)
	}

	accountId, err := userDao.AddAccount(tx, &module.UserAccount{Mobile: req.MobileNo,
		LoginPwd: req.LoginPassword, PayPwd: req.TradePassword})

	register := &module.UserRegisterInfo{
		UserID:     accountId,
		TerminalID: header.Terminalid,
		Devicecode: header.Devicecode,
		Version:    header.Version,
	}
	registerId, err := userDao.AddRegisterInfo(tx, register)
	if nil != err {
		common.Logger.Error("Failed to Register: Create registerInfo is failed, err", err)
		return 0, err
	}
	tx.Commit()
	common.Logger.Info("Register Success", "accountId", accountId, "registerId", registerId)
	return accountId, nil
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
