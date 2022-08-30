package handler

import (
	"encoding/json"
	"fmt"
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

func NewAdminHandler(usersrv user.UserService,session *SessionHandler) *AdminHandler{
	return &AdminHandler{usersrv:usersrv, session:session }
}


func(adminHan *AdminHandler)CreateAdmin(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	fmt.Println("Admin")
	if r.Method == "POST"{
		response := &struct{
			Status string
			Admin *entity.User
		}{
			Status: "Invalid Input",
		}
		Input:=&struct{
			Email string `json:"email"`
			Password string `json:"password"`
			Confirm_password string `json:"confirm_password"`
		}{}
		
		read,e:=ioutil.ReadAll(r.Body)

		if e!=nil{
			log.Fatal(err)
		}
		
		er:= json.Unmarshal(read,&Input)

		if er != nil || Input.Email == "" || Input.Password == "" || Input.Confirm_password==""{
			w.Write(helper.MarshalResponse(response))
			return
		}

		if len(Input.Password)<8 {
			response.Status = "password lenght must be more than 8 character"
			w.Write(helper.MarshalResponse(response))
			return
		}
		if Input.Password != Input.Confirm_password {
			response.Status = "confirm password is not the same"
			w.Write(helper.MarshalResponse(response))
			return			
		}

		if adminHan.usersrv.GetUserByEmail(Input.Email) != nil {
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

		admin := &entity.User{
			Email:Input.Email,
			Password:Input.Password,
			Role:"Admin",
			}		
		
		adm,err:=adminHan.usersrv.CreateUser(*admin)
		if err!=nil|| adm == nil{
			response.Status ="Internal server error"
			w.Write(helper.MarshalResponse(response))
			return
		}
		response.Status = "successfully created"
		response.Admin = adm
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
	
	userUpdated,errs:=adminHan.usersrv.UpdateUser(*admin)

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
		Status string
		AdminId uint
	}{
		Status: "login faild",
	}
	input :=&struct{
		Email string
		Password string
	}{}
	fmt.Println("here")
	read,_:= ioutil.ReadAll(r.Body)

	if er := json.Unmarshal(read,&input);er != nil{
		w.Write(helper.MarshalResponse(response))
		return
	}

	if len(input.Password)<4 || input.Password ==""{
		response.Status= "invalid password input "
		w.Write(helper.MarshalResponse(response))
		return
	}
	if input.Email == ""{
		response.Status= "invalud input email"
		w.Write(helper.MarshalResponse(response))
		return
	}

	admin := adminHan.usersrv.GetUserByEmail(input.Email)
	if admin == nil{
		response.Status = "this email is not registered"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if !helper.MatchPassword(admin.Password,input.Password) {
		response.Status = "Incorrect password"
		w.Write(helper.MarshalResponse(response))
		return
		
	}
	
	sess := &entity.Session{
		UserID:strconv.Itoa(int(admin.ID)),
		Email: admin.Email,	
		Role:"Admin",	
	}

	if !adminHan.session.CreateSession(sess,w){
		response.Status = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Status = "Login successful"
	response.AdminId = admin.ID
	w.Write(helper.MarshalResponse(response))
}

func(adminHan *AdminHandler)Logout(w http.ResponseWriter,r *http.Request){
	sess := adminHan.session.GetSession(r)
	response := &struct{
		Success bool
		Message string
		AdminId string
	}{
		Success: false,
		Message: "UnAuthorized User...",
	}
	if sess == nil || sess.Role == "Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	if adminHan.session.DeleteSession(w)!= nil{
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Message = "Logout successful"
	response.Success = true
	response.AdminId = sess.UserID
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
	respose.AdminId = admin.ID
		
	w.Write(helper.MarshalResponse(respose))
}
