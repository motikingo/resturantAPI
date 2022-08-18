package entity

import (
	//"image"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
)

const(
	SessionName = "resturant"
)

type User struct{
	gorm.Model
	Email string  	`json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string `json:"password"  gorm:"type:varchar(255);not null"`
	Role  string
}

type Comment struct{
	gorm.Model
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	UserID uint			`json:"user_id"`

}


type Item struct{
	gorm.Model
	Name string 			`json:"name"  	 gorm:"type:varchar(255);not null"`
	Price float64 		  	`json:"price" 	 gorm:"type:varchar(255);not null"`
	Description string 	 	`json:"description" gorm:type:"varchar(255);not null"`	
	Image string 			`json:"image" gorm:"type:varchar(255);not null"`	
	Number int 				`json:"number" gorm:"type:varchar(255);not null"`
	Catagories []string 	`json:"catagories" gorm:"type:varchar(255);not null"`
	Ingridients []string 	`json:"ingridients" gorm:"type:varchar(255);not null"`

}

type Order struct{
	gorm.Model
	PlaceAt time.Time 	`json :"placeAt" gorm:"type:time;not null"`
	ItemID uint 		`json :"itemsid"`
	CatagoryID uint 	`json :"catagoryid"`
	UserID uint 		`json :"userid"`
	Number int 			`json : "count"`
	Orderbill float64 	`json: "order_bill"`
}

type Catagory struct{
	gorm.Model
	Name string		 `json:"name" "gorm:type:varchar(255);not null"`
	ImageUrl string  `json:"image" "gorm:type:varchar(255);not null"`
	Items []string 	 `json:"items" gorm:"type:varchar(255);"`

}

type Ingredient struct{
	gorm.Model
	Name string 		`json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null"`

}

// type Role struct{
// 	gorm.Model
// 	Name string `json:"name" gorm:"type:varchar(255);not null"`

// }

type Session struct {
	UserID string
	Email string
	Role  string
	jwt.RegisteredClaims
}


