package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/helper"
	"github.com/motikingo/resturant-api/user"
)

type AdminHandler struct{
	usersrv user.UserService
	session *SessionHandler
}

func NewAdminHandler(usersrv user.UserService) *AdminHandler{
	return &AdminHandler{usersrv: usersrv}
}

func(adminHan *AdminHandler)CreateAdmin(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	
	if r.Method == "POST"{
		response := &struct{
			status string
			admin *entity.User
		}{
			status: "Invalid Input",
		}
		input:=&struct{
			email string
			password string
			confirm_password string
		}{}
		
		read,e:=ioutil.ReadAll(r.Body)

		if e!=nil{
			log.Fatal(err)
		}
		
		er:= json.Unmarshal(read,&input)

		if er != nil || input.email == "" || input.password == "" || input.confirm_password==""{
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

		if adminHan.usersrv.GetUserByEmail(input.email) != nil {
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

		admin := &entity.User{
			Email:input.email,
			Password:input.password,
			Role:"Admin",
			}		
		
		adm,err:=adminHan.usersrv.CreateUser(admin)
		if err!=nil|| adm == nil{
			response.status ="Internal server error"
			w.Write(helper.MarshalResponse(response))
			return
		}
		response.status = "successfully created"
		response.admin = adm
		w.Write(helper.MarshalResponse(response))
			
	}
	w.Write([]byte("server Error"))

}

func(adminHan *AdminHandler)ChangeAdminPassword(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := adminHan.session.GetSession(r)
	if sess ==nil || sess.Role != "Admin" {
		w.Write([]byte("UnAuthorized user"))
		return
	}

	var oldPassword,newPassword string

	id,er := strconv.Atoi(sess.UserID)
	ids := uint(id)
	if er !=nil{
		return
	}

	admin,err := adminHan.usersrv.GetUserByID(ids)

	if len(err)>0{
		w.Write([]byte("no user with given Id"))
		return
	}
	if !helper.MatchPassword(admin.Password,oldPassword){
		w.Write([]byte("wrong password"))
		adminHan.session.DeleteSession(w)
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
	er = json.Unmarshal(read,&admin)
	if er!=nil{
		log.Fatal(err)
	}
	
	userUpdated,errs:=adminHan.usersrv.UpdateUser(admin)

	if errs!=nil{
		log.Fatal(err)
	}
	userMash,errr:=json.MarshalIndent(userUpdated,"","/r/r")

	if errr!=nil{
		log.Fatal(err)
	}

	w.Write(userMash)

}

func(adminHan *AdminHandler)Login(w http.ResponseWriter,r *http.Request){

	response := &struct{
		status string
		adminId uint
	}{
		status: "login faild",
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
		response.status= "invalid password input "
		w.Write(helper.MarshalResponse(response))
		return
	}
	if input.email == ""{
		response.status= "invalud input email"
		w.Write(helper.MarshalResponse(response))
		return
	}

	admin := adminHan.usersrv.GetUserByEmail(input.email)
	if admin == nil{
		response.status = "this email is not registered"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if !helper.MatchPassword(admin.Password,input.password) {
		response.status = "Incorrect password"
		w.Write(helper.MarshalResponse(response))
		return
		
	}
	sess := &entity.Session{
		UserID: string(admin.ID),
		Email: admin.Email,	
		Role:"Admin",	
	}

	if !adminHan.session.CreateSession(sess,w){
		response.status = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.status = "Login successful"
	response.adminId = admin.ID
	w.Write(helper.MarshalResponse(response))
}

func(adminHan *AdminHandler)Logout(w http.ResponseWriter,r *http.Request){
	sess := adminHan.session.GetSession(r)
	response := &struct{
		success bool
		message string
		adminId string
	}{
		success: false,
		message: "UnAuthorized User...",
	}
	if sess == nil || sess.Role == "Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	if adminHan.session.DeleteSession(w)!= nil{
		response.message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.message = "Logout successful"
	response.success = true
	response.adminId = sess.UserID
	w.Write(helper.MarshalResponse(response))

}

func(adminHan *AdminHandler)DeleteAdmin(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := adminHan.session.GetSession(r)

	respose:= &struct{
		message string
		AdminId  uint
	}{
	
		message: "Unauthprized user",
	}
	
	if sess == nil || sess.Role != "Admin"{
		w.Write(helper.MarshalResponse(respose))
		return
	}
	id,_ := strconv.Atoi(sess.UserID)
	admin,err:=adminHan.usersrv.GetUserByID(uint(id))
	if err!=nil || admin == nil{
		w.Write(helper.MarshalResponse(respose))
		log.Fatal(err)
		return
	}
	
	id,_ = strconv.Atoi(r.FormValue("user_id"))
	admin,err =adminHan.usersrv.DeleteUser(uint(id))
	if len(err)>0{
		respose.message= "no such user"
		w.Write(helper.MarshalResponse(respose))
		log.Fatal(err)
		return
	}
	respose.message = "Admin Account Delete successful"
	respose.AdminId = uint(id)
		
	w.Write(helper.MarshalResponse(respose))
}
