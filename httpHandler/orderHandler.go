package handler

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

	"fmt"
	"log"

	//"fmt"
	//"io/ioutil"
	"net/http"
	//"strconv"

	//"github.com/gorilla/mux"

	"github.com/gorilla/mux"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/helper"
	"github.com/motikingo/resturant-api/menu"
	"github.com/motikingo/resturant-api/user"

	Order "github.com/motikingo/resturant-api/order"
)

type OrderHandler struct {
	odrSrv  Order.OrderService
	itemsrv menu.ItemService
	userSrv user.UserService
	session *SessionHandler
}

func NewOrderHandler(odrSrv Order.OrderService, itemsrv menu.ItemService, userSrv user.UserService, session *SessionHandler) OrderHandler {
	return OrderHandler{odrSrv: odrSrv, itemsrv: itemsrv, userSrv: userSrv, session: session}
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

func (odrHan *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sess := odrHan.session.GetSession(r)
	if sess == nil {
		w.Write(helper.MarshalResponse("UnAuthorized user"))
		return
	}
	id := mux.Vars(r)["id"]
	ids, e := strconv.Atoi(id)

	if e != nil {
		w.Write(helper.MarshalResponse("Internal Server Error"))
		return
	}

	ord, err := odrHan.odrSrv.Order(uint(ids))
	if err != nil || ord == nil {
		w.Write(helper.MarshalResponse("No such Order"))
		log.Fatal(err)
	}

	w.Write(helper.MarshalResponse(ord))
}

func (odrHan *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sess := odrHan.session.GetSession(r)
	response := &struct {
		Status string
		Order  *entity.Order
	}{
		Status: "order create faild",
	}
	input := &struct {
		ItemId int
		Number int
	}{}
	if sess == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}

	read, er := ioutil.ReadAll(r.Body)
	if er != nil {
		w.Write(helper.MarshalResponse(response))
		return
	}

	if e := json.Unmarshal(read, &input); e != nil || input.ItemId == 0 || input.Number < 1 {
		fmt.Println("here")
		w.Write(helper.MarshalResponse(response))
		return
	}
	userId, _ := strconv.Atoi(sess.UserID)
	item, ers := odrHan.itemsrv.Item(uint(input.ItemId))
	if len(ers) > 0 {
		w.Write(helper.MarshalResponse(response))
		return
	}
	if item.Number < input.Number {
		response.Status = "Sorry we don't have enought item come back latter"
		w.Write(helper.MarshalResponse(response))
		return
	}
	bill := item.Price * float64(input.Number)
	order := entity.Order{
		PlaceAt:   time.Now(),
		ItemID:    uint(input.ItemId),
		UserID:    uint(userId),
		Number:    input.Number,
		Orderbill: bill,
	}

	user, _ := odrHan.userSrv.GetUserByID(uint(userId))
	user.Orders = append(user.Orders, order)
	_, ers = odrHan.userSrv.UpdateUser(*user)

	if len(ers) > 0 {
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	ord, err := odrHan.odrSrv.CreateOrder(order)
	if len(err) > 0 {
		user.Orders = user.Orders[:len(user.Orders)-1]
		_, _ = odrHan.userSrv.UpdateUser(*user)
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	item.Number -= input.Number
	_, ers = odrHan.itemsrv.UpdateItem(*item)
	if len(ers) > 0 {
		user.Orders = user.Orders[:len(user.Orders)-1]
		_, _ = odrHan.userSrv.UpdateUser(*user)
		_, _ = odrHan.odrSrv.DeleteOrder(order.ID)
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Status = "successfully ordered"
	response.Order = ord
	w.Write(helper.MarshalResponse(response))

}

func (odrHan *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sess := odrHan.session.GetSession(r)
	response := &struct {
		Status string
		Order  *entity.Order
	}{
		Status: "order create faild",
	}
	input := &struct {
		Number int
	}{}
	if sess == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}

	read, er := ioutil.ReadAll(r.Body)
	if er != nil {
		w.Write(helper.MarshalResponse(response))
		return
	}

	if e := json.Unmarshal(read, &input); e != nil || input.Number < 1 {
		fmt.Println("here")
		w.Write(helper.MarshalResponse(response))
		return
	}
	ordId, _ := strconv.Atoi(mux.Vars(r)["id"])
	ordr, _ := odrHan.odrSrv.Order(uint(ordId))

	userId, _ := strconv.Atoi(sess.UserID)
	item, ers := odrHan.itemsrv.Item(uint((*ordr).ItemID))
	if len(ers) > 0 {
		w.Write(helper.MarshalResponse(response))
		return
	}
	if item.Number < input.Number {
		response.Status = "Sorry we don't have enought item come back latter"
		w.Write(helper.MarshalResponse(response))
		return
	}
	bill := item.Price * float64(input.Number)
	order := entity.Order{
		PlaceAt:   time.Now(),
		ItemID:    uint(ordr.ItemID),
		UserID:    uint(userId),
		Number:    input.Number,
		Orderbill: bill,
	}
	order.ID = ordr.ID
	ord, err := odrHan.odrSrv.UpdateOrder(order)
	if len(err) > 0 {
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	item.Number = item.Number + ordr.Number - input.Number
	_, ers = odrHan.itemsrv.UpdateItem(*item)
	if len(ers) > 0 {
		_, _ = odrHan.odrSrv.UpdateOrder(order)
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Status = " ordered successfully updated"
	response.Order = ord
	w.Write(helper.MarshalResponse(response))

}
func (odrHan *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	sess := odrHan.session.GetSession(r)
	response := &struct {
		Status string
		Order  *entity.Order
	}{
		Status: "order Delete faild",
	}

	if sess == nil {
		w.Write(helper.MarshalResponse(response))
		return
	}

	ordId, _ := strconv.Atoi(mux.Vars(r)["id"])
	ordr, ers := odrHan.odrSrv.DeleteOrder(uint(ordId))

	if len(ers) > 0 {
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}
	item, ers := odrHan.itemsrv.Item(uint((*ordr).ItemID))
	if len(ers) > 0 {
		w.Write(helper.MarshalResponse(response))
		return
	}

	item.Number = item.Number + ordr.Number
	_, ers = odrHan.itemsrv.UpdateItem(*item)
	if len(ers) > 0 {
		_, _ = odrHan.odrSrv.CreateOrder(*ordr)
		response.Status = "Internal Server Error"
		w.Write(helper.MarshalResponse(response))
		return
	}

	response.Status = " ordered successfully deleted"
	response.Order = ordr
	w.Write(helper.MarshalResponse(response))
}
