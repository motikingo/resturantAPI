package handler

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	//"fmt"
	"log"

	//"fmt"
	//"io/ioutil"
	"net/http"
	//"strconv"

	//"github.com/gorilla/mux"

	"github.com/gorilla/mux"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/helper"

	"github.com/motikingo/resturant-api/comment"
)
var err error
var errs []error
var comm entity.Comment
type CommentHandler struct{
	comSrv comment.CommentService
	session *SessionHandler
}

func NewCommentHandler(comSrv comment.CommentService, session *SessionHandler)CommentHandler{
	return CommentHandler{comSrv:comSrv, session:session}
}

func(comHandler *CommentHandler)GetComments(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := comHandler.session.GetSession(r)
	response := &struct{
		status string
		comments []entity.Comment
	}{
		status: "Unauthorized user",
	}
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}

	comments,err := comHandler.comSrv.Comments()
	if len(err)>0 || len(comments )>0{
		response.status = "No comment found"
		w.Write(helper.MarshalResponse(response))
		return
	}
	response.status = "successfully retrieved comment" 
	w.Write(helper.MarshalResponse(response))
}

func(comHandler *CommentHandler)GetComment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := comHandler.session.GetSession(r)
	response := &struct{
		status string
		comment *entity.Comment
	}{
		status: "Unauthorized user",
	}
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}

	id := mux.Vars(r)["id"]
	ids,_:=strconv.Atoi(id)

	comment,err := comHandler.comSrv.Comment(uint(ids))
	if err!=nil {
		response.status = "No such comment"
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	w.Write(helper.MarshalResponse(comment))
}

func(comHandler *CommentHandler)CreateComment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	response := &struct{
		status string
		comment *entity.Comment
	}{
		status: "Unauthorized user",
	}
	input := &struct{
		description string
	}{}
	session := comHandler.session.GetSession(r)
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	read,er:=ioutil.ReadAll(r.Body)

	if er!=nil {
		w.Write(helper.MarshalResponse(response))
		return
	}

	er = json.Unmarshal(read,&input)
	if er!=nil || input.description == ""{
		response.status = "Invalid Input"
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	com,ers :=comHandler.comSrv.CreateComment(&comm)

	if ers!=nil {
		response.status = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	w.Write(helper.MarshalResponse(com))
}

func(comHandler *CommentHandler)UpdateComment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := comHandler.session.GetSession(r)
	response := &struct{
		status string
		comment entity.Comment
	}{
		status: "UnAuthorized user",
	}
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	input := &struct{
		description string
	}{}
	var comm entity.Comment
	id:= mux.Vars(r)["id"] 
	ids,_:= strconv.Atoi(id)
	
	read,e:= ioutil.ReadAll(r.Body)

	if e!= nil{
		response.status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	err = json.Unmarshal(read,&input)
	if err != nil {
		response.status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	if input.description == "" {
		response.status = "Nothing added"
		w.Write(helper.MarshalResponse(response))
		return
	}
	comUpdated,er:=comHandler.comSrv.UpdateComment(uint(ids),comm)

	if len(er)>0 {
		w.Write(helper.MarshalResponse(response))
		return
	}
	

	w.Write(helper.MarshalResponse(comUpdated))
}

func(comHandler *CommentHandler)DeleteComment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := comHandler.session.GetSession(r)
	if session == nil{
		w.Write([]byte("Unauthorized user"))
		return
	}
	id := mux.Vars(r)["id"]
	ids,e:=strconv.Atoi(id)

	if e!=nil {
		log.Fatal(err)
	}

	comment,err := comHandler.comSrv.DeleteComment(uint(ids))
	if err!=nil {
		log.Fatal(err)
	}
	commentsJson,er:=json.MarshalIndent(comment,"","/t/")
	if er!=nil {
		log.Fatal(err)
	}
	w.Write(commentsJson)

}