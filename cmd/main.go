package main

import (
	//"fmt"
	"log"
	"net/http"


	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	//"github.com/motikingo/resturant-api/entity"

	handler "github.com/motikingo/resturant-api/httpHandler"
	menurepository "github.com/motikingo/resturant-api/menu/repository"
	menuService "github.com/motikingo/resturant-api/menu/service"

	CommentRepository "github.com/motikingo/resturant-api/comment/repository"
	CommentService "github.com/motikingo/resturant-api/comment/service"
	OrderRespository "github.com/motikingo/resturant-api/order/repository"
	OrderService "github.com/motikingo/resturant-api/order/service"
	UserReposirory "github.com/motikingo/resturant-api/user/repository"
	UserService "github.com/motikingo/resturant-api/user/service"
	"github.com/motikingo/resturant-api/middleware"
	"github.com/motikingo/resturant-api/db"


)

var db *gorm.DB

func init(){
	db = database.Connect()
	if db != nil{
		database.Migrate(db)
	}
}

func main(){
	
	r := mux.NewRouter()

	session := handler.NewSessionHandler()

	middleWareHan := middleware.NewMiddlewareHandler(session)
	userRepo:= UserReposirory.NewUserGormRepository(db)
	userSrv := UserService.NewUserGormService(userRepo)
	userHan:= handler.NewUserHandler(userSrv,session)

	r.HandleFunc("/users", middleWareHan.Authenticate(userHan.GetUsers)).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}",middleWareHan.Authenticate(userHan.GetUser)).Methods("GET")
	r.HandleFunc("/signUp",userHan.CreateUser).Methods("POST")
	r.HandleFunc("/Login",userHan.Login).Methods("POST")
	r.HandleFunc("/Logout",middleWareHan.Authenticate(userHan.Logout)).Methods("GET")
	r.HandleFunc("/my_orders",middleWareHan.Authenticate(userHan.MyOrders)).Methods("GET")
	r.HandleFunc("/my_comments",middleWareHan.Authenticate(userHan.MyComments)).Methods("GET")
	r.HandleFunc("/update/user/{id:[0-9]+}",middleWareHan.Authenticate(userHan.ChangeUserPassword)).Methods("PUT")
	r.HandleFunc("/delete/user/{id:[0-9]+}",middleWareHan.Authenticate(userHan.DeleteUser)).Methods("DELETE")	

	adminRepo:= UserReposirory.NewUserGormRepository(db)
	adminSrv := UserService.NewUserGormService(adminRepo)
	adminHan:= handler.NewAdminHandler(adminSrv,session)

	// r.HandleFunc("/admin",adminHan.GetUsers).Methods("GET")
	// r.HandleFunc("/admin/{id:[0-9]+}",adminHan.GetUser).Methods("GET")
	r.HandleFunc("admin/signUp",adminHan.CreateAdmin).Methods("POST")
	r.HandleFunc("admin/Login",adminHan.Login).Methods("POST")
	r.HandleFunc("admin/Logout", middleWareHan.Authenticate(adminHan.Logout)).Methods("GET")
	r.HandleFunc("/Change_password/admin/{id:[0-9]+}",middleWareHan.Authenticate(adminHan.ChangeAdminPassword)).Methods("PUT")
	r.HandleFunc("/delete/admin/{id:[0-9]+}",middleWareHan.Authenticate(adminHan.DeleteAdmin)).Methods("DELETE")	
	log.Fatal(http.ListenAndServe(":80",r))

	cataRepo:= menurepository.NewCatagoryGormRepository(db)
	catSrvc := menuService.NewCatagoryGormService(cataRepo)
	cataHandler := handler.NewCatagoryHandler(catSrvc,session)

	r.HandleFunc("/",middleWareHan.Authenticate(cataHandler.GetCatagories)).Methods("GET")
	r.HandleFunc("/catagory/{id:[0-9]+}",middleWareHan.Authenticate(cataHandler.GetCatagory)).Methods("GET")
	r.HandleFunc("/create/catagory/",middleWareHan.Authenticate(cataHandler.CreateCatagory)).Methods("POST")
	r.HandleFunc("/update/catagory/{id:[0-9]+}",middleWareHan.Authenticate(cataHandler.UpdateCatagory)).Methods("PUT")
	r.HandleFunc("/delete/catagory/{id:[0-9]+}",middleWareHan.Authenticate(cataHandler.DeleteCatagory)).Methods("DELETE")

	commRepo := CommentRepository.NewCommentRepo(db)
	commSrv := CommentService.NewCommentService(commRepo)
	commHandler := handler.NewCommentHandler(commSrv,session)

	r.HandleFunc("/comments",middleWareHan.Authenticate(commHandler.GetComments)).Methods("GET")
	r.HandleFunc("/comments/{id:[0-9]+}",middleWareHan.Authenticate(commHandler.GetComment)).Methods("GET")
	r.HandleFunc("/create/comment/",middleWareHan.Authenticate(commHandler.CreateComment)).Methods("POST")
	r.HandleFunc("/update/comment/{id:[0-9]+}",middleWareHan.Authenticate(commHandler.UpdateComment)).Methods("PUT")
	r.HandleFunc("/delete/comment/{id:[0-9]+}",middleWareHan.Authenticate(commHandler.DeleteComment)).Methods("DELETE")

	ordRepo:= OrderRespository.NewOrderGormRespository(db)
	ordService:=OrderService.NewOrderGormService(ordRepo)
	ordHandler:= handler.NewOrderHandler(ordService,session)

	//r.HandleFunc("/orders",middleWareHan.Authenticate(ordHandler.)).Methods("GET")
	r.HandleFunc("/orders/{id:[0-9]+}",middleWareHan.Authenticate(ordHandler.GetOrder)).Methods("GET")
	r.HandleFunc("/create/order/",middleWareHan.Authenticate(ordHandler.CreateOrder)).Methods("POST")
	r.HandleFunc("/update/order/{id:[0-9]+}",middleWareHan.Authenticate(ordHandler.UpdateOrder)).Methods("PUT")
	r.HandleFunc("/delete/order/{id:[0-9]+}",middleWareHan.Authenticate(ordHandler.DeleteOrder)).Methods("DELETE")

	itemRepo := menurepository.NewItemRepository(db)
	itemSrv := menuService.NewItemGormService(itemRepo)
	itemHandler:=handler.NewItemHandler(itemSrv,session)

	r.HandleFunc("/Items",middleWareHan.Authenticate(itemHandler.GetItems)).Methods("GET")
	r.HandleFunc("/Items/{id:[0-9]+}",middleWareHan.Authenticate(itemHandler.GetItem)).Methods("GET")
	r.HandleFunc("/create/Item/",middleWareHan.Authenticate(itemHandler.CreateItem)).Methods("POST")
	r.HandleFunc("/update/Items/{id:[0-9]+}",middleWareHan.Authenticate(itemHandler.UpdateItem)).Methods("PUT")
	r.HandleFunc("/detelet/Items/{id:[0-9]+}",middleWareHan.Authenticate(itemHandler.DeleteItem)).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":80",r))

}

