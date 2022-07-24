package OrderRespository

import (
	"github.com/jinzhu/gorm"
	"github.com/motikingo/resturant-api/Order"
	"github.com/motikingo/resturant-api/entity"
)


type OrderGormRespository struct{
	db *gorm.DB
}

func NewOrderGormRespository(db *gorm.DB) Order.OrderRespository{
	return &OrderGormRespository{db: db}
}

func(odRepo *OrderGormRespository) Orders()([]entity.Order,[]error){
	orders:= []entity.Order{}
	err:= odRepo.db.Find(&orders).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return orders,nil

}

func(odRepo *OrderGormRespository)Order(id uint)(*entity.Order,[]error){
	var order entity.Order
	err:= odRepo.db.First(&order,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return &order,nil
}

func(odRepo *OrderGormRespository)UpdateOrder(id uint)(*entity.Order,[]error){

	orders,err:=odRepo.Order(id)
	if len(err)>0{
		return nil,err
	}
	err = odRepo.db.Save(&orders).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return orders,nil
}

func(odRepo *OrderGormRespository)DeleteOrder(id uint)(*entity.Order,[]error){
	orders,err:=odRepo.Order(id)
	if len(err)>0{
		return nil,err
	}
	err = odRepo.db.Delete(&orders,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return orders,nil
	
}

func(odRepo *OrderGormRespository)CreateOrder(order *entity.Order)(*entity.Order,[]error){

	ord := order 
	err := odRepo.db.Create(&ord).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return order,nil
}

func(odRepo *OrderGormRespository)CustomerOrders(customer entity.User) ([]entity.Order,[]error){

	orders := []entity.Order{}
	err := odRepo.db.Model(customer).Related(&orders,"Orders").GetErrors()
	if len(err)>0{
		return nil,err
	}
	return orders,nil
}
