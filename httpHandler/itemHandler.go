package handler

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"log"

	//"reflect"

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
	//userV user.UserService
}

func NewItemHandler(itemSrv menu.ItemService,session *SessionHandler,ingrd menu.IngredientService,catsrv menu.CatagorySrv)ItemHandler{

	return ItemHandler{itemSrv:itemSrv,session:session,ingrd:ingrd,catsrv:catsrv}
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
	response := struct{
		Message string
		Item entity.Item
	}{
		Message: "UnAuthorized User",
	}
	
	session := itemHa.session.GetSession(r)
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	id:= mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)

	if e!=nil{
		response.Message = "They given Id is not Integer."
		w.Write(helper.MarshalResponse(response))

	}

	item,err:= itemHa.itemSrv.Item(uint(ids))
	if err!=nil{
		response.Message = "No such Item"
		w.Write(helper.MarshalResponse(response))
	}

	response.Message = "get Item successfully"
	response.Item = *item 
	//a,_:=json.Marshal(response)
	w.Write(helper.MarshalResponse(response))
}

func (itemHa *ItemHandler)CreateItem(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	session := itemHa.session.GetSession(r)
	response := &struct{
		Message string
		Item *entity.Item
	}{
		Message: "UnAuthorized User",
	}
	input := &struct{
		Name string `json:"name"`
		Price float64 `json:"price"`
		Description string `json:"description"`
		Imageurl string `json:"image"`
		Number int `json:"number"`
		Ingredient []int `json:"ingredients"`
	}{}
	if session == nil || session.Role !="Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}
	read,er:= ioutil.ReadAll(r.Body)

	if er!=nil{
		response.Message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
	}
	e:= json.Unmarshal(read,&input)

	if e!=nil{
		response.Message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	if input.Name == "" || input.Price == 0 || input.Description == "" || input.Imageurl == "" || len(input.Ingredient)<1|| input.Number<1{
		response.Message = "Invalid input"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if itemHa.itemSrv.IsItemNameExist(input.Name){
		response.Message = "Item name already exist"
		w.Write(helper.MarshalResponse(response))
		return
	}

	var igrds []entity.Ingredient

	for _,ig:=range input.Ingredient{
		igrd,ers := itemHa.ingrd.Ingredient(uint(ig))
		if len(ers)>0{
			response.Message = "Internal server Error"
			w.Write(helper.MarshalResponse(response))
			return
		}
		igrds = append(igrds, *igrd)
	}

	//fmt.Println(reflect.TypeOf(input.Ingredient[0]))
	itm := entity.Item{
		Name: input.Name,
		Price: input.Price,
		Description: input.Description,
		Image: input.Imageurl,
		Number: input.Number,
		Ingridients: igrds,
	}
	
	item,err:= itemHa.itemSrv.CreateItem(itm)
	if err!=nil{
		response.Message = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Message = "create successful"
	response.Item = item

	w.Write(helper.MarshalResponse(response))
}

func (itemHa *ItemHandler)UpdateItem(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	response := &struct{
		Message string
		Item *entity.Item
	}{
		Message: "UnAuthorized User",
	}
	input := &struct{
		Name string `json:"name"`
		Price float64 `json:"price"`
		Description string  `json:"description"`
		Imageurl string  `json:"imageurl"`
		Number int 		`json:"number"`
	}{}
	session := itemHa.session.GetSession(r)
	if session == nil|| session.Role != "Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	id,e:=strconv.Atoi( mux.Vars(r)["id"])

	if e!= nil{
		response.Message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	read,er:=ioutil.ReadAll(r.Body)

	if er!= nil{
		response.Message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	// if input.Name == "" || input.Price == 0 || input.Description == "" || input.Imageurl == "" || input.Number<1{
	// 	response.Message = "Invalid input"
	// 	w.Write(helper.MarshalResponse(response))
	// 	return
	// }
	if itemHa.itemSrv.IsItemNameExist(input.Name){
		response.Message = "This Item does't exist"
		w.Write(helper.MarshalResponse(response))
		return
	}
	if err:=json.Unmarshal(read,&input); err!= nil{
		response.Message = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	itm := entity.Item{
		Name: input.Name,
		Price: input.Price,
		Description: input.Description,
		Image: input.Imageurl,
		Number: input.Number,
	}
	itm.ID = uint(id)
	itemUpdate,ers:=itemHa.itemSrv.UpdateItem(itm)

	if ers!= nil{
		log.Fatal(ers)
	}

	response.Message = "successfully updated"
	response.Item = itemUpdate
	w.Write(helper.MarshalResponse(response))
}

func(itemHa *ItemHandler) AddIngredient(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := itemHa.session.GetSession(r)
	response := &struct{
		Message string
		Ingrd *entity.Ingredient
	}{
		Message: "UnAuthorized user",
	}
	input := &struct{
		Item_id int `json:"item_id"`
		Ingredient_id int `json:"ingredient_id"` 
	}{}
 	if session == nil || session.Role!="Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}

	// igrdId,e := strconv.Atoi(r.FormValue("ingredient_id"))
	// if e != nil{
	// 	w.Write(helper.MarshalResponse(response))
	// 	return
	// }
	// itemId,e := strconv.Atoi(r.FormValue("item_id"))

	// if e != nil{
	// 	w.Write(helper.MarshalResponse(response))
	// 	return
	// }
	read,ers := ioutil.ReadAll(r.Body)
	if ers != nil {
		response.Message  = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	if er := json.Unmarshal(read,&input);er!=nil {
		response.Message  = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	item,er := itemHa.itemSrv.Item(uint(input.Item_id))
	if item == nil || len(er)>0{
		response.Message  = "No such Item"
		w.Write(helper.MarshalResponse(response))
		return

	}
	igrd,e := itemHa.ingrd.Ingredient(uint(input.Ingredient_id))

	if igrd == nil || len(e)>0{
		response.Message  = "No such ingredient"
		w.Write(helper.MarshalResponse(response))
		return

	}

	for _,ig:= range item.Ingridients{
		if ig.ID == igrd.ID {
			response.Message  = "this ingredient is already added"
			w.Write(helper.MarshalResponse(response))
			return
		}
	}

	item.Ingridients = append(item.Ingridients, *igrd)
	item,er = itemHa.itemSrv.UpdateItem(*item)
	if item == nil || len(er)>0{
		response.Message  = "Internal server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Message = "Ingredient added successfully"
	response.Ingrd = igrd
	w.Write(helper.MarshalResponse(response))
}

func(itemHa * ItemHandler)RemoveIngrediend(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := itemHa.session.GetSession(r)
	response := &struct{
		Message string
		Ingrd *entity.Ingredient
	}{
		Message: "UnAuthorized user",
	}
	input := &struct{
		Ingredient_id int `json:"ingredient_id"`
		Item_id int 	`json:"item_id"`
	}{}
	if session == nil || session.Role!="Admin"{
		w.Write(helper.MarshalResponse(response))
		return
	}
	read,_ := ioutil.ReadAll(r.Body)
	if er:=json.Unmarshal(read,&input); er!=nil{
		response.Message  = "Internal server Error"
			w.Write(helper.MarshalResponse(response))
			return
	}

	item,ers := itemHa.itemSrv.Item(uint(input.Item_id))
	if item == nil || len(ers)>0{
		response.Message  = "No such Item"
		w.Write(helper.MarshalResponse(response))
		return

	}
	igrd,ers := itemHa.ingrd.Ingredient(uint(input.Ingredient_id))

	if igrd == nil || len(ers)>0{
		response.Message  = "No such ingredient"
		w.Write(helper.MarshalResponse(response))
		return

	}
	check:= false
	var ingrds[]entity.Ingredient
	for _,ig:= range item.Ingridients{
		if ig.ID != uint(input.Ingredient_id) {
			fmt.Println(ig.ID)
			ingrds = append(ingrds, ig)
			
		}else{
			check = true
		}
	}
	if check{
		copy(item.Ingridients , ingrds)
		item,ers = itemHa.itemSrv.UpdateItem(*item)
		if item == nil || len(ers)>0{
			response.Message  = "Internal server Error"
			w.Write(helper.MarshalResponse(response))
			return

		}

		response.Message = "Ingredient deleted successfully"
		response.Ingrd = igrd
		w.Write(helper.MarshalResponse(ingrds))
	}
	

	response.Message = "No such Ingredient "
	response.Ingrd = igrd
	w.Write(helper.MarshalResponse(response))
}

func (itemHa *ItemHandler)GetMyIngregients(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	response := &struct{
		Message string
		Igrds []entity.Ingredient
	}{
		Message: "Unauthorized user",
	}
	session := itemHa.session.GetSession(r)
	if session == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	id:= mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)

	if e!=nil{
		response.Message = "The given Id is not Integer."
		w.Write(helper.MarshalResponse(response))

	}

	item,err:= itemHa.itemSrv.Item(uint(ids))
	if err!=nil || item == nil{
		response.Message = "No such Item"
		w.Write(helper.MarshalResponse(response))
		return
	}
	for _,ig := range item.Ingridients{
		igrd,err := itemHa.ingrd.Ingredient(ig.ID)

		if err!=nil || igrd == nil{
			response.Message = "No such Ingredient"
			w.Write(helper.MarshalResponse(response))
			return
		}
		response.Igrds = append(response.Igrds, *igrd)
	}

	response.Message = "successfully retrieved Ingredients"
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

	for _,ca := range item.Catagories{
		cat,ers := itemHa.catsrv.Catagory(ca.ID)
		if ers!=nil{
			return
		}

		for _,c := range cat.Items{
			if c.ID != item.ID{
				cat.Items = append(cat.Items, c)
			}
		}
		_,ers  = itemHa.catsrv.UpdateCatagory(*cat)

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

