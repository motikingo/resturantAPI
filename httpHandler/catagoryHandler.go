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
	//"github.com/jinzhu/gorm"

	//"github.com/motikingo/resturant-api/menu/service"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/helper"
	"github.com/motikingo/resturant-api/menu"
)

// var tmpl *template.Template

// func init(){
// 	tmpl = template.Must(template.ParseGlob("../template/menu/*.html"))
// }

type CatagoryHandler struct {
	catSrv  menu.CatagorySrv
	session *SessionHandler
	itemSrv menu.ItemService
}

func NewCatagoryHandler(catSrv menu.CatagorySrv, session *SessionHandler, itemSrv menu.ItemService) CatagoryHandler {
	return CatagoryHandler{catSrv: catSrv, session: session, itemSrv: itemSrv}
}

func (catHandler *CatagoryHandler) GetCatagories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	w.Header().Set("Content-Type", "application/json")
	session := catHandler.session.GetSession(r)

	fmt.Println("here")
	if session == nil {
		w.Write([]byte("Unauthorized user"))
		return
	}
	catagories, err := catHandler.catSrv.Catagories()

	if len(err) > 0 {
		w.Write([]byte("Internal server error"))
		return
	}

	w.Write(helper.MarshalResponse(catagories))
}

func (catHandler *CatagoryHandler) GetCatagory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	session := catHandler.session.GetSession(r)
	if session == nil {
		w.Write([]byte("Unauthorized user"))
		return
	}
	fmt.Println("here")
	id := mux.Vars(r)["id"]

	ids, err := strconv.Atoi(id)
	fmt.Println(ids)
	if err != nil {
		//fmt.Println(ids)
		w.Write([]byte(err.Error()))
		return
	}

	cat, er := catHandler.catSrv.Catagory(uint(ids))

	if len(er) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// catagories,err := json.MarshalIndent(cat,"","t/t/")

	// if err != nil{
	// 	http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
	// 	return
	// }

	w.Write(helper.MarshalResponse(cat))
}

func (catHandler *CatagoryHandler) CreateCatagory(w http.ResponseWriter, r *http.Request) {
	var cat entity.Catagory
	session := catHandler.session.GetSession(r)
	if session == nil || session.Role != "Admin" {
		w.Write([]byte("Unauthorized user"))
		return
	}

	respose := &struct {
		Success bool
		Message string
		Cat     *entity.Catagory
	}{
		Success: false,
		Message: "failed to create catagory",
	}
	input := &struct {
		Name     string `json:"name"`
		Imageurl string `json:"imageurl"`
		ItemId   []int  `json:"item_Id"`
	}{}
	read, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write(helper.MarshalResponse(respose))
		return
	}
	err = json.Unmarshal(read, &input)

	if err != nil {
		w.Write(helper.MarshalResponse(respose))
		log.Fatal(err)

		return
	}

	if catHandler.catSrv.IsCatagoryNameExist(input.Name) {
		respose.Message = "catagory with this name already exist"
		w.Write(helper.MarshalResponse(respose))
		return

	}
	if input.Name == "" || input.Imageurl == "" || len(input.ItemId) < 1 {
		respose.Message = "Invalid Input"
		w.Write(helper.MarshalResponse(respose))
		return

	}
	cat.Name = input.Name
	cat.ImageUrl = input.Imageurl
	for _, id := range input.ItemId {
		item, ers := catHandler.itemSrv.Item(uint(id))
		if len(ers) > 0 {
			respose.Message = "Internal server error"
			w.Write(helper.MarshalResponse(respose))
			return
		}
		cat.Items = append(cat.Items, *item)
	}
	newcat, errs := catHandler.catSrv.CreateCatagory(cat)
	if len(errs) > 0 {
		respose.Message = "Internal server error"
		w.Write(helper.MarshalResponse(respose))
		return
	}

	// for _, it := range cat.Items {
	// 	item, ers := catHandler.itemSrv.Item(it.ID)
	// 	if len(ers) > 0 || item == nil {
	// 		respose.Message = "Internal server error"
	// 		w.Write(helper.MarshalResponse(respose))
	// 		return
	// 	}
	// item.Catagories = append(item.Catagories, string(cat.ID))
	// item,ers = catHandler.itemSrv.UpdateItem(*item)
	// if len(ers)>0 || item == nil{
	// 	respose.message = "Internal server error"
	// 	w.Write(helper.MarshalResponse(respose))
	// 	return
	// }
	// }
	//updated,_ = ioutil.ReadAll(updatedcat)
	// c,_:=json.MarshalIndent(newcat,"","/t/t")
	// //err = json.NewDecoder().Decode(&cat)
	// if err != nil{
	// 	fmt.Println("not Unmarshal!")

	// 	return
	// }

	respose.Message = "catagory successfully created"
	respose.Success = true
	respose.Cat = newcat
	w.Write(helper.MarshalResponse(respose))

}

