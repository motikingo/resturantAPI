package user

import (
	"github.com/motikingo/resturant-api/entity"
)

type UserRepo interface {
	Users() ([]entity.User, []error)
	GetUserByID(id uint) (*entity.User, []error)
	GetUserByEmail(email string) *entity.User
	GetUserByEmailAndID(id uint, email string) *entity.User
	UpdateUser(user entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	CreateUser(user entity.User) (*entity.User, []error)
}
