package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var err error
var db *gorm.DB
func main(){
	err = godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
		return
	}

	dialect:= os.Getenv("dialect")
	user := os.Getenv("user")
	host := os.Getenv("host")
	dbport := os.Getenv("dbport")
	dbname := os.Getenv("dbname")
	passord := os.Getenv("password")

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s port =%s sslmode=disable password=%s" ,host,user,dbname,dbport,passord)

	db,err = gorm.Open(dialect,dbURL)

	if err != nil{
		log.Fatal(err)
		return
	}
	defer db.Close()
	
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080",r))
}