package menuService

import (
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/menu"
)

type IngridientGormService struct{
	repo menu.IngridientRepo
}

func NewIngridientGormRepository(repo menu.IngridientRepo)menu.IngridientService{
	return &IngridientGormService{repo:repo}
}



func(ingrd *IngridientGormService) Ingridients()([]entity.Ingridient,[]error){

	ingridients,err:= ingrd.repo.Ingridients()
	if len(err)>0{
		return nil,err
	}
	return ingridients,nil
}

func(ingrd *IngridientGormService) Ingridient(id uint)(*entity.Ingridient,[]error){

	ingridient,err:= ingrd.repo.Ingridient(id)
	if len(err)>0{
		return nil,err
	}
	return ingridient,nil
}
func(ingrd *IngridientGormService)UpdateIngridient(id uint)(*entity.Ingridient,[]error){

	ingridients,err := ingrd.repo.UpdateIngridient(id)
	if len(err)>0{
		return nil,err
	}
	return ingridients,nil
}
func(ingrd *IngridientGormService)DeleteIngridient(id uint)(*entity.Ingridient,[]error){
	
	ingridient,err := ingrd.repo.DeleteIngridient(id)
	if len(err)>0{
		return nil,err
	}
	return ingridient,nil
}
func(ingrd *IngridientGormService)CreateIngridient(id uint)(*entity.Ingridient,[]error){
	
	ingridient,err := ingrd.repo.CreateIngridient(id)
	if len(err)>0{
		return nil,err
	}
	return ingridient,nil
}

