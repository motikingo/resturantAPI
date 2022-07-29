package UserService

import (
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/user"
)

type UserGormService struct{
	repo user.UserRepo
}
func NewUserGormService(repo user.UserRepo)user.UserService{
	return &UserGormService{repo:repo}
}


func(userRepo *UserGormService) Users()([]entity.User,[]error){

	users,err:= userRepo.repo.Users()
	if len(err)>0{
		return nil,err
	}
	return users,nil

}

func(userRepo *UserGormService)User(id uint)(*entity.User,[]error){

	user,err:= userRepo.repo.User(id)
	if len(err)>0{
		return nil,err
	}
	return user,nil

}

func(userRepo *UserGormService)UpdateUser(id uint,usr entity.User)(*entity.User,[]error){

	user,err := userRepo.repo.UpdateUser(id,usr)
	if len(err)>0{
		return nil,err
	}
	return user,nil
}

func(userRepo *UserGormService)DeleteUser(id uint)(*entity.User,[]error){

	
	user,err := userRepo.repo.DeleteUser(id)
	if len(err)>0{
		return nil,err
	}
	return user,nil
}
func(userRepo *UserGormService)CreateUser(user *entity.User)(*entity.User,[]error){

	usr := user
	usr,err:= userRepo.repo.CreateUser(user)
	if len(err)>0{
		return nil,err
	}
	return usr,nil
}

