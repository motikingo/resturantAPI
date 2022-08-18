package handler

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

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
	Order "github.com/motikingo/resturant-api/order"
	"github.com/motikingo/resturant-api/menu"

)

type OrderHandler struct{
	odrSrv Order.OrderService
	itemsrv menu.ItemService
	session *SessionHandler
}

func NewOrderHandler(odrSrv Order.OrderService,session *SessionHandler)OrderHandler{
	return OrderHandler{odrSrv: odrSrv,session:session}
}

// func(odrHan *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	ords,err:= odrHan.odrSrv.Orders()
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// 	ordsMar,er:= json.MarshalIndent(ords,"","t/t/")
// 	if er!=nil{
// 		log.Fatal(err)
// 	}

// 	w.Write(ordsMar)
// }

func(odrHan *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := odrHan.session.GetSession(r)
	if sess == nil {
		w.Write(helper.MarshalResponse("UnAuthorized user"))
		return
	}
	id:=mux.Vars(r)["id"]
	ids,e:=strconv.Atoi(id)
	
	if e!= nil{
		w.Write(helper.MarshalResponse("Internal Server Error"))
		return
	}
		
	ord,err:= odrHan.odrSrv.Order(uint(ids))
	if err!=nil || ord == nil{
		w.Write(helper.MarshalResponse("No such Order"))
		log.Fatal(err)
	}
	
	w.Write(helper.MarshalResponse(ord))
}

func(odrHan *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := odrHan.session.GetSession(r)
	response := &struct{
		status string
		order * entity.Order
	}{
		status: "order create faild",
	}
	input := &struct{
		itemId string
		number int
	}{}
	if sess == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}

	read,er:=ioutil.ReadAll(r.Body)
	if er!=nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	e:=json.Unmarshal(read,&input)
	if e!=nil|| input.itemId == "" || input.number < 1 {
		w.Write(helper.MarshalResponse(response))
		return
	}
	itemId,_ := strconv.Atoi(input.itemId)
	userId,_ := strconv.Atoi(sess.UserID)
	item,ers := odrHan.itemsrv.Item(uint(itemId))
	if len(ers)>0{
		w.Write(helper.MarshalResponse(response))
		return
	}
	if item.Number < input.number{
		response.status = "Sorry we don't have enought item come back latter"
		w.Write(helper.MarshalResponse(response))
		return
	}
	bill := item.Price * float64(input.number)
	order := entity.Order{
		PlaceAt: time.Now(),
		ItemID: uint(itemId),
		UserID: uint(userId),
		Number : input.number,
		Orderbill:bill,

	}
	
	ord,err:= odrHan.odrSrv.CreateOrder(order)
	if err!=nil{
		w.Write(helper.MarshalResponse(response))
		return
	}

	item.Number = item.Number - ord.Number

	item,ers = odrHan.itemsrv.UpdateItem(*item)
	if er!=nil{
		response.status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	response.status = "successfully ordered"
	response.order = ord 

	w.Write(helper.MarshalResponse(response))
}

func(odrHan *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := odrHan.session.GetSession(r)
	
	response := &struct{
		status string
		order * entity.Order
	}{
		status: "order update faild",
	}
	input := &struct{
		count int
	}{}

	if sess==nil {
		w.Write(helper.MarshalResponse(response))
		return
	}
	ids:=mux.Vars(r)["id"]
	id,_:= strconv.Atoi(ids)
	odr,ers := odrHan.odrSrv.Order(uint(id))

	if odr==nil || len(ers)>0{
		w.Write(helper.MarshalResponse(response))
		return
	}
	read,er:=ioutil.ReadAll(r.Body)

	if er!=nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	er = json.Unmarshal(read,&input) 

	if er!=nil || input.count < 1{
		w.Write(helper.MarshalResponse(response))
		return
	}

	if odr.Number == input.count{
		response.status = "Nothing change"
		w.Write(helper.MarshalResponse(response))
		return
	}
	odr = &entity.Order{
		PlaceAt: time.Now(),
		Number: input.count,
	}
	odr.ID = uint(id)
	ordernew,errs:=odrHan.odrSrv.UpdateOrder(*odr)

	if len(errs)>0{
		response.status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	response.status = "Successfully updated"
	response.order = ordernew
	w.Write(helper.MarshalResponse(response))
}

func(odrHan *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	sess := odrHan.session.GetSession(r)
	response := &struct{
		status string
		orderId uint
	}{
		status: "Delete order faild",
	}
	if sess == nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	id:=mux.Vars(r)["id"]
	ids,_:=strconv.Atoi(id)
	ord,err:= odrHan.odrSrv.Order(uint(ids))
	if ord==nil || len(err)>0{
		response.status = "No such order"
		w.Write(helper.MarshalResponse(response))
		return
	}
	ord,err = odrHan.odrSrv.DeleteOrder(uint(ids))
	if err!=nil{
		w.Write(helper.MarshalResponse(response))
		return
	}
	
	w.Write(helper.MarshalResponse(ord))
}