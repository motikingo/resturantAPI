package UserService

import (
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/user"
)

type UserGormService struct {
	repo user.UserRepo
}

func NewUserGormService(repo user.UserRepo) user.UserService {
	return &UserGormService{repo: repo}
}

func (usersrv *UserGormService) Users() ([]entity.User, []error) {

	users, err := usersrv.repo.Users()
	if len(err) > 0 {
		return nil, err
	}
	return users, nil

}

func (usersrv *UserGormService) GetUserByID(id uint) (*entity.User, []error) {

	user, err := usersrv.repo.GetUserByID(id)
	if len(err) > 0 {
		return nil, err
	}
	return user, nil

}

func (usersrv *UserGormService) GetUserByEmail(email string) *entity.User {
	user := usersrv.repo.GetUserByEmail(email)

	return user
}

func (usersrv *UserGormService) GetUserByEmailAndID(id uint, email string) *entity.User {

	user := usersrv.repo.GetUserByEmailAndID(id, email)
	return user
}

func (usersrv *UserGormService) UpdateUser(usr entity.User) (*entity.User, []error) {

	user, err := usersrv.repo.UpdateUser(usr)
	if len(err) > 0 {
		return nil, err
	}
	return user, nil
}

func (usersrv *UserGormService) DeleteUser(id uint) (*entity.User, []error) {

	user, err := usersrv.repo.DeleteUser(id)
	if len(err) > 0 {
		return nil, err
	}
	return user, nil
}
func (usersrv *UserGormService) CreateUser(user entity.User) (*entity.User, []error) {

	usr, err := usersrv.repo.CreateUser(user)
	if len(err) > 0 {
		return nil, err
	}
	return usr, nil
}
