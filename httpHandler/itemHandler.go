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

	//"github.com/motikingo/resturant-api/menu"
	"github.com/motikingo/resturant-api/menu"
)

type ItemHandler struct{

	itemSrv menu.ItemService
}

func NewItemHandler(itemSrv menu.ItemService)ItemHandler{

	return ItemHandler{itemSrv:itemSrv}
}

func (itemHa *ItemHandler)GetItems(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	items,err:= itemHa.itemSrv.Items()
	if err!=nil{
		log.Fatal(err)
	}

	itemsMar,er:= json.MarshalIndent(items,"","/t/")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(itemsMar)


}

func (itemHa *ItemHandler)GetItem(w http.ResponseWriter,r *http.Request){


	w.Header().Set("Content-Type","application/json")
	id:= mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)

	if e!=nil{
		log.Fatal(err)
	}

	items,err:= itemHa.itemSrv.Item(uint(ids))
	if err!=nil{
		log.Fatal(err)
	}

	itemsMar,er:= json.MarshalIndent(items,"","/t/")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(itemsMar)
}

func (itemHa *ItemHandler)CreateItem(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	var item entity.Item
	read,er:= ioutil.ReadAll(r.Body)

	if er!=nil{
		log.Fatal(err)
	}
	e:= json.Unmarshal(read,&item)

	if e!=nil{
		log.Fatal(err)
	}
	items,err:= itemHa.itemSrv.CreateItem(&item)
	if err!=nil{
		log.Fatal(err)
	}

	itemsMar,er:= json.MarshalIndent(items,"","/t/")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(itemsMar)
}

func (itemHa *ItemHandler)UpdateItem(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	var item entity.Item
	ids:= mux.Vars(r)["id"]

	id,e:=strconv.Atoi(ids)

	if e!= nil{
		log.Fatal(err)
	}
	read,er:=ioutil.ReadAll(r.Body)

	if er!= nil{
		log.Fatal(err)
	}

	err:=json.Unmarshal(read,&item)
	if err!= nil{
		log.Fatal(err)
	}

	itemUpdate,ers:=itemHa.itemSrv.UpdateItem(uint(id),item)

	if ers!= nil{
		log.Fatal(ers)
	}

	itemMar,errr:=json.MarshalIndent(itemUpdate,"","'r/r/")


	if errr!= nil{
		log.Fatal(errr)
	}
	w.Write(itemMar)
}

func (itemHa *ItemHandler)DeleteItem(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	id:= mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)

	if e!=nil{
		log.Fatal(err)
	}

	items,err:= itemHa.itemSrv.DeleteItem(uint(ids))
	if err!=nil{
		log.Fatal(err)
	}

	itemsMar,er:= json.MarshalIndent(items,"","/t/")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(itemsMar)
}