package menu

import (
	"github.com/motikingo/resturant-api/entity"
)

type CatagorySrv interface{
	Catagories()([]entity.Catagory,[]error)
	Catagory(id uint)(*entity.Catagory,[]error)
	IsCatagoryNameExist(name string)bool
	UpdateCatagory(id uint,ct entity.Catagory)(*entity.Catagory,[]error)
	DeleteCatagory(id uint)(*entity.Catagory,[]error)
	CreateCatagory(ct entity.Catagory)(*entity.Catagory,[]error)
}

type ItemService interface{
	Items()([]entity.Item,[]error)
	Item(id uint)(*entity.Item,[]error)
	IsItemNameExist(name string)bool
	UpdateItem(item entity.Item)(*entity.Item,[]error)
	DeleteItem(id uint)(*entity.Item,[]error)
	CreateItem(item entity.Item)(*entity.Item,[]error)
}

type IngredientService interface{
	Ingredients()([]entity.Ingredient,[]error)
	Ingredient(id uint)(*entity.Ingredient,[]error)
	UpdateIngredient(ingrd entity.Ingredient)(*entity.Ingredient,[]error)
	DeleteIngredient(id uint)(*entity.Ingredient,[]error)
	CreateIngredient(ingrd entity.Ingredient)(*entity.Ingredient,[]error)
}