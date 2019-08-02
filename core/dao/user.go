package dao

import (
	"God/core/common"
	"God/core/module"

	"github.com/jinzhu/gorm"
)

type UserDao struct {
}

func (self *UserDao) AddUserInfo(tx *gorm.DB, user *module.UserInfo) (int, error) {
	if err := tx.Create(user).Error; nil == err {
		return 0, err
	}
	return user.ID, nil
}

func (self *UserDao) AddRegisterInfo(tx *gorm.DB, register *module.UserRegisterInfo) (int, error) {
	if err := tx.Create(register).Error; nil == err {
		return 0, err
	}
	return register.ID, nil
}

func (self *UserDao) GetUserInfo(id uint32) (*module.UserInfo, error) {

	user := &module.UserInfo{}
	err := common.DB.Model(user).Where("id = ? ", id).Error
	if nil != err {
		return nil, err
	}
	return user, err
}
