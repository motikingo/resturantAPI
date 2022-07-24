package menurepository

import (
	"github.com/jinzhu/gorm"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/menu"
)

type IngridientGormRepository struct{
	db *gorm.DB
}

func NewIngridientGormRepository(db *gorm.DB)menu.IngridientRepo{
	return &IngridientGormRepository{db:db}
}



func(ingrd *IngridientGormRepository) Ingridients()([]entity.Ingridient,[]error){

	ingridients := []entity.Ingridient{}
	err:= ingrd.db.Find(&ingridients).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return ingridients,nil
}

func(ingrd *IngridientGormRepository) Ingridient(id uint)(*entity.Ingridient,[]error){

	var ingridient entity.Ingridient
	err:= ingrd.db.First(&ingridient,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return &ingridient,nil
}
func(ingrd *IngridientGormRepository)UpdateIngridient(id uint)(*entity.Ingridient,[]error){

	ingridients,err := ingrd.Ingridient(id)
	if len(err)>0{
		return nil,err
	}
	err= ingrd.db.Find(&ingridients).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return ingridients,nil
}
func(ingrd *IngridientGormRepository)DeleteIngridient(id uint)(*entity.Ingridient,[]error){

	ingridient,err := ingrd.Ingridient(id)
	if len(err)>0{
		return nil,err
	}
	
	err = ingrd.db.Delete(&ingridient,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return ingridient,nil
}
func(ingrd *IngridientGormRepository)CreateIngridient(id uint)(*entity.Ingridient,[]error){

	ingridient,err := ingrd.Ingridient(id)
	if len(err)>0{
		return nil,err
	}
	err = ingrd.db.Create(&ingridient).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return ingridient,nil
}

