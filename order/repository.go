package Order

import (
	"github.com/motikingo/resturant-api/entity"
)

type OrderRespository interface{
	Orders()([]entity.Order,[]error)
	Order(id uint)(* entity.Order,[]error)
	UpdateOrder(id uint)(*entity.Order,[]error)
	DeleteOrder(id uint)(*entity.Order,[]error)
	CreateOrder(ord entity.Order)(*entity.Order, []error)
	CustomerOrders(customer entity.User) ([]entity.Order,[]error)
}

