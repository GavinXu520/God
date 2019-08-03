package dao

import (
	"God/core/common"
	"God/core/module"

	"github.com/jinzhu/gorm"
)

type UserDao struct {
}

func (self *UserDao) HasAccountByMobile(tx *gorm.DB, mobileNo string) (bool, error) {
	var account module.UserAccount
	var num int
	err := tx.Model(&account).Select("id").Where("mobile = ?", mobileNo).Count(&num).Error
	if nil != err {
		return false, err
	}
	if num != 0 {
		return true, nil
	}
	return false, nil
}

func (self *UserDao) QueryAccountByMobile(tx *gorm.DB, mobileNo string) (*module.UserAccount, error) {

	return nil, nil
}

func (self *UserDao) AddAccount(tx *gorm.DB, account *module.UserAccount) (int, error) {
	if err := tx.Create(account).Error; nil != err {
		return 0, err
	}
	return account.ID, nil
}

func (self *UserDao) AddUserBase(tx *gorm.DB, user *module.UserBase) (int, error) {
	if err := tx.Create(user).Error; nil != err {
		return 0, err
	}
	return user.ID, nil
}

func (self *UserDao) AddRegisterInfo(tx *gorm.DB, register *module.UserRegisterInfo) (int, error) {
	if err := tx.Create(register).Error; nil != err {
		return 0, err
	}
	return register.ID, nil
}

func (self *UserDao) GetUserBase(id uint32) (*module.UserBase, error) {

	user := &module.UserBase{}
	err := common.DB.Model(user).Where("id = ? ", id).Find(user).Error
	if nil != err {
		return nil, err
	}
	return user, err
}
