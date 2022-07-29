package UserReposirory

import (
	"github.com/jinzhu/gorm"
	"github.com/motikingo/resturant-api/user"
	"github.com/motikingo/resturant-api/entity"
)

type UserGormRepository struct{
	db *gorm.DB
}
func NewUserGormRepository(db *gorm.DB)user.UserRepo{
	return &UserGormRepository{db:db}
}


func(userRepo *UserGormRepository) Users()([]entity.User,[]error){
	users:=[]entity.User{}
	err:= userRepo.db.Find(&users).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return users,nil

}

func(userRepo *UserGormRepository)User(id uint)(*entity.User,[]error){
	var user entity.User
	err:= userRepo.db.First(&user,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return &user,nil

}

func(userRepo *UserGormRepository)UpdateUser(id uint,usr entity.User)(*entity.User,[]error){

	user,err:=userRepo.User(id)

	if len(err)>0{
		return nil,err
	}

	user.Username = usr.Username
	user.Password = usr.Password
	user.Orders = usr.Orders
	user.Comments = usr.Comments
	user.Roles = usr.Roles
	
	err = userRepo.db.Save(&user).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return user,nil
}

func(userRepo *UserGormRepository)DeleteUser(id uint)(*entity.User,[]error){

	user,err:=userRepo.User(id)

	if len(err)>0{
		return nil,err
	}
	err = userRepo.db.Delete(&user,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return user,nil
}
func(userRepo *UserGormRepository)CreateUser(user *entity.User)(*entity.User,[]error){

	usr := user
	err:= userRepo.db.Create(&usr).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return usr,nil
}

