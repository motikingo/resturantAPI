package entity

import (
	//"image"
	"time"

	"github.com/jinzhu/gorm"
)

// import(
// 	""
// )

type User struct{
	gorm.Model
	username string `gorm:"type:varchar(255);not null " `
	password string `gorm:"type:varchar(255);not null"`
	orders []Order  `gorm:"many2many:user_orders"`
	comments []Comment `gorm:"many2many:user_comments"`
	roles []Role 		`gorm:"many2many:user_roles"`

}

type Comment struct{
	gorm.Model
	description string `gorm:"type:varchar(255);not null"`
	userID uint 

}

type Item struct{
	gorm.Model
	name string `gorm:type:varchar(255);not null`
	catagorys []Catagory `gorm:"many2many:item_catagoies"`
	price float64
	description string 
	Image string 			
	ingridients []Ingridient 	`gorm:many2many:item_ingridients`

}

type Order struct{
	gorm.Model
	placeAt time.Time 
	ItemsID uint `json :"itemsid"`
	catagoryID uint `json :"catagoryid"`
	userID uint `json :"userid"`

}

type Catagory struct{
	gorm.Model
	name string
	userID uint

}

type Ingridient struct{
	gorm.Model
	name string 	`gorm:"type:varchar(255);not null"`
	description string `gorm:type:varchar(255);not null`


}

type Role struct{
	gorm.Model
	name string

}