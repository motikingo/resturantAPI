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
		Status string
		Comments []entity.Comment
	}{
		Status: "Unauthorized user",
	}
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}

	comments,err := comHandler.comSrv.Comments()
	if len(err)>0 || len(comments )>0{
		response.Status = "No comment found"
		w.Write(helper.MarshalResponse(response))
		return
	}
	response.Status = "successfully retrieved comment" 
	w.Write(helper.MarshalResponse(response))
}

func(comHandler *CommentHandler)GetComment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := comHandler.session.GetSession(r)
	response := &struct{
		Status string
		Comment *entity.Comment
	}{
		Status: "Unauthorized user",
	}
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}

	id := mux.Vars(r)["id"]
	ids,_:=strconv.Atoi(id)

	comment,err := comHandler.comSrv.Comment(uint(ids))
	if err!=nil {
		response.Status = "No such comment"
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	w.Write(helper.MarshalResponse(comment))
}

func(comHandler *CommentHandler)CreateComment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	response := &struct{
		Status string
		Comment *entity.Comment
	}{
		Status: "Unauthorized user",
	}
	input := &struct{
		Description string
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
	if er!=nil || input.Description == ""{
		response.Status = "Invalid Input"
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	com,ers :=comHandler.comSrv.CreateComment(&comm)

	if ers!=nil {
		response.Status = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	w.Write(helper.MarshalResponse(com))
}

func(comHandler *CommentHandler)UpdateComment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := comHandler.session.GetSession(r)
	response := &struct{
		Status string
		Comment entity.Comment
	}{
		Status: "UnAuthorized user",
	}
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	input := &struct{
		Description string
	}{}
	var comm entity.Comment
	id:= mux.Vars(r)["id"] 
	ids,_:= strconv.Atoi(id)
	
	read,e:= ioutil.ReadAll(r.Body)

	if e!= nil{
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	err = json.Unmarshal(read,&input)
	if err != nil {
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	if input.Description == "" {
		response.Status = "Nothing added"
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