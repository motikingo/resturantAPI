package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/joho/godotenv"
	"github.com/motikingo/resturant-api/entity"
	handler "github.com/motikingo/resturant-api/httpHandler"
	menurepository "github.com/motikingo/resturant-api/menu/repository"
	menuService "github.com/motikingo/resturant-api/menu/service"

	CommentRepository "github.com/motikingo/resturant-api/comment/repository"
	CommentService "github.com/motikingo/resturant-api/comment/service"
	OrderRespository "github.com/motikingo/resturant-api/order/repository"
	OrderService "github.com/motikingo/resturant-api/order/service"
	UserReposirory "github.com/motikingo/resturant-api/user/repository"
	UserService "github.com/motikingo/resturant-api/user/service"
)

var err error
var db *gorm.DB
func main(){
	err = godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
		return
	}

	dialect := os.Getenv("dialect")
	user := os.Getenv("user")
	host := os.Getenv("host")
	//dbport := os.Getenv("dbport")
	dbname := os.Getenv("dbname")
	passord := os.Getenv("password")

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s  sslmode=disable password=%s" ,host,user,dbname,passord)

	//db,err = gorm.Open("postgres","user = postgres password = moti dbname = resturant sslmode=disable")


	db,err = gorm.Open(dialect,dbURL)

	if err != nil{
		fmt.Println("opening db error")
		log.Fatal(err)
		return
	}
	defer db.Close()
	db.Debug().AutoMigrate(&entity.Catagory{})
	db.Debug().AutoMigrate(&entity.Comment{})
	db.Debug().AutoMigrate(&entity.User{})
	db.Debug().AutoMigrate(&entity.Ingridient{})
	db.Debug().AutoMigrate(&entity.Item{})
	db.Debug().AutoMigrate(&entity.Order{})
	db.Debug().AutoMigrate(&entity.Role{})


	r := mux.NewRouter()


	cataRepo:= menurepository.NewCatagoryGormRepository(db)
	catSrvc := menuService.NewCatagoryGormService(cataRepo)
	cataHandler := handler.NewCatagoryHandler(catSrvc)

	r.HandleFunc("/",cataHandler.GetCatagories).Methods("GET")
	r.HandleFunc("/catagory/{id:[0-9]+}",cataHandler.GetCatagory).Methods("GET")
	r.HandleFunc("/create/catagory/",cataHandler.CreateCatagory).Methods("POST")
	r.HandleFunc("/update/catagory/{id:[0-9]+}",cataHandler.UpdateCatagory).Methods("PUT")
	r.HandleFunc("/delete/catagory/{id:[0-9]+}",cataHandler.DeleteCatagory).Methods("DELETE")


	commRepo := CommentRepository.NewCommentRepo(db)
	commSrv := CommentService.NewCommentService(commRepo)
	commHandler := handler.NewCommentHandler(commSrv)

	r.HandleFunc("/comments",commHandler.GetComments).Methods("GET")
	r.HandleFunc("/comments/{id:[0-9]+}",commHandler.GetComment).Methods("GET")
	r.HandleFunc("/create/comment/",commHandler.CreateComment).Methods("POST")
	r.HandleFunc("/update/comment/{id:[0-9]+}",commHandler.UpdateComment).Methods("PUT")
	r.HandleFunc("/delete/comment/{id:[0-9]+}",commHandler.DeleteComment).Methods("DELETE")

	ordRepo:= OrderRespository.NewOrderGormRespository(db)
	ordService:=OrderService.NewOrderGormService(ordRepo)
	ordHandler:= handler.NewOrderHandler(ordService)

	r.HandleFunc("/user/orders",ordHandler.GetOrders).Methods("GET")
	r.HandleFunc("/user/orders/{id:[0-9]+}",ordHandler.GetOrder).Methods("GET")
	r.HandleFunc("/create/user/order/",ordHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/update/user/order/{id:[0-9]+}",ordHandler.UpdateOrder).Methods("PUT")
	r.HandleFunc("/delete/user/order/{id:[0-9]+}",ordHandler.DeleteOrder).Methods("DELETE")

	itemRepo := menurepository.NewItemRepository(db)
	itemSrv := menuService.NewItemGormService(itemRepo)
	itemHandler:=handler.NewItemHandler(itemSrv)

	r.HandleFunc("/Items",itemHandler.GetItems).Methods("GET")
	r.HandleFunc("/Items/{id:[0-9]+}",itemHandler.GetItem ).Methods("GET")
	r.HandleFunc("/create/Item/",itemHandler.CreateItem ).Methods("POST")
	r.HandleFunc("/update/Items/{id:[0-9]+}",itemHandler.UpdateItem ).Methods("PUT")
	r.HandleFunc("/detelet/Items/{id:[0-9]+}",itemHandler.DeleteItem ).Methods("DELETE")
	
	userRepo:= UserReposirory.NewUserGormRepository(db)
	userSrv := UserService.NewUserGormService(userRepo)
	userHan:= handler.NewUserHandler(userSrv)

	r.HandleFunc("/users",userHan.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}",userHan.GetUser).Methods("GET")
	r.HandleFunc("/create/user/",userHan.CreateUser).Methods("POST")
	r.HandleFunc("/update/user/{id:[0-9]+}",userHan.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete/user/{id:[0-9]+}",userHan.DeleteUser).Methods("DELETE")	
	log.Fatal(http.ListenAndServe(":80",r))

}

