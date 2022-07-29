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
	Username string  `json:"username" gorm:"type:varchar(255);not null " `
	Password string `json:"password"  gorm:"type:varchar(255);not null"`
	Orders []Order  `json:"orders"  	gorm:"one2many:user_orders"`
	Comments []Comment `json:"comments" gorm:"many2many:user_comments"`
	Roles []Role 		`json:"roles"  gorm:"many2many:user_roles"`

}

type Comment struct{
	gorm.Model
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	UserID uint `json:"user_id"`

}


type Item struct{
	gorm.Model
	Name string `json:"name" gorm:type:varchar(255);not null`
	Catagories []Catagory `json:"catagories" gorm:"many2many:item_catagoies"`
	Price float64 		  `json:"price" gorm:type:varchar(255);not null`
	Description string 	 `json:"description" gorm:type:varchar(255);not null`	
	Image string 		`json:"image" `		
	Ingridients []Ingridient 	`json:"ingridient" gorm:many2many:item_ingridients`

}

type Order struct{
	gorm.Model
	PlaceAt time.Time `json :"placeAt" gorm:not null`
	ItemID uint `json :"itemsid"`
	CatagoryID uint `json :"catagoryid"`
	UserID uint `json :"userid"`

}

type Catagory struct{
	gorm.Model
	Name string `json:"name" gorm:type:varchar(255);not null`
	Items []Item `json:"items"`

}

type Ingridient struct{
	gorm.Model
	Name string 	`json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:type:varchar(255);not null`

}

type Role struct{
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);not null"`

}