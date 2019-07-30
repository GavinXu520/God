package service

import (
	"God/core/dao"
	"God/core/module"
)

type UserService struct {
}

var userDao = &dao.UserDao{}

func (self *UserService) GetUserInfo(id uint32) (*module.UserInfo, error) {
	return userDao.GetUserInfo(id)
}
