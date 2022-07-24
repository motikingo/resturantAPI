package menu

import (
	"github.com/motikingo/resturant-api/entity"
)


type CatagoryRepo interface{
	Catagories()([]entity.Catagory,[]error)
	Catagory(id uint)(*entity.Catagory,[]error)
	UpdateCatagory(id uint)(*entity.Catagory,[]error)
	DeleteCatagory(id uint)(*entity.Catagory,[]error)
	CreateCatagory(ct entity.Catagory)(*entity.Catagory,[]error)
}

type ItemRepo interface{
	Items()([]entity.Item,[]error)
	Item(id uint)(*entity.Item,[]error)

	UpdateItem(id uint)(*entity.Item,[]error)

	DeleteItem(id uint)(*entity.Item,[]error)
	CreateItem(item *entity.Item)(*entity.Item,[]error)


}

type IngridientRepo interface{
	Ingridients()([]entity.Ingridient,[]error)
	Ingridient(id uint)(*entity.Ingridient,[]error)
	UpdateIngridient(id uint)(*entity.Ingridient,[]error)
	DeleteIngridient(id uint)(*entity.Ingridient,[]error)
	CreateIngridient(id uint)(*entity.Ingridient,[]error)
}