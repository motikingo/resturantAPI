package user

import (
	"github.com/motikingo/resturant-api/entity"
)

type UserService interface{
	Users()([]entity.User,[]error)
	User(id uint)(*entity.User,[]error)
	UpdateUser(id uint)(*entity.User,[]error)
	DeleteUser(id uint)(*entity.User,[]error)
	CreateUser(user *entity.User)(*entity.User,[]error)
}