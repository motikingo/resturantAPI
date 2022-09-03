package UserReposirory

import (
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/user"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) user.UserRepo {
	return &UserGormRepository{db: db}
}

func (userRepo *UserGormRepository) Users() ([]entity.User, []error) {
	users := []entity.User{}
	err := userRepo.db.Find(&users).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return users, nil

}

func (userRepo *UserGormRepository) GetUserByID(id uint) (*entity.User, []error) {
	var user entity.User
	err := userRepo.db.Preload("Orders").Preload("Comments").First(&user, id).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return &user, nil

}

func (userRepo *UserGormRepository) GetUserByEmail(email string) *entity.User {
	var user entity.User
	//err := userRepo.db.First(&user,email).GetErrors()
	err := userRepo.db.Where("email = ?", email).First(&user).GetErrors()
	if len(err) > 0 {
		return nil
	}
	return &user
}
func (userRepo *UserGormRepository) GetUserByEmailAndID(id uint, email string) *entity.User {

	var user entity.User
	err := userRepo.db.First(&user, id, email).GetErrors()
	if len(err) > 0 {
		return nil
	}
	return &user

}

func (userRepo *UserGormRepository) UpdateUser(usr entity.User) (*entity.User, []error) {

	user, err := userRepo.GetUserByID(usr.ID)

	if len(err) > 0 {
		return nil, err
	}

	user.Email = func() string {
		if user.Email != usr.Email {
			return user.Email
		}
		return user.Email
	}()
	user.Password = func() string {
		if user.Password != usr.Password {
			return usr.Password
		}
		return user.Password
	}()

	user.Orders = func() []entity.Order {
		if !reflect.DeepEqual(user.Orders, usr.Orders) {
			userRepo.db.Model(&user).Association("orders").Clear()
			return usr.Orders
		}
		return user.Orders
	}()

	user.Comments = func() []entity.Comment {
		if !reflect.DeepEqual(user.Orders, usr.Orders) {
			userRepo.db.Model(&user).Association("comments").Clear()
			return usr.Comments
		}
		return user.Comments
	}()

	err = userRepo.db.Save(&user).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return user, nil
}

func (userRepo *UserGormRepository) DeleteUser(id uint) (*entity.User, []error) {

	user, err := userRepo.GetUserByID(id)

	if len(err) > 0 {
		return nil, err
	}
	err = userRepo.db.Delete(&user, id).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return user, nil
}
func (userRepo *UserGormRepository) CreateUser(user entity.User) (*entity.User, []error) {
	err := userRepo.db.Create(&user).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return &user, nil
}
