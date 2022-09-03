package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/helper"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/motikingo/resturant-api/comment"
	"github.com/motikingo/resturant-api/user"
)

var err error
var comm entity.Comment

type CommentHandler struct {
	comSrv  comment.CommentService
	userSrv user.UserService
	session *SessionHandler
}

func NewCommentHandler(comSrv comment.CommentService, userSrv user.UserService, session *SessionHandler) CommentHandler {
	return CommentHandler{comSrv: comSrv, userSrv: userSrv, session: session}
}

func (comHandler *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session := comHandler.session.GetSession(r)
	response := &struct {
		Status   string
		Comments []entity.Comment
	}{
		Status: "Unauthorized user",
	}
	if session == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}

	comments, err := comHandler.comSrv.Comments()
	if len(err) > 0 || len(comments) > 0 {
		response.Status = "No comment found"
		w.Write(helper.MarshalResponse(response))
		return
	}
	response.Status = "successfully retrieved comment"
	w.Write(helper.MarshalResponse(response))
}

func (comHandler *CommentHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session := comHandler.session.GetSession(r)
	response := &struct {
		Status  string
		Comment *entity.Comment
	}{
		Status: "Unauthorized user",
	}
	if session == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}

	id := mux.Vars(r)["id"]
	ids, _ := strconv.Atoi(id)

	comment, err := comHandler.comSrv.Comment(uint(ids))
	if err != nil {
		response.Status = "No such comment"
		w.Write(helper.MarshalResponse(response))
		return
	}

	w.Write(helper.MarshalResponse(comment))
}

func (comHandler *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := &struct {
		Status  string
		Comment *entity.Comment
	}{
		Status: "Unauthorized user",
	}
	input := &struct {
		Description string
	}{}
	session := comHandler.session.GetSession(r)
	if session == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}
	read, er := ioutil.ReadAll(r.Body)

	if er != nil {
		w.Write(helper.MarshalResponse(response))
		return
	}

	if er := json.Unmarshal(read, &input); er != nil || input.Description == "" {
		response.Status = "Invalid Input"
		w.Write(helper.MarshalResponse(response))
		return
	}

	userId, _ := strconv.Atoi(session.UserID)
	user, ers := comHandler.userSrv.GetUserByID(uint(userId))
	if len(ers) > 0 {
		response.Status = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	comm = entity.Comment{
		Description: input.Description,
		UserID:      user.ID,
	}
	user.Comments = append(user.Comments, comm)
	_, ers = comHandler.userSrv.UpdateUser(*user)
	if len(ers) > 0 {
		response.Status = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	com, ers := comHandler.comSrv.CreateComment(&comm)

	if len(ers) > 0 {
		response.Status = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	w.Write(helper.MarshalResponse(com))
}

func (comHandler *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session := comHandler.session.GetSession(r)
	response := &struct {
		Status  string
		Comment entity.Comment
	}{
		Status: "UnAuthorized user",
	}
	if session == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}
	input := &struct {
		Description string
	}{}
	id := mux.Vars(r)["id"]
	ids, _ := strconv.Atoi(id)

	read, e := ioutil.ReadAll(r.Body)

	if e != nil {
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	if err := json.Unmarshal(read, &input); err != nil {
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	if input.Description == "" {
		response.Status = "Nothing added"
		w.Write(helper.MarshalResponse(response))
		return
	}
	com, err := comHandler.comSrv.Comment(uint(ids))
	if len(err) > 0 {
		response.Status = "No such Comment"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if com.UserID != uint(ids) {
		w.Write(helper.MarshalResponse(response))
		return
	}
	fmt.Println(com.UserID)
	com.Description = input.Description
	com, err = comHandler.comSrv.UpdateComment(*com)

	if len(err) > 0 {
		w.Write(helper.MarshalResponse(response))
		return
	}

	w.Write(helper.MarshalResponse(com))
}

func (comHandler *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session := comHandler.session.GetSession(r)
	response := &struct {
		Status string
		Com    entity.Comment
	}{
		Status: "Unauthorized user",
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if session == nil {

		w.Write(helper.MarshalResponse(response))
		return
	}
	comment, err := comHandler.comSrv.DeleteComment(uint(id))
	if err != nil {
		response.Status = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if comment.UserID != uint(id) {
		w.Write(helper.MarshalResponse(response))
		return
	}
	response.Status = "Comment Deleted successfully"
	response.Com = *comment
	w.Write(helper.MarshalResponse(response))

}
