package menurepository

import (
	"github.com/jinzhu/gorm"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/menu"
)

type IngredientGormRepository struct{
	db *gorm.DB
}

func NewIngredientGormRepository(db *gorm.DB)menu.IngredientRepo{
	return &IngredientGormRepository{db:db}
}



func(ingrd *IngredientGormRepository) Ingredients()([]entity.Ingredient,[]error){

	ingredients := []entity.Ingredient{}
	err:= ingrd.db.Find(&ingredients).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return ingredients,nil
}

func(ingrd *IngredientGormRepository) Ingredient(id uint)(*entity.Ingredient,[]error){

	var ingredient entity.Ingredient
	err:= ingrd.db.First(&ingredient,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return &ingredient,nil
}
func(ingrd *IngredientGormRepository)UpdateIngredient(id uint,Ingr entity.Ingredient)(*entity.Ingredient,[]error){

	ingredients,err := ingrd.Ingredient(id)
	if len(err)>0{
		return nil,err
	}
	err= ingrd.db.Find(&ingredients).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return ingredients,nil
}
func(ingrd *IngredientGormRepository)DeleteIngredient(id uint)(*entity.Ingredient,[]error){

	ingredient,err := ingrd.Ingredient(id)
	if len(err)>0{
		return nil,err
	}
	
	err = ingrd.db.Delete(&ingredient,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return ingredient,nil
}
func(ingrdRepo *IngredientGormRepository)CreateIngredient(igrd entity.Ingredient)(*entity.Ingredient,[]error){

	ingredient :=igrd

	err := ingrdRepo.db.Create(&ingredient).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return &ingredient,nil
}

