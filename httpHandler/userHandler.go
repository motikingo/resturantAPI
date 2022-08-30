package handler

import (
	"encoding/json"
	"fmt"
	"log"

	//"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/helper"
	"github.com/motikingo/resturant-api/order"
	"github.com/motikingo/resturant-api/user"
	"github.com/motikingo/resturant-api/comment"

)

type UserHandler struct{
	userSrvc user.UserService
	orderserv Order.OrderService
	commSrv comment.CommentService
	session *SessionHandler
}


func NewUserHandler(userSrvc user.UserService,session *SessionHandler)UserHandler{
	return UserHandler{userSrvc:userSrvc,session:session}
}

func(usrHan *UserHandler)GetUsers(w http.ResponseWriter,r *http.Request){
	session := usrHan.session.GetSession(r)
	if session == nil||session.Role != "Admin"{
		w.Write([]byte("Uauthorized user"))
		return
	}
	
	//w.Header().Set("Content-Type","application/json")
	users,err:=usrHan.userSrvc.Users()
	if err!=nil{
		log.Fatal(err)
	}
	usr,er:=json.MarshalIndent(users,"","/t/t")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(usr)
}

func(usrHan *UserHandler)GetUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := usrHan.session.GetSession(r)
	if session == nil{
		w.Write([]byte("Uauthorized user"))
		return
	}
	id:=mux.Vars(r)["id"]
	ids,e:=strconv.Atoi(id)
	if e!=nil{
		log.Fatal(err)
	}

	user,err:=usrHan.userSrvc.GetUserByID(uint(ids))
	if err!=nil{
		log.Fatal(err)
	}
	urs,er:=json.MarshalIndent(user,"","/t/t")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(urs)
}

func(usrHan *UserHandler)CreateUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	
	if r.Method == "POST"{
		response := &struct{
			Status string
			User *entity.User
		}{
			Status: "Invalid Input",
		}
		Input:=&struct{
			Email string `json:"email"`
			Password string 	`json:"password"`
			Confirm_password string `json:"confirm_password"`
		}{}
		var user entity.User
		
		read,e:=ioutil.ReadAll(r.Body)

		if e!=nil{
			log.Fatal(err)
		}
		
		
		if er:= json.Unmarshal(read,&Input); er !=nil || Input.Email == "" || Input.Password == "" || Input.Confirm_password==""{
			
			w.Write(helper.MarshalResponse(response))
			return
		}
		
		if len(Input.Password)<2 {
			response.Status = "password lenght must be more than 8 character"
			w.Write(helper.MarshalResponse(response))
			return
		}
		if Input.Password != Input.Confirm_password {
			response.Status = "confirm password is not the same"
			w.Write(helper.MarshalResponse(response))
			return
			
		}

		if usrHan.userSrvc.GetUserByEmail(Input.Email) != nil {
			response.Status ="this Email is already exist"
			w.Write(helper.MarshalResponse(response))
			return
			
		}
		Input.Password = helper.HashPassword(Input.Password)
		if Input.Password == ""{
			response.Status ="Internal server error"
			w.Write(helper.MarshalResponse(response))
			return
		}
		
		user = entity.User{Email:Input.Email,Password:Input.Password}

		fmt.Println(user.Email,user.Password)

		urs,err:=usrHan.userSrvc.CreateUser(user)
		fmt.Println(urs)

		if err!=nil|| urs == nil{
			fmt.Println("here I am")
			response.Status ="Internal server error"
			w.Write(helper.MarshalResponse(response))
			return
		}
		
		response.Status = "successfully created"
		response.User = urs
		w.Write(helper.MarshalResponse(response))
			
	}
	w.Write([]byte("server Error"))

}
func(usrHan *UserHandler)ChangeUserPassword(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := usrHan.session.GetSession(r)
	if sess ==nil {
		w.Write([]byte("UnAuthorized user"))
		return
	}

	var oldPassword,newPassword string
	var err []error
	userId,_ := strconv.Atoi(sess.UserID)
	user,err := usrHan.userSrvc.GetUserByID(uint(userId))

	if len(err)>0{
		w.Write([]byte("no user with given Id"))
		return
	}
	if !helper.MatchPassword(user.Password,oldPassword){
		w.Write([]byte("wrong password"))
		usrHan.session.DeleteSession(w)
		return
	}
	if len(newPassword)<8{
		w.Write([]byte("short password length"))
		return
	}
	
	read,er:= ioutil.ReadAll(r.Body)
	if er!=nil{
		log.Fatal(err)
	}
	er = json.Unmarshal(read,&user)
	if er!=nil{
		log.Fatal(err)
	}
	id,er := strconv.Atoi(sess.UserID)
	user.ID = uint(id)
	if er !=nil{
		return
	}
	userUpdated,errs:=usrHan.userSrvc.UpdateUser(*user)

	if errs!=nil{
		log.Fatal(err)
	}
	userMash,errr:=json.MarshalIndent(userUpdated,"","/r/r")

	if errr!=nil{
		log.Fatal(err)
	}

	w.Write(userMash)

}

