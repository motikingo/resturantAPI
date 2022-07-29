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

	Order "github.com/motikingo/resturant-api/order"
)

type OrderHandler struct{
	odrSrv Order.OrderService
}

func NewOrderHandler(odrSrv Order.OrderService)OrderHandler{
	return OrderHandler{odrSrv: odrSrv}
}

func(odrHan *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	ords,err:= odrHan.odrSrv.Orders()
	if err!=nil{
		log.Fatal(err)
	}
	ordsMar,er:= json.MarshalIndent(ords,"","t/t/")
	if er!=nil{
		log.Fatal(err)
	}

	w.Write(ordsMar)
}

func(odrHan *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	id:=mux.Vars(r)["id"]
	ids,e:=strconv.Atoi(id)
	
	if e!= nil{
		log.Fatal(e)
	}
		

	ord,err:= odrHan.odrSrv.Order(uint(ids))
	if err!=nil{
		log.Fatal(err)
	}
	ordsMar,er:= json.MarshalIndent(ord,"","t/t/")
	if er!=nil{
		log.Fatal(err)
	}

	w.Write(ordsMar)
}

func(odrHan *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var order entity.Order
	read,er:=ioutil.ReadAll(r.Body)
	if er!=nil{
		log.Fatal(err)
	}
	e:=json.Unmarshal(read,&order)
	if e!=nil{
		log.Fatal(err)
	}
	ord,err:= odrHan.odrSrv.CreateOrder(order)
	if err!=nil{
		log.Fatal(err)
	}
	ordsMar,er:= json.MarshalIndent(ord,"","t/t/")
	if er!=nil{
		log.Fatal(err)
	}

	w.Write(ordsMar)
}

func(odrHan *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var order entity.Order
	ids:=mux.Vars(r)["id"]
	id,e:= strconv.Atoi(ids)

	if e!=nil{
		log.Fatal(err)
	}
	read,er:=ioutil.ReadAll(r.Body)

	if er!=nil{
		log.Fatal(err)
	}
	er = json.Unmarshal(read,&order) 

	if er!=nil{
		log.Fatal(err)
	}

	ordernew,errs:=odrHan.odrSrv.UpdateOrder(uint(id),order)

	if errs!=nil{
		log.Fatal(err)
	}
	orderMarsh,errr:= json.MarshalIndent(ordernew,"","/r/r")

	if errr!=nil{
		log.Fatal(err)
	}

	w.Write(orderMarsh)


}

func(odrHan *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type","application/json")
	id:=mux.Vars(r)["id"]
	ids,e:=strconv.Atoi(id)
	if e!=nil{
		log.Fatal(err)
	}
	ords,err:= odrHan.odrSrv.DeleteOrder(uint(ids))
	if err!=nil{
		log.Fatal(err)
	}
	ordsMar,er:= json.MarshalIndent(ords,"","t/t/")
	if er!=nil{
		log.Fatal(err)
	}

	w.Write(ordsMar)
}

// func(odrHan *OrderHandler) {
	
// }