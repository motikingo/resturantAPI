package menu

import (
	"github.com/motikingo/resturant-api/entity"
)


type CatagoryRepo interface{
	Catagories()([]entity.Catagory,[]error)
	Catagory(id uint)(*entity.Catagory,[]error)
	IsCatagoryNameExist(name string)bool
	UpdateCatagory(ct entity.Catagory)(*entity.Catagory,[]error)
	DeleteCatagory(id uint)(*entity.Catagory,[]error)
	CreateCatagory(ct entity.Catagory)(*entity.Catagory,[]error)
}

type ItemRepo interface{
	Items()([]entity.Item,[]error)
	Item(id uint)(*entity.Item,[]error)
	IsItemNameExist(name string)bool
	UpdateItem(item entity.Item)(*entity.Item,[]error)
	DeleteItem(id uint)(*entity.Item,[]error)
	CreateItem(item entity.Item)(*entity.Item,[]error)


}

type IngredientRepo interface{
	Ingredients()([]entity.Ingredient,[]error)
	Ingredient(id uint)(*entity.Ingredient,[]error)
	UpdateIngredient(ingrd entity.Ingredient)(*entity.Ingredient,[]error)
	DeleteIngredient(id uint)(*entity.Ingredient,[]error)
	CreateIngredient(ingrd entity.Ingredient)(*entity.Ingredient,[]error)
}