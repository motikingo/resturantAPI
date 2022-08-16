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
	"github.com/motikingo/resturant-api/helper"

	//"github.com/motikingo/resturant-api/menu"
	"github.com/motikingo/resturant-api/menu"
)

type ItemHandler struct{

	itemSrv menu.ItemService
	session *SessionHandler
	ingrd menu.IngredientService
	catsrv menu.CatagorySrv
}

func NewItemHandler(itemSrv menu.ItemService,session *SessionHandler)ItemHandler{

	return ItemHandler{itemSrv:itemSrv,session: session}
}

func (itemHa *ItemHandler)GetItems(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	session := itemHa.session.GetSession(r)
	if session == nil || session.Role !="Admin"{
		w.Write([]byte("Unauthorized user"))
		return
	}
	items,err:= itemHa.itemSrv.Items()
	if err!=nil{
		log.Fatal()
	}	

	w.Write(helper.MarshalResponse(items))

}

func (itemHa *ItemHandler)GetItem(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	response := &struct{
		message string
		item *entity.Item
	}{
		message: "Unauthorized user",
	}
	session := itemHa.session.GetSession(r)
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	id:= mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)

	if e!=nil{
		response.message = "They given Id is not Integer."
		w.Write(helper.MarshalResponse(response))

	}

	item,err:= itemHa.itemSrv.Item(uint(ids))
	if err!=nil{
		response.message = "No such Item"
		w.Write(helper.MarshalResponse(response))
	}

	response.message = "get Item successfully"
	response.item = item 

	w.Write(helper.MarshalResponse(response))
}

