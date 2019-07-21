package dao

import (
	"God/core/common"
	"God/core/entity"
)

type UserDao struct {
}

func (self *UserDao) GetUserInfo(id uint32) (*entity.UserInfo, error) {

	user := &entity.UserInfo{}
	err := common.DB.Model(user).Where("id = ? ", id).Error
	if nil != err {
		return nil, err
	}
	return user, err
}
