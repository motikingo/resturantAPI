package handler

import (
	"encoding/json"
	"log"

	//"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/user"
)

type UserHandler struct{
	userSrvc user.UserService
}


func NewUserHandler(userSrvc user.UserService)UserHandler{
	return UserHandler{userSrvc:userSrvc}
}

func(usrHan *UserHandler)GetUsers(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	users,err:=usrHan.userSrvc.Users()
	if err!=nil{
		log.Fatal(err)
	}
	urs,er:=json.MarshalIndent(users,"","/t/t")

	if er!=nil{
		log.Fatal(err)
	}
	w.Write(urs)
}

func(usrHan *UserHandler)GetUser(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
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
	var user entity.User
	read,e:=ioutil.ReadAll(r.Body)

	if e!=nil{
		log.Fatal(err)
	}

	er:=json.Unmarshal(read,&user)
	
	if er!=nil{
		log.Fatal(err)
	}
	userMar,err:=usrHan.userSrvc.CreateUser(&user)
	if err!=nil{
		log.Fatal(err)
	}
	urs,er:=json.MarshalIndent(userMar,"","/t/t")

	if er!=nil{
		log.Fatal(err)
	}
	w.Write(urs)
}
func(usrHan *UserHandler)UpdateUser(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	var user entity.User
	ids := mux.Vars(r)["id"]
	id,e:=strconv.Atoi(ids)
	if e!=nil{
		log.Fatal(err)
	}
	read,er:= ioutil.ReadAll(r.Body)
	if er!=nil{
		log.Fatal(err)
	}
	er = json.Unmarshal(read,&user)
	if er!=nil{
		log.Fatal(err)
	}

	userUpdated,errs:=usrHan.userSrvc.UpdateUser(uint(id),user)

	if errs!=nil{
		log.Fatal(err)
	}
	userMash,errr:=json.MarshalIndent(userUpdated,"","/r/r")

	if errr!=nil{
		log.Fatal(err)
	}

	w.Write(userMash)

}
func(usrHan *UserHandler)DeleteUser(w http.ResponseWriter,r *http.Request){


	w.Header().Set("Content-Type","application/json")
	id:=mux.Vars(r)["id"]
	ids,e:=strconv.Atoi(id)
	if e!=nil{
		log.Fatal(err)
	}

	user,err:=usrHan.userSrvc.DeleteUser(uint(ids))
	if err!=nil{
		log.Fatal(err)
	}
	urs,er:=json.MarshalIndent(user,"","/t/t")

	if er!=nil{
		log.Fatal(err)
	}
	w.Write(urs)
}