package handler

import (
	"encoding/json"
	"go/types"
	//"fmt"
	"log"

	//"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/helper"
	"github.com/motikingo/resturant-api/user"
)

type UserHandler struct{
	userSrvc user.UserService
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

	user,err:=usrHan.userSrvc.User(uint(ids))
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
			status string
			user *entity.User
		}{
			status: "Invalid Input",
		}
		input:=&struct{
			email string
			password string
			confirm_password string
		}{}
		var user entity.User
		
		read,e:=ioutil.ReadAll(r.Body)

		if e!=nil{
			log.Fatal(err)
		}
		
		er:= json.Unmarshal(read,&input)
		if er !=nil || input.email = "" || input.password = "" || input.confirm_password==""{
			w.Write(helper.MarshalResponse(response))
			return
		}

		if len(input.password)<8 {
			response.status = "password lenght must be more than 8 character"
			w.Write(helper.MarshalResponse(response))
			return
		}
		if input.password != input.confirm_password {
			response.status = "confirm password is not the same"
			w.Write(helper.MarshalResponse(response))
			return
			
		}

		if usrHan.userSrvc.GetUserByEmail(input.email) != nil {
			response.status ="this Email is already exist"
			w.Write(helper.MarshalResponse(response))
			return
			
		}
		input.password = helper.HashPassword(input.password)
		if input.password == ""{
			response.status ="Internal server error"
			w.Write(helper.MarshalResponse(response))
			return
		}

		user = entity.User{Email:input.email,Password:input.password}		
		
		urs,err:=usrHan.userSrvc.CreateUser(&user)
		if err!=nil|| urs == nil{
			response.status ="Internal server error"
			w.Write(helper.MarshalResponse(response))
			return
		}
		response.status = "successfully created"
		response.user = urs
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
	var user entity.User
	var oldPassword,newPassword string
	var err []error
	user,err = usrHan.userSrvc.GetUserByID(sess.UserID)

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
	userUpdated,errs:=usrHan.userSrvc.UpdateUser(user)

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
		success bool
		message string
		userId uint
	}{
		success: false,
		message: "login faild",
	}
	input :=&struct{
		email string
		password string
	}{}

	read,er:= ioutil.ReadAll(r.Body)
	er = json.Unmarshal(read,&input)
	if er != nil{
		w.Write(helper.MarshalResponse(response))
		return
	}

	if len(input.password)<4 || input.password ==""{
		response.message= "invalud input password"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if input.email == ""{
		response.message= "invalud input email"
		w.Write(helper.MarshalResponse(response))
		return
	}

	user := usrHan.userSrvc.GetUserByEmail(input.email)
	if user == nil{
		response.message = "this email is not registered"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if !helper.MatchPassword(user.Password,input.password) {
		response.message = "Incorrect password"
		w.Write(helper.MarshalResponse(response))
		return
		
	}
	sess := &entity.Session{
		UserID: string(user.ID),
		Email: user.Email,		
	}

	if !usrHan.session.CreateSession(sess,w){
		response.message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.message = "Login successful"
	response.success = true
	response.userId = user.ID
	w.Write(helper.MarshalResponse(response))
}

func(usrHan *UserHandler)Logout(w http.ResponseWriter,r *http.Request){
	sess := usrHan.session.GetSession(r)
	response := &struct{
		success bool
		message string
		userId string
	}{
		success: false,
		message: "Session Error...",
	}
	if sess == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	if usrHan.session.DeleteSession(w)!= nil{
		response.message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.message = "Logout successful"
	response.success = true
	response.userId = sess.UserID
	w.Write(helper.MarshalResponse(response))

}

func(usrHan *UserHandler)DeleteUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := usrHan.session.GetSession(r)

	respose:= &struct{
		status bool
		message string
		useId  uint
	}{
		status:false,
		message: "Unauthprized user",
	}
	
	if sess == nil{
		w.Write(helper.MarshalResponse(respose))
		return
	}
	user,err:=usrHan.userSrvc.GetUserByID(sess.UserID)
	if err!=nil || user == nil{
		w.Write(helper.MarshalResponse(respose))
		log.Fatal(err)
		return
	}

	if sess.Role == "Admin" {
		id,_ := strconv.Atoi(r.FormValue("user_id"))
		user,err =usrHan.userSrvc.DeleteUser(uint(id))
		if len(err)>0{
			respose.message= "no such user"
			w.Write(helper.MarshalResponse(respose))
			log.Fatal(err)
			return
		}
		respose.message = "Admin Account Delete successful"
		respose.status = true
		respose.useId = uint(id)
		
	}else{
		id,_ := strconv.Atoi(sess.UserID)
		user,err = usrHan.userSrvc.DeleteUser(uint(id))
		if len(err)>0{
			log.Fatal(err)
		}
		usrHan.session.DeleteSession(w)
		respose.message = " Account Delete successful"
		respose.status = true
		respose.useId =  user.ID
	}
	w.Write(helper.MarshalResponse(respose))
}