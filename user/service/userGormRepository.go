package UserService

import (
	"github.com/motikingo/resturant-api/user"
	"github.com/motikingo/resturant-api/entity"
)

type UserGormRepository struct{
	repo user.UserRepo
}
func NewUserGormRepository(repo user.UserRepo)user.UserRepo{
	return &UserGormRepository{repo:repo}
}


func(userRepo *UserGormRepository) Users()([]entity.User,[]error){

	users,err:= userRepo.repo.Users()
	if len(err)>0{
		return nil,err
	}
	return users,nil

}

func(userRepo *UserGormRepository)User(id uint)(*entity.User,[]error){

	user,err:= userRepo.repo.User(id)
	if len(err)>0{
		return nil,err
	}
	return user,nil

}

func(userRepo *UserGormRepository)UpdateUser(id uint)(*entity.User,[]error){

	user,err := userRepo.repo.UpdateUser(id)
	if len(err)>0{
		return nil,err
	}
	return user,nil
}

func(userRepo *UserGormRepository)DeleteUser(id uint)(*entity.User,[]error){

	
	user,err := userRepo.repo.DeleteUser(id)
	if len(err)>0{
		return nil,err
	}
	return user,nil
}
func(userRepo *UserGormRepository)CreateUser(user *entity.User)(*entity.User,[]error){

	usr := user
	usr,err:= userRepo.repo.CreateUser(user)
	if len(err)>0{
		return nil,err
	}
	return usr,nil
}

