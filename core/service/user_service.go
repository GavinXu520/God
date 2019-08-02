package service

import (
	"God/core/common"
	"God/core/dao"
	"God/core/entity"
	"God/core/module"
)

type UserService struct {
}

var userDao = &dao.UserDao{}

func (self *UserService) Register(header *entity.ReqHeader, req *entity.RegisterReq) error {

	tx := common.DB.Begin()
	if nil != tx.Error {
		return tx.Error
	}
	defer tx.Rollback()

	userId, err := userDao.AddUserInfo(tx, &module.UserInfo{})
	if nil != err {
		common.Logger.Error("Failed to Register: Create userInfo is failed, err", err)
		return err
	}
	register := &module.UserRegisterInfo{
		UserID:     userId,
		TerminalID: header.Terminalid,
		Devicecode: header.Devicecode,
		Version:    header.Version,
	}
	registerId, err := userDao.AddRegisterInfo(tx, register)
	if nil != err {
		common.Logger.Error("Failed to Register: Create registerInfo is failed, err", err)
		return err
	}

	_ = registerId
	tx.Commit()
	return nil
}

func (self *UserService) GetUserInfo(id uint32) (*module.UserInfo, error) {
	return userDao.GetUserInfo(id)
}
