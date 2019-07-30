package dao

import (
	"God/core/common"
	"God/core/module"
)

type UserDao struct {
}

func (self *UserDao) GetUserInfo(id uint32) (*module.UserInfo, error) {

	user := &module.UserInfo{}
	err := common.DB.Model(user).Where("id = ? ", id).Error
	if nil != err {
		return nil, err
	}
	return user, err
}
