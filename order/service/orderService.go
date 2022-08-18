package OrderService

import (
	"github.com/motikingo/resturant-api/order"
	"github.com/motikingo/resturant-api/entity"
)


type OrderGormService struct{
	repo Order.OrderRespository
}

func NewOrderGormService(repo Order.OrderRespository) Order.OrderService{
	return &OrderGormService{repo: repo}
}

func(odRepo *OrderGormService) Orders()([]entity.Order,[]error){

	orders,err:= odRepo.repo.Orders()
	if len(err)>0{
		return nil,err
	}
	return orders,nil

}

func(odRepo *OrderGormService)Order(id uint)(*entity.Order,[]error){

	order,err:= odRepo.repo.Order(id)
	if len(err)>0{
		return nil,err
	}
	return order,nil
}

func(odRepo *OrderGormService)UpdateOrder(ord entity.Order)(*entity.Order,[]error){

	
	order,err := odRepo.repo.UpdateOrder(ord)
	if len(err)>0{
		return nil,err
	}
	return order,nil
}

func(odRepo *OrderGormService)DeleteOrder(id uint)(*entity.Order,[]error){
	
	orders,err := odRepo.repo.DeleteOrder(id)
	if len(err)>0{
		return nil,err
	}
	return orders,nil
	
}

func(odRepo *OrderGormService)CreateOrder(order entity.Order)(*entity.Order,[]error){

	
	ord ,err := odRepo.repo.CreateOrder(order)
	if len(err)>0{
		return nil,err
	}
	return ord,nil
}

func(odRepo *OrderGormService)CustomerOrders(customer entity.User) ([]entity.Order,[]error){

	orders,err := odRepo.repo.CustomerOrders(customer)
	if len(err)>0{
		return nil,err
	}
	return orders,nil
}