func (itemHa *ItemHandler)CreateItem(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	session := itemHa.session.GetSession(r)
	response := &struct{
		message string
		item *entity.Item
	}{
		message: "UnAuthorized User",
	}
	input := &struct{
		name string
		price float64
		description string
		imageurl string
		ingredient []string
	}{}
	if session == nil || session.Role !="Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}
	read,er:= ioutil.ReadAll(r.Body)

	if er!=nil{
		response.message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
	}
	e:= json.Unmarshal(read,&input)

	if e!=nil{
		response.message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if input.name == "" || input.price == 0 || input.description == "" || input.imageurl == "" || len(input.ingredient)<1{
		response.message = "Invalid input"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if itemHa.itemSrv.IsItemNameExist(input.name){
		response.message = "Item name already exist"
		w.Write(helper.MarshalResponse(response))
		return
	}
	itm := entity.Item{
		Name: input.name,
		Price: input.price,
		Description: input.description,
		Image: input.imageurl,
		Ingridients: input.ingredient,
	}
	item,err:= itemHa.itemSrv.CreateItem(itm)
	if err!=nil{
		response.message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.message = "create successful"
	response.item = item

	w.Write(helper.MarshalResponse(response))
}

func (itemHa *ItemHandler)UpdateItem(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	response := &struct{
		message string
		item *entity.Item
	}{
		message: "UnAuthorized User",
	}
	input := &struct{
		name string
		price float64
		description string
		imageurl string
	}{}
	session := itemHa.session.GetSession(r)
	if session == nil|| session.Role == "Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}
	ids:= mux.Vars(r)["id"]
	id,e:=strconv.Atoi(ids)

	if e!= nil{
		response.message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	read,er:=ioutil.ReadAll(r.Body)

	if er!= nil{
		response.message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	err:=json.Unmarshal(read,&input)
	if err!= nil{
		response.message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	itm := entity.Item{
		Name: input.name,
		Price: input.price,
		Description: input.description,
		Image: input.imageurl,
		Ingridients: input.ingredient,
	}
	itm.ID = uint(id)
	itemUpdate,ers:=itemHa.itemSrv.UpdateItem(itm)

	if ers!= nil{
		log.Fatal(ers)
	}

	response.message = "successfully updated"
	response.item = itemUpdate
	w.Write(helper.MarshalResponse(response))
}

func(itemHa *ItemHandler) AddIngredient(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := itemHa.session.GetSession(r)
	response := &struct{
		message string
		ingrd *entity.Ingredient
	}{
		message: "UnAuthorized user",
	}

	if session == nil || session.Role!="Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}

	igrdId,e := strconv.Atoi(r.FormValue("ingredient_id"))
	if e != nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	itemId,e := strconv.Atoi(r.FormValue("item_id"))

	if e != nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	item,ers := itemHa.itemSrv.Item(uint(itemId))
	if item == nil || len(ers)>0{
		response.message  = "No such Item"
		w.Write(helper.MarshalResponse(response))
		return

	}
	igrd,ers := itemHa.ingrd.Ingredient(uint(igrdId))

	if igrd == nil || len(ers)>0{
		response.message  = "No such ingredient"
		w.Write(helper.MarshalResponse(response))
		return

	}

	for _,id:= range item.Ingridients{
		if id == string(igrd.ID) {
			response.message  = "this ingredient is already added"
			w.Write(helper.MarshalResponse(response))
			return
		}
	}

	item.Ingridients = append(item.Ingridients, string(igrd.ID))
	item,ers = itemHa.itemSrv.UpdateItem(*item)
	if item == nil || len(ers)>0{
		response.message  = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.message = "Ingredient added successfully"
	response.ingrd = igrd
	w.Write(helper.MarshalResponse(response))
}

func(itemHa * ItemHandler)RemoveIngrediend(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := itemHa.session.GetSession(r)
	response := &struct{
		message string
		ingrd *entity.Ingredient
	}{
		message: "UnAuthorized user",
	}

	if session == nil || session.Role!="Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}
	igrdId,e := strconv.Atoi(r.FormValue("ingredient_id"))
	if e != nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	itemId,e := strconv.Atoi(r.FormValue("item_id"))

	if e != nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	item,ers := itemHa.itemSrv.Item(uint(itemId))
	if item == nil || len(ers)>0{
		response.message  = "No such Item"
		w.Write(helper.MarshalResponse(response))
		return

	}
	igrd,ers := itemHa.ingrd.Ingredient(uint(igrdId))

	if igrd == nil || len(ers)>0{
		response.message  = "No such ingredient"
		w.Write(helper.MarshalResponse(response))
		return

	}

	for _,id:= range item.Ingridients{
		if id == string(igrd.ID) {
			response.message  = "this ingredient is already added"
			w.Write(helper.MarshalResponse(response))
			return
		}
	}

	item.Ingridients = append(item.Ingridients, string(igrd.ID))
	item,ers = itemHa.itemSrv.UpdateItem(*item)
	if item == nil || len(ers)>0{
		response.message  = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return

	}

	response.message = "Ingredient added successfully"
	response.ingrd = igrd
	w.Write(helper.MarshalResponse(response))
}

func (itemHa *ItemHandler)GetMyIngregients(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	response := &struct{
		message string
		igrds []entity.Ingredient
	}{
		message: "Unauthorized user",
	}
	session := itemHa.session.GetSession(r)
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	id:= mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)

	if e!=nil{
		response.message = "The given Id is not Integer."
		w.Write(helper.MarshalResponse(response))

	}

	item,err:= itemHa.itemSrv.Item(uint(ids))
	if err!=nil || item == nil{
		response.message = "No such Item"
		w.Write(helper.MarshalResponse(response))
		return
	}
	for _,id := range item.Ingridients{
		idIgrd,_ := strconv.Atoi(id)
		igrd,err := itemHa.ingrd.Ingredient(uint(idIgrd))

		if err!=nil || igrd == nil{
			response.message = "No such Ingredient"
			w.Write(helper.MarshalResponse(response))
			return
		}
		response.igrds = append(response.igrds, *igrd)
	}

	response.message = "successfully retrieved Ingredients"
	w.Write(helper.MarshalResponse(response))
}

func (itemHa *ItemHandler)DeleteItem(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")

	session := itemHa.session.GetSession(r)
	if session == nil{
		w.Write([]byte("Unauthorized user"))
		return
	}
	id:= mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)

	if e!=nil{
		log.Fatal(err)
	}
	item,err:= itemHa.itemSrv.Item(uint(ids))

	if err!=nil{
		log.Fatal(err)
	}

	for _,id := range item.Catagories{
		idCat,_ := strconv.Atoi(id)
		cat,ers := itemHa.catsrv.Catagory(uint(idCat))
		if ers!=nil{
			return
		}

		for _,c := range cat.Items{
			if c != string(item.ID){
				cat.Items = append(cat.Items, c)
			}
		}
		_,ers  = itemHa.catsrv.UpdateCatagory(cat)

		if ers!=nil{
			return
		}	 
	}

	item,err = itemHa.itemSrv.DeleteItem(uint(ids))
	if err!=nil{
		log.Fatal(err)
	}

	w.Write(helper.MarshalResponse(item))
}

