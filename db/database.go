package db

import (
	"github.com/jinzhu/gorm"
	"github.com/motikingo/resturant-api/entity"
)

func Migrate(db *gorm.DB){
	db.Debug().AutoMigrate(&entity.Catagory{})
	db.Debug().AutoMigrate(&entity.Comment{})
	db.Debug().AutoMigrate(&entity.User{})
	db.Debug().AutoMigrate(&entity.Ingredient{})
	db.Debug().AutoMigrate(&entity.Item{})
	db.Debug().AutoMigrate(&entity.Order{})
}