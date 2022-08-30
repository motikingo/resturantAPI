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
	database "github.com/motikingo/resturant-api/db"
	"github.com/motikingo/resturant-api/middleware"
	OrderRespository "github.com/motikingo/resturant-api/order/repository"
	OrderService "github.com/motikingo/resturant-api/order/service"
	UserReposirory "github.com/motikingo/resturant-api/user/repository"
	UserService "github.com/motikingo/resturant-api/user/service"
)

var db *gorm.DB

func init(){
	db = database.Connect()
	if db != nil{
		database.Migrate(db)
	}
}

func main(){
	defer db.Close()
	r := mux.NewRouter()

	session := handler.NewSessionHandler()

	middleWareHan := middleware.NewMiddlewareHandler(session)
	userRepo:= UserReposirory.NewUserGormRepository(db)
	userSrv := UserService.NewUserGormService(userRepo)
	userHan:= handler.NewUserHandler(userSrv,session)

	adminHan:= handler.NewAdminHandler(userSrv,session)

	cataRepo:= menurepository.NewCatagoryGormRepository(db)
	catSrvc := menuService.NewCatagoryGormService(cataRepo)
	cataHandler := handler.NewCatagoryHandler(catSrvc,session)

	commRepo := CommentRepository.NewCommentRepo(db)
	commSrv := CommentService.NewCommentService(commRepo)
	commHandler := handler.NewCommentHandler(commSrv,session)

	ordRepo:= OrderRespository.NewOrderGormRespository(db)
	ordService:=OrderService.NewOrderGormService(ordRepo)
	ordHandler:= handler.NewOrderHandler(ordService,session)

	ingrdRepo := menurepository.NewIngredientGormRepository(db)
	ingrdSrv := menuService.NewIngredientGormService(ingrdRepo)

	itemRepo := menurepository.NewItemRepository(db)
	itemSrv := menuService.NewItemGormService(itemRepo)

	ingrdHa := handler.NewIngredientHandler(ingrdSrv,itemSrv,session)
	itemHandler := handler.NewItemHandler(itemSrv,session,ingrdSrv,catSrvc)

	r.HandleFunc("/users", middleWareHan.Authenticate(userHan.GetUsers)).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}",middleWareHan.Authenticate(userHan.GetUser)).Methods("GET")
	r.HandleFunc("/signUp",middleWareHan.sessioShouldNotExist(userHan.CreateUser)).Methods("POST")
	r.HandleFunc("/Login", middleWareHan.sessioShouldNotExist(userHan.Login)).Methods("POST")
	r.HandleFunc("/Logout",middleWareHan.Authenticate(userHan.Logout)).Methods("GET")
	r.HandleFunc("/my_orders",middleWareHan.Authenticate(userHan.MyOrders)).Methods("GET")
	r.HandleFunc("/my_comments",middleWareHan.Authenticate(userHan.MyComments)).Methods("GET")
	r.HandleFunc("/update/user/{id:[0-9]+}",middleWareHan.Authenticate(userHan.ChangeUserPassword)).Methods("PUT")
	r.HandleFunc("/delete/user/{id:[0-9]+}",middleWareHan.Authenticate(userHan.DeleteUser)).Methods("DELETE")	

	// adminRepo:= UserReposirory.NewUserGormRepository(db)
	// adminSrv := UserService.NewUserGormService(adminRepo)

	// r.HandleFunc("/admin",adminHan.GetUsers).Methods("GET")
	// r.HandleFunc("/admin/{id:[0-9]+}",adminHan.GetUser).Methods("GET")
	r.HandleFunc("/admin/signUp",middleWareHan.sessioShouldNotExist(adminHan.CreateAdmin)).Methods("POST")
	r.HandleFunc("/admin/Login",middleWareHan.sessioShouldNotExist(adminHan.Login)).Methods("POST")
	r.HandleFunc("/admin/Logout", middleWareHan.Authenticate(adminHan.Logout)).Methods("GET")
	r.HandleFunc("/Change_password/admin/{id:[0-9]+}",middleWareHan.Authenticate(adminHan.ChangeAdminPassword)).Methods("PUT")
	r.HandleFunc("/delete/admin/{id:[0-9]+}",middleWareHan.Authenticate(adminHan.DeleteAdmin)).Methods("DELETE")	

	
	r.HandleFunc("/catagories",middleWareHan.Authenticate(cataHandler.GetCatagories)).Methods("GET")
	r.HandleFunc("/catagory/{id:[0-9]+}",middleWareHan.Authenticate(cataHandler.GetCatagory)).Methods("GET")
	r.HandleFunc("/create/catagory/",middleWareHan.OnlyAdminAuth(cataHandler.CreateCatagory)).Methods("POST")
	r.HandleFunc("/update/catagory/{id:[0-9]+}",middleWareHan.OnlyAdminAuth(cataHandler.UpdateCatagory)).Methods("PUT")
	r.HandleFunc("/delete/catagory/{id:[0-9]+}",middleWareHan.OnlyAdminAuth(cataHandler.DeleteCatagory)).Methods("DELETE")

	
	r.HandleFunc("/comments",middleWareHan.Authenticate(commHandler.GetComments)).Methods("GET")
	r.HandleFunc("/comments/{id:[0-9]+}",middleWareHan.Authenticate(commHandler.GetComment)).Methods("GET")
	r.HandleFunc("/create/comment/",middleWareHan.Authenticate(commHandler.CreateComment)).Methods("POST")
	r.HandleFunc("/update/comment/{id:[0-9]+}",middleWareHan.Authenticate(commHandler.UpdateComment)).Methods("PUT")
	r.HandleFunc("/delete/comment/{id:[0-9]+}",middleWareHan.Authenticate(commHandler.DeleteComment)).Methods("DELETE")

	
	//r.HandleFunc("/orders",middleWareHan.Authenticate(ordHandler.)).Methods("GET")
	r.HandleFunc("/orders/{id:[0-9]+}",middleWareHan.Authenticate(ordHandler.GetOrder)).Methods("GET")
	r.HandleFunc("/create/order/",middleWareHan.Authenticate(ordHandler.CreateOrder)).Methods("POST")
	r.HandleFunc("/update/order/{id:[0-9]+}",middleWareHan.Authenticate(ordHandler.UpdateOrder)).Methods("PUT")
	r.HandleFunc("/delete/order/{id:[0-9]+}",middleWareHan.Authenticate(ordHandler.DeleteOrder)).Methods("DELETE")


	r.HandleFunc("/Ingredients",middleWareHan.Authenticate(ingrdHa.GetIngredients)).Methods("GET")
	r.HandleFunc("/Ingredient/{id:[0-9]+}",middleWareHan.Authenticate(ingrdHa.GetIngredient)).Methods("GET")
	r.HandleFunc("/create/Ingredient/",middleWareHan.OnlyAdminAuth(ingrdHa.CreateIngredient)).Methods("POST")
	r.HandleFunc("/update/Ingredient/{id:[0-9]+}",middleWareHan.OnlyAdminAuth(ingrdHa.UpdateIngredient)).Methods("PUT")
	r.HandleFunc("/detelet/Ingredients/{id:[0-9]+}",middleWareHan.OnlyAdminAuth(ingrdHa.DeleteIngredient)).Methods("DELETE")


	r.HandleFunc("/Items",middleWareHan.Authenticate(itemHandler.GetItems)).Methods("GET")
	r.HandleFunc("/Items/{id:[0-9]+}",middleWareHan.Authenticate(itemHandler.GetItem)).Methods("GET")
	r.HandleFunc("/create/Item/",middleWareHan.OnlyAdminAuth(itemHandler.CreateItem)).Methods("POST")
	r.HandleFunc("/update/Items/{id:[0-9]+}",middleWareHan.OnlyAdminAuth(itemHandler.UpdateItem)).Methods("PUT")
	r.HandleFunc("/detelet/Items/{id:[0-9]+}",middleWareHan.OnlyAdminAuth(itemHandler.DeleteItem)).Methods("DELETE")
	r.HandleFunc("/AddIngredient/Item/",middleWareHan.OnlyAdminAuth(itemHandler.AddIngredient)).Methods("POST")
	r.HandleFunc("/RemoveIngrediend/Item/",middleWareHan.OnlyAdminAuth(itemHandler.RemoveIngrediend)).Methods("PUT")
	
	log.Fatal(http.ListenAndServe(":80",r))
	
}