func (catHander *CatagoryHandler) UpdateCatagory(w http.ResponseWriter, r *http.Request) {
	var cat entity.Catagory
	w.Header().Set("Content-Type", "application/json")
	session := catHander.session.GetSession(r)
	response := &struct {
		Message  string
		Catagory *entity.Catagory
	}{
		Message: "Unauthorized user",
	}

	input := &struct {
		Name     string
		Imageurl string
	}{}
	if session == nil || session.Role != "Admin" {
		w.Write(helper.MarshalResponse(response))
		return

	}

	id := mux.Vars(r)["id"]
	ids, e := strconv.Atoi(id)
	if e != nil {
		log.Fatal(e)
		return
	}
	read, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	if err := json.Unmarshal(read, &input); err != nil {
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		log.Fatal(err)
		return
	}

	if input.Name == "" || input.Imageurl == "" {
		response.Message = "Invalid input"
		w.Write(helper.MarshalResponse(response))
		return

	}
	c, ers := catHander.catSrv.Catagory(uint(ids))
	if len(ers) > 0 {
		response.Message = "No such catagory"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if c.Name != input.Name {
		if catHander.catSrv.IsCatagoryNameExist(input.Name) {
			response.Message = "catagory with this name already exist"
			w.Write(helper.MarshalResponse(response))
			return
		}
	}

	cat = entity.Catagory{
		Name:     input.Name,
		ImageUrl: input.Imageurl,
	}
	cat.ID = uint(ids)

	catupdated, er := catHander.catSrv.UpdateCatagory(cat)

	if er != nil {
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		log.Fatal(err)
		return
	}
	response.Message = "Update successful"
	response.Catagory = catupdated
	w.Write(helper.MarshalResponse(response))
}

func (catHandler *CatagoryHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session := catHandler.session.GetSession(r)
	response := &struct {
		Message  string
		Catagory *entity.Catagory
	}{
		Message: "UnAuthorized user",
	}
	input := &struct {
		Item_id int
	}{}

	if &input.Item_id == nil {
		response.Message = "Invalid input"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if session == nil || session.Role != "Admin" {
		w.Write(helper.MarshalResponse(response))
		return
	}
	read, _ := ioutil.ReadAll(r.Body)
	if er := json.Unmarshal(read, input); er != nil {
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	item, er := catHandler.itemSrv.Item(uint(input.Item_id))
	if item == nil || len(er) > 0 {
		response.Message = "No such Item"
		w.Write(helper.MarshalResponse(response))
		return
	}

	Cat_id, _ := strconv.Atoi(mux.Vars(r)["id"])

	cat, err := catHandler.catSrv.Catagory(uint(Cat_id))
	if cat == nil || len(err) > 0 {
		response.Message = "No such Catagory"
		w.Write(helper.MarshalResponse(response))
		return
	}

	for _, it := range cat.Items {
		if it.ID == item.ID {
			response.Message = "Item already exist"
			w.Write(helper.MarshalResponse(response))
			return
		}
	}
	cat.Items = append(cat.Items, *item)
	cat, err = catHandler.catSrv.UpdateCatagory(*cat)
	// item.Catagories = append(item.Catagories, cat_id)
	// item,er = catHandler.itemSrv.UpdateItem(*item)

	if cat == nil || len(err) > 0 {
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	response.Message = "Add Item successful"
	response.Catagory = cat

	w.Write(helper.MarshalResponse(response))

}

func (catHandler *CatagoryHandler) GetMyItems(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	sess := catHandler.session.GetSession(r)

	response := &struct {
		Message string
		Items   []*entity.Item
	}{
		Message: "UnAuthorized user",
	}

	if sess == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["catagory_id"])

	cat, ers := catHandler.catSrv.Catagory(uint(id))

	if cat == nil || len(ers) > 0 {
		response.Message = "no such catagory"
		w.Write(helper.MarshalResponse(response))
		return
	}

	for _, itm := range cat.Items {
		item, ers := catHandler.itemSrv.Item(itm.ID)
		if item == nil || len(ers) > 0 {
			response.Message = "no such catagory"
			w.Write(helper.MarshalResponse(response))
			return
		}
		response.Items = append(response.Items, item)

	}

	response.Message = "successfully retrive Items"
	w.Write(helper.MarshalResponse(response))

}

func (catHandler *CatagoryHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	session := catHandler.session.GetSession(r)
	response := &struct {
		Message  string
		Catagory *entity.Catagory
	}{
		Message: "UnAuthorized user",
	}
	input := &struct {
		Item_id int `json:"item_id"`
	}{}
	if session == nil || session.Role != "Admin" {
		w.Write(helper.MarshalResponse(response))
		return
	}
	read, _ := ioutil.ReadAll(r.Body)

	if er := json.Unmarshal(read, input); er != nil {
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	item, er := catHandler.itemSrv.Item(uint(input.Item_id))
	if item == nil || len(er) > 0 {
		response.Message = "No such Item"
		w.Write(helper.MarshalResponse(response))
		return
	}
	id_cat, _ := strconv.Atoi(mux.Vars(r)["id"])
	cat, err := catHandler.catSrv.Catagory(uint(id_cat))
	if cat == nil || len(err) > 0 {
		response.Message = "No such Catagory"
		w.Write(helper.MarshalResponse(response))
		return
	}
	check := false
	var itms []entity.Item
	for _, it := range cat.Items {
		if it.ID != item.ID {
			itms = append(itms, it)
			fmt.Println(it.ID)
		} else {
			check = true
		}
	}
	if check {
		copy(cat.Items, itms)
		cat, err = catHandler.catSrv.UpdateCatagory(*cat)
		fmt.Println(itms)
		if cat == nil || len(err) > 0 {
			response.Message = "Internal server Error"
			w.Write(helper.MarshalResponse(response))
			return
		}
		response.Message = "Delete Item from catagory successfully"
		response.Catagory = cat
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Message = "Given item is not in the catagory"
	response.Catagory = cat
	w.Write(helper.MarshalResponse(response))

}

func (catHandler *CatagoryHandler) DeleteCatagory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	session := catHandler.session.GetSession(r)
	response := &struct {
		Message  string
		Catagory *entity.Catagory
	}{
		Message: "Unauthorized user",
	}
	if session == nil || session.Role != "Admin" {
		w.Write(helper.MarshalResponse(response))
		return
	}
	id := mux.Vars(r)["id"]
	ids, _ := strconv.Atoi(id)

	cata, er := catHandler.catSrv.Catagory(uint(ids))
	if len(er) > 0 || cata == nil {
		response.Message = "no such catagory exist"
		w.Write(helper.MarshalResponse(response))
		return
	}

	// for _,it:= range cata.Items{
	// 	item,er:= catHandler.itemSrv.Item(it.ID)

	// 	if len(er)>0 || item ==nil{
	// 		response.message = "no such Item exist"
	// 		w.Write(helper.MarshalResponse(response))
	// 		return
	// 	}

	// 	// for _,c_id:= range item.Catagories{
	// 	// 	if c_id != id {
	// 	// 		item.Catagories = append(item.Catagories, c_id)
	// 	// 	}
	// 	// }
	// 	// item,er = catHandler.itemSrv.UpdateItem(*item)

	// 	// if len(er)>0 || item ==nil{
	// 	// 	response.message = "Internal Server Error"
	// 	// 	w.Write(helper.MarshalResponse(response))
	// 	// 	return
	// 	// }

	// }
	cata, er = catHandler.catSrv.DeleteCatagory(uint(ids))
	if len(er) > 0 || cata == nil {
		response.Message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	response.Message = "Delete catagory successful"
	response.Catagory = cata

	w.Write(helper.MarshalResponse(response))
}
