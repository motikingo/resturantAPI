package menuService

import (
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/menu"
)

type IngredientGormService struct{
	repo menu.IngredientRepo
}

func NewIngredientGormRepository(repo menu.IngredientRepo)menu.IngredientService{
	return &IngredientGormService{repo:repo}
}



func(ingrdServ *IngredientGormService) Ingredients()([]entity.Ingredient,[]error){

	ingredients,err:= ingrdServ.repo.Ingredients()
	if len(err)>0{
		return nil,err
	}
	return ingredients,nil
}

func(ingrdServ *IngredientGormService) Ingredient(id uint)(*entity.Ingredient,[]error){

	ingredient,err:= ingrdServ.repo.Ingredient(id)
	if len(err)>0{
		return nil,err
	}
	return ingredient,nil
}
func(ingrdServ *IngredientGormService)UpdateIngredient(id uint,ingr entity.Ingredient)(*entity.Ingredient,[]error){

	ingredient,err := ingrdServ.repo.UpdateIngredient(id,ingr)
	if len(err)>0{
		return nil,err
	}
	return ingredient,nil
}
func(ingrdServ *IngredientGormService)DeleteIngredient(id uint)(*entity.Ingredient,[]error){
	
	ingredient,err := ingrdServ.repo.DeleteIngredient(id)
	if len(err)>0{
		return nil,err
	}
	return ingredient,nil
}
func(ingrdServ *IngredientGormService)CreateIngredient(ingrd entity.Ingredient)(*entity.Ingredient,[]error){
	
	ingredient,err := ingrdServ.repo.CreateIngredient(ingrd)
	if len(err)>0{
		return nil,err
	}
	return ingredient,nil
}

