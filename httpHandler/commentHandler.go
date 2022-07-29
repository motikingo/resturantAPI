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

	"github.com/motikingo/resturant-api/comment"
)
var err error
var errs []error
var comm entity.Comment
type CommentHandler struct{
	comSrv comment.CommentService
}

func NewCommentHandler(comSrv comment.CommentService)CommentHandler{
	return CommentHandler{comSrv:comSrv}
}

func(comHandler *CommentHandler)GetComments(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	comments,err := comHandler.comSrv.Comments()
	if err!=nil {
		log.Fatal(err)
	}
	commentsJson,er:=json.MarshalIndent(comments,"","/t/")
	if er!=nil {
		log.Fatal(err)
	}
	w.Write(commentsJson)
}

func(comHandler *CommentHandler)GetComment(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	id := mux.Vars(r)["id"]
	ids,e:=strconv.Atoi(id)

	if e!=nil {
		log.Fatal(err)
	}

	comment,err := comHandler.comSrv.Comment(uint(ids))
	if err!=nil {
		log.Fatal(err)
	}
	commentsJson,er:=json.MarshalIndent(comment,"","/t/")
	if er!=nil {
		log.Fatal(err)
	}
	w.Write(commentsJson)
}

func(comHandler *CommentHandler)CreateComment(w http.ResponseWriter, r *http.Request){
	var comm entity.Comment
	w.Header().Set("Content-Type","application/json")
	read,er:=ioutil.ReadAll(r.Body)

	if er!=nil {
		log.Fatal(err)
	}

	er = json.Unmarshal(read,&comm)
	if er!=nil {
		log.Fatal(err)
	}
	com,ers :=comHandler.comSrv.CreateComment(&comm)

	if ers!=nil {
		log.Fatal(err)
	}
	comMarshal,e:=json.MarshalIndent(com,"","/t/t")

	if e!=nil {
		log.Fatal(err)
	}

	w.Write(comMarshal)
}

func(comHandler *CommentHandler)UpdateComment(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	var comm entity.Comment
	id:= mux.Vars(r)["id"] 
	ids,err:= strconv.Atoi(id)
	if err != nil{
		log.Fatal(err)
	}
	read,e:= ioutil.ReadAll(r.Body)

	if e!= nil{
		log.Fatal(err)
	}
	err = json.Unmarshal(read,&comm)
	if err != nil{
		log.Fatal(err)
	}

	comUpdated,er:=comHandler.comSrv.UpdateComment(uint(ids),comm)

	if er != nil{
		log.Fatal(err)
	}
	comMarsh,errr:= json.MarshalIndent(comUpdated,"","/t/t")
	if errr != nil{
		log.Fatal(err)
	}

	w.Write(comMarsh)
}

func(comHandler *CommentHandler)DeleteComment(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
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