func(usrHan *UserHandler)Login(w http.ResponseWriter,r *http.Request){

	response := &struct{
		Success bool
		Message string
		UserId uint
	}{
		Success: false,
		Message: "login faild",
	}
	input :=&struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}{}

	read,_:= ioutil.ReadAll(r.Body)

	
	if er := json.Unmarshal(read,&input); er != nil{
		w.Write(helper.MarshalResponse(response))
		return
	}

	if len(input.Password)<4 || input.Password ==""{
		response.Message= "invalud input password"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if input.Email == ""{
		response.Message= "invalud input email"
		w.Write(helper.MarshalResponse(response))
		return
	}

	user := usrHan.userSrvc.GetUserByEmail(input.Email)
	if user == nil{
		response.Message = "this email is not registered"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if !helper.MatchPassword(user.Password,input.Password) {
		response.Message = "Incorrect password"
		w.Write(helper.MarshalResponse(response))
		return
		
	}
	sess := &entity.Session{
		UserID: strconv.Itoa(int(user.ID)),
		Email: user.Email,		
	}

	if !usrHan.session.CreateSession(sess,w){
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Message = "Login successful"
	response.Success = true
	response.UserId = user.ID
	w.Write(helper.MarshalResponse(response))
}

func(usrHan *UserHandler)Logout(w http.ResponseWriter,r *http.Request){
	sess := usrHan.session.GetSession(r)
	response := &struct{
		Success bool
		Message string
		UserId string
	}{
		Success: false,
		Message: "Session Error...",
	}
	if sess == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}
	if usrHan.session.DeleteSession(w)!= nil{
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Message = "Logout successful"
	response.Success = true
	response.UserId = sess.UserID
	w.Write(helper.MarshalResponse(response))

}
func(usrHan *UserHandler)MyOrders(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := usrHan.session.GetSession(r)
	response := &struct{
		Status string
		Orders []entity.Order
	}{
		Status: "UnAthorized User",
	}
	if sess == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}
	orders,ers := usrHan.orderserv.Orders()
	if len(orders)<1 || len(ers)>0{
		response.Status = "No order found"
		w.Write(helper.MarshalResponse(response))
		return
	}
	id,_ := strconv.Atoi(sess.UserID) 
	for _,order := range orders{
		if order.UserID == uint(id){
			response.Orders = append(response.Orders, order)
		}
	}
	if len(response.Orders)<1{
		response.Status = "No order found"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Status = "Orders retrieved"
	w.Write(helper.MarshalResponse(response))

}

func(usrHan *UserHandler)MyComments(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := usrHan.session.GetSession(r)
	response := &struct{
		Status string
		Comments []entity.Comment
	}{
		Status: "UnAthorized User",
	}
	if sess == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}
	comments,ers := usrHan.commSrv.Comments()
	if len(comments)<1 || len(ers)>0{
		response.Status = "No comment found"
		w.Write(helper.MarshalResponse(response))
		return
	}
	id,_ := strconv.Atoi(sess.UserID) 
	for _,comment := range comments{
		if comment.UserID == uint(id){
			response.Comments = append(response.Comments, comment)
		}
	}
	if len(response.Comments)<1{
		response.Status = "No comment found"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Status = "comment retrieved"
	w.Write(helper.MarshalResponse(response))

}


func(usrHan *UserHandler)DeleteUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := usrHan.session.GetSession(r)

	respose:= &struct{
		Status bool
		Message string
		UseId  uint
	}{
		Status:false,
		Message: "Unauthprized user",
	}
	
	if sess == nil{
		w.Write(helper.MarshalResponse(respose))
		return
	}
	id,_ := strconv.Atoi(sess.UserID) 
	user,err:=usrHan.userSrvc.GetUserByID(uint(id))
	if err!=nil || user == nil{
		w.Write(helper.MarshalResponse(respose))
		log.Fatal(err)
		return
	}

	if sess.Role == "Admin" {
		id,_ := strconv.Atoi(r.FormValue("user_id"))
		user,err =usrHan.userSrvc.DeleteUser(uint(id))
		if len(err)>0{
			respose.Message= "no such user"
			w.Write(helper.MarshalResponse(respose))
			log.Fatal(err)
			return
		}
		respose.Message = "Admin Account Delete successful"
		respose.Status = true
		respose.UseId =user.ID
		
	}else{
		id,_ := strconv.Atoi(sess.UserID)
		user,err = usrHan.userSrvc.DeleteUser(uint(id))
		if len(err)>0{
			log.Fatal(err)
		}
		usrHan.session.DeleteSession(w)
		respose.Message = " Account Delete successful"
		respose.Status = true
		respose.UseId =  user.ID
	}
	w.Write(helper.MarshalResponse(respose))
}