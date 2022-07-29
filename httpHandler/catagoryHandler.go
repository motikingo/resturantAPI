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
	"github.com/motikingo/resturant-api/menu"
)

type CatagoryHandler struct{
	catSrv menu.CatagorySrv
}

func NewCatagoryHandler(catSrv menu.CatagorySrv)CatagoryHandler{
	return CatagoryHandler{catSrv: catSrv}
}

func(catHandler *CatagoryHandler) GetCatagories(w http.ResponseWriter,r *http.Request,){
	w.Header().Set("Content-Type","application/json")
	handler,err:= catHandler.catSrv.Catagories()
	
	if len(err)>0{
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	
	// catagories:=[] entity.Catagory{
	// 	entity.Catagory{gorm.Model{ID: 1},"moti",[]entity.Item{}},
	// }
	catagor,er := json.MarshalIndent(handler,"","t/t/")
	if er != nil{
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}

	
	w.Write(catagor)
}

func(catHandler *CatagoryHandler) GetCatagory(w http.ResponseWriter,r *http.Request){
	
	w.Header().Set("Content-Type","application/json")
	id:=mux.Vars(r)["id"]
	
	ids,err := strconv.Atoi(id)
	fmt.Println(ids)
	if  err != nil {
		//fmt.Println(ids)
		w.Write([]byte(err.Error()))
		return
	}

	cat,er:= catHandler.catSrv.Catagory(uint(ids))
	
	if len(er)>0{
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	catagories,err := json.MarshalIndent(cat,"","t/t/")

	if err != nil{
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}

	w.Write(catagories)
}



func(catHandler *CatagoryHandler) CreateCatagory(w http.ResponseWriter,r *http.Request){
	var cat entity.Catagory
	read,err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("not read!")
		return 
	}
	err = json.Unmarshal(read,&cat)
	if err != nil{
		fmt.Println("not Unmarshal!")
		log.Fatal(err)

		return
	}

	updatedcat,errs:= catHandler.catSrv.CreateCatagory(cat)
	if errs != nil{
		return
	}
	//updated,_ = ioutil.ReadAll(updatedcat)
	c,_:=json.MarshalIndent(updatedcat,"","/t/t") 
	//err = json.NewDecoder().Decode(&cat)
	if err != nil{
		fmt.Println("not Unmarshal!")

		return
	}

	w.Write(c)


}

func(catHander *CatagoryHandler) UpdateCatagory(w http.ResponseWriter, r *http.Request){

	var cat entity.Catagory
	w.Header().Set("Content-Type","application/json")
	id := mux.Vars(r)["id"]
	ids,e:= strconv.Atoi(id)
	if e!= nil{
		log.Fatal(e)
		return
	}
	read,err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("not read!")
		return 
	}

	err=json.Unmarshal(read,&cat)
	if err != nil{
		fmt.Println("not Unmarshal!")
		log.Fatal(err)

		return
	}

	
	catupdated,er:= catHander.catSrv.UpdateCatagory(uint(ids),cat)
	
	if er!= nil{
		log.Fatal(e)
		return
	}

	catMarsh,e := json.MarshalIndent(catupdated,"","/t/")

	if e != nil{
		log.Fatal(e)
	}

	w.Write(catMarsh)
}


func(catHandler *CatagoryHandler) DeleteCatagory(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	id := mux.Vars(r)["id"]
	ids,_ := strconv.Atoi(id)
	cata,er:= catHandler.catSrv.DeleteCatagory(uint(ids)) 
	if er != nil{
		log.Fatal(er)
	}

	read,e:= json.MarshalIndent(cata,"","/t/t")
	if e != nil{
		log.Fatal(er)
	}
	w.Write(read)
}