package dao

import (
	"gitee.com/molonglove/goboot/gorm"
	"go-demo/app/models/entity"
	"go-demo/core"
)

var User = NewUserDao()

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (u *UserDao) GetUserByUserName(userName string) (user entity.User, err error) {
	err = core.DB.
		Builder().
		Select().
		From("sys_user").
		Where(gorm.Eq("user_name", userName)).
		QExecute(&user).
		Error
	return
}
