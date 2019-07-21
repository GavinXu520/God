package service

import (
	"God/core/dao"
	"God/core/entity"
)

type UserService struct {
}

var userDao = &dao.UserDao{}

func (self *UserService) GetUserInfo(id uint32) (*entity.UserInfo, error) {
	return userDao.GetUserInfo(id)
}
