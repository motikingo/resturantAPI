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
	"github.com/motikingo/resturant-api/menu"
)

type IngredientHandler struct{

	ingrdSrv menu.IngredientService
	itemsrv  menu.ItemService
	session *SessionHandler
}

func NewIngredientHandler(ingrdSrv menu.IngredientService,session *SessionHandler)IngredientHandler{

	return IngredientHandler{ingrdSrv:ingrdSrv,session:session}
}

func (ingrdHandler *IngredientHandler)GetIngredients(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	session := ingrdHandler.session.GetSession(r)
	if session == nil{
		w.Write([]byte("Unauthorized user"))
		return
	}
	igrd,err:= ingrdHandler.ingrdSrv.Ingredients()
	if err!=nil{
		log.Fatal(err)
	}

	igrdMar,er:= json.MarshalIndent(igrd,"","/t/")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(igrdMar)


}

func (ingrdHandler *IngredientHandler)GetIngredient(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := ingrdHandler.session.GetSession(r)
	if session == nil{
		w.Write([]byte("Unauthorized user"))
		return
	}
	id:= mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)

	if e!=nil{
		log.Fatal(err)
	}

	igrd,err:= ingrdHandler.ingrdSrv.Ingredient(uint(ids))
	if err!=nil{
		log.Fatal(err)
	}

	igrdMar,er:= json.MarshalIndent(igrd,"","/t/")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(igrdMar)
}

func (ingrdHandler *IngredientHandler)CreateIngredients(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := ingrdHandler.session.GetSession(r)
	if session == nil || session.Role != "Admin"{
		w.Write([]byte("Unauthorized user"))
		return
	}
	var ingrd entity.Ingredient
	read,er:= ioutil.ReadAll(r.Body)

	if er!=nil{
		log.Fatal(err)
	}
	e:= json.Unmarshal(read,&ingrd)

	if e!=nil{
		log.Fatal(err)
	}
	ingrdAdd,errs:= ingrdHandler.ingrdSrv.CreateIngredient(ingrd)
	if errs!=nil{
		log.Fatal(err)
	}

	ingrdMar,err:= json.MarshalIndent(ingrdAdd,"","/t/")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(ingrdMar)
}

func (ingrdHandler *IngredientHandler)UpdateIngredient(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	session := ingrdHandler.session.GetSession(r)
	if session == nil{
		w.Write([]byte("Unauthorized user"))
		return
	}
	var ingrd entity.Ingredient
	ids:= mux.Vars(r)["id"]

	id,e:=strconv.Atoi(ids)

	if e!= nil{
		log.Fatal(err)
	}
	read,er:=ioutil.ReadAll(r.Body)

	if er!= nil{
		log.Fatal(err)
	}

	err:=json.Unmarshal(read,&ingrd)
	if err!= nil{
		log.Fatal(err)
	}

	ingrdUpdate,ers:=ingrdHandler.ingrdSrv.UpdateIngredient(uint(id),ingrd)

	if ers!= nil{
		log.Fatal(ers)
	}

	itemMar,errr:=json.MarshalIndent(ingrdUpdate,"","'r/r/")


	if errr!= nil{
		log.Fatal(errr)
	}
	w.Write(itemMar)
}

func (ingrdHandler *IngredientHandler)DeleteIngredient(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	session := ingrdHandler.session.GetSession(r)
	if session == nil{
		w.Write([]byte("Unauthorized user"))
		return
	}
	igrdId:= mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)

	if e!=nil{
		log.Fatal(err)
	}
	items,ers := ingrdHandler.itemsrv.Items()

	if ers!=nil{
		log.Fatal(err)
		return
	}
	for _,item := range items{
		check := false
		for _,i := range item.Ingridients{
			if igrdId != i{
				item.Ingridients = append(item.Ingridients, i)
			}else{
				check = true
			}
		}
		if check {
			item,ers := ingrdHandler.itemsrv.UpdateItem(item)
			if ers!=nil || item==nil{
				log.Fatal(err)
				return
			}

		}
	} 

	ingrd,err:= ingrdHandler.ingrdSrv.DeleteIngredient(uint(ids))
	if err!=nil{
		log.Fatal(err)
	}


	ingrdMar,er:= json.MarshalIndent(ingrd,"","/t/")

	if er!=nil{
		log.Fatal(err)
	}

	w.Write(ingrdMar)
}