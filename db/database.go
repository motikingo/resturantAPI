package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/motikingo/resturant-api/entity"
)

func Connect() *gorm.DB{

	err := godotenv.Load("../.env")
	if err != nil{
		log.Fatal(err)
		return nil
	}

	dialect := os.Getenv("dialect")
	user := os.Getenv("user")
	host := os.Getenv("host")
	//dbport := os.Getenv("dbport")
	dbname := os.Getenv("dbname")
	passord := os.Getenv("password")
	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s  sslmode=disable password=%s" ,host,user,dbname,passord)

	//db,err = gorm.Open("postgres","user = postgres password = moti dbname = resturant sslmode=disable")


	db,err := gorm.Open(dialect,dbURL)

	if err != nil||  db == nil{
		fmt.Println("opening db error")
		return nil
	}
	
	return db
}

func Migrate(db *gorm.DB){
	db.Debug().AutoMigrate(&entity.Catagory{})
	db.Debug().AutoMigrate(&entity.Comment{})
	db.Debug().AutoMigrate(&entity.User{})
	db.Debug().AutoMigrate(&entity.Ingredient{})
	db.Debug().AutoMigrate(&entity.Item{})
	db.Debug().AutoMigrate(&entity.Order{})
}